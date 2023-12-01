import { encoder } from "./keys.js";
import bls from "@chainsafe/bls";
import { keys } from "../constants.js";
import stringify from "json-canon";

export const sign = (data) => {
  const json = stringify(data);
  const buffer = Buffer.from(json, "utf8");
  return encoder.encode(keys.secretKey.sign(buffer).toBytes());
};

export const attest = (request) => {
  const signer = encoder.encode(keys.publicKey.toBytes());
  const signature = sign(request);
  return { request, signature, signer };
};

export const verify = ({ signer, signature, request }) => {
  const message = Buffer.from(stringify(request), "utf8");
  const publicKey = bls.PublicKey.fromBytes(
    Buffer.from(encoder.decode(signer))
  );
  const decodedSignature = bls.Signature.fromBytes(
    Buffer.from(encoder.decode(signature))
  );
  return decodedSignature.verify(publicKey, message);
};

export const aggregate = (aggregated, signature) =>
  encoder.encode(
    bls.Signature.aggregate([
      bls.Signature.fromBytes(Buffer.from(encoder.decode(aggregated))),
      bls.Signature.fromBytes(Buffer.from(encoder.decode(signature))),
    ]).toBytes()
  );
