import { logger } from "../../logger/index.js";
import { safeReadConfig } from "../../utils/config.js";
import { loadKeys } from "../../crypto/bls/keys.js";
import { encoder } from "../../crypto/base58/index.js";
import { toMurmur } from "../../crypto/murmur/index.js";
import { stringify } from "yaml";
import { version } from "../../constants.js";

import getos from "getos";
import ping from "ping";
import os from "os";
import osName from "os-name";
import prettyBytes from "pretty-bytes";

const getOsName = async () =>
  new Promise((resolve) => {
    if (os.platform() === "linux") {
      getos((err, result) => {
        if (err) {
          resolve(osName());
        } else {
          const { dist, codename, release } = result as getos.LinuxOs;
          resolve([dist, codename, release].filter(Boolean).join(" "));
        }
      });
    } else {
      resolve(osName());
    }
  });

const systemInfo = async () => ({
  os: {
    platform: os.platform(),
    name: await getOsName(),
    release: os.release(),
  },
  cpu: {
    arch: os.arch(),
    cores: os.cpus().length,
    parallelism: os.availableParallelism(),
  },
  memory: {
    total: prettyBytes(os.totalmem()),
    free: prettyBytes(os.freemem()),
  },
  node: { version: process.version },
});

const networkQuality = async (host: string) => {
  const cleanHost = host.replace(/^(ws|http)s?:\/\//, "").replace(/\/.*$/, "");
  try {
    const res = await ping.promise.probe(cleanHost);
    return {
      ping: {
        time: res.time,
        alive: res.alive,
        loss:
          res.packetLoss === "unknown"
            ? res.packetLoss
            : parseFloat(res.packetLoss),
      },
    };
  } catch (error: any) {
    return { ping: { time: "unknown", alive: "unknown", loss: "unknown" } };
  }
};

export const diagnoseAction = async (configFile: string) => {
  const config = safeReadConfig(configFile);
  if (!config) {
    return process.exit(1);
  }

  if (!config.secretKey) {
    logger.error("No secret key supplied");
    return process.exit(1);
  }

  const keys = loadKeys(config.secretKey);
  const publicKey = encoder.encode(keys.publicKey.toBytes());
  const murmur = await toMurmur(publicKey);

  const host = Array.isArray(config.rpc.ethereum)
    ? config.rpc.ethereum[0]
    : config.rpc.ethereum;

  const pingResult = await networkQuality(host);
  const systemInfoResults = await systemInfo();

  const info = {
    ...systemInfoResults,
    ...pingResult,
    unchained: { version, publicKey, murmur },
  };

  const stringified = stringify(info);
  console.log(stringified);

  return process.exit(0);
};
