import { makeSpinner } from "../spinner.js";
import { topic, state, nameRegex } from "../constants.js";
import { logger } from "../logger/index.js";
import { processRpc } from "../rpc/index.js";
import { processGossip } from "../gossip/index.js";
import { sockets } from "../constants.js";
import { parse } from "../utils/json.js";

import HyperSwarm from "hyperswarm";

const swarm = new HyperSwarm();
const spinner = makeSpinner("Looking for peers");

swarm.on("connection", async (socket, info) => {
  if (spinner.isEnabled) {
    spinner.succeed("Found peers");
    spinner.isEnabled = false;
    state.connected = true;
  }

  const peerAddr = info.publicKey.toString("hex");
  const peer = `[${peerAddr.slice(0, 4)}···${peerAddr.slice(-4)}]`;
  const meta = { socket, peer, peerAddr, name: peer };

  sockets.set(peerAddr, meta);
  logger.info(`Connected to a new peer: ${peerAddr}`);

  socket.on("data", async (data) => {
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
        await socket.write(JSON.stringify(result));
      } catch (error) {
        const info = error.code || error.errno || error.message;
        logger.error(`Socket error with peer ${meta.name}: ${info}`);
      }
    } else if (message.type === "gossip") {
      await processGossip(message, meta);
    }
  });

  socket.on("error", (error) => {
    const info = error.code || error.errno || error.message;
    logger.error(`Socket error with peer ${meta.name}: ${info}`);
  });

  socket.on("timeout", () => {
    logger.error(`Socket error with peer ${meta.name}: E_TIMEOUT`);
  });

  socket.on("close", () => {
    sockets.delete(peerAddr);
  });

  try {
    const introducePayload = JSON.stringify({
      type: "call",
      request: { method: "introduce", args: {} },
    });
    await socket.write(introducePayload);
  } catch (error) {}
});

export const discover = () => {
  if (state.connected) {
    logger.debug("Running the peer discovery mechanism");
  }
  const discovery = swarm.join(topic);
  discovery.flushed().then(() => {
    setTimeout(discover, 30000);
  });
};

export const start = () => {
  logger.info(`Starting the node: ${swarm.keyPair.publicKey.toString("hex")}`);
  spinner.start();
  discover();
};
