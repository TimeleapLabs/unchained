import { murmur332 } from "@multiformats/murmur3";
import { encoder } from "../base58/index.js";
import { cache } from "../../utils/cache.js";
import { minutes } from "../../utils/time.js";

const murmurCache = cache<string, string>(minutes(15));

export const toMurmur = async (data: string) =>
  encoder.encode(await murmur332.encode(Buffer.from(data)));

export const uint8ArrayToMurmur = async (data: Uint8Array) =>
  encoder.encode(await murmur332.encode(data));

export const toMurmurCached = async (data: string) => {
  const cached = murmurCache.get(data);
  if (cached) {
    return cached;
  }
  const murmur = await toMurmur(data);
  murmurCache.set(data, murmur);
  return murmur;
};
