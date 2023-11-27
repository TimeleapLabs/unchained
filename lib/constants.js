import crypto from "crypto";

export const version = "0.2.5";
export const protocolVersion = "0.2.2";

export const topic = crypto
  .createHash("sha256")
  .update(`Kenshi.Unchained.Testnet.Topic.V${protocolVersion}`)
  .digest();

export const sockets = new Map();
export const keys = {};
export const config = {};
