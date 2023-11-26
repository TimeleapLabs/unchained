import HyperSwarm from "hyperswarm";
import { makeSpinner } from "../spinner.js";
import { topic } from "../constants.js";
import { logger } from "../logger/index.js";
import { rpc } from "../rpc/index.js";
import { sockets } from "../constants.js";

const swarm = new HyperSwarm();
const spinner = makeSpinner("Looking for peers");

swarm.on("connection", (socket, info) => {
  if (spinner.isEnabled) {
    spinner.succeed("Found peers");
    spinner.isEnabled = false;
  }

  const peerAddr = info.publicKey.toString("hex");
  const peer = `[${peerAddr.slice(0, 4)}...${peerAddr.slice(-4)}]`;

  sockets.set(peerAddr, socket);
  logger.info(`Connected to a new peer: ${peerAddr}`);

  socket.on("data", async (data) => {
    const message = JSON.parse(data.toString());
    logger.debug(`Received a packet from: ${peerAddr}`);
    logger.silly(message);

    if (message.error) {
      logger.error(`Received an error from peer ${peer}: ${message.error}`);
    } else if (message.result) {
      //logger.info(`Received result from peer ${peer}: ${message.result}`);
    } else if (message.method) {
      const result = await rpc(message);
      try {
        await socket.write(JSON.stringify(result));
      } catch (error) {
        logger.error(`Socket error with peer ${peer}: ${error.code}`);
      }
      // TODO: debug code
      if (message.method === "attest" && result) {
        const { signers, data } = message.args[1];
        logger.info(
          `${signers.length}x attestation at block ${data.block}: $${data.price}`
        );
      }
    }
  });

  socket.on("error", (error) => {
    logger.error(`Socket error with peer ${peer}: ${error.code}`);
  });

  socket.on("timeout", () => {
    logger.error(`Socket error with peer ${peer}: E_TIMEOUT`);
  });

  socket.on("close", () => {
    sockets.delete(peerAddr);
  });
});

export const start = () => {
  logger.info(`Starting the node: ${swarm.keyPair.publicKey.toString("hex")}`);
  spinner.start();
  swarm.join(topic);
};
