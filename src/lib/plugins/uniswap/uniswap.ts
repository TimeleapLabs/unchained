import { ethers } from "ethers";
import { gossipMethods, config, keys, sockets } from "../../constants.js";
import { logger } from "../../logger/index.js";
import { state } from "../../constants.js";
import { WS } from "iso-websocket";
import { attest as blsAttest, verify } from "../../crypto/bls/index.js";
import { aggregate, verifyAggregate } from "../../crypto/bls/threads/index.js";
import { encoder } from "../../crypto/base58/index.js";
import { WebSocketLike } from "ethers";
import { WebSocket } from "unws";
import { addOnePoint } from "../../score/index.js";
import { debounceAsync } from "../../utils/debounce.js";
import { db } from "../../db/db.js";
import { datasets } from "../../network/index.js";
import { cache } from "../../utils/cache.js";

import type { WantPacket, WantAnswer, Dataset } from "../../network/index.js";

import {
  Config,
  GossipMethod,
  AssetPriceMetric,
  GossipRequest,
  SignatureItem,
  AssetPriceValue,
} from "../../types.js";

import { Attestation, PriceSignatureInput } from "./types.js";
import { minutes } from "../../utils/time.js";
import { hashObject } from "../../utils/hash.js";
import { toMurmur } from "../../crypto/murmur/index.js";

const blockCache = new Map<number, number>();
const attestations = new Map<number, Attestation>();
const pendingAttestations = new Map<number, SignatureItem[]>();
const keyToIdCache = new Map<string, number>();
const wantCache = cache<string, any>(minutes(15));

let provider: ethers.Provider | undefined;
let getNewRpc: boolean = false;

const CACHE_SIZE = 50;

const ws = (endpoint: string): WebSocketLike => {
  const socket = new WS(endpoint, {
    ws: WebSocket,
    retry: { forever: true },
  }) as WebSocketLike;

  socket.onerror = () => {
    getNewRpc = true;
  };

  return socket;
};

let currentProvider: number = 0;

const getNextConnectionUrl = (config: Config): string => {
  if (typeof config.rpc.ethereum === "string") {
    return config.rpc.ethereum;
  } else {
    if (currentProvider > config.rpc.ethereum.length) {
      currentProvider = 0;
    }
    return config.rpc.ethereum[currentProvider++];
  }
};

const getProvider = (config: Config) => {
  if (getNewRpc || !provider) {
    if (getNewRpc) {
      provider?.destroy();
      getNewRpc = false;
    }
    const endpoint = getNextConnectionUrl(config);
    provider = endpoint.startsWith("wss://")
      ? new ethers.WebSocketProvider(ws(endpoint))
      : new ethers.JsonRpcProvider(endpoint);
  }
  return provider;
};

const setCache = (block: number, price: number) => {
  blockCache.set(block, price);
  for (const key of blockCache.keys()) {
    if (key < block - CACHE_SIZE) {
      blockCache.delete(key);
    }
  }
};

const addPendingAttestation = (
  block: number,
  signer: string,
  signature: string
) => {
  const pendingSignatures = pendingAttestations.get(block) || [];
  const confirmedSigners = attestations.get(block)?.signers;

  const alreadyAdded =
    pendingSignatures.some((item) => item.signer === signer) ||
    confirmedSigners?.some((cSigner) => cSigner === signer);

  if (!alreadyAdded) {
    pendingAttestations.set(block, [
      ...pendingSignatures,
      { signer, signature },
    ]);
    if (blockCache.has(block)) {
      processAttestations({ key: block, args: [block] }).catch((err: Error) => {
        logger.error(
          `Encountered an error while processing attestations: ${err.message}`
        );
      });
    }
  }
  for (const key of pendingAttestations.keys()) {
    if (key < block - CACHE_SIZE) {
      // FIXME: Security problem where a validator can reset another
      // FIXME: peer's cache by sending a big block number
      pendingAttestations.delete(key);
    }
  }
  return !alreadyAdded;
};

