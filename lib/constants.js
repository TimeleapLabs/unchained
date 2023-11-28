import crypto from "crypto";

export const version = "0.3.3";
export const protocolVersion = "0.3.0";

export const topic = crypto
  .createHash("sha256")
  .update(`Kenshi.Unchained.Testnet.Topic.V${protocolVersion}`)
  .digest();

export const sockets = new Map();
export const keys = {};
export const config = {};

export const rpcMethods = {};
export const gossipMethods = {};

export const errors = {
  E_NOT_FOUND: 404,
  E_INTERNAL: 500,
};

export const state = {};
