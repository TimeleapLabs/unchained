import crypto from "crypto";
import serialize from "json-canon";

export const sha = (input: string | Buffer) =>
  crypto.createHash("sha256").update(input).digest();

export const hashObject = (object: any) =>
  sha(serialize(object)).toString("hex");
