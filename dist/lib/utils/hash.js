import crypto from "crypto";
import serialize from "json-canon";
export const sha = (input) => crypto.createHash("sha256").update(input).digest();
export const hashObject = (object) => sha(serialize(object)).toString("hex");
