import { epoch } from "./time.js";

export interface Cache<K, V> {
  set(key: K, value: V): Map<K, V>;
  get(key: K): V | undefined;
  has(key: K): boolean;
  entries(): IterableIterator<[K, V]>;
}

export const cache = <K, V>(ttl: number): Cache<K, V> => {
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
    set(key: K, value: V): Map<K, V> {
      addTtlAndCleanup(key);
      return map.set(key, value);
    },
    get(key: K): V | undefined {
      return map.get(key);
    },
    has(key: K): boolean {
      return map.has(key);
    },
    entries(): IterableIterator<[K, V]> {
      return map.entries();
    },
  };
};
