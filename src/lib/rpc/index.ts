import { encoder } from "../bls/keys.js";
import { config, keys, rpcMethods, errors } from "../constants.js";
import { NodeSystemError } from "../types.js";
import assert from "assert";

interface RpcRequest {
  method: string;
  args?: any; // Replace 'any' with a more specific type if available
}

const defaultMethods = {
  timestamp: (): number => new Date().valueOf(),
  introduce: async (): Promise<{ name: string; publicKey: string }> => {
    assert(keys.publicKey !== undefined, "Public key not found");
    return {
      name: config.name,
      publicKey: encoder.encode(keys.publicKey.toBytes()),
    };
  },
};

Object.assign(rpcMethods, defaultMethods);

const thisArg = {};

export const processRpc = async (message: {
  request: RpcRequest;
}): Promise<{ result?: any; error?: string | number }> => {
  try {
    const { method, args } = message.request;

    if (!(method in rpcMethods)) {
      return { error: errors.E_NOT_FOUND };
    }

    const result = await rpcMethods[method].call(thisArg, args);
    return { result };
  } catch (error) {
    const errno = (error as NodeSystemError).code;
    return { error: errno || errors.E_INTERNAL };
  }
};
