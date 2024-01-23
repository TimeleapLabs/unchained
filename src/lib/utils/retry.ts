import { seconds } from "./time.js";

export const TIMEOUT = Symbol("TIMEOUT");
export const CALLERROR = Symbol("CALLERROR");

const timeout = (ms: number): Promise<Symbol> =>
  new Promise((resolve) => setTimeout(resolve, ms, TIMEOUT));

export const withTimeout = async <T>(
  promise: Promise<T>,
  ms: number
): Promise<T | Symbol> => {
  try {
    return Promise.race([promise, timeout(ms)]);
  } catch (error) {
    return CALLERROR;
  }
};

const retryTimes: { [key: number]: number } = {
  3: 5000,
  2: 15000,
  1: 45000,
};

export const runWithRetries = async <T>(
  fn: (...args: any[]) => Promise<T>,
  args: any[],
  timeoutMs: number = seconds(1),
  retries: number = 3
): Promise<T | Symbol> => {
  while (true) {
    const promise = fn.apply(null, args);
    const result = await withTimeout<T>(promise, timeoutMs);
    const needsRetry = result === TIMEOUT || result === CALLERROR;
    if (needsRetry && retries) {
      const nextRun = retryTimes[retries--];
      await timeout(nextRun);
    } else {
      return result;
    }
  }
};
