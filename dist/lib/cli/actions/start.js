import { logger } from "../../logger/index.js";
import { start } from "../../swarm/index.js";
import { parse, stringify } from "yaml";
import { readFileSync, writeFileSync } from "fs";
import { checkForUpdates } from "../../update.js";
import { makeKeys, encodeKeys, loadKeys, encoder } from "../../bls/keys.js";
import { keys, config as globalConfig, nameRegex } from "../../constants.js";
import { run } from "../../daemon/index.js";
import assert from "node:assert";
export const startAction = async (configFile, options) => {
    const configContent = readFileSync(configFile).toString();
    const config = configContent ? { ...parse(configContent) } : null;
    if (!config) {
        logger.error("Invalid config file");
        return process.exit(1);
    }
    logger.level = options.log || config.log || "info";
    if (!config.secretKey && !options.generate) {
        logger.error("No secret key supplied");
        logger.warn("Run me with --generate to generate a new secret for you");
        return process.exit(1);
    }
    else if (options.generate) {
        const newKeys = makeKeys();
        const encodedKeys = encodeKeys(newKeys);
        config.secretKey = encodedKeys.secretKey;
        const serialized = stringify(config);
        writeFileSync(configFile, serialized);
    }
    await checkForUpdates();
    Object.assign(keys, loadKeys(config));
    Object.assign(globalConfig, config);
    assert(keys.publicKey !== undefined, "No public key available");
    if (!config.name) {
        logger.warn("Node name not found in config");
        logger.warn("Using the first 8 letters of your public key");
        config.name = encoder.encode(keys.publicKey.toBytes()).slice(0, 8);
    }
    else if (config.name.length > 24) {
        logger.error("Node name cannot be more than 24 characters");
        return process.exit(1);
    }
    else if (!config.name.match(nameRegex)) {
        logger.error("Only English letters, numbers, and @._'- are allowed in the name");
        return process.exit(1);
    }
    run();
    start();
};
