import { sha } from "./utils/hash.js";
import {
  State,
  KeyPair,
  Config,
  StringAnyObject,
  StringGossipMethodObject,
  MetaData,
  Murmur,
} from "./types.js";

export const version = "0.10.0";
export const protocolVersion = "0.10.0";

export const topic = sha(`Kenshi.Unchained.Testnet.Topic.V${protocolVersion}`);

export const sockets = new Map<string, MetaData>();
export const keys: KeyPair = Object({});

export const config: Config = {
  name: "Change Me",
  log: "info",
  rpc: {
    ethereum: "https://ethereum.publicnode.com",
  },
  database: {
    url: "",
  },
  secretKey: "",
  publicKey: "",
  lite: false,
  gossip: 48,
  peers: {
    max: 128,
    parallel: 16,
  },
};

export const rpcMethods: StringAnyObject = {};
export const gossipMethods: StringGossipMethodObject<any, any> = {};

export const errors = {
  E_NOT_FOUND: 404,
  E_INTERNAL: 500,
};

export const state: State = {
  connected: false,
};

export const nameRegex = /^[a-z0-9 @\._'-]+$/i;
export const murmur: Murmur = { address: "" };
