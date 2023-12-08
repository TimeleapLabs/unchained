import { Collection, Db } from "mongodb";

export interface AssetPrice {
  timestamp: Date;
  block: number;
  price: BigInt;
  asset: string;
  source: string;
  signature: string;
  signers: string[];
}

export let assetPrices: Collection<AssetPrice>;

export const initCollection = (db: Db) => {
  assetPrices = db.collection<AssetPrice>("assetPrices");
};
