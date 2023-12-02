import { sha } from "./utils/hash.js";

export const version = "0.5.3";
export const protocolVersion = "0.5.0";

export const topic = sha(`Kenshi.Unchained.Testnet.Topic.V${protocolVersion}`);

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

export const nameRegex = /^[a-z0-9 @\._'-]+$/i;
