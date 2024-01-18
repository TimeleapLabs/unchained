import { attest } from "../crypto/bls/index.js";
import { WaveRequest } from "../types.js";
import { keys, sockets } from "../constants.js";
import { debounce } from "../utils/debounce.js";
import { getMode } from "../utils/mode.js";
import { Table } from "console-table-printer";
import { logger } from "../logger/index.js";
import { cache } from "../utils/cache.js";
import { db } from "../db/db.js";
import { WantAnswer, WantPacket, datasets } from "../network/index.js";
import { getSprint, minutes, seconds } from "../utils/time.js";
import { isEqual } from "../utils/uint8array.js";
import { encoder } from "../crypto/base58/index.js";
import { toMurmurCached } from "../crypto/murmur/index.js";
import { hashObject } from "../utils/hash.js";

import type { ScoreMetric, ScoreSignatureInput, ScoreValues } from "./types.js";
import { hashUint8Array } from "../utils/uint8array.js";
import { number } from "zod";

export interface ScoreMap {
  [key: string]: ScoreValues;
}

export interface Score {
  score: number;
  peer: Uint8Array;
}

const DATASET = "scores::peers::validations";

const scoreCache = cache<number, ScoreMap>(minutes(15)); // 15 minutes
const upsertCache = cache<number, boolean>(minutes(15));
const waveCache = cache<string, any>(minutes(15));
const peerScoreMap = new Map<string, Score>();
// TODO: Better share this with the UniSwap plugin (and others)
const keyToIdCache = new Map<string, number>();
const keyToNameCache = new Map<string, string | null>();

export const addOnePoint = async (peer: Uint8Array) => {
  const hash = await hashUint8Array(peer);
  const current = peerScoreMap.get(hash) || { score: 0, peer };
  peerScoreMap.set(hash, { ...current, score: current.score + 1 });
};

export const resetAllScores = (
  map: Map<string, Score> = peerScoreMap
): Map<string, Score> => {
  const clone = new Map(map.entries());
  map.clear();
  return clone;
};

export const getAllScores = (map: Map<string, Score> = peerScoreMap) =>
  map.entries();

export const getScoresPayload = async (
  map: Map<string, Score> = peerScoreMap
): Promise<WaveRequest<ScoreMetric, ScoreValues>> => {
  const sprint = getSprint();
  const value: ScoreValues = [];
  for (const score of map.values()) {
    value.push(score);
  }
  const payload: ScoreSignatureInput = { metric: { sprint }, value };
  const signed = attest(payload);
  const data = {
    method: "scoreAttest",
    metric: { sprint },
    dataset: DATASET,
    ...signed,
    payload,
  };
  await scoreAttest([data]);
  return data;
};

const getScoresForAllPeers = (sprintScores: ScoreMap) => {
  const rawScores: { scores: number[]; peer: Uint8Array }[] = [];

  for (const scoreArr of Object.values(sprintScores)) {
    for (const { score, peer } of scoreArr) {
      const index = rawScores.findIndex((item) => isEqual(item.peer, peer));
      if (index > -1) {
        rawScores[index].scores.push(score);
      } else {
        rawScores.push({ peer, scores: [score] });
      }
    }
  }

  return rawScores;
};

const getScoresForPeer = (sprint: number, publicKey: Uint8Array) => {
  const sprintScores = scoreCache.get(sprint) as ScoreMap;
  const rawScores: number[] = [];

  for (const scoreArr of Object.values(sprintScores)) {
    for (const { score, peer } of scoreArr) {
      if (isEqual(publicKey, peer)) {
        rawScores.push(score);
        continue;
      }
    }
  }

  return rawScores;
};

const printMyScore = debounce((sprint: number, publicKey: Uint8Array) => {
  const rawScores = getScoresForPeer(sprint, publicKey);

  if (!rawScores.length) {
    return;
  }

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
}, seconds(5));

export const storeSprintScores = async () => {
  const previousSprint = getSprint() - 1;
  const sprintScores = scoreCache.get(previousSprint);

  if (!sprintScores) {
    return;
  }

  if (upsertCache.get(previousSprint)) {
    return;
  }

  upsertCache.set(previousSprint, true);

  const signerNames = new Map<string, string>();

  for (const peer of sockets.values()) {
    if (peer.publicKey) {
      const hash = peer.murmurAddr || (await hashUint8Array(peer.publicKey));
      signerNames.set(hash, peer.name);
    }
  }

  const allScores = getScoresForAllPeers(sprintScores);

  for (const { peer, scores } of allScores) {
    const hash = await hashUint8Array(peer);
    const oldKey = encoder.encode(peer);

    if (!keyToIdCache.has(hash)) {
      const key = Buffer.from(oldKey);
      const name = signerNames.get(hash);
      const signer = await db.signer.upsert({
        where: { oldKey },
        // see https://github.com/prisma/prisma/issues/18883
        update: { oldKey, key, name },
        create: { oldKey, key, name },
        select: { id: true },
      });
      keyToIdCache.set(hash, signer.id);
    }

    // TODO: We're not doing any verification on this data
    const score = getMode(scores);
    const signerId = keyToIdCache.get(hash) as number;

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

  if (!Object.keys(requests[0]?.payload?.value || {}).length) {
    return null;
  }

  const currentSprint = getSprint();
  const payloadSprint = requests[0].metric.sprint;

  if (currentSprint !== payloadSprint) {
    return null;
  }

  if (!scoreCache.has(payloadSprint)) {
    scoreCache.set(payloadSprint, {});
  }

  const hash = await toMurmurCached(hashObject(requests[0].metric));

  if (!waveCache.has(hash)) {
    waveCache.set(hash, {
      ...requests[0].metric,
      have: new Map<string, any>(),
    });
  }

  const cache = waveCache.get(hash);
  const sprintScores = scoreCache.get(payloadSprint) as ScoreMap;
  const publicKey = keys.publicKey.toBytes();

  let scoreUpdated = false;

  for (const request of requests) {
    if (!request.payload) {
      continue;
    }

    const murmur = await hashUint8Array(request.signer);
    if (cache.have.has(murmur)) {
      continue;
    }

    if (Array.isArray(sprintScores[murmur])) {
      continue;
    }

    cache.have.set(murmur, request);

    for (const { score, peer } of request.payload.value) {
      sprintScores[murmur] = [];

      if (isEqual(peer, publicKey)) {
        scoreUpdated = true;
      }

      sprintScores[murmur].push({ score, peer });
    }
  }

  if (scoreUpdated) {
    printMyScore({ key: payloadSprint, args: [payloadSprint, publicKey] });
  }
};

const have = async (data: WantAnswer) => {
  const cache = waveCache.get(data.want);
  if (!cache) {
    return;
  }
  await scoreAttest(data.have);
};

const want = async (data: WantPacket) => {
  const cache = waveCache.get(data.want);
  if (!cache) {
    return [];
  }
  return [...cache.have.entries()]
    .filter(([murmur]: [string, any]) => !data.have.includes(murmur))
    .map(([_, item]: [string, any]) => item);
};

datasets.set(DATASET, { have, want });

export const getHave = async (want: string) => {
  const cache = waveCache.get(want);
  if (!cache) {
    return [];
  }
  return [...cache.have.keys()];
};