const addPendingAttestations = async (
  cache: any,
  block: number,
  signatures: { signer: string; signature: string }[]
) => {
  const pending = pendingAttestations.get(block) || [];
  const confirmed = attestations.get(block)?.signers;
  const murmurMap = new Map(
    [...sockets.values()].map((meta) => [meta.publicKey, meta.murmurAddr])
  );

  let newSigners = false;

  for (const { signature, signer } of signatures) {
    const alreadyAdded =
      pending.some((item) => item.signer === signer) ||
      confirmed?.some((cSigner) => cSigner === signer);

    if (alreadyAdded) {
      continue;
    }

    pendingAttestations.set(block, [...pending, { signer, signature }]);
    newSigners = true;
    const murmur = murmurMap.get(signer) || (await toMurmur(signer));
    cache.have = [...cache.have, { signer, signature, murmur }];
  }

  if (newSigners) {
    processAttestations({ key: block, args: [block] }).catch((err: Error) => {
      logger.error(
        `Encountered an error while processing attestations: ${err.message}`
      );
    });
  }

  for (const key of pendingAttestations.keys()) {
    if (key < block - CACHE_SIZE) {
      // FIXME: Security problem where a validator can reset another
      // FIXME: peer's cache by sending a big block number
      pendingAttestations.delete(key);
    }
  }
};

const updateAssetPrice = debounceAsync(
  async (
    block: number,
    price: number,
    signature: string,
    signers: string[]
  ) => {
    const dataset = await db.dataSet.upsert({
      where: { name: "uniswap::ethereum::ethereum" },
      update: {},
      create: { name: "uniswap::ethereum::ethereum" },
      select: { id: true },
    });

    const assetPrice = await db.assetPrice.upsert({
      where: { dataSetId_block: { dataSetId: dataset.id, block } },
      update: { signature },
      create: { dataSetId: dataset.id, block, price, signature },
      select: { id: true },
    });

    const signerNames = new Map(
      [...sockets.values()].map((item) => [item.publicKey, item.name])
    );

    for (const key of signers) {
      if (!keyToIdCache.has(key)) {
        const name = signerNames.get(key);
        const signer = await db.signer.upsert({
          where: { key },
          // see https://github.com/prisma/prisma/issues/18883
          update: { key, name },
          create: { key, name },
          select: { id: true },
        });
        keyToIdCache.set(key, signer.id);
      }
    }

    for (const key of signers) {
      const signerId = keyToIdCache.get(key) as number;
      const combo = { signerId, assetPriceId: assetPrice.id };

      // TODO: Add a upsert tracker/cache
      // Create relation in SignersOnAssetPrice
      await db.signersOnAssetPrice.upsert({
        where: { signerId_assetPriceId: combo },
        // see https://github.com/prisma/prisma/issues/18883
        update: combo,
        create: combo,
      });
    }
  },
  500
);

const printAttestations = (
  size: number,
  block: number,
  price: number,
  signersSet: string[]
) => {
  logger.info(`${size}x validations at block ${block}: $${price}`);
  if (logger.isVerboseEnabled()) {
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
        .filter(
          (peer) => peer?.publicKey && signersSet.includes(peer.publicKey)
        )
        .map((peer) => peer?.name || "?"),
      missing: allPeers
        .filter(
          (peer) => peer?.publicKey && !signersSet.includes(peer.publicKey)
        )
        .map((peer) => peer?.name || "?"),
    };
    logger.verbose(`Received signatures: ${peerStates.signed.join(", ")}`);
    logger.verbose(
      `Missing signatures: ${peerStates.missing.join(", ") || "N/A"}`
    );
  }
};

