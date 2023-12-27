import { runPrismaCommand } from "../../../db/manage/index.js";
import { logger } from "../../../logger/index.js";
import { config as globalConfig } from "../../../constants.js";
import { parse } from "yaml";
import { readFileSync } from "fs";
const safeReadConfig = (configFile) => {
    try {
        const configContent = readFileSync(configFile).toString();
        return configContent;
    }
    catch (error) {
        return null;
    }
};
export const generateDbAction = async (configFile) => {
    const configContent = safeReadConfig(configFile);
    const config = configContent ? { ...parse(configContent) } : null;
    if (!config) {
        logger.error("Invalid config file");
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
    }
    else {
        logger.error("Failed to generation database client.");
    }
    process.exit(exitCode || 0);
};
