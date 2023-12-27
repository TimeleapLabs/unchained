import { ethers } from "ethers";
import { gossipMethods, config, keys, sockets } from "../../constants.js";
import { logger } from "../../logger/index.js";
import { state } from "../../constants.js";
import { WS } from "iso-websocket";
import { attest as blsAttest, aggregate, verify, verifyAggregate, } from "../../bls/index.js";
import { encoder } from "../../bls/keys.js";
import { WebSocket } from "unws";
import { addOnePoint } from "../../score/index.js";
import { debounce } from "../../utils/debounce.js";
import { db } from "../../db/db.js";
const cache = new Map();
const attestations = new Map();
const pendingAttestations = new Map();
let provider;
let getNewRpc = false;
const CACHE_SIZE = 50;
const ws = (endpoint) => new WS(endpoint, {
    ws: WebSocket,
    retry: { forever: true },
});
let currentProvider = 0;
const getNextConnectionUrl = (config) => {
    if (typeof config.rpc.ethereum === "string") {
        return config.rpc.ethereum;
    }
    else {
        if (currentProvider > config.rpc.ethereum.length) {
            currentProvider = 0;
        }
        return config.rpc.ethereum[currentProvider++];
    }
};
const getProvider = (config) => {
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
const setCache = (block, price) => {
    cache.set(block, price);
    for (const key of cache.keys()) {
        if (key < block - CACHE_SIZE) {
            cache.delete(key);
        }
    }
};
const addPendingAttestation = (block, signer, signature) => {
    const pendingSignatures = pendingAttestations.get(block) || [];
    const confirmedSignatures = attestations.get(block)?.signatures;
    const alreadyAdded = pendingSignatures.some((item) => item.signer === signer) ||
        confirmedSignatures?.some((item) => item.signer === signer);
    if (!alreadyAdded) {
        pendingAttestations.set(block, [
            ...pendingSignatures,
            { signer, signature },
        ]);
        if (cache.has(block)) {
            processAttestations({ key: block, args: [block] });
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
const updateAssetPrice = debounce(async (block, price, signature, signers) => {
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
    const signerNames = new Map([...sockets.values()].map((item) => [item.publicKey, item.name]));
    for (const key of signers) {
        const name = signerNames.get(key);
        // Find or create each signer
        const signer = await db.signer.upsert({
            where: { key },
            // see https://github.com/prisma/prisma/issues/18883
            update: { key },
            create: { key, name },
            select: { id: true },
        });
        // Create relation in SignersOnAssetPrice
        await db.signersOnAssetPrice.upsert({
            where: {
                signerId_assetPriceId: {
                    signerId: signer.id,
                    assetPriceId: assetPrice.id,
                },
            },
            // see https://github.com/prisma/prisma/issues/18883
            update: { signerId: signer.id, assetPriceId: assetPrice.id },
            create: { signerId: signer.id, assetPriceId: assetPrice.id },
        });
    }
}, 500);
const printAttestations = (size, block, price, signersSet) => {
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
                .filter((peer) => peer?.publicKey && signersSet.has(peer.publicKey))
                .map((peer) => peer?.name || "?"),
            missing: allPeers
                .filter((peer) => peer?.publicKey && !signersSet.has(peer.publicKey))
                .map((peer) => peer?.name || "?"),
        };
        logger.verbose(`Received signatures: ${peerStates.signed.join(", ")}`);
        logger.verbose(`Missing signatures: ${peerStates.missing.join(", ") || "N/A"}`);
    }
};
const processAttestations = debounce(async (block) => {
    if (!cache.has(block)) {
        return;
    }
    const price = cache.get(block);
    if (typeof price !== "number") {
        return;
    }
    const data = { metric: { block }, value: { price } };
    const stored = attestations.get(block) || { signatures: [] };
    const pending = pendingAttestations.get(block) || [];
    if (!pending.length) {
        return;
    }
    // reset pending attestations
    pendingAttestations.set(block, []);
    const currentSignatures = stored.signatures.map((item) => item.signature);
    const newSignatureSet = pending.filter(({ signer }) => !currentSignatures.includes(signer));
    if (!newSignatureSet.length) {
        return;
    }
    // verify aggregated pending signatures
    const pendingSigners = pending.map((item) => item.signer);
    const pendingAggregated = aggregate(pending.map((item) => item.signature));
    const aggregatedVerify = verifyAggregate(pendingSigners, pendingAggregated, data);
    const pendingSignatures = aggregatedVerify
        ? pending
        : pending.filter((item) => verify({ ...item, data }));
    const allSignatures = [...stored.signatures, ...pendingSignatures];
    const uniqueSignatures = [];
    const signersSet = new Set();
    for (const item of allSignatures) {
        if (!signersSet.has(item.signer)) {
            signersSet.add(item.signer);
            uniqueSignatures.push(item);
        }
    }
    // add peer scores
    for (const { signer } of newSignatureSet) {
        addOnePoint(signer);
    }
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
        updateAssetPrice({
            key: block,
            args: [block, price, aggregated, [...signersSet]],
        });
    }
    const { size } = signersSet;
    if (size > 1) {
        printAttestations(size, block, price, signersSet);
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
export const work = async (_parameters, poolAddress, decimals, inverse) => {
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
            logger.warn(`Detected high latency with the Ethereum RPC node: ${took}ms`);
        }
        else if (state.connected) {
            logger.debug(`Request to Ethereum RPC node took ${took}ms`);
        }
        if (cache.has(block)) {
            return null;
        }
        setCache(block, price);
        const data = { metric: { block }, value: { price } };
        const signed = blsAttest(data);
        addPendingAttestation(block, signed.signer, signed.signature);
        // TODO: we need to properly handle `dataset`
        return {
            method: "uniswapAttest",
            metric: { block },
            dataset: "ethereum::uniswap::ethereum",
            ...signed,
        };
    }
    catch (error) {
        logger.warn("Could not get the Ethereum price.");
        logger.warn("Getting a new RPC connection.");
        getNewRpc = true;
        throw error;
    }
};
export const attest = async (request) => {
    const { metric, signer, signature } = request;
    const added = addPendingAttestation(metric.block, signer, signature);
    return added ? request : null;
};
Object.assign(gossipMethods, { uniswapAttest: attest });