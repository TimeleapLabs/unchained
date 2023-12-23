import { config } from "../constants.js";
import { MongoClient, Db } from "mongodb";
import { logger } from "../logger/index.js";

import assert from "node:assert";
import * as assetPrices from "./collections/AssetPrice.js";
import * as nodeNames from "./collections/NodeNames.js";

export let db: Db;
export let client: MongoClient;

export const initDB = async () => {
  assert(config.database, "Database settings not set");
  client = new MongoClient(config.database.url);
  try {
    await client.connect();
    logger.info("Successfully connected to the database");
  } catch (error) {
    logger.error("Could not connect to the database!");
    return process.exit(1);
  }
  db = client.db(config.database.name);
  await assetPrices.initCollection(db);
  await nodeNames.initCollection(db);
};
