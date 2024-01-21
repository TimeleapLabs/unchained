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
  max: z.number().gte(8).optional(),
  parallel: z.number().gte(4).optional(),
});

const jailConfigSchema = z.object({
  duration: z.number().min(1).max(10).optional(),
  strikes: z.number().min(5).optional(),
});

const jitterConfigSchema = z
  .object({
    max: z.number().gte(0),
    min: z.number().gte(0),
  })
  .refine(
    (schema) => schema.max > schema.min,
    "Max jitter should be bigger than min"
  );

const wavesConfigSchema = z.object({
  count: z.number().gt(5).optional(),
  select: z.number().gte(25).lte(100).optional(),
  group: z.number().gte(4).optional(),
  jitter: jitterConfigSchema.optional(),
});

export const userConfigSchema = z.object({
  name: z
    .string()
    .max(24)
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
  waves: wavesConfigSchema.optional(),
});
