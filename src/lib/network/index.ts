import { sockets, rpcMethods, config } from "../constants.js";
import { MetaData } from "../types.js";
import { chunks } from "../utils/array.js";
import { jitter } from "../utils/time.js";
import { randomDistinct } from "../utils/random.js";
import { compress } from "snappy";
import { serialize } from "../utils/sia.js";

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
  const payload = await compress(serialize(call));
  await writePayload(nodes, payload);
};

const isFree = (node: MetaData) => !node.needsDrain;

export const queryNetworkFor = async (
  want: string,
  dataset: string,
  getHave: (want: string) => Promise<any>
) => {
  const nodes = [...sockets.values()].filter(isFree);
  const count = Math.floor((nodes.length * config.waves.select) / 100);
  const selected =
    count >= nodes.length
      ? randomDistinct(nodes.length, count).map((index) => nodes[index])
      : nodes;
  const groups = chunks(selected, config.waves.group);
  for (const group of groups) {
    const have = await getHave(want);
    const packet: WantPacket = { want, dataset, have };
    await wantRpcCall(group, packet);
  }
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

const have = async (data: WantAnswer) => {
  const dataset = datasets.get(data.dataset);
  if (!dataset) {
    return;
  }
  await dataset.have(data);
};

Object.assign(rpcMethods, { want, have });
