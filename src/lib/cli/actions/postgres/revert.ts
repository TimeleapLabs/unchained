import { runPrismaCommand } from "../../../db/manage/index.js";
import { logger } from "../../../logger/index.js";
import { config as globalConfig } from "../../../constants.js";
import { safeReadConfig } from "../../../utils/config.js";

export const revertDbAction = async (migration: string, configFile: string) => {
  const config = safeReadConfig(configFile);
  if (!config) {
    return process.exit(1);
  }

  if (config.lite || !config.database?.url) {
    logger.error("Database config is invalid.");
    process.exit(1);
  }

  Object.assign(globalConfig, config);

  const exitCode = await runPrismaCommand([
    "migrate",
    "resolve",
    "--rolled-back ",
    migration,
  ]);

  if (!exitCode) {
    logger.info("Database migration reverted successfully.");
  } else {
    logger.error("Failed to revert the database migration.");
  }

  process.exit(exitCode || 0);
};
