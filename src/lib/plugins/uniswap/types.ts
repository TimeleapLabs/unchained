import {
  SignatureInput,
  AssetPriceMetric,
  AssetPriceValue,
} from "../../types.js";

export interface Attestation {
  signers: Uint8Array[];
  aggregated?: Uint8Array;
}

export type PriceSignatureInput = SignatureInput<
  AssetPriceMetric,
  AssetPriceValue
>;
