import crypto from "crypto";

export const topic = crypto
  .createHash("sha256")
  .update("Kenshi.Unchained.Topic")
  .digest();
