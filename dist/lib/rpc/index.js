import { encoder } from "../bls/keys.js";
import { config, keys, rpcMethods, errors } from "../constants.js";
import assert from "assert";
const defaultMethods = {
    timestamp: () => new Date().valueOf(),
    introduce: async () => {
        assert(keys.publicKey !== undefined, "Public key not found");
        return {
            name: config.name,
            publicKey: encoder.encode(keys.publicKey.toBytes()),
        };
    },
};
Object.assign(rpcMethods, defaultMethods);
const thisArg = {};
export const processRpc = async (message) => {
    try {
        const { method, args } = message.request;
        if (!(method in rpcMethods)) {
            return { error: errors.E_NOT_FOUND };
        }
        const result = await rpcMethods[method].call(thisArg, args);
        return { result };
    }
    catch (error) {
        const errno = error.code;
        return { error: errno || errors.E_INTERNAL };
    }
};
