import { Collection, Db } from "mongodb";

export interface AssetPrice {
  timestamp: Date;
  block: number;
  price: number;
  asset: string;
  source: string;
  signature: string;
  signers: string[];
}

export let assetPrices: Collection<AssetPrice>;

export const initCollection = async (db: Db) => {
  assetPrices = db.collection<AssetPrice>("assetPrices");
  await assetPrices.createIndex({ block: -1, asset: 1 });
};
