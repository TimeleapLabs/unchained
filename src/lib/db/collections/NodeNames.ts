import { Collection, Db } from "mongodb";

export interface NodeNames {
  address: "string";
  name: "string";
}

export let nodeNames: Collection<NodeNames>;

export const initCollection = async (db: Db) => {
  nodeNames = db.collection<NodeNames>("nodeNames");
  await nodeNames.createIndex({ address: 1, name: 1 });
};
