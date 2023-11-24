import * as Y from "yjs";
import { LeveldbPersistence } from "y-leveldb";

const y = {};

export const addToNode = async (node, data) => {
  y.nodes.getArray(node).push([data]);
  await y.persistence.storeUpdate("nodes", Y.encodeStateAsUpdate(y.nodes));
};

export const replaceNodes = async (update) => {
  await y.persistence.storeUpdate("nodes", update);
  y.nodes = await y.persistence.getYDoc("nodes");
};

export const setup = async (store) => {
  y.persistence = new LeveldbPersistence(store);
  y.nodes = await y.persistence.getYDoc("nodes");
};
