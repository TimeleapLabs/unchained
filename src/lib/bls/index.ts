import { encoder } from "./keys.js";
import bls from "@chainsafe/bls";
import { keys } from "../constants.js";
import stringify from "json-canon";
import assert from "node:assert";

import { GossipSignatureInput, SignatureItem } from "../types.js";

export const sign = (data: any): string => {
  assert(keys.secretKey !== undefined, "No secret key in config");
  const json = stringify(data);
  const buffer = Buffer.from(json, "utf8");
  return encoder.encode(keys.secretKey.sign(buffer).toBytes());
};

export const attest = (
  payload: GossipSignatureInput<any, any>
): SignatureItem => {
  assert(keys.publicKey !== undefined, "No public key in config");
  const signer = encoder.encode(keys.publicKey.toBytes());
  const signature = sign(payload);
  return { signer, signature };
};

export const verify = ({
  signer,
  signature,
  data,
}: {
  signer: string;
  signature: string;
  data: any;
}): boolean => {
  const message = Buffer.from(stringify(data), "utf8");
  const publicKey = bls.PublicKey.fromBytes(
    Buffer.from(encoder.decode(signer))
  );
  const decodedSignature = bls.Signature.fromBytes(
    Buffer.from(encoder.decode(signature))
  );
  return decodedSignature.verify(publicKey, message);
};

export const verifyAggregate = (
  signers: string[],
  signature: string,
  data: any
): boolean => {
  const message = Buffer.from(stringify(data), "utf8");
  const decodedSignature = bls.Signature.fromBytes(
    Buffer.from(encoder.decode(signature))
  );
  const publicKeys = signers.map((signer) =>
    bls.PublicKey.fromBytes(Buffer.from(encoder.decode(signer)))
  );
  return decodedSignature.verifyAggregate(publicKeys, message);
};

export const aggregate = (signatures: string[]): string =>
  encoder.encode(
    bls.Signature.aggregate(
      signatures.map((signature) =>
        bls.Signature.fromBytes(Buffer.from(encoder.decode(signature)))
      )
    ).toBytes()
  );
