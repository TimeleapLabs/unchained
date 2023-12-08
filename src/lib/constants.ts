import { sha } from "./utils/hash.js";
import { State, Keys, Config, StringAnyObject } from "./types.js";

export const version = "0.6.0";
export const protocolVersion = "0.5.0";

export const topic = sha(`Kenshi.Unchained.Testnet.Topic.V${protocolVersion}`);

export const sockets = new Map();
export const keys: Keys = {};

export const config: Config = {
  name: "Change Me",
  log: "info",
  store: "~/.unchained",
  rpc: {
    ethereum: "https://ethereum.publicnode.com",
  },
  database: {
    url: "",
    name: "unchained",
  },
  secretKey: "",
};

export const rpcMethods: StringAnyObject = {};
export const gossipMethods: StringAnyObject = {};

export const errors = {
  E_NOT_FOUND: 404,
  E_INTERNAL: 500,
};

export const state: State = {
  connected: false,
};

export const nameRegex = /^[a-z0-9 @\._'-]+$/i;
