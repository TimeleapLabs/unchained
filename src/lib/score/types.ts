import { SignatureInput } from "../types.js";

export type ScoreSignatureInput = SignatureInput<ScoreMetric, ScoreValues>;

export interface ScoreMetric {
  sprint: number;
}

export type ScoreValues = {
  [key: string]: number;
};
