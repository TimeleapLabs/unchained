import { sockets, rpcMethods, config } from "../constants.js";
import { MetaData } from "../types.js";
import { chunks } from "../utils/array.js";
import { epoch, jitter, seconds } from "../utils/time.js";
import { randomDistinct } from "../utils/random.js";
import { compress } from "snappy";
import { serialize } from "../utils/sia.js";
import { logger } from "../logger/index.js";

export interface WantPacket {
  dataset: string;
  want: string;
  have: string[];
}

export interface WantAnswer {
  dataset: string;
  want: string;
  have: any[];
}

export interface Dataset {
  want: (data: WantPacket) => any[] | Promise<any[]>;
  have: (data: WantAnswer) => void | Promise<void>;
}

export const datasets = new Map<string, Dataset>();

const writePayload = async (nodes: MetaData[], payload: Buffer) => {
  for (const node of nodes) {
    if (!node.socket.closed) {
      node.lastSocketWrite = epoch();
      const sent = node.socket.write(payload);
      // TODO: Maybe add back pressure config to allow N cached messages
      if (!sent) {
        node.needsDrain = true;
      }
      if (config.waves.jitter.max) {
        await jitter(config.waves.jitter.max, config.waves.jitter.min);
      }
    }
  }
};

const wantRpcCall = async (nodes: MetaData[], data: WantPacket) => {
  const call = { type: "call", request: { method: "want", args: data } };
  const payload = await compress(serialize(call));
  await writePayload(nodes, payload);
};

const haveRpcCall = async (nodes: MetaData[], data: WantAnswer) => {
  const call = { type: "call", request: { method: "have", args: data } };
  const serialized = serialize(call);
  const compressed = await compress(serialized);
  await writePayload(nodes, compressed);
};

const notBusy = (node: MetaData) => !node.needsDrain;

const notWaitingFor = (identifier: string) => (node: MetaData) =>
  !node.rpcRequests.has(identifier);

// TODO: Should this be configurable?
const notOverwhelmed = (node: MetaData) => node.rpcRequests.size < Infinity;

// TODO: Should this be configurable?
const notTooFast = (node: MetaData) =>
  !node.lastSocketWrite || node.lastSocketWrite + seconds(0.125) < epoch();

export const queryNetworkFor = async (
  want: string,
  dataset: string,
  getHave: (want: string) => Promise<any>
): Promise<boolean> => {
  const id = `${dataset}::${want}`;

  const nodes = [...sockets.values()]
    .filter(notBusy)
    .filter(notTooFast)
    .filter(notOverwhelmed)
    .filter(notWaitingFor(id));

  if (!nodes.length) {
    return false;
  }

  const count = Math.floor((sockets.size * config.waves.select) / 100);

  const selected =
    count >= nodes.length * 2
      ? randomDistinct(nodes.length, count).map((index) => nodes[index])
      : nodes;

  const groups = chunks(selected, config.waves.group);

  for (const group of groups) {
    const have = await getHave(want);
    const packet: WantPacket = { want, dataset, have };

    for (const node of group) {
      node.rpcRequests.add(id);
    }

    await wantRpcCall(group, packet);
  }

  return true;
};

const want = async (data: WantPacket, sender: MetaData) => {
  const dataset = datasets.get(data.dataset);
  if (!dataset) {
    return;
  }
  const have = await dataset.want(data);
  if (!have.length) {
    return;
  }
  const packet: WantAnswer = { dataset: data.dataset, want: data.want, have };
  await haveRpcCall([sender], packet);
};

const have = async (data: WantAnswer, sender: MetaData) => {
  const dataset = datasets.get(data.dataset);
  if (!dataset) {
    return;
  }
  sender.rpcRequests.delete(`${data.dataset}::${data.want}`);
  logger.debug(`Peer ${sender.name} has fulfilled ${data.want}`);
  await dataset.have(data);
};

Object.assign(rpcMethods, { want, have });
