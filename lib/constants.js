import crypto from "crypto";

export const topic = crypto
  .createHash("sha256")
  .update("Kenshi.Unchained.Testnet.Topic")
  .digest();

export const version = "0.1.3";

export const sockets = new Map();
