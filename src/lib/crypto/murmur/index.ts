import { murmur332 } from "@multiformats/murmur3";
import { encoder } from "../base58/index.js";

export const toMurmur = async (data: string) =>
  encoder.encode(await murmur332.encode(Buffer.from(data)));
