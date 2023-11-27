import { sockets } from "../constants.js";
import crypto from "crypto";

const randomIndex = (length) => {
  if (length <= 0) {
    throw new Error("Array length must be greater than 0.");
  }

  let index, randomByte;

  do {
    randomByte = crypto.randomBytes(1)[0];
    index = randomByte % length;
  } while (randomByte - index >= 256 - (256 % length));

  return index;
};

const gossipTo = async (nodes, data) => {
  const promises = [];
  for (const { socket } of nodes) {
    promises.push(socket.write(JSON.stringify(data)));
  }
  await Promise.all(promises).catch(() => null);
};

export const gossip = async (data) => {
  if (sockets.size === 0) {
    return;
  }
  const { signers } = data.args[1];
  const nodes = [...sockets.values()].filter(
    (node) => !signers.includes(node.publicKey)
  );
  if (nodes.length <= 3) {
    await gossipTo(nodes, data);
  } else {
    const random = new Array(3).fill(sockets.size).map(randomIndex);
    const chosen = [...new Set(random)].map((index) => nodes[index]);
    await gossipTo(chosen, data);
  }
};
