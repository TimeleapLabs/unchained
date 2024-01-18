import bls from "@chainsafe/bls";
import { keys } from "../../constants.js";
import stringify from "json-canon";
import assert from "node:assert";

import { SignatureInput, SignatureItem } from "../../types.js";

export const sign = (data: any): Uint8Array => {
  assert(keys.secretKey !== undefined, "No secret key in config");
  const json = stringify(data);
  const buffer = Buffer.from(json, "utf8");
  return keys.secretKey.sign(buffer).toBytes();
};

export const attest = (payload: SignatureInput<any, any>): SignatureItem => {
  assert(keys.publicKey !== undefined, "No public key in config");
  const signer = keys.publicKey.toBytes();
  const signature = sign(payload);
  return { signer, signature };
};

export const verify = ({
  signer,
  signature,
  data,
}: {
  signer: Uint8Array;
  signature: Uint8Array;
  data: any;
}): boolean => {
  const message = Buffer.from(stringify(data), "utf8");
  const publicKey = bls.PublicKey.fromBytes(signer);
  const decodedSignature = bls.Signature.fromBytes(signature);
  return decodedSignature.verify(publicKey, message);
};

export const verifyAggregate = (
  signers: Uint8Array[],
  signature: Uint8Array,
  data: any
): boolean => {
  const message = Buffer.from(stringify(data), "utf8");
  const decodedSignature = bls.Signature.fromBytes(signature);
  const publicKeys = signers.map((key) => bls.PublicKey.fromBytes(key));
  return decodedSignature.verifyAggregate(publicKeys, message);
};

export const aggregate = (signatures: Uint8Array[]): Uint8Array =>
  bls.Signature.aggregate(
    signatures.map((signature) => bls.Signature.fromBytes(signature))
  ).toBytes();
