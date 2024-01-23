import { ethers } from "ethers";
import { config, sockets } from "../../constants.js";
import { logger } from "../../logger/index.js";
import { state } from "../../constants.js";
import { WS } from "iso-websocket";
import { attest as blsAttest, verify } from "../../crypto/bls/index.js";
import { aggregate, verifyAggregate } from "../../crypto/bls/threads/index.js";
import { WebSocketLike } from "ethers";
import { WebSocket } from "unws";
import { addOnePoint } from "../../score/index.js";
import { debounceAsync } from "../../utils/debounce.js";
import { db } from "../../db/db.js";
import { datasets } from "../../network/index.js";
import { cache } from "../../utils/cache.js";
import { isEqual, hashUint8Array } from "../../utils/uint8array.js";

import type { WantPacket, WantAnswer } from "../../network/index.js";

import {
  Config,
  AssetPriceMetric,
  WaveRequest,
  SignatureItem,
  AssetPriceValue,
} from "../../types.js";

import { Attestation, PriceSignatureInput } from "./types.js";
import { minutes } from "../../utils/time.js";
import { hashObject } from "../../utils/hash.js";
import { toMurmurCached } from "../../crypto/murmur/index.js";
import { encoder } from "../../crypto/base58/index.js";

const blockCache = new Map<number, number>();
const attestations = new Map<number, Attestation>();
const pendingAttestations = new Map<number, SignatureItem[]>();
const keyToIdCache = new Map<string, number>();
const waveCache = cache<string, any>(minutes(15));

let provider: ethers.Provider | undefined;
let getNewRpc: boolean = false;

const CACHE_SIZE = 50;

const ws = (endpoint: string): WebSocketLike => {
  const socket = new WS(endpoint, {
    ws: WebSocket,
    retry: { forever: true },
    automaticOpen: false,
  });

  socket.addEventListener("error", () => {
    getNewRpc = true;
  });

  socket.open();

  return socket as WebSocketLike;
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

const addPendingAttestations = async (
  cache: any,
  block: number,
  signatures: SignatureItem[]
) => {
  if (!pendingAttestations.has(block)) {
    pendingAttestations.set(block, []);
  }

  const pending = pendingAttestations.get(block) as SignatureItem[];
  const confirmed = attestations.get(block)?.signers;

  let newSigners = false;

  for (const { signature, signer } of signatures) {
    if (!signature || !signer) {
      continue;
    }

    if (!(signature instanceof Uint8Array)) {
      continue;
    }

    if (!(signer instanceof Uint8Array)) {
      continue;
    }

    const alreadyAdded =
      pending.some((item) => isEqual(item.signer, signer)) ||
      confirmed?.some((cSigner) => isEqual(cSigner, signer));

    if (alreadyAdded) {
      continue;
    }

    pending.push({ signer, signature });
    newSigners = true;

    const murmur = await hashUint8Array(signer);
    cache.have = [...cache.have, { signer, signature, murmur }];
  }

  if (newSigners) {
    processAttestations({ key: block, args: [block] }).catch((err: Error) => {
      if (err) {
        logger.error(
          `Encountered an error while processing attestations: ${err.message}`
        );
      }
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

interface Signer {
  publicKey: Buffer;
  name?: string;
}

const updateAssetPrice = debounceAsync(
  async (
    block: number,
    price: number,
    uintSignature: Uint8Array,
    uintSigners: Uint8Array[]
  ) => {
    const signature = Buffer.from(uintSignature);
    const signers = uintSigners.map((arr) => Buffer.from(arr));

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

    const signerNames = new Map<string, string>();

    for (const peer of sockets.values()) {
      if (peer.publicKey) {
        const hash = peer.murmurAddr || (await hashUint8Array(peer.publicKey));
        signerNames.set(hash, peer.name);
      }
    }

    for (const key of signers) {
      const hash = await hashUint8Array(key);
      const oldKey = encoder.encode(key);

      if (!keyToIdCache.has(hash)) {
        const name = signerNames.get(hash);

        const signer = await db.signer.upsert({
          where: { oldKey },
          // see https://github.com/prisma/prisma/issues/18883
          update: { oldKey, key, name },
          create: { oldKey, key, name },
          select: { id: true },
        });
        keyToIdCache.set(hash, signer.id);
      }

      const signerId = keyToIdCache.get(hash) as number;
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

const printAttestations = async (
  size: number,
  block: number,
  price: number,
  signersSet: Uint8Array[]
) => {
  logger.info(`${size}x validations at block ${block}: $${price}`);
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
    ({ signer }) => !currentSigners.some((item) => isEqual(item, signer))
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
      : await aggregate(signatureList as Uint8Array[]);

  attestations.set(block, { ...stored, aggregated, signers });

  if (!config.lite) {
    updateAssetPrice({
      key: block,
      args: [block, price, aggregated, [...signers]],
    }).catch((err: Error) => {
      if (err) {
        logger.error(
          `Error encountered while updating asset prices in the database: ${err.message}`
        );
      }
    });
  }

  const { length } = signers;
  if (length > 1) {
    printAttestations(length, block, price, signers).catch(() => null);
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
): Promise<WaveRequest<AssetPriceMetric, AssetPriceValue> | null> => {
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
    const hash = await toMurmurCached(hashObject(data.metric));
    if (!waveCache.has(hash)) {
      waveCache.set(hash, { block, have: [] });
    }
    const cache = waveCache.get(hash);
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

const have = (data: WantAnswer) => {
  const cache = waveCache.get(data.want);
  if (!cache) {
    return;
  }
  addPendingAttestations(cache, cache.block, data.have);
};

const want = async (data: WantPacket) => {
  const cache = waveCache.get(data.want);
  if (!cache) {
    return [];
  }
  return cache.have.filter((item: any) => !data.have.includes(item.murmur));
};

datasets.set("ethereum::uniswap::ethereum", { have, want });

export const getHave = async (want: string) => {
  const cache = waveCache.get(want);
  if (!cache) {
    return [];
  }
  return cache.have.map((item: { murmur: string }) => item.murmur);
};
