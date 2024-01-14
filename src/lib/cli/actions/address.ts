import { logger } from "../../logger/index.js";
import { stringify } from "yaml";
import { writeFileSync } from "fs";
import { makeKeys, encodeKeys, loadKeys } from "../../crypto/bls/keys.js";
import { encoder } from "../../crypto/base58/index.js";
import { keys } from "../../constants.js";
import { safeReadConfig } from "../../utils/config.js";
import { murmur } from "../../constants.js";
import { toMurmur } from "../../crypto/murmur/index.js";
import assert from "node:assert";

interface AddressOptions {
  generate?: boolean;
  ci?: boolean;
}

export const addressAction = async (
  configFile: string,
  options: AddressOptions
) => {
  const config = safeReadConfig(configFile);
  if (!config) {
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
  murmur.address = await toMurmur(address);

  if (options.ci) {
    console.log(address);
    console.log(murmur.address);
  } else {
    logger.info(`Unchained public address is ${address}`);
    logger.info(`Unchained gossip address is ${murmur.address}`);
  }

  return process.exit(0);
};
