import HyperSwarm from "hyperswarm";
import { topic } from "../constants.js";
import { logger } from "../logger/index.js";
import { rpc } from "../rpc/index.js";
import { makeSpinner } from "../spinner.js";

const swarm = new HyperSwarm();
const spinner = makeSpinner("Looking for peers");

const dummyCommand = {
  method: "uniSpot",
  args: ["0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640", [18, 6], true],
};

swarm.on("connection", (socket, info) => {
  if (spinner.isEnabled) {
    spinner.succeed("Found peers");
  }

  logger.info(`Connected to a new peer: ${info.publicKey.toString("hex")}`);

  socket.on("data", async (data) => {
    const message = JSON.parse(data.toString());

    logger.debug(`Received a packet from: ${info.publicKey.toString("hex")}`);

    const peer = `[${info.publicKey.toString("hex").slice(-4)}]`;

    if (message.error) {
      logger.error(`Received an error from peer ${peer}: ${message.error}`);
    } else if (message.result) {
      logger.info(`Recevied result from peer ${peer}: ${message.result}`);
    } else if (message.method) {
      const result = await rpc(message);
      try {
        socket.write(JSON.stringify(result));
      } catch (error) {
        logger.error(`Network error: ${error.code}`);
      }
    }
  });

  const interval = setInterval(() => {
    if (!socket.destroyed) {
      logger.debug('Calling the "uniSpot" method on peer.');
      try {
        socket.write(JSON.stringify(dummyCommand));
      } catch (error) {
        logger.error(`Network error: ${error.code}`);
      }
    } else {
      clearInterval(interval);
    }
  }, 5000 + Math.random() * 5000);
});

export const start = () => {
  logger.info(`Starting the node: ${swarm.keyPair.publicKey.toString("hex")}`);
  spinner.start();
  swarm.join(topic);
};
