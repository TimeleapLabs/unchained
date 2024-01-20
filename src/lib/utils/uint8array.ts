import { uint8ArrayToMurmur } from "../crypto/murmur/index.js";

export const isEqual = (lhs: Uint8Array, rhs: Uint8Array) =>
  lhs.every((v, i) => v === rhs[i]);

export const hashUint8Array = (array: Uint8Array) => uint8ArrayToMurmur(array);

export const copyUint8Array = (array: Uint8Array) => Uint8Array.from(array);
