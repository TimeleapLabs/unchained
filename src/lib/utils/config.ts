import { readFileSync } from "fs";
import { logger } from "../logger/index.js";
import { Config } from "../types.js";
import { parse } from "yaml";

const readFileSafe = (file: string) => {
  try {
    return readFileSync(file).toString();
  } catch (error) {
    return null;
  }
};

export const safeReadConfig = (configFile: string): Config | null => {
  try {
    const configContent = readFileSafe(configFile);

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
    logger.error("Invalid config file");
    return null;
  }
};
