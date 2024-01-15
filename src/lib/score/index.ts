import { attest, sign } from "../crypto/bls/index.js";
import { WaveRequest } from "../types.js";
import { encodeKeys } from "../crypto/bls/keys.js";
import { keys, sockets } from "../constants.js";
import { debounce } from "../utils/debounce.js";
import { getMode } from "../utils/mode.js";
import { Table } from "console-table-printer";
import { logger } from "../logger/index.js";
import { cache } from "../utils/cache.js";
import { db } from "../db/db.js";
import { WantAnswer, WantPacket, datasets } from "../network/index.js";
import { minutes } from "../utils/time.js";

import { toMurmur } from "../crypto/murmur/index.js";
import { hashObject } from "../utils/hash.js";

import type {
  ScoreMetric,
  ScoreSignatureInput,
  ScoreValue,
  ScoreValues,
} from "./types.js";

export interface ScoreMap {
  [key: string]: { [key: string]: number };
}

const scoreCache = cache<number, ScoreMap>(15 * 60 * 1000); // 15 minutes
const upsertCache = cache<number, boolean>(15 * 60 * 1000);
const wantCache = cache<string, any>(minutes(15));
const peerScoreMap = new Map<string, number>();
// TODO: Better share this with the UniSwap plugin (and others)
const keyToIdCache = new Map<string, number>();
const keyToNameCache = new Map<string, string | null>();

export const addOnePoint = (peer: string) => {
  const current = peerScoreMap.get(peer) || 0;
  peerScoreMap.set(peer, current + 1);
};

export const resetAllScores = (
  map: Map<string, number> = peerScoreMap
): Map<string, number> => {
  const clone = new Map(map.entries());
  map.clear();
  return clone;
};

export const getAllScores = (map: Map<string, number> = peerScoreMap) =>
  map.entries();

export const getScoresPayload = async (
  map: Map<string, number> = peerScoreMap
): Promise<WaveRequest<ScoreMetric, ScoreValues>> => {
  const sprint = Math.ceil(new Date().valueOf() / 300000);
  const value: ScoreValues = [];
  for (const [peer, score] of map.entries()) {
    const toSign = { peer, score, sprint };
    const signature = sign(toSign);
    const data: ScoreValue = { ...toSign, signature };
    value.push(data);
  }
  const payload: ScoreSignatureInput = { metric: { sprint }, value };
  const signed = attest(payload);
  const data = {
    method: "scoreAttest",
    metric: { sprint },
    dataset: "scores::peers::validations",
    ...signed,
    payload,
  };
  await scoreAttest([data]);
  return data;
};

const printMyScore = debounce((sprint: number, publicKey: string) => {
  const sprintScores = scoreCache.get(sprint) as ScoreMap;
  if (!sprintScores[publicKey]) {
    return;
  }
  const rawScores = Object.values(sprintScores[publicKey]);
  const score = getMode(rawScores);
  const min = Math.min(...rawScores);
  const max = Math.max(...rawScores);
  const average = rawScores.reduce((a, b) => a + b) / rawScores.length;

  const table = new Table({
    columns: [
      { name: "sprint", title: "Sprint", alignment: "center", color: "blue" },
      {
        name: "attestations",
        title: "Attestations",
        alignment: "center",
        color: "green",
      },
      { name: "score", title: "Score", alignment: "center", color: "green" },
      {
        name: "average",
        title: "Average",
        alignment: "center",
        color: "yellow",
      },
      { name: "max", title: "Max", alignment: "center", color: "green" },
      { name: "min", title: "Min", alignment: "center", color: "red" },
    ],
  });

  table.addRow({
    sprint,
    attestations: rawScores.length,
    score,
    min,
    max,
    average: average.toFixed(4),
  });

  logger.info("Score received from peers");
  table.printTable();
}, 1000);

export const storeSprintScores = async () => {
  const previousSprint = Math.ceil(new Date().valueOf() / 300000) - 1;
  const sprintScores = scoreCache.get(previousSprint);
  if (!sprintScores) {
    return;
  }
  if (upsertCache.get(previousSprint)) {
    return;
  }
  upsertCache.set(previousSprint, true);

  const signerNames = new Map(
    [...sockets.values()].map((item) => [item.publicKey, item.name])
  );

  for (const key of Object.keys(sprintScores)) {
    const name = signerNames.get(key);
    if (!keyToIdCache.has(key) || keyToNameCache.get(key) !== name) {
      const signer = await db.signer.upsert({
        where: { key },
        // see https://github.com/prisma/prisma/issues/18883
        update: { key, name },
        create: { key, name },
        select: { id: true, name: true },
      });
      keyToIdCache.set(key, signer.id);
      keyToNameCache.set(key, signer.name);
    }
  }

  // Update node names

  for (const peer of Object.keys(sprintScores)) {
    // TODO: We're not doing any verification on this data
    const scores = Object.values(sprintScores[peer]);
    const score = getMode(scores);
    const signerId = keyToIdCache.get(peer) as number;
    await db.sprintPoint.upsert({
      where: { signerId_sprint: { signerId, sprint: previousSprint } },
      update: { score },
      create: { signerId, sprint: previousSprint, score },
    });
    await db.signer.update({
      where: { id: signerId },
      data: { points: { increment: score } },
    });
  }
};

const scoreAttest = async (
  requests: WaveRequest<ScoreMetric, ScoreValues>[]
) => {
  if (!requests.length) {
    return null;
  }

  if (!requests.every((req) => req.payload)) {
    return null;
  }

  if (!requests[0]?.payload?.value.length) {
    return null;
  }

  const currentSprint = Math.ceil(new Date().valueOf() / 300000);
  const payloadSprint = requests[0].payload.value[0].sprint;

  if (currentSprint !== payloadSprint) {
    return null;
  }

  if (!scoreCache.has(payloadSprint)) {
    scoreCache.set(payloadSprint, {});
  }

  const hash = await toMurmur(hashObject(requests[0].metric));

  if (!wantCache.has(hash)) {
    wantCache.set(hash, { ...requests[0].metric, have: [] });
  }

  const murmurMap = new Map(
    [...sockets.values()].map((meta) => [meta.publicKey, meta.murmurAddr])
  );

  const cache = wantCache.get(hash);
  const sprintScores = scoreCache.get(payloadSprint) as ScoreMap;

  for (const request of requests) {
    if (!request.payload) {
      continue;
    }

    const murmur =
      murmurMap.get(request.signer) || (await toMurmur(request.signer));

    cache.have.push({ murmur, request });

    for (const entry of request.payload.value) {
      sprintScores[entry.peer] ||= {};
      sprintScores[entry.peer][request.signer] = entry.score;
    }
  }

  const { publicKey } = encodeKeys(keys);
  printMyScore({ key: payloadSprint, args: [payloadSprint, publicKey] });
};

const have = async (data: WantAnswer) => {
  const cache = wantCache.get(data.want);
  if (!cache) {
    return;
  }
  await scoreAttest(data.have.map((item) => item.request));
};

const want = async (data: WantPacket) => {
  const cache = wantCache.get(data.want);
  if (!cache) {
    return [];
  }
  return cache.have.filter((item: any) => !data.have.includes(item.murmur));
};

datasets.set("scores::peers::validations", { have, want });

export const getHave = async (want: string) => {
  const cache = wantCache.get(want);
  if (!cache) {
    return [];
  }
  return cache.have.map((item: { murmur: string }) => item.murmur);
};
