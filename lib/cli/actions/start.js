import { logger } from "../../logger/index.js";
import { start } from "../../swarm/index.js";
import { parse, stringify } from "yaml";
import { readFileSync, writeFileSync } from "fs";
import { untilde } from "../../utils/untilde.js";
import { checkForUpdates } from "../../update.js";
import { makeKeys, encodeKeys, loadKeys, encoder } from "../../bls/keys.js";
import { keys, config as globalConfig, nameRegex } from "../../constants.js";
import { run } from "../../daemon/index.js";

import * as nodes from "../../repo/nodes.js";

export const startAction = async (configFile, options) => {
  const config = { ...parse(readFileSync(configFile).toString()) };

  if (!config) {
    logger.error("Invalid config file");
    return process.exit(1);
  }

  logger.level = options.log || config.log || "info";

  if (!config.secretKey && !options.generate) {
    logger.error("No secret key supplied");
    logger.warn("Run me with --generate to generate a new secret for you");
    return process.exit(1);
  } else if (options.generate) {
    const keys = encodeKeys(makeKeys());
    config.secretKey = keys.secretKey;
    const serialized = stringify(config);
    writeFileSync(configFile, serialized);
  }

  await checkForUpdates();
  await nodes.setup(untilde(options.store || config.store));

  Object.assign(keys, loadKeys(config));
  Object.assign(globalConfig, config);

  if (!config.name) {
    logger.warn("Node name not found in config");
    logger.warn("Using the first 8 letters of your public key");
    config.name = encoder.encode(keys.publickKey.toBytes()).slice(0, 8);
  } else if (config.name.length > 24) {
    logger.error("Node name cannot be more than 24 characters");
    return process.exit(1);
  } else if (!config.name.match(nameRegex)) {
    logger.error(
      "Only English letters, numbers, and @._'- are allowed in the name"
    );
    return process.exit(1);
  }

  run();
  start();
};
