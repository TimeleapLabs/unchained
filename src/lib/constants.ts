import { minutes } from "./utils/time.js";
import {
  State,
  KeyPair,
  Config,
  StringAnyObject,
  MetaData,
  Murmur,
} from "./types.js";

export const version = "0.10.27";
export const protocolVersion = "0.10.23";

export const sockets = new Map<string, MetaData>();
export const keys: KeyPair = Object({});

export const config: Config = {
  name: "Change Me",
  log: "info",
  network: `Kenshi::Unchained::Testnet::V${protocolVersion}`,
  rpc: {
    ethereum: "https://ethereum.publicnode.com",
  },
  database: {
    url: "",
  },
  secretKey: "",
  publicKey: "",
  lite: false,
  peers: {
    max: 128,
    parallel: 16,
  },
  jail: {
    duration: minutes(5),
    strikes: 5,
  },
  waves: {
    count: 4,
    group: 16,
    select: 35,
    jitter: {
      min: 10,
      max: 25,
    },
  },
};

export const rpcMethods: StringAnyObject = {};

export const errors = {
  E_NOT_FOUND: 404,
  E_INTERNAL: 500,
  E_DUPLICATE: 409,
  E_TOO_MANY_REQUESTS: 429,
};

export const state: State = {
  connected: false,
};

export const nameRegex = /^[a-z0-9 @\._'-]+$/i;
export const murmur: Murmur = { address: "" };
