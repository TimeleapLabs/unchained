const TIMEOUT = Symbol("TIMEOUT");
const CALLERROR = Symbol("CALLERROR");

const timeout = (ms) =>
  new Promise((resolve) => setTimeout(resolve, ms, TIMEOUT));

export const withTimeout = async (promise, ms) =>
  Promise.race([promise, timeout(ms)]);

const retryTimes = {
  3: 5000,
  2: 15000,
  1: 45000,
};

export const runWithRetries = async (
  fn,
  args,
  timeoutMs = 5000,
  retries = 3
) => {
  while (true) {
    const promise = fn.apply(null, args);
    const result = await withTimeout(promise, timeoutMs).catch(() => CALLERROR);
    const needsRetry = result === TIMEOUT || result === CALLERROR;
    if (needsRetry && retries) {
      const nextRun = retryTimes[retries--];
      await timeout(nextRun);
    } else {
      return result;
    }
  }
};
