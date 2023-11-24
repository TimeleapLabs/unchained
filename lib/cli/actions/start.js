import { logger } from "../../logger/index.js";
import { start } from "../../swarm/index.js";
import { ask } from "../../ask.js";
import { parse } from "yaml";
import { readFileSync } from "fs";
import { setup } from "../../repo/nodes.js";
import { untilde } from "../../utils/untilde.js";

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

  await setup(untilde(config.store));
  start();
};
