import { logger } from "../../logger/index.js";
import { stringify } from "yaml";
import { writeFileSync } from "fs";
import { makeKeys, encodeKeys, loadKeys } from "../../crypto/bls/keys.js";
import { encoder } from "../../crypto/base58/index.js";
import { keys } from "../../constants.js";
import { safeReadConfig } from "../../utils/config.js";
import { murmur } from "../../constants.js";
import { hashUint8Array } from "../../utils/uint8array.js";
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

  const bytes = keys.publicKey.toBytes();
  const address = encoder.encode(bytes);
  murmur.address = await hashUint8Array(bytes);

  if (options.ci) {
    console.log(address);
    console.log(murmur.address);
  } else {
    logger.info(`Unchained public address is ${address}`);
    logger.info(`Unchained wave address is ${murmur.address}`);
  }

  return process.exit(0);
};
