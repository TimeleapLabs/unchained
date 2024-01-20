import { getAllScores } from "./index.js";
import { logger } from "../logger/index.js";
import { Table } from "console-table-printer";
import { sockets, keys } from "../constants.js";
import { getSprint } from "../utils/time.js";
import { hashUint8Array, isEqual } from "../utils/uint8array.js";
import { Score } from "./types.js";

export const printScores = async (map: Map<string, Score>) => {
  const table = new Table({
    columns: [
      { name: "peer", title: "Peer", alignment: "left", color: "blue" },
      { name: "name", title: "Name", alignment: "left", color: "yellow" },
      { name: "score", title: "Score", alignment: "center", color: "green" },
    ],
  });

  const thisNode = keys.publicKey.toBytes();
  const rows = [];
  const signerNames = new Map<string, string>();

  for (const peer of sockets.values()) {
    if (peer.publicKey) {
      const hash = peer.murmurAddr || (await hashUint8Array(peer.publicKey));
      signerNames.set(hash, peer.name);
    }
  }

  for (const [murmur, { peer, score }] of getAllScores(map)) {
    if (isEqual(peer, thisNode)) {
      continue;
    }
    const name = signerNames.get(murmur) || "?";
    rows.push({ peer: murmur, score, name });
  }

  if (!rows.length) {
    return;
  }

  table.addRows(rows.sort((a, b) => b.score - a.score));

  const sprint = getSprint();
  logger.info(`Scores for sprint ${sprint} are:`);
  table.printTable();
};
