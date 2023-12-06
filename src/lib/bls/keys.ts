import bls from "@chainsafe/bls";
import { Base58 } from "base-ex";

import { SecretKey, PublicKey } from "@chainsafe/bls/types";

interface KeyPair {
  secretKey: SecretKey;
  publicKey: PublicKey;
}

interface EncodedKeyPair {
  secretKey: string;
  publicKey: string;
}

const getPair = (secretKey: SecretKey): KeyPair => {
  const publicKey = secretKey.toPublicKey();
  return { secretKey, publicKey };
};

export const makeKeys = (): KeyPair => {
  const secretKey = bls.SecretKey.fromKeygen();
  return getPair(secretKey);
};

export const loadKeys = (config: EncodedKeyPair): KeyPair => {
  const decoded = decodeKeys(config);
  const secretKey = bls.SecretKey.fromBytes(decoded.secretKey);
  return getPair(secretKey);
};

export const encoder = new Base58("bitcoin");

export const encodeKeys = (pair: KeyPair): EncodedKeyPair => {
  return {
    secretKey: encoder.encode(pair.secretKey.toBytes()),
    publicKey: encoder.encode(pair.publicKey.toBytes()),
  };
};

export const decodeKeys = (
  pair: EncodedKeyPair
): { secretKey: Buffer; publicKey: Buffer } => {
  return {
    secretKey: Buffer.from(encoder.decode(pair.secretKey)),
    publicKey: Buffer.from(encoder.decode(pair.publicKey)),
  };
};
