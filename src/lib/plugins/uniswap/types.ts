import {
  GossipSignatureInput,
  AssetPriceMetric,
  AssetPriceValue,
} from "../../types.js";

export interface Attestation {
  signers: string[];
  aggregated?: string;
}

export type PriceSignatureInput = GossipSignatureInput<
  AssetPriceMetric,
  AssetPriceValue
>;
