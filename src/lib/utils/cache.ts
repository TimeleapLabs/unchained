import { epoch } from "./time.js";

export const cache = <K, V>(ttl: number) => {
  const map = new Map<K, V>();
  const timeouts = new Map<K, number>();

  const addTtlAndCleanup = (key: K) => {
    const now = epoch();
    timeouts.set(key, now);
    for (const key of timeouts.keys()) {
      if ((timeouts.get(key) as number) + ttl < now) {
        timeouts.delete(key);
        map.delete(key);
      }
    }
  };

  return {
    set(key: K, value: V) {
      addTtlAndCleanup(key);
      return map.set(key, value);
    },
    get(key: K): V | undefined {
      return map.get(key);
    },
    has(key: K): boolean {
      return map.has(key);
    },
  };
};
