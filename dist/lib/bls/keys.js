import bls from "@chainsafe/bls";
import { Base58 } from "base-ex";
const getPair = (secretKey) => {
    const publicKey = secretKey.toPublicKey();
    return { secretKey, publicKey };
};
export const makeKeys = () => {
    const secretKey = bls.SecretKey.fromKeygen();
    return getPair(secretKey);
};
export const loadKeys = (encodedSecretKey) => {
    const decoded = Buffer.from(encoder.decode(encodedSecretKey));
    const secretKey = bls.SecretKey.fromBytes(decoded);
    return getPair(secretKey);
};
export const encoder = new Base58("bitcoin");
export const encodeKeys = (pair) => {
    return {
        secretKey: encoder.encode(pair.secretKey.toBytes()),
        publicKey: encoder.encode(pair.publicKey.toBytes()),
    };
};
export const decodeKeys = (pair) => {
    return {
        secretKey: Buffer.from(encoder.decode(pair.secretKey)),
        publicKey: Buffer.from(encoder.decode(pair.publicKey)),
    };
};
