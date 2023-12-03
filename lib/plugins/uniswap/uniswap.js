import { ethers } from "ethers";
import { attest as blsAttest } from "../../bls/index.js";
import { encoder } from "../../bls/keys.js";
import { keys, gossipMethods, config } from "../../constants.js";
import { logger } from "../../logger/index.js";
import { state } from "../../constants.js";
import { WS } from "iso-websocket";
import { hashObject } from "../../utils/hash.js";

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

const setEarlyAttestations = (request, signer) => {
  const hash = hashObject(request);
  const current = earlyAttestations.get(hash) || { signers: [] };
  if (!current.signers.includes(signer)) {
    earlyAttestations.set(hash, {
      signers: [...current.signers, signer],
      request,
    });
  }
  for (const [key, value] of earlyAttestations.entries()) {
    if (value.block < request.data.block - CACHE_SIZE) {
      // FIXME: Security problem where a validator can reset another
      // FIXME: peer's cache by sending a big block number
      attestations.delete(key);
    }
  }
};

const setAttestations = (request, signer) => {
  const hash = hashObject(request);
  const stored = attestations.get(hash) || { signers: [], request };
  const early = earlyAttestations.get(hash) || { signers: [] };
  const currentSigners = [...new Set([...stored.signers, ...early.signers])];
  if (!currentSigners.includes(signer)) {
    attestations.set(hash, { ...stored, signers: [...currentSigners, signer] });
    const total = currentSigners.length + 1;
    if (total > 1) {
      const { block, price } = request.data;
      logger.info(`${total}x validations at block ${block}: $${price}`);
    }
  }
  for (const [key, value] of attestations.entries()) {
    // FIXME: Security problem where a validator can reset another
    // FIXME: peer's cache by sending a big block number
    if (value.block < request.data.block - CACHE_SIZE) {
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
    const data = { block, price };
    const request = { method: "uniswapAttest", data, parameters };
    setAttestations(request, [encoder.encode(keys.publicKey.toBytes())]);
    return blsAttest(request);
  } catch (error) {
    logger.warn("Could not get the Ethereum price. Check your RPC.");
    throw error;
  }
};

export const attest = async (incoming) => {
  const { data } = incoming.request;

  if (!cache.has(data.block)) {
    setEarlyAttestations(incoming.request, incoming.signer);
    return null;
  }
  const price = cache.get(data.block);
  if (price !== data.price) {
    return false;
  }
  setAttestations(incoming.request, incoming.signer);
  return true;
};

Object.assign(gossipMethods, { uniswapAttest: attest });
