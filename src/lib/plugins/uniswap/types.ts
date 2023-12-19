import {
  GossipSignatureInput,
  AssetPriceMetric,
  AssetPriceValue,
  SignatureItem,
} from "../../types.js";

export interface Attestation {
  signatures: SignatureItem[];
  aggregated?: string;
}

export type PriceSignatureInput = GossipSignatureInput<
  AssetPriceMetric,
  AssetPriceValue
>;
