import { runPrismaCommand } from "../../../db/manage/index.js";
import { logger } from "../../../logger/index.js";
import { config as globalConfig } from "../../../constants.js";
import { parse } from "yaml";
import { readFileSync } from "fs";
import { Config } from "../../../types.js";

const safeReadConfig = (configFile: string): string | null => {
  try {
    const configContent = readFileSync(configFile).toString();
    return configContent;
  } catch (error) {
    return null;
  }
};

export const initDbAction = async (configFile: string) => {
  const configContent = safeReadConfig(configFile);
  if (!configContent) {
    logger.error("Failed to read the config file");
    return process.exit(1);
  }

  const config: Config = configContent ? { ...parse(configContent) } : null;
  if (!config) {
    logger.error("Invalid config file");
    return process.exit(1);
  }

  if (config.lite || !config.database?.url) {
    logger.error("Database config is invalid.");
    process.exit(1);
  }

  Object.assign(globalConfig, config);

  const exitCode = await runPrismaCommand(["migrate", "deploy"]);

  if (!exitCode) {
    logger.info("Database migrations were successful.");
  } else {
    logger.error("Failed to run database migrations.");
  }

  process.exit(exitCode || 0);
};
