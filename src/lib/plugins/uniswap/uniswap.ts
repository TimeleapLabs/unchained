import { ethers } from "ethers";
import { gossipMethods, config, keys, sockets } from "../../constants.js";
import { logger } from "../../logger/index.js";
import { state } from "../../constants.js";
import { WS } from "iso-websocket";
import { attest as blsAttest, aggregate, verify } from "../../bls/index.js";
import { encoder } from "../../bls/keys.js";
import { WebSocketLike } from "ethers";
import { WebSocket } from "unws";
import { assetPrices } from "../../db/collections/AssetPrice.js";

import assert from "assert";

import {
  Config,
  GossipMethod,
  AssetPriceMetric,
  GossipRequest,
  SignatureItem,
} from "../../types.js";

import { Attestation, PriceSignatureInput } from "./types.js";

const cache = new Map<number, number>();
const attestations = new Map<number, Attestation>();
const earlyAttestations = new Map<number, SignatureItem[]>();

let provider: ethers.Provider | undefined;

const CACHE_SIZE = 50;

const ws = (endpoint: string): WebSocketLike =>
  new WS(endpoint, {
    ws: WebSocket,
    retry: { forever: true },
  }) as WebSocketLike;

const getProvider = (config: Config) => {
  if (!provider) {
    const endpoint = config.rpc.ethereum;
    provider = endpoint.startsWith("wss://")
      ? new ethers.WebSocketProvider(ws(endpoint))
      : new ethers.JsonRpcProvider(endpoint);
  }
  return provider;
};

const setCache = (block: number, price: number) => {
  cache.set(block, price);
  for (const key of cache.keys()) {
    if (key < block - CACHE_SIZE) {
      cache.delete(key);
    }
  }
};

const setEarlyAttestations = (
  block: number,
  signer: string,
  signature: string
) => {
  const attestations = earlyAttestations.get(block) || [];
  const alreadyAdded = attestations.some((item) => item.signer === signer);
  if (!alreadyAdded) {
    earlyAttestations.set(block, [...attestations, { signer, signature }]);
  }
  for (const [key] of earlyAttestations.entries()) {
    if (key < block - CACHE_SIZE) {
      // FIXME: Security problem where a validator can reset another
      // FIXME: peer's cache by sending a big block number
      earlyAttestations.delete(key);
    }
  }
};

const updateAssetPrice = async (
  block: number,
  price: number,
  aggregated: string,
  signers: string[]
) => {
  await assetPrices.updateOne(
    {
      block,
      asset: "ethereum",
      source: "uniswap-ethereum",
    },
    {
      $set: {
        price,
        signature: aggregated,
        signers,
      },
      $setOnInsert: {
        timestamp: new Date(), // FIXME
        asset: "ethereum",
        source: "uniswap-ethereum",
        block,
      },
    },
    { upsert: true }
  );
};

const printAttestations = (
  size: number,
  block: number,
  price: number,
  signersSet: Set<string>
) => {
  logger.info(`${size}x validations at block ${block}: $${price}`);
  const allPeers = [
    keys.publicKey
      ? {
          name: "@",
          publicKey: encoder.encode(Buffer.from(keys.publicKey.toBytes())),
        }
      : null,
    ...sockets.values(),
  ].filter(Boolean);
  const peerStates = {
    signed: allPeers
      .filter((peer) => signersSet.has(peer.publicKey))
      .map((peer) => peer.name),
    missing: allPeers
      .filter((peer) => !signersSet.has(peer.publicKey))
      .map((peer) => peer.name),
  };
  logger.verbose(`Received signatures: ${peerStates.signed.join(", ")}`);
  logger.verbose(
    `Missing signatures: ${peerStates.missing.join(", ") || "N/A"}`
  );
};

const setAttestations = async (
  block: number,
  signer: string,
  signature: string
) => {
  const price = cache.get(block);
  assert(typeof price === "number", "Attempting to attest an unchaed block");
  const data: PriceSignatureInput = { metric: { block }, value: { price } };

  if (!verify({ signer, signature, data })) {
    return false;
  }

  const stored = attestations.get(block) || { signatures: [] };

  if (stored.signatures.some((item) => item.signer === signer)) {
    return false;
  }

  const early = earlyAttestations.get(block) || [];
  const earlySignatures = early.filter((item) => verify({ ...item, data }));
  const allSignatures = [
    ...stored.signatures,
    ...earlySignatures,
    { signer, signature },
  ];

  const uniqueSignatures = [];
  const signersSet = new Set<string>();

  for (const item of allSignatures) {
    if (!signersSet.has(item.signer)) {
      signersSet.add(item.signer);
      uniqueSignatures.push(item);
    }
  }

  const currentSignatures = stored.signatures.map((item) => item.signature);
  const newSignatureSet = [...early, { signer, signature }].filter(
    ({ signer }) => !currentSignatures.includes(signer)
  );

  const newSignatures = newSignatureSet.map((item) => item.signature);
  const currentAggregation = stored.aggregated || "";
  const signatureList = [currentAggregation, ...newSignatures].filter(Boolean);
  const aggregated = aggregate(signatureList);

  attestations.set(block, {
    ...stored,
    aggregated,
    signatures: [...newSignatureSet, ...stored.signatures],
  });

  if (!config.lite) {
    await updateAssetPrice(block, price, aggregated, [...signersSet]);
  }

  const { size } = signersSet;
  if (size > 1) {
    printAttestations(size, block, price, signersSet);
  }

  for (const [key] of attestations.entries()) {
    // FIXME: Security problem where a validator can reset another
    // FIXME: peer's cache by sending a big block number
    if (key < block - CACHE_SIZE) {
      attestations.delete(key);
    }
  }
  return true;
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

export const work = async (
  _parameters: any,
  poolAddress: string,
  decimals: [number, number],
  inverse: boolean
): Promise<GossipRequest<AssetPriceMetric> | null> => {
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
    if (took > 2000 && state.connected) {
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
    const data: PriceSignatureInput = { metric: { block }, value: { price } };
    const signed = blsAttest(data);
    await setAttestations(block, signed.signer, signed.signature);
    // TODO: we need to properly handle `dataset`
    return {
      method: "uniswapAttest",
      metric: { block },
      dataset: "ethereum::uniswap::ethereum",
      ...signed,
    };
  } catch (error) {
    logger.warn("Could not get the Ethereum price. Check your RPC.");
    throw error;
  }
};

export const attest: GossipMethod<AssetPriceMetric> = async (
  request: GossipRequest<AssetPriceMetric>
) => {
  const { metric, signer, signature } = request;

  if (!cache.has(metric.block)) {
    setEarlyAttestations(metric.block, signer, signature);
    return null;
  } else {
    const valid = await setAttestations(metric.block, signer, signature);
    if (valid) {
      return request;
    }
  }

  return null;
};

Object.assign(gossipMethods, { uniswapAttest: attest });
