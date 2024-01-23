import { encoder } from "../crypto/base58/index.js";
import { config, keys, rpcMethods, errors, version } from "../constants.js";
import { MetaData, NodeSystemError, IntroducePayload } from "../types.js";
import { murmur } from "../constants.js";
import assert from "assert";

interface RpcRequest {
  method: string;
  args?: any; // Replace 'any' with a more specific type if available
}

const defaultMethods = {
  timestamp: (): number => new Date().valueOf(),
  introduce: (): IntroducePayload => {
    assert(keys.publicKey !== undefined, "Public key not found");
    return {
      name: config.name,
      publicKey: keys.publicKey.toBytes(),
      murmurAddr: murmur.address,
      client: {
        version,
        peers: config.peers,
        waves: config.waves,
      },
    };
  },
};

Object.assign(rpcMethods, defaultMethods);

export const processRpc = async (
  message: { request: RpcRequest },
  sender: MetaData
): Promise<{ result?: any; error?: string | number }> => {
  try {
    if (sender.needsDrain) {
      return { error: errors.E_TOO_MANY_REQUESTS };
    }

    const { method, args } = message.request;

    if (!(method in rpcMethods)) {
      return { error: errors.E_NOT_FOUND };
    }

    const result = await rpcMethods[method].call(null, args, sender);
    return { result };
  } catch (error) {
    const errno = (error as NodeSystemError).code;
    return { error: errno || errors.E_INTERNAL };
  }
};
