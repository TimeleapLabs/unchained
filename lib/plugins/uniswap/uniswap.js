import { ethers } from "ethers";

let provider;

const getProvider = (config) => {
  if (!provider) {
    provider = new ethers.JsonRpcProvider(config.rpc.ethereum);
  }
  return provider;
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

export const getPrice = async (config, poolAddress, decimals, inverse) => {
  const provider = getProvider(config);
  const pool = new ethers.Contract(poolAddress, poolABI, provider);
  const { sqrtPriceX96 } = await pool.slot0();
  const delta = BigInt(decimals[0] - decimals[1]);
  const price = (10n ** 18n * sqrtPriceX96 ** 2n) / (10n ** delta * 2n ** 192n);
  return inverse ? 1e18 / Number(price) : Number(price) / 1e18;
};
