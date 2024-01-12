export const epoch = () => new Date().valueOf();

export const sleep = (ms: number) =>
  new Promise((resolve) => setTimeout(resolve, ms));

export const seconds = (s: number) => s * 1000;
export const minutes = (m: number) => m * seconds(60);
