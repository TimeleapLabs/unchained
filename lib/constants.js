import crypto from "crypto";

export const version = "0.2.0";

export const topic = crypto
  .createHash("sha256")
  .update(`Kenshi.Unchained.Testnet.Topic.V${version}`)
  .digest();

export const sockets = new Map();

export const keys = {};

// TODO: config should go here
