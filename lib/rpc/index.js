import { loadKeys } from "../bls/keys.js";
import * as uniswap from "../plugins/uniswap/uniswap.js";
import { gossip } from "../gossip/index.js";

const errors = {
  E_NOT_FOUND: 404,
  E_INTERNAL: 500,
};

const methods = {
  timestamp: () => new Date().valueOf(),
  attest: async (_type, payload) => {
    // TODO: we are ignoring type for now
    const result = await uniswap.attest(
      payload.data,
      payload.signature,
      payload.signers
    );
    if (result) {
      await gossip(
        JSON.stringify({
          method: "attest",
          args: ["ethPrice", result],
        })
      );
    }
    return result;
  },
};

const thisArg = {};
const bls = {};

export const rpc = async (message) => {
  const { method, args } = message;

  if (!(method in methods)) {
    return { error: errors.E_NOT_FOUND };
  }

  try {
    const result = await methods[method].apply(thisArg, args);
    return { result };
  } catch (error) {
    return { error: error.code || errors.E_INTERNAL };
  }
};

export const setup = (config) => {
  thisArg.config = config;
  bls.keys = loadKeys(config);
};
