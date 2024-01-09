import { logger } from "../../logger/index.js";
import { parse, stringify } from "yaml";
import { readFileSync, writeFileSync } from "fs";
import { makeKeys, encodeKeys, loadKeys } from "../../crypto/bls/keys.js";
import { encoder } from "../../crypto/base58/index.js";
import { keys } from "../../constants.js";
import { Config } from "../../types.js";
import assert from "node:assert";

interface AddressOptions {
  generate?: boolean;
  ci?: boolean;
}

const safeReadConfig = (configFile: string): string | null => {
  try {
    const configContent = readFileSync(configFile).toString();
    return configContent;
  } catch (error) {
    return null;
  }
};

export const addressAction = async (
  configFile: string,
  options: AddressOptions
) => {
  const configContent = safeReadConfig(configFile);
  const config: Config = configContent ? { ...parse(configContent) } : null;

  if (!config) {
    logger.error("Invalid config file");
    return process.exit(1);
  }

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

  Object.assign(keys, loadKeys(config.secretKey));
  assert(keys.publicKey !== undefined, "No public key available"); // Likely always passes

  const address = encoder.encode(keys.publicKey.toBytes());

  if (options.ci) {
    console.log(address);
  } else {
    logger.info(`Unchained public address is ${address}`);
  }

  return process.exit(0);
};
