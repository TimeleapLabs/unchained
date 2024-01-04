import { makeSpinner } from "../spinner.js";
import { topic, state, nameRegex } from "../constants.js";
import { logger } from "../logger/index.js";
import { processRpc } from "../rpc/index.js";
import { processGossip } from "../gossip/index.js";
import { sockets } from "../constants.js";
import { parse } from "../utils/json.js";
import { Duplex } from "stream";
import { MetaData, NodeSystemError, PeerInfo } from "../types.js";
import { config } from "../constants.js";
import { isJailed, strike } from "./jail.js";

import HyperSwarm from "hyperswarm";

let swarm: HyperSwarm;
const spinner = makeSpinner("Looking for peers");

const setupEventListeners = () => {
  swarm.on("connection", async (socket: Duplex, info: PeerInfo) => {
    if (spinner.isEnabled) {
      spinner.succeed("Found peers");
      spinner.isEnabled = false;
      state.connected = true;
    }

    if (sockets.size >= config.peers.max) {
      return socket.end();
    }

    const peerAddr = info.publicKey.toString("hex");
    const peer = `[${peerAddr.slice(0, 4)}···${peerAddr.slice(-4)}]`;
    const meta: MetaData = {
      socket,
      peer,
      peerAddr,
      name: peer,
      isSocketBusy: false,
    };

    if (isJailed(meta.name, info)) {
      return socket.end();
    }

    let timeout: NodeJS.Timeout;

    socket.on("error", (error: NodeSystemError) => {
      const code = error.code || error.errno || error.message;
      logger.debug(`Socket error with peer ${meta.name}: ${code}`);
      strike(meta.name, info);
    });

    socket.on("timeout", () => {
      logger.debug(`Socket error with peer ${meta.name}: ETIMEDOUT`);
      strike(meta.name, info);
    });

    socket.on("close", () => {
      clearTimeout(timeout);
      sockets.delete(peerAddr);
    });

    socket.on("drain", () => {
      meta.isSocketBusy = false;
    });

    sockets.set(peerAddr, meta);
    logger.info(`Connected to a new peer: ${peerAddr}`);

    const warnNoData = () => {
      timeout = setTimeout(() => {
        logger.warn(`No data from ${meta.name} in the last 60 seconds`);
        strike(meta.name, info);
        warnNoData();
      }, 60000);
    };

    warnNoData();

    socket.on("data", async (data) => {
      clearTimeout(timeout);
      warnNoData();

      const message = parse(data.toString());

      if (message instanceof Error) {
        return logger.debug(`Received a faulty packet from: ${peerAddr}`);
      }

      logger.debug(`Received a packet from: ${peerAddr}`);
      logger.silly(message);

      if (message.error) {
        logger.error(
          `Received an error from peer ${meta.name}: ${message.error}`
        );
      } else if (message.result) {
        // TODO: this needs to be handled properly
        if (message.result.name && typeof message.result.name === "string") {
          if (!message.result.name.match(nameRegex)) {
            return logger.warn(`Received an illegal name from ${meta.name}`);
          }
          const oldName = meta.name;
          meta.name = message.result.name.slice(0, 24);
          // TODO: verify the validity of the public key
          meta.publicKey = message.result.publicKey;
          logger.info(`Peer ${oldName} is ${meta.name}`);
        }
      } else if (message.type === "call") {
        const result = await processRpc(message);
        try {
          socket.write(JSON.stringify(result));
        } catch (error) {
          const err = error as NodeSystemError;
          const info = err.code || err.errno || err.message;
          logger.error(`Socket error with peer ${meta.name}: ${info}`);
        }
      } else if (message.type === "gossip") {
        await processGossip(message);
      }
    });

    try {
      const introducePayload = JSON.stringify({
        type: "call",
        request: { method: "introduce", args: {} },
      });
      socket.write(introducePayload);
    } catch (error) {}
  });
};

export const discover = (): void => {
  if (state.connected) {
    logger.debug("Running the peer discovery mechanism");
  }
  const discovery = swarm.join(topic);
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
