import { readFileSync } from "fs";
import { logger } from "../logger/index.js";
import { Config } from "../types.js";
import { parse } from "yaml";
import { userConfigSchema } from "../schema.js";

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
      logger.error("Config isn't valid YAML");
      return null;
    }

    const schemaCheck = userConfigSchema.safeParse(config);
    if (!schemaCheck.success) {
      for (const issue of schemaCheck.error.issues) {
        logger.error(
          `Config file errors at ${issue.path.join(".")}: ${issue.message}`
        );
      }
      logger.warn("See https://kenshi.io/r/conf for the correct config format");
      return null;
    }

    return config;
  } catch (error) {
    logger.error("Invalid config file");
    return null;
  }
};
