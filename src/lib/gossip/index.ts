import {
  gossipMethods,
  errors,
  sockets,
  config,
  murmur,
} from "../constants.js";
import { Gossip, GossipRequest, MetaData, NodeSystemError } from "../types.js";
import { brotliCompressSync } from "zlib";

import crypto from "crypto";
import assert from "node:assert";

const randomIndex = (length: number): number => {
  if (length <= 0) {
    throw new Error("Array length must be greater than 0.");
  }

  let index: number, randomByte: number;

  do {
    randomByte = crypto.randomBytes(1)[0];
    index = randomByte % length;
  } while (randomByte - index >= 256 - (256 % length));

  return index;
};

const gossipTo = async (nodes: MetaData[], data: any): Promise<void> => {
  const payload = brotliCompressSync(JSON.stringify(data));
  for (const node of nodes) {
    if (!node.socket.closed && !node.isSocketBusy) {
      const sent = node.socket.write(payload);
      if (!sent) {
        node.isSocketBusy = true;
      }
    }
  }
};

const randomDistinct = (length: number, count: number): number[] => {
  const set = new Set<number>();
  while (set.size < count) {
    const random = randomIndex(length);
    set.add(random);
  }
  return [...set];
};

export const gossip = (
  request: GossipRequest<any, any>,
  seen: string[]
): void => {
  if (sockets.size === 0) {
    return;
  }
  assert(murmur.address !== "", "No murmur address found");
  if (seen.includes(murmur.address)) {
    return;
  }
  const payload = { type: "gossip", request, seen: [...seen, murmur.address] };
  const values = [...sockets.values()] as MetaData[];
  const nodes = values
    .filter((node) => !node.isSocketBusy)
    .filter((node) => node.murmurAddr && !seen.includes(node.murmurAddr));
  if (!nodes.length) {
    return;
  }
  if (nodes.length <= config.gossip) {
    gossipTo(nodes, payload);
  } else {
    const indexes = randomDistinct(nodes.length, config.gossip);
    const chosen = indexes.map((index) => nodes[index]);
    gossipTo(chosen, payload);
  }
};

export const processGossip = async (
  incoming: Gossip<unknown, unknown>
): Promise<void | { error?: string | number }> => {
  try {
    // TODO: We should detect and slash nodes if they send wrong data
    const { method: methodName } = incoming.request;

    /**
     * INFO: we're using `hasOwnProperty` instead of `in` to prevent attacks
     * where a peer can access the internal properties of `gossipMethods`.
     */
    if (!gossipMethods.hasOwnProperty(methodName)) {
      return { error: errors.E_NOT_FOUND };
    }

    const method = gossipMethods[methodName];
    const payload = await method(incoming.request);
    if (payload) {
      gossip(payload, incoming.seen);
    }
  } catch (error) {
    const systemError = error as NodeSystemError;
    const message =
      systemError.code || systemError.errno || systemError.message;
    return { error: message || errors.E_INTERNAL };
  }
};
