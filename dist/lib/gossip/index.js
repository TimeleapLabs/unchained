import { gossipMethods, errors, keys, sockets, config } from "../constants.js";
import { encoder } from "../bls/keys.js";
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
    const payload = JSON.stringify(data);
    for (const node of nodes) {
        if (!node.socket.closed && !node.isSocketBusy) {
            const sent = node.socket.write(payload);
            if (!sent) {
                node.isSocketBusy = true;
            }
        }
    }
};
const randomDistinct = (length, count) => {
    const set = new Set();
    while (set.size < count) {
        const random = randomIndex(length);
        set.add(random);
    }
    return [...set];
};
export const gossip = async (request, seen) => {
    if (sockets.size === 0) {
        return;
    }
    assert(keys.publicKey !== undefined, "No public key found");
    const publicKey = encoder.encode(keys.publicKey.toBytes());
    if (seen.includes(publicKey)) {
        return;
    }
    const payload = { type: "gossip", request, seen: [...seen, publicKey] };
    const values = [...sockets.values()];
    const nodes = values
        .filter((node) => node.publicKey && !seen.includes(node.publicKey))
        .filter((node) => !node.isSocketBusy);
    if (!nodes.length) {
        return;
    }
    if (nodes.length <= config.gossip) {
        await gossipTo(nodes, payload);
    }
    else {
        const indexes = randomDistinct(nodes.length, config.gossip);
        const chosen = indexes.map((index) => nodes[index]);
        await gossipTo(chosen, payload);
    }
};
export const processGossip = async (incoming) => {
    try {
        // TODO: We should detect and slash nodes if they send wrong data
        const { method: methodName } = incoming.request;
        /**
         * INFO: we're using `hasOwnProperty` instead of `in` to prevent attacks
         * where a peer can access the internal properties of `gossipMethods`.
         */
        if (!gossipMethods.hasOwnProperty(methodName)) {
            return { error: errors.E_NOT_FOUND };
        }
        const method = gossipMethods[methodName];
        const payload = await method(incoming.request);
        if (payload) {
            await gossip(payload, incoming.seen);
        }
    }
    catch (error) {
        const systemError = error;
        const message = systemError.code || systemError.errno || systemError.message;
        return { error: message || errors.E_INTERNAL };
    }
};
