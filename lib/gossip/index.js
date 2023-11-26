import { sockets } from "../constants.js";
import crypto from "crypto";

const randomIndex = (length) => {
  if (length <= 0) {
    throw new Error("Array length must be greater than 0.");
  }

  let index, randomByte;

  do {
    randomByte = crypto.randomBytes(1)[0];
    index = randomByte % length;
  } while (randomByte - index >= 256 - (256 % length));

  return index;
};

export const gossip = async (data) => {
  if (sockets.size === 0) {
    return;
  }
  const promises = [];
  const random = new Array(3).fill(sockets.size).map(randomIndex);
  const chosen = [...new Set(random)];
  const values = [...sockets.values()];
  for (const index of chosen) {
    const socket = values[index];
    promises.push(socket.write(data));
  }
  await Promise.all(promises).catch(() => null);
};
