export const epoch = () => new Date().valueOf();

export const sleep = (ms: number) =>
  new Promise((resolve) => setTimeout(resolve, ms));
