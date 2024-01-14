import { gossipMethods, errors, sockets, config } from "../constants.js";
import { Gossip, GossipRequest, MetaData, NodeSystemError } from "../types.js";
import { brotliCompressSync } from "zlib";
import { hashObject } from "../utils/hash.js";
import { cache } from "../utils/cache.js";
import { toMurmur } from "../crypto/murmur/index.js";
import { minutes } from "../utils/time.js";
import { randomDistinct } from "../utils/random.js";

const gossipCache = cache<string, boolean>(minutes(15));

const gossipTo = async (
  nodes: MetaData[],
  data: Gossip<any, any>
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
};

const notSeen = (seen: string[]) => (meta: MetaData) =>
  meta.murmurAddr && !seen.includes(meta.murmurAddr);

const isFree = (meta: MetaData) => !meta.needsDrain;

const randomNodes = (nodes: MetaData[]) =>
  randomDistinct(nodes.length, config.gossip.infect).map(
    (index) => nodes[index]
  );

export const gossip = async (
  request: GossipRequest<any, any>,
  seen: string[]
): Promise<void> => {
  const values = [...sockets.values()] as MetaData[];
  const nodes = seen
    ? values.filter(isFree).filter(notSeen(seen))
    : values.filter(isFree);

  if (!nodes.length) {
    return;
  }

  const selected =
    nodes.length > config.gossip.infect ? randomNodes(nodes) : nodes;

  const selectedMurmurs = selected
    .map((meta) => meta.murmurAddr as string)
    .filter(Boolean);

  // Optimistic:
  const payload = {
    type: "gossip" as const,
    request,
    seen: [...seen, ...selectedMurmurs],
  };

  await gossipTo(selected, payload);
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

    const hash = await toMurmur(hashObject(incoming.request));
    const alreadySeen = gossipCache.get(hash);

    if (alreadySeen) {
      return { error: errors.E_DUPLICATE };
    }

    gossipCache.set(hash, true);
    const method = gossipMethods[methodName];
    await method(incoming.request);
    await gossip(incoming.request, incoming.seen || []);
  } catch (error) {
    const systemError = error as NodeSystemError;
    const message =
      systemError.code || systemError.errno || systemError.message;
    return { error: message || errors.E_INTERNAL };
  }
};
