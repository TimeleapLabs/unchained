const errors = {
  E_NOT_FOUND: 404,
};

const methods = {
  timestamp: () => new Date().valueOf(),
};

export const rpc = async (message) => {
  const { method, args } = message;

  if (!(method in methods)) {
    return { error: errors.E_NOT_FOUND };
  }

  return { result: await methods[method].apply(null, args) };
};
