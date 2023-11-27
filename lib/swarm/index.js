import HyperSwarm from "hyperswarm";
import { makeSpinner } from "../spinner.js";
import { topic } from "../constants.js";
import { logger } from "../logger/index.js";
import { rpc } from "../rpc/index.js";
import { sockets } from "../constants.js";

const swarm = new HyperSwarm();
const spinner = makeSpinner("Looking for peers");

swarm.on("connection", async (socket, info) => {
  if (spinner.isEnabled) {
    spinner.succeed("Found peers");
    spinner.isEnabled = false;
  }

  const peerAddr = info.publicKey.toString("hex");
  const peer = `[${peerAddr.slice(0, 4)}...${peerAddr.slice(-4)}]`;

  const meta = { peer, peerAddr, name: peer };

  sockets.set(peerAddr, { socket });
  logger.info(`Connected to a new peer: ${peerAddr}`);

  socket.on("data", async (data) => {
    const message = JSON.parse(data.toString());
    logger.debug(`Received a packet from: ${peerAddr}`);
    logger.silly(message);

    if (message.error) {
      logger.error(
        `Received an error from peer ${meta.name}: ${message.error}`
      );
    } else if (message.result) {
      //logger.info(`Received result from peer ${meta.name}: ${message.result}`);
      if (message.result.name && typeof message.result.name === "string") {
        if (!message.result.name.match(/[a-z0-9 @._-]+/i)) {
          return;
        }
        const mapdata = sockets.get(peerAddr);
        mapdata.name = message.result.name.slice(0, 24);
        // TODO: verify the validity of the public key
        mapdata.publicKey = message.result.publicKey;
        sockets.set(peerAddr, mapdata);
        logger.info(`Peer ${meta.name} is ${mapdata.name}`);
        meta.name = mapdata.name;
      }
    } else if (message.method) {
      const result = await rpc(message);
      try {
        await socket.write(JSON.stringify(result));
      } catch (error) {
        logger.error(`Socket error with peer ${meta.name}: ${error.code}`);
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
    logger.error(`Socket error with peer ${meta.name}: ${error.code}`);
  });

  socket.on("timeout", () => {
    logger.error(`Socket error with peer ${meta.name}: E_TIMEOUT`);
  });

  socket.on("close", () => {
    sockets.delete(peerAddr);
  });

  try {
    const introducePayload = JSON.stringify({ method: "introduce", args: [] });
    await socket.write(introducePayload);
  } catch (error) {}
});

export const start = () => {
  logger.info(`Starting the node: ${swarm.keyPair.publicKey.toString("hex")}`);
  spinner.start();
  swarm.join(topic);
};
