import { logger } from "../../logger/index.js";
import { startSwarm } from "../../swarm/index.js";
import { parse, stringify } from "yaml";
import { readFileSync, writeFileSync } from "fs";
import { checkForUpdates } from "../../update.js";
import { makeKeys, encodeKeys, loadKeys, encoder } from "../../bls/keys.js";
import { keys, config as globalConfig, nameRegex } from "../../constants.js";
import { runTasks } from "../../daemon/index.js";
import { Config } from "../../types.js";
import { initDB } from "../../db/db.js";
import assert from "node:assert";

interface StartOptions {
  log?: string;
  lite?: boolean;
  generate?: boolean;
  gossip?: string;
  maxPeers?: string;
  parallelPeers?: string;
}

const safeReadConfig = (configFile: string): string | null => {
  try {
    const configContent = readFileSync(configFile).toString();
    return configContent;
  } catch (error) {
    return null;
  }
};

export const startAction = async (
  configFile: string,
  options: StartOptions
) => {
  const configContent = safeReadConfig(configFile);
  const config: Config = configContent ? { ...parse(configContent) } : null;

  if (!config) {
    logger.error("Invalid config file");
    return process.exit(1);
  }

  logger.level = options.log || config.log || "info";
  config.lite = options.lite || config.lite || false;

  config.gossip =
    parseInt(options.gossip || "0") || config.gossip || globalConfig.gossip;

  config.peers ||= globalConfig.peers;
  config.peers.max = parseInt(options.maxPeers || "0") || config.peers.max;
  config.peers.parallel =
    parseInt(options.parallelPeers || "0") || config.peers.parallel;

  if (!config.secretKey && !options.generate) {
    logger.error("No secret key supplied");
    logger.warn("Run me with --generate to generate a new secret for you");
    return process.exit(1);
  }

  if (!config.secretKey && options.generate) {
    const newKeys = makeKeys();
    const encodedKeys = encodeKeys(newKeys);
    config.secretKey = encodedKeys.secretKey;
    config.publicKey = encodedKeys.publicKey;
    const serialized = stringify(config);
    writeFileSync(configFile, serialized);
  }

  await checkForUpdates();

  Object.assign(keys, loadKeys(config.secretKey));
  assert(keys.publicKey !== undefined, "No public key available");

  const address = encoder.encode(keys.publicKey.toBytes());
  logger.info(`Unchained public address is ${address}`);

  if (!config.name) {
    logger.warn("Node name not found in config");
    logger.warn("Using the first 8 letters of your public key");
    config.name = address.slice(0, 8);
  } else if (config.name.length > 24) {
    logger.error("Node name cannot be more than 24 characters");
    return process.exit(1);
  } else if (!config.name.match(nameRegex)) {
    logger.error(
      "Only English letters, numbers, and @._'- are allowed in the name"
    );
    return process.exit(1);
  }

  if (!config.lite && !config.database?.url) {
    logger.error("Database URL is not provided.");
    return process.exit(1);
  }

  if (!config.lite && !config.database?.name) {
    config.database = { url: config.database?.url || "", name: "unchained" };
  }

  Object.assign(globalConfig, config);

  if (!config.lite) {
    await initDB();
  }

  runTasks();
  startSwarm();
};
