import { runPrismaCommand } from "../../../db/manage/index.js";
import { logger } from "../../../logger/index.js";
import { config as globalConfig } from "../../../constants.js";
import { safeReadConfig } from "../../../utils/config.js";

export const generateDbAction = async (configFile: string) => {
  const config = safeReadConfig(configFile);
  if (!config) {
    return process.exit(1);
  }

  if (config.lite || !config.database?.url) {
    logger.error("Database config is invalid.");
    process.exit(1);
  }

  Object.assign(globalConfig, config);
  const exitCode = await runPrismaCommand(["generate"]);

  if (!exitCode) {
    logger.info("Database client generation was successful.");
  } else {
    logger.error("Failed to generation database client.");
  }

  process.exit(exitCode || 0);
};
