import { logger } from "../../logger/index.js";
import { start } from "../../swarm/index.js";
import { ask } from "../../ask.js";
import { parse } from "yaml";
import { readFileSync } from "fs";
import { untilde } from "../../utils/untilde.js";
import { checkForUpdates } from "../../update.js";

import * as nodes from "../../repo/nodes.js";
import * as rpc from "../../rpc/index.js";

export const startAction = async (configFile, options) => {
  const config = { ...parse(readFileSync(configFile).toString()), ...options };

  if (!config) {
    logger.error("Invalid config file");
    return process.exit(1);
  }

  logger.level = config.log || "info";

  if (options.ask) {
    try {
      config.privateKey = await ask();
    } catch (error) {}
  }

  if (!config.privateKey) {
    logger.error("No private key supplied");
    return process.exit(1);
  }

  await checkForUpdates();
  await nodes.setup(untilde(config.store));
  rpc.setup(config);
  start();
};
