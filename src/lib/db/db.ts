import { config } from "../constants.js";
import { logger } from "../logger/index.js";
import { PrismaClient } from "../../../prisma/client/index.js";

import assert from "node:assert";

export let db: PrismaClient;

export const initDB = async () => {
  assert(config.database, "Database settings not set");
  db = new PrismaClient({ datasourceUrl: config.database.url });

  try {
    await db.$connect();
    logger.info("Successfully connected to the database");
  } catch (error) {
    logger.info("Failed to connect to the database");
    process.exit(1);
  }
};
