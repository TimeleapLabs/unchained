import { sha } from "./utils/hash.js";
export const version = "0.8.10";
export const protocolVersion = "0.8.0";
export const topic = sha(`Kenshi.Unchained.Testnet.Topic.V${protocolVersion}`);
export const sockets = new Map();
export const keys = Object({});
export const config = {
    name: "Change Me",
    log: "info",
    rpc: {
        ethereum: "https://ethereum.publicnode.com",
    },
    database: {
        url: "",
        name: "unchained",
    },
    secretKey: "",
    publicKey: "",
    lite: false,
    gossip: 24,
    peers: {
        max: 512,
        parallel: 16,
    },
};
export const rpcMethods = {};
export const gossipMethods = {};
export const errors = {
    E_NOT_FOUND: 404,
    E_INTERNAL: 500,
};
export const state = {
    connected: false,
};
export const nameRegex = /^[a-z0-9 @\._'-]+$/i;
