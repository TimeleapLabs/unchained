import { gossipMethods, errors, keys, sockets } from "../constants.js";
import { attest, verify } from "../bls/index.js";
import { encoder } from "../bls/keys.js";
import { logger } from "../logger/index.js";
import crypto from "crypto";
import assert from "node:assert";
const randomIndex = (length) => {
    if (length <= 0) {
        throw new Error("Array length must be greater than 0.");
    }
    let index, randomByte;
    do {
        randomByte = crypto.randomBytes(1)[0];
        index = randomByte % length;
    } while (randomByte - index >= 256 - (256 % length));
    return index;
};
const gossipTo = async (nodes, data) => {
    const promises = nodes.map(({ socket }) => socket.write(JSON.stringify(data)));
    await Promise.all(promises).catch(() => null);
};
export const gossip = async ({ request, signature, signer, seen, }) => {
    if (sockets.size === 0) {
        return;
    }
    assert(keys.publicKey !== undefined, "No public key found");
    const publicKey = encoder.encode(keys.publicKey.toBytes());
    if (seen.includes(publicKey)) {
        return;
    }
    const payload = {
        type: "gossip",
        request,
        signature,
        signer,
        seen: [...seen, publicKey],
    };
    const values = Array.from(sockets.values());
    const nodes = values.filter((node) => node.publicKey && !seen.includes(node.publicKey));
    if (!nodes.length) {
        return;
    }
    if (nodes.length <= 3) {
        await gossipTo(nodes, payload);
    }
    else {
        const random = new Array(3).fill(null).map(() => randomIndex(nodes.length));
        const chosen = [...new Set(random)].map((index) => nodes[index]);
        await gossipTo(chosen, payload);
    }
};
const verifyGossip = (incoming, connection) => {
    if (!connection.publicKey) {
        return false;
    }
    if (verify(incoming)) {
        logger.debug(`Successfully verified packet sent by ${connection.publicKey}`);
        return true;
    }
    else {
        logger.warn(`Couldn't verify packet sent by ${connection.publicKey}`);
        return false;
    }
};
export const processGossip = async (incoming, connection) => {
    try {
        if (!verifyGossip(incoming, connection)) {
            return;
        }
        const { method } = incoming.request;
        if (!(method in gossipMethods)) {
            return { error: errors.E_NOT_FOUND };
        }
        const ok = await gossipMethods[method].call(null, incoming, connection);
        if (!ok) {
            return;
        }
        const payload = attest(incoming.request);
        if (payload) {
            await gossip({ ...payload, seen: incoming.seen });
        }
    }
    catch (error) {
        const systemError = error;
        const message = systemError.code || systemError.errno || systemError.message;
        return { error: message || errors.E_INTERNAL };
    }
};
