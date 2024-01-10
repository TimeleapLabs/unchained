import {
  gossipMethods,
  errors,
  sockets,
  murmur,
  rpcMethods,
} from "../constants.js";
import { Gossip, GossipRequest, MetaData, NodeSystemError } from "../types.js";
import { brotliCompressSync } from "zlib";
import { hashObject } from "../utils/hash.js";
import { cache } from "../utils/cache.js";
import { Duplex } from "stream";
import { toMurmur } from "../crypto/murmur/index.js";
import { logger } from "../logger/index.js";

const ackCache = cache<string, Set<string>>(5 * 60 * 1000);

const gossipTo = async (
  nodes: MetaData[],
  data: Gossip<any, any>
): Promise<void> => {
  const payload = brotliCompressSync(JSON.stringify(data));
  const payloadHash = await toMurmur(hashObject(data.request));
  for (const node of nodes) {
    if (!node.socket.closed && !node.isSocketBusy) {
      const sent = node.socket.write(payload);
      if (!sent) {
        node.isSocketBusy = true;
      }
    }
  }
  ackCache.set(payloadHash, new Set(data.seen));
  setTimeout(processAck, 60 * 1000, data.request, payloadHash);
};

const filterSeen = (seen: string[]) => (meta: MetaData) =>
  meta.murmurAddr && !seen.includes(meta.murmurAddr);

export const gossip = (
  request: GossipRequest<any, any>,
  seen: string[]
): void => {
  const payload = { type: "gossip" as const, request, seen };
  const values = [...sockets.values()] as MetaData[];
  const nodes = values.filter(filterSeen(seen));
  if (nodes.length) {
    gossipTo(nodes, payload);
  } else {
    logger.info("No peers to gossip to.");
  }
};

export const processGossip = async (
  incoming: Gossip<unknown, unknown>,
  socket: Duplex
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
    await method(incoming.request);

    const ackPayload = brotliCompressSync(
      JSON.stringify({
        type: "call",
        request: {
          method: "ack",
          args: {
            murmurAddr: murmur.address,
            payloadHash: await toMurmur(hashObject(incoming.request)),
          },
        },
      })
    );
    socket.write(ackPayload);
  } catch (error) {
    const systemError = error as NodeSystemError;
    const message =
      systemError.code || systemError.errno || systemError.message;
    return { error: message || errors.E_INTERNAL };
  }
};

const processAck = (request: GossipRequest<any, any>, payloadHash: string) => {
  logger.info(`Processing ack ${payloadHash}`);
  const seen = ackCache.get(payloadHash);
  if (!seen) {
    return;
  }
  const distRequest = {
    method: "dist",
    args: {
      seen: [...seen.values(), murmur.address],
      request,
    },
  };
  const distPayload = brotliCompressSync(
    JSON.stringify({
      type: "call",
      request: distRequest,
    })
  );
  for (const node of sockets.values()) {
    if (!node.socket.closed && !node.isSocketBusy) {
      const sent = node.socket.write(distPayload);
      if (!sent) {
        node.isSocketBusy = true;
      }
    }
  }
};

type AckArgs = {
  payloadHash: string;
  murmurAddr: string;
};

type DistArgs = {
  request: GossipRequest<any, any>;
  seen: string[];
};

const ack = ({ payloadHash, murmurAddr }: AckArgs) => {
  logger.info(`Ack ${payloadHash} received from ${murmurAddr}`);
  const set = ackCache.get(payloadHash);
  if (set) {
    set.add(murmurAddr);
  }
};

const dist = async ({ request, seen }: DistArgs) => {
  const payloadHash = await toMurmur(hashObject(request));
  logger.info(
    `Dist received for ${payloadHash}. Seen: ${JSON.stringify(seen)}`
  );
  gossip(request, seen);
};

Object.assign(rpcMethods, { ack, dist });
