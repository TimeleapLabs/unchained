import { getAllScores } from "./index.js";
import { logger } from "../logger/index.js";
import { Table } from "console-table-printer";
import { sockets, keys } from "../constants.js";
import { encodeKeys } from "../crypto/bls/keys.js";
import { getSprint } from "../utils/time.js";

export const printScores = (map: Map<string, number>) => {
  const table = new Table({
    columns: [
      { name: "peer", title: "Peer", alignment: "left", color: "blue" },
      { name: "name", title: "Name", alignment: "left", color: "yellow" },
      { name: "score", title: "Score", alignment: "center", color: "green" },
    ],
  });

  const { publicKey: thisNode } = encodeKeys(keys);

  const rows = [];
  const publicKeyMap = new Map(
    [...sockets.entries()].map(([_, { publicKey, name }]) => [publicKey, name])
  );

  for (const [peer, score] of getAllScores(map)) {
    if (peer === thisNode) {
      continue;
    }
    const name = publicKeyMap.get(peer) || "?";
    rows.push({ peer, score, name });
  }

  if (!rows.length) {
    return;
  }

  table.addRows(rows.sort((a, b) => b.score - a.score));

  const sprint = getSprint();
  logger.info(`Scores for sprint ${sprint} are:`);
  table.printTable();
};
