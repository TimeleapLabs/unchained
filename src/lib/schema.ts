import { z } from "zod";
import { nameRegex } from "./constants.js";

const POSTGRES_REGEX = /^postgres(ql)?:\/\/[^:]+:[^@]+@[^\/]+\/.+$/;

const rpcListSchema = z.object({
  ethereum: z.union([z.string().url(), z.array(z.string().url())]),
});

const databaseConfigSchema = z.object({
  url: z.string().regex(POSTGRES_REGEX, "Not a valid Postgres URI"),
});

const peerConfigSchema = z.object({
  max: z.number().gt(8).optional(),
  parallel: z.number().gt(4).optional(),
});

const jailConfigSchema = z.object({
  duration: z.number().optional(),
  strikes: z.number().optional(),
});

export const userConfigSchema = z.object({
  name: z
    .string()
    .regex(
      nameRegex,
      "Only English letters, numbers, and @._'- are allowed in the name"
    ),
  log: z
    .enum(["error", "warn", "info", "verbose", "debug", "silly"])
    .optional(),
  rpc: rpcListSchema,
  lite: z.boolean().optional(),
  database: databaseConfigSchema.optional(),
  secretKey: z.string().optional(),
  publicKey: z.string().optional(),
  peers: peerConfigSchema.optional(),
  jail: jailConfigSchema.optional(),
  waves: z.number().gte(5).optional(),
});
