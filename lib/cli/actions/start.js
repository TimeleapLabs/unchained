import { logger } from "../../logger/index.js";
import { start } from "../../swarm/index.js";
import { parse, stringify } from "yaml";
import { readFileSync, writeFileSync } from "fs";
import { untilde } from "../../utils/untilde.js";
import { checkForUpdates } from "../../update.js";
import { makeKeys, encodeKeys, loadKeys } from "../../bls/keys.js";
import { keys } from "../../constants.js";

import * as nodes from "../../repo/nodes.js";
import * as rpc from "../../rpc/index.js";
import * as dameon from "../../daemon/index.js";

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

  rpc.setup(config);
  dameon.setup(config);
  start();
};
