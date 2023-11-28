import { ethers } from "ethers";
import { attest as blsAttest } from "../../bls/index.js";
import { encoder } from "../../bls/keys.js";
import { keys, gossipMethods, config } from "../../constants.js";
import { logger } from "../../logger/index.js";

const cache = new Map();
const attestations = new Map();
const earlyAttestations = new Map();

let provider;

const CACHE_SIZE = 50;

const getProvider = (config) => {
  if (!provider) {
    provider = new ethers.JsonRpcProvider(config.rpc.ethereum);
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

const setEarlyAttestations = (block, signers, price) => {
  const current = earlyAttestations.get(block) || [];
  const newSigners = [];
  for (const signer of signers) {
    if (!current.includes(signer)) {
      newSigners.push(signer);
    }
  }
  if (newSigners.length) {
    // TODO: this can be improved
    earlyAttestations.set(`${block}-${price}`, [...current, ...newSigners]);
  }
  for (const key of earlyAttestations.keys()) {
    if (Number(key.split("-").shift()) < block - CACHE_SIZE) {
      attestations.delete(key);
    }
  }
};

const setAttestations = (block, signers, price) => {
  const stored = attestations.get(block) || [];
  const early = earlyAttestations.get(`${block}-${price}`) || [];
  const current = [...new Set([...stored, ...early])];
  const newSigners = [];
  for (const signer of signers) {
    if (!current.includes(signer)) {
      newSigners.push(signer);
    }
  }
  if (newSigners.length) {
    attestations.set(block, [...current, ...newSigners]);
    const total = current.length + newSigners.length;
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
  if (cache.has(block)) {
    return null;
  }
  setCache(block, price);
  setAttestations(block, [encoder.encode(keys.publicKey.toBytes())], price);
  return blsAttest(
    { method: "uniswapAttest", data: { block, price }, parameters },
    null,
    []
  );
};

export const attest = async ({ request, signers }) => {
  const { data } = request;
  if (!cache.has(data.block)) {
    setEarlyAttestations(data.block, signers, data.price);
    return null;
  }
  const price = cache.get(data.block);
  if (price !== data.price) {
    return false;
  }
  setAttestations(data.block, signers, price);
  return true;
};

Object.assign(gossipMethods, { uniswapAttest: attest });
