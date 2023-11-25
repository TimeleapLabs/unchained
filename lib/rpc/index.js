import * as uniswap from "../plugins/uniswap/index.js";

const errors = {
  E_NOT_FOUND: 404,
  E_INTERNAL: 500,
};

const methods = {
  timestamp: () => new Date().valueOf(),
  ...uniswap.plugin,
};

const thisArg = {};

export const rpc = async (message) => {
  const { method, args } = message;

  if (!(method in methods)) {
    return { error: errors.E_NOT_FOUND };
  }

  try {
    const result = await methods[method].apply(thisArg, args);
    return { result };
  } catch (error) {
    return { error: error.code || E_INTERNAL };
  }
};

export const setup = (config) => {
  thisArg.config = config;
};
