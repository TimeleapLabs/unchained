import { sockets } from "../constants.js";
import { gossipMethods, errors } from "../constants.js";
import { attest } from "../bls/index.js";
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

export const gossip = async ({ request, signature, signers }) => {
  if (sockets.size === 0) {
    return;
  }
  const payload = { type: "gossip", request, signature, signers };
  const nodes = [...sockets.values()].filter(
    (node) => !signers.includes(node.publicKey)
  );
  if (!nodes.length) {
    return;
  }
  if (nodes.length <= 3) {
    await gossipTo(nodes, payload);
  } else {
    const random = new Array(3).fill(nodes.length).map(randomIndex);
    const chosen = [...new Set(random)].map((index) => nodes[index]);
    await gossipTo(chosen, payload);
  }
};

export const processGossip = async (incoming) => {
  const { method } = incoming.request;

  if (!(method in gossipMethods)) {
    return { error: errors.E_NOT_FOUND };
  }

  try {
    const ok = await gossipMethods[method].call(null, incoming);
    if (!ok) {
      return;
    }
    const payload = attest(
      incoming.request,
      incoming.signature,
      incoming.signers
    );
    if (payload) {
      await gossip(payload);
    }
  } catch (error) {
    console.trace(error);
    return { error: error.code || errors.E_INTERNAL };
  }
};