const processAttestations = debounceAsync(async (block: number) => {
  if (!blockCache.has(block)) {
    return;
  }
  const price = blockCache.get(block);
  if (typeof price !== "number") {
    return;
  }
  const data: PriceSignatureInput = { metric: { block }, value: { price } };
  const stored = attestations.get(block) || { signers: [] };
  const pending = pendingAttestations.get(block) || [];

  if (!pending.length) {
    return;
  }

  // reset pending attestations
  pendingAttestations.set(block, []);

  const currentSigners = stored.signers;
  const newSignatureSet = pending.filter(
    ({ signer }) => !currentSigners.includes(signer)
  );

  if (!newSignatureSet.length) {
    return;
  }

  // verify aggregated new signatures
  const newSigners = newSignatureSet.map((item) => item.signer);
  const signers = [...newSigners, ...stored.signers];

  const newAggregated = await aggregate(
    newSignatureSet.map((item) => item.signature)
  );
  const isValid = await verifyAggregate(newSigners, newAggregated, data);

  const validNewSigs = isValid
    ? newSignatureSet
    : newSignatureSet.filter((item) => verify({ ...item, data }));

  // add peer scores
  for (const { signer } of validNewSigs) {
    addOnePoint(signer);
  }

  const newSignatures = isValid
    ? [newAggregated]
    : newSignatureSet.map((item) => item.signature);

  const signatureList = [stored.aggregated, ...newSignatures].filter(Boolean);

  const aggregated =
    signatureList.length === 1
      ? signatureList[0]
      : await aggregate(signatureList as string[]);

  attestations.set(block, { ...stored, aggregated, signers });

  if (!config.lite) {
    updateAssetPrice({
      key: block,
      args: [block, price, aggregated, [...signers]],
    }).catch((err: Error) => {
      logger.error(
        `Error encountered while updating asset prices in the database: ${err.message}`
      );
    });
  }

  const { length } = signers;
  if (length > 1) {
    printAttestations(length, block, price, signers);
  }

  for (const key of attestations.keys()) {
    // FIXME: Security problem where a validator can reset another
    // FIXME: peer's cache by sending a big block number
    if (key < block - CACHE_SIZE) {
      attestations.delete(key);
    }
  }
  return true;
}, 500);

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
): Promise<GossipRequest<AssetPriceMetric, AssetPriceValue> | null> => {
  try {
    const start = new Date();
    const provider = getProvider(config);
    const pool = new ethers.Contract(poolAddress, poolABI, provider);
    const block = await provider.getBlockNumber();
    if (blockCache.has(block)) {
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
    if (blockCache.has(block)) {
      return null;
    }
    setCache(block, price);
    const data: PriceSignatureInput = { metric: { block }, value: { price } };
    const signed = blsAttest(data);
    const hash = await toMurmur(hashObject(data.metric));
    if (!wantCache.has(hash)) {
      wantCache.set(hash, { block, have: [] });
    }
    const cache = wantCache.get(hash);
    await addPendingAttestations(cache, block, [signed]);
    // TODO: we need to properly handle `dataset`
    return {
      method: "uniswapAttest",
      metric: { block },
      dataset: "ethereum::uniswap::ethereum",
      ...signed,
    };
  } catch (error) {
    logger.warn("Could not get the Ethereum price.");
    logger.warn("Getting a new RPC connection.");
    getNewRpc = true;
    throw error;
  }
};

export const attest: GossipMethod<AssetPriceMetric, AssetPriceValue> = async (
  request: GossipRequest<AssetPriceMetric, AssetPriceValue>
) => {
  const { metric, signer, signature } = request;
  const added = addPendingAttestation(metric.block, signer, signature);
  return added ? request : null;
};

const have = (data: WantAnswer) => {
  const cache = wantCache.get(data.want);
  if (!cache) {
    return;
  }
  addPendingAttestations(cache, cache.block, data.have);
};

const want = async (data: WantPacket) => {
  const cache = wantCache.get(data.want);
  if (!cache) {
    return [];
  }
  const have = [];
  for (const item of cache.have) {
    if (!data.have.includes(item.murmur)) {
      have.push(item);
    }
  }
  return have;
};

Object.assign(gossipMethods, { uniswapAttest: attest });
datasets.set("ethereum::uniswap::ethereum", { have, want });

export const getHave = async (want: string) => {
  const cache = wantCache.get(want);
  if (!cache) {
    return [];
  }
  return cache.have.map((item: { murmur: string }) => item.murmur);
};
