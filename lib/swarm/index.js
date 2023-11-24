import HyperSwarm from "hyperswarm";
import { topic } from "../constants.js";
import { logger } from "../logger/index.js";
import { rpc } from "../rpc/index.js";

const swarm = new HyperSwarm();

swarm.on("connection", (socket, info) => {
  logger.info(`Connected to a new peer: ${info.publicKey.toString("hex")}`);

  socket.on("data", async (data) => {
    const message = JSON.parse(data.toString());

    logger.debug(`Received a packet from: ${info.publicKey.toString("hex")}`);

    if (message.error) {
      logger.error(`Received error: ${message.error}`);
    } else if (message.result) {
      logger.info(`Recevied result: ${message.result}`);
    } else if (message.method) {
      const result = await rpc(message);
      socket.write(JSON.stringify(result));
    }
  });

  const interval = setInterval(() => {
    if (!socket.destroyed) {
      logger.debug('Calling the "timestamp" method.');
      socket.write(JSON.stringify({ method: "timestamp", args: [] }));
    } else {
      clearInterval(interval);
    }
  }, 1000 + Math.random() * 5000);
});

export const start = () => {
  logger.info(`Starting the swarm: ${swarm.keyPair.publicKey.toString("hex")}`);
  swarm.join(topic);
};
