import { sha } from "./utils/hash.js";
import { minutes } from "./utils/time.js";
import {
  State,
  KeyPair,
  Config,
  StringAnyObject,
  MetaData,
  Murmur,
} from "./types.js";

export const version = "0.10.16";
export const protocolVersion = "0.10.15";

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
  peers: {
    max: 128,
    parallel: 16,
  },
  jail: {
    duration: minutes(5),
    strikes: 5,
  },
  waves: {
    count: 10,
    group: 8,
    select: 35,
    jitter: {
      min: 5,
      max: 15,
    },
  },
};

export const rpcMethods: StringAnyObject = {};

export const errors = {
  E_NOT_FOUND: 404,
  E_INTERNAL: 500,
  E_DUPLICATE: 409,
};

export const state: State = {
  connected: false,
};

export const nameRegex = /^[a-z0-9 @\._'-]+$/i;
export const murmur: Murmur = { address: "" };
