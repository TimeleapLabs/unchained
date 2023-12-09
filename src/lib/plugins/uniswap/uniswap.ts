import { ethers } from "ethers";
import { gossipMethods, config, keys, sockets } from "../../constants.js";
import { logger } from "../../logger/index.js";
import { state } from "../../constants.js";
import { WS } from "iso-websocket";
import { hashObject } from "../../utils/hash.js";
import { attest as blsAttest, aggregate } from "../../bls/index.js";
import { encoder } from "../../bls/keys.js";
import { WebSocketLike } from "ethers";
import { Config } from "../../types.js";
import { WebSocket } from "unws";
import { assetPrices } from "../../db/collections/AssetPrice.js";

const cache = new Map<number, number>();
const attestations = new Map<string, any>(); // Replace 'any' with a more specific type if available
const earlyAttestations = new Map<string, any>(); // Replace 'any' with a more specific type if available

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
  request: any,
  signer: string,
  signature: string
) => {
  const hash = hashObject(request);
  const current = earlyAttestations.get(hash) || { signers: [] };
  const signersList = current.signers.map((sig: any) => sig.signer);
  if (!signersList.includes(signer)) {
    const signers = [...current.signers, { signer, signature }];
    earlyAttestations.set(hash, { signers, request });
  }
  for (const [key, value] of earlyAttestations.entries()) {
    if (value.block < request.data.block - CACHE_SIZE) {
      // FIXME: Security problem where a validator can reset another
      // FIXME: peer's cache by sending a big block number
      attestations.delete(key);
    }
  }
};

const setAttestations = async (
  request: any,
  signer: string,
  signature: string
) => {
  const hash = hashObject(request);
  const stored = attestations.get(hash) || { signers: [], request };
  const early = earlyAttestations.get(hash) || { signers: [] };
  const signersList = [
    ...new Set([
      ...stored.signers,
      ...early.signers.map((sig: any) => sig.signer),
    ]),
  ];
  if (!signersList.includes(signer)) {
    const newSignatures = [...early.signers, { signer, signature }]
      .filter((sig: any) => !stored.signers.includes(sig.signer))
      .map((sig: any) => sig.signature);
    const signatureList = stored.signature
      ? [stored.signature, ...newSignatures]
      : newSignatures;
    const aggregatedSignature = aggregate(signatureList);
    const signers = [...signersList, signer];
    attestations.set(hash, {
      ...stored,
      signers,
      signature: aggregatedSignature,
    });
    if (!config.lite) {
      assetPrices.updateOne(
        {
          block: request.data.block,
          asset: "ethereum",
          source: "uniswap-ethereum",
        },
        {
          $set: {
            price: request.data.price,
            signature: aggregatedSignature,
            signers,
          },
          $setOnInsert: {
            timestamp: new Date(), // FIXME
            asset: "ethereum",
            source: "uniswap-ethereum",
            block: request.data.block,
          },
        },
        { upsert: true }
      );
    }
    const { length } = signers;
    if (length > 1) {
      const { block, price } = request.data;
      logger.info(`${length}x validations at block ${block}: $${price}`);
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
          .filter((peer) => signers.includes(peer.publicKey))
          .map((peer) => peer.name),
        missing: allPeers
          .filter((peer) => !signers.includes(peer.publicKey))
          .map((peer) => peer.name),
      };
      logger.verbose(`Received signatures: ${peerStates.signed.join(", ")}`);
      logger.verbose(
        `Missing signatures: ${peerStates.missing.join(", ") || "N/A"}`
      );
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

export const work = async (
  parameters: any,
  poolAddress: string,
  decimals: [number, number],
  inverse: boolean
): Promise<any | null> => {
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
    const payload = blsAttest(request);
    await setAttestations(request, payload.signer, payload.signature);
    return payload;
  } catch (error) {
    logger.warn("Could not get the Ethereum price. Check your RPC.");
    throw error;
  }
};

export const attest = async ({
  request,
  signer,
  signature,
}: {
  request: any;
  signer: string;
  signature: string;
}): Promise<boolean | null> => {
  const { data } = request;

  if (!cache.has(data.block)) {
    setEarlyAttestations(request, signer, signature);
    return null;
  }
  const price = cache.get(data.block);
  if (price !== data.price) {
    return false;
  }
  await setAttestations(request, signer, signature);
  return true;
};

Object.assign(gossipMethods, { uniswapAttest: attest });
