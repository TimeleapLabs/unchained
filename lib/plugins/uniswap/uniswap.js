import { ethers } from "ethers";
import { attest as blsAttest } from "../../bls/index.js";

export const cache = new Map();
let provider;

const getProvider = (config) => {
  if (!provider) {
    provider = new ethers.JsonRpcProvider(config.rpc.ethereum);
  }
  return provider;
};

const setCache = (block, price) => {
  cache.set(block, price);
  for (const key of cache.keys()) {
    if (key < block - 12) {
      cache.delete(key);
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
export const work = async (config, poolAddress, decimals, inverse) => {
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
  setCache(block, price);
  return blsAttest({ block, price }, null, []);
};

export const attest = async (data, signature, signers) => {
  if (!cache.has(data.block)) {
    return null;
  }
  const price = cache.get(data.block);
  if (price !== data.price) {
    return false;
  }
  return blsAttest(data, signature, signers);
};
