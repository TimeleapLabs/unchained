import { makeSpinner } from "../spinner.js";
import { state, nameRegex, sockets, config } from "../constants.js";
import { logger } from "../logger/index.js";
import { processRpc } from "../rpc/index.js";
import { serialize, parse } from "../utils/sia.js";
import { Duplex } from "stream";
import {
  IntroducePayload,
  MetaData,
  NodeSystemError,
  PeerInfo,
} from "../types.js";
import { isJailed, strike } from "./jail.js";
import { compress, uncompress } from "snappy";
import { copyUint8Array, isUint8Array } from "../utils/uint8array.js";
import { sha } from "../utils/hash.js";
import HyperSwarm from "hyperswarm";
import { minutes } from "../utils/time.js";

let swarm: HyperSwarm;
const spinner = makeSpinner("Looking for peers");

const introducePayload = await compress(
  serialize({ type: "call", request: { method: "introduce", args: {} } })
);

const safeClearTimeout = (timeout: NodeJS.Timeout | null) =>
  timeout && clearTimeout(timeout);

class safeClosedError extends Error {}

const safeCloseSocket = (socket: Duplex) => {
  try {
    if (!socket.closed) {
      socket.pause();
      while (socket.read());
      socket.destroy(new safeClosedError());
    }
  } catch (_err) {}
};

const safeDecompressAndParse = async (packet: Buffer) => {
  try {
    const uncompressed = await uncompress(packet, { asBuffer: true });
    return parse(uncompressed as Buffer) as any;
  } catch (error) {
    return error;
  }
};

const setupEventListeners = () => {
  swarm.on("connection", async (socket: Duplex, info: PeerInfo) => {
    if (spinner.isEnabled) {
      spinner.succeed("Found peers");
      spinner.isEnabled = false;
      state.connected = true;
    }

    const peerAddr = info.publicKey.toString("hex");
    const peer = `[${peerAddr.slice(0, 4)}···${peerAddr.slice(-4)}]`;
    const meta: MetaData = {
      socket,
      peer,
      peerAddr,
      name: peer,
      rpcRequests: new Set(),
    };

    let strikeOnClose = true;
    let timeout: NodeJS.Timeout | null = null;

    socket.on("error", (error: NodeSystemError) => {
      if (error instanceof safeClosedError) {
        return;
      }
      strike(meta);
      const code = error.code || error.errno || error.message;
      logger.debug(`Socket error with peer ${meta.name}: ${code}`);
    });

    socket.on("timeout", () => {
      logger.debug(`Socket error with peer ${meta.name}: ETIMEDOUT`);
      const jailed = strike(meta);
      if (jailed) {
        strikeOnClose = false;
        safeCloseSocket(socket);
      }
    });

    socket.on("close", () => {
      if (strikeOnClose) {
        strike(meta);
      }
      safeClearTimeout(timeout);
      sockets.delete(peerAddr);
    });

    sockets.set(peerAddr, meta);

    if (sockets.size > config.peers.max || isJailed(meta)) {
      sockets.delete(peerAddr);
      strikeOnClose = false;
      return safeCloseSocket(socket);
    }

    socket.on("drain", () => {
      meta.needsDrain = false;
    });

    logger.debug(`Connected to a new peer: ${peerAddr}`);

    const warnNoData = () => {
      timeout = setTimeout(() => {
        logger.warn(
          `No data from ${meta.name}@${meta.client?.version} in the last 60 seconds`
        );
        const jailed = strike(meta);
        if (jailed) {
          strikeOnClose = false;
          return safeCloseSocket(socket);
        }
        warnNoData();
      }, minutes(1));
    };

    warnNoData();

    socket.on("data", async (data) => {
      safeClearTimeout(timeout);
      warnNoData();

      const message = await safeDecompressAndParse(data);

      if (message instanceof Error) {
        return logger.debug(`Received a faulty packet from: ${peerAddr}`);
      }

      logger.debug(`Received a packet from: ${peerAddr}`);
      logger.silly(message);

      if (message.error) {
        // TODO: Give a score to each socket based on the number of
        // TODO: messages they can handle and take it into account
        if (message.error !== 429) {
          logger.error(
            `Received an error from peer ${meta.name}: ${message.error}`
          );
        }
      } else if (message.result) {
        // TODO: this needs to be handled properly
        if (message.result.name && typeof message.result.name === "string") {
          const result = message.result as IntroducePayload;
          if (!result.name.match(nameRegex)) {
            return logger.warn(`Received an illegal name from ${meta.name}`);
          }
          meta.name = result.name.slice(0, 24);
          // TODO: verify the validity of the public key
          meta.publicKey = isUint8Array(result.publicKey)
            ? copyUint8Array(result.publicKey)
            : undefined;
          meta.murmurAddr = result.murmurAddr;
          meta.client = result.client;
          logger.info(`Connected to peer ${meta.name}: ${meta.murmurAddr}`);
        }
      } else if (message.type === "call") {
        const result = await processRpc(message, meta);
        if (result.result || result.error) {
          try {
            const payload = await compress(serialize(result));
            socket.write(payload);
          } catch (error) {
            const err = error as NodeSystemError;
            const info = err.code || err.errno || err.message;
            logger.error(`Socket error with peer ${meta.name}: ${info}`);
          }
        }
      }
    });

    try {
      socket.write(introducePayload);
    } catch (error) {}
  });
};

export const discover = (): void => {
  if (state.connected) {
    logger.debug("Running the peer discovery mechanism");
  }
  const discovery = swarm.join(sha(config.network));
  discovery.flushed().then(() => {
    setTimeout(discover, 30000);
  });
};

export const startSwarm = (): void => {
  swarm = new HyperSwarm({
    maxParallel: config.peers.parallel,
    maxPeers: config.peers.max,
  });
  logger.verbose(
    `Noise public key is: ${swarm.keyPair.publicKey.toString("hex")}`
  );
  spinner.start();
  setupEventListeners();
  discover();
};
