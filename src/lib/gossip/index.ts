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
import { minutes, seconds } from "../utils/time.js";

const ackCache = cache<string, Set<string>>(minutes(15));
const ackTimeoutCache = cache<string, NodeJS.Timeout>(minutes(15));
const gossipCache = cache<string, boolean>(minutes(15));
const ACK_TIMEOUT = seconds(5);

const gossipTo = async (
  nodes: MetaData[],
  data: Gossip<any, any>,
  payloadHash: string
): Promise<void> => {
  const payload = brotliCompressSync(JSON.stringify(data));

  for (const node of nodes) {
    if (!node.socket.closed) {
      const sent = node.socket.write(payload);
      // TODO: Maybe add back pressure config to allow N cached messages
      if (!sent) {
        node.needsDrain = true;
      }
    }
  }
  const ackSeen = ackCache.get(payloadHash)?.values() || [];
  const aggregatedSeen = new Set([...data.seen, ...ackSeen]);
  ackCache.set(payloadHash, aggregatedSeen);
  const maybeTimeout = ackTimeoutCache.get(payloadHash);
  if (maybeTimeout) {
    maybeTimeout?.unref();
    clearTimeout(maybeTimeout);
  }
  ackTimeoutCache.set(
    payloadHash,
    setTimeout(processAck, ACK_TIMEOUT, data.request, payloadHash)
  );
};

const notSeen = (seen: string[]) => (meta: MetaData) =>
  meta.murmurAddr && !seen.includes(meta.murmurAddr);

const isFree = (meta: MetaData) => !meta.needsDrain;

export const gossip = async (
  request: GossipRequest<any, any>,
  seen: string[]
): Promise<void> => {
  const payload = { type: "gossip" as const, request, seen };
  const payloadHash = await toMurmur(hashObject(request));
  const values = [...sockets.values()] as MetaData[];
  const ackSeen = ackCache.get(payloadHash)?.values() || [];
  const aggregatedSeen = [...new Set([...seen, ...ackSeen])];
  const nodes = values.filter(isFree).filter(notSeen(aggregatedSeen));
  if (nodes.length) {
    await gossipTo(nodes, payload, payloadHash);
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

    const hash = await toMurmur(hashObject(incoming.request));
    const alreadySeen = gossipCache.get(hash);

    if (!alreadySeen) {
      gossipCache.set(hash, true);
      const method = gossipMethods[methodName];
      await method(incoming.request);
    }

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

const processAck = async (
  request: GossipRequest<any, any>,
  payloadHash: string
) => {
  const maybeTimeout = ackTimeoutCache.get(payloadHash);
  if (maybeTimeout) {
    maybeTimeout?.unref();
    clearTimeout(maybeTimeout);
  }
  const seen = ackCache.get(payloadHash);
  if (!seen) {
    return;
  }
  const distRequest = {
    method: "dist",
    args: {
      seen: [...seen.values()],
      request,
    },
  };
  if (!seen.has(murmur.address)) {
    distRequest.args.seen.push(murmur.address);
  }
  const distPayload = brotliCompressSync(
    JSON.stringify({
      type: "call",
      request: distRequest,
    })
  );
  const nodes = [...sockets.values()].filter(isFree);
  for (const node of nodes) {
    if (!node.socket.closed) {
      const sent = node.socket.write(distPayload);
      if (!sent) {
        node.needsDrain = true;
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
  const set = ackCache.get(payloadHash);
  if (set) {
    set.add(murmurAddr);
  }
};

const dist = async ({ request, seen }: DistArgs) => {
  await gossip(request, seen);
};

Object.assign(rpcMethods, { ack, dist });
