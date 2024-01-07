import { attest, sign } from "../bls/index.js";
import { GossipRequest, GossipMethod } from "../types.js";
import { encodeKeys } from "../bls/keys.js";
import { keys, gossipMethods, sockets } from "../constants.js";
import { debounce } from "../utils/debounce.js";
import { getMode } from "../utils/mode.js";
import { Table } from "console-table-printer";
import { logger } from "../logger/index.js";
import { cache } from "../utils/cache.js";
import { db } from "../db/db.js";

import {
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
const peerScoreMap = new Map<string, number>();
// TODO: Better share this with the UniSwap plugin (and others)
const keyToIdCache = new Map<string, number>();

export const addOnePoint = (peer: string) => {
  const current = peerScoreMap.get(peer) || 0;
  peerScoreMap.set(peer, current + 1);
};

export const resetScore = (
  peer: string,
  map: Map<string, number> = peerScoreMap
) => map.set(peer, 0);

export const resetAllScores = (
  map: Map<string, number> = peerScoreMap
): Map<string, number> => {
  const clone = new Map(map.entries());
  map.clear();
  return clone;
};

export const getScoreOf = (
  peer: string,
  map: Map<string, number> = peerScoreMap
): number => map.get(peer) || 0;

export const getAllScores = (map: Map<string, number> = peerScoreMap) =>
  map.entries();

export const getScoresPayload = (
  map: Map<string, number> = peerScoreMap
): GossipRequest<ScoreMetric, ScoreValues> => {
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
  return {
    method: "scoreAttest",
    metric: { sprint },
    dataset: "scores::peers::validations",
    ...signed,
    payload,
  };
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
}, 2500);

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
    if (!keyToIdCache.has(key)) {
      const name = signerNames.get(key);
      const signer = await db.signer.upsert({
        where: { key },
        // see https://github.com/prisma/prisma/issues/18883
        update: { key },
        create: { key, name },
        select: { id: true },
      });
      keyToIdCache.set(key, signer.id);
    }
  }

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

export const scoreAttest: GossipMethod<ScoreMetric, ScoreValues> = async (
  request: GossipRequest<ScoreMetric, ScoreValues>
) => {
  if (!request.payload) {
    return null;
  }

  if (!request.payload.value.length) {
    return null;
  }

  const currentSprint = Math.ceil(new Date().valueOf() / 300000);
  const payloadSprint = request.payload.value[0].sprint;

  if (currentSprint !== payloadSprint) {
    return null;
  }

  if (!scoreCache.has(payloadSprint)) {
    scoreCache.set(payloadSprint, {});
  }

  const sprintScores = scoreCache.get(payloadSprint) as ScoreMap;
  const { publicKey } = encodeKeys(keys);

  if (sprintScores[publicKey]?.[request.signer]) {
    return null;
  }

  for (const entry of request.payload.value) {
    sprintScores[entry.peer] ||= {};
    sprintScores[entry.peer][request.signer] = entry.score;
  }

  printMyScore({ key: payloadSprint, args: [payloadSprint, publicKey] });

  return request;
};

Object.assign(gossipMethods, { scoreAttest });
