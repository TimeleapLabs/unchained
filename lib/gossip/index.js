import { gossipMethods, errors, keys, sockets } from "../constants.js";
import { attest } from "../bls/index.js";
import { encoder } from "../bls/keys.js";
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

export const gossip = async ({ request, signature, signer, seen }) => {
  if (sockets.size === 0) {
    return;
  }
  const publicKey = encoder.encode(keys.publicKey.toBytes());
  if (seen.includes(publicKey)) {
    return;
  }
  const payload = {
    type: "gossip",
    request,
    signature,
    signer,
    seen: [...seen, publicKey],
  };
  const values = [...sockets.values()];
  const nodes = values.filter((node) => !seen.includes(node.publicKey));
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

export const processGossip = async (incoming, connection) => {
  try {
    const { method } = incoming.request;

    if (!(method in gossipMethods)) {
      return { error: errors.E_NOT_FOUND };
    }

    const ok = await gossipMethods[method].call(null, incoming, connection);
    if (!ok) {
      return;
    }

    const payload = attest(incoming.request);

    if (payload) {
      await gossip({ ...payload, seen: incoming.seen });
    }
  } catch (error) {
    return { error: error.code || errors.E_INTERNAL };
  }
};
