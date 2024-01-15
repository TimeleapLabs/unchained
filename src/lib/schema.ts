import { z } from "zod";
import { nameRegex } from "./constants.js";

const POSTGRES_REGEX = /^postgres(ql)?:\/\/[^:]+:[^@]+@[^\/]+\/.+$/;

export const rPCListSchema = z.object({
  ethereum: z.union([z.string().url(), z.array(z.string().url())]),
});

export const databaseConfigSchema = z.object({
  url: z.string().regex(POSTGRES_REGEX, "Not a valid Postgres URI"),
});

export const peerConfigSchema = z.object({
  max: z.number().gt(8).optional(),
  parallel: z.number().gt(4).optional(),
});

export const jailConfigSchema = z.object({
  duration: z.number().optional(),
  strikes: z.number().optional(),
});

export const gossipConfigSchema = z.object({
  infect: z.number().optional(),
  die: z.number().optional(),
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
  rpc: rPCListSchema,
  lite: z.boolean().optional(),
  database: databaseConfigSchema.optional(),
  secretKey: z.string().optional(),
  publicKey: z.string().optional(),
  peers: peerConfigSchema.optional(),
  jail: jailConfigSchema.optional(),
  waves: z.number().gte(5).optional(),
});
