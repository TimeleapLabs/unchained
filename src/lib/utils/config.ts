import { readFileSync } from "fs";
import { logger } from "../logger/index.js";
import { Config } from "../types.js";
import { parse } from "yaml";

export const safeReadConfig = (configFile: string): Config | null => {
  try {
    const configContent = readFileSync(configFile).toString();

    if (!configContent) {
      logger.error("Failed to read the config file");
      return null;
    }

    const config: Config = configContent ? { ...parse(configContent) } : null;
    if (!config) {
      logger.error("Invalid config file");
      return null;
    }

    return config;
  } catch (error) {
    return null;
  }
};
