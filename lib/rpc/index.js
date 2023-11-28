import { encoder } from "../bls/keys.js";
import { config, keys, rpcMethods, errors } from "../constants.js";

const defaultMethods = {
  timestamp: () => new Date().valueOf(),
  introduce: async () => {
    return {
      name: config.name,
      publicKey: encoder.encode(keys.publicKey.toBytes()),
    };
  },
};

Object.assign(rpcMethods, defaultMethods);

const thisArg = {};

export const processRpc = async (message) => {
  const { method, args } = message.request;

  if (!(method in rpcMethods)) {
    return { error: errors.E_NOT_FOUND };
  }

  try {
    const result = await rpcMethods[method].call(thisArg, args);
    return { result };
  } catch (error) {
    return { error: error.code || errors.E_INTERNAL };
  }
};
