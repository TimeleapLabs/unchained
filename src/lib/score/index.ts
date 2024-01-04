import { attest, sign } from "../bls/index.js";
import { GossipRequest, GossipMethod } from "../types.js";
import { encodeKeys } from "../bls/keys.js";
import { keys, gossipMethods } from "../constants.js";
import { debounce } from "../utils/debounce.js";
import { getMode } from "../utils/mode.js";
import { Table } from "console-table-printer";
import { logger } from "../logger/index.js";

import {
  ScoreMetric,
  ScoreSignatureInput,
  ScoreValue,
  ScoreValues,
} from "./types.js";

export interface ScoreMap {
  [key: string]: number;
}

const peerScoreMap = new Map<string, number>();
const myScoreMap = new Map<number, ScoreMap>();

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

const printMyScore = debounce((sprint: number) => {
  const sprintScores = myScoreMap.get(sprint) || {};
  const rawScores = Object.values(sprintScores);
  const score = getMode(rawScores);
  const min = Math.min(...rawScores);
  const max = Math.max(...rawScores);
  const average = rawScores.reduce((a, b) => a + b) / rawScores.length;

  console.log({ rawScores, score, min, max });

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

export const scoreAttest: GossipMethod<ScoreMetric, ScoreValues> = async (
  request: GossipRequest<ScoreMetric, ScoreValues>
) => {
  if (!request.payload) {
    return null;
  }

  const { publicKey } = encodeKeys(keys);
  const sprint = Math.ceil(new Date().valueOf() / 300000);

  if (!myScoreMap.has(sprint)) {
    myScoreMap.set(sprint, {});
  }

  const sprintScores = myScoreMap.get(sprint) || {};

  for (const entry of request.payload.value) {
    if (entry.peer === publicKey) {
      if (typeof sprintScores[request.signer] === "undefined") {
        sprintScores[request.signer] = entry.score;
        printMyScore({ key: sprint, args: [sprint] });
        break;
      }
    }
  }

  return request;
};

Object.assign(gossipMethods, { scoreAttest });
