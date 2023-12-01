import { ethers } from "ethers";
import { attest as blsAttest, verify } from "../../bls/index.js";
import { encoder } from "../../bls/keys.js";
import { keys, gossipMethods, config } from "../../constants.js";
import { logger } from "../../logger/index.js";
import { state } from "../../constants.js";
import { WS } from "iso-websocket";

import WebSocket from "ws";

const cache = new Map();
const attestations = new Map();
const earlyAttestations = new Map();

let provider;

const CACHE_SIZE = 50;

const ws = (endpoint) =>
  new WS(endpoint, { ws: WebSocket, retry: { forever: true } });

const getProvider = (config) => {
  if (!provider) {
    const endpoint = config.rpc.ethereum;
    provider = endpoint.startsWith("wss://")
      ? new ethers.WebSocketProvider(ws(endpoint))
      : new ethers.JsonRpcProvider(endpoint);
  }
  return provider;
};

const setCache = (block, price) => {
  cache.set(block, price);
  for (const key of cache.keys()) {
    if (key < block - CACHE_SIZE) {
      cache.delete(key);
    }
  }
};

const setEarlyAttestations = (block, signer, price) => {
  const current = earlyAttestations.get(block) || [];
  if (!current.includes(signer)) {
    earlyAttestations.set(`${block}-${price}`, [...current, signer]);
  }
  for (const key of earlyAttestations.keys()) {
    if (Number(key.split("-").shift()) < block - CACHE_SIZE) {
      attestations.delete(key);
    }
  }
};

const setAttestations = (block, signer, price) => {
  const stored = attestations.get(block) || [];
  const early = earlyAttestations.get(`${block}-${price}`) || [];
  const current = [...new Set([...stored, ...early])];
  if (!current.includes(signer)) {
    attestations.set(block, [...current, signer]);
    const total = current.length + 1;
    if (total > 1) {
      logger.info(`${total}x validations at block ${block}: $${price}`);
    }
  }
  for (const key of attestations.keys()) {
    if (key < block - CACHE_SIZE) {
      attestations.delete(key);
    }
  }
};

const poolABI = [
  `function slot0() external view returns
      (uint160 sqrtPriceX96,
      int24 tick,
      uint16 observationIndex,
      uint16 observationCardinality,
      uint16 observationCardinalityNext,
      uint8 feeProtocol,
      bool unlocked)`,
];

// TODO: our cache works only with one asset
export const work = async (parameters, poolAddress, decimals, inverse) => {
  try {
    const start = new Date();
    const provider = getProvider(config);
    const pool = new ethers.Contract(poolAddress, poolABI, provider);
    const block = await provider.getBlockNumber();
    if (cache.has(block)) {
      return null;
    }
    const { sqrtPriceX96 } = await pool.slot0();
    const delta = BigInt(decimals[0] - decimals[1]);
    const raw = (10n ** 18n * sqrtPriceX96 ** 2n) / (10n ** delta * 2n ** 192n);
    const price = inverse ? 1e18 / Number(raw) : Number(raw) / 1e18;
    const end = new Date();
    const took = end.valueOf() - start.valueOf();
    if (took > 2000) {
      logger.warn(
        `Detected high latency with the Ethereum RPC node: ${took}ms`
      );
    } else if (state.connected) {
      logger.debug(`Request to Ethereum RPC node took ${took}ms`);
    }
    if (cache.has(block)) {
      return null;
    }
    setCache(block, price);
    setAttestations(block, [encoder.encode(keys.publicKey.toBytes())], price);
    return blsAttest({
      method: "uniswapAttest",
      data: { block, price },
      parameters,
    });
  } catch (error) {
    logger.warn("Could not get the Ethereum price. Check your RPC.");
    throw error;
  }
};

export const verifyGossip = (incoming, connection) => {
  // TODO: the request packet needs to be reconstructed, hashed, and compared
  // TODO: to make sure the packet has the correct format

  if (verify(incoming)) {
    logger.verbose(`Successfully verified packet sent by ${connection.name}`);
    return true;
  } else {
    logger.warn(`Couldn't verify packet sent by ${connection.name}`);
    return false;
  }
};

export const attest = async (incoming, connection) => {
  if (!verifyGossip(incoming, connection)) {
    return false;
  }

  const { data } = incoming.request;

  if (!cache.has(data.block)) {
    setEarlyAttestations(data.block, incoming.signer, data.price);
    return null;
  }
  const price = cache.get(data.block);
  if (price !== data.price) {
    return false;
  }
  setAttestations(data.block, incoming.signer, price);
  return true;
};

Object.assign(gossipMethods, { uniswapAttest: attest });
