import { z } from "zod";
import { nameRegex } from "./constants.js";

export const rPCListSchema = z.object({
  ethereum: z.union([z.string().url(), z.array(z.string().url())]),
});

export const databaseConfigSchema = z.object({
  url: z.string(),
});

export const peerConfigSchema = z.object({
  max: z.number(),
  parallel: z.number(),
});

export const jailConfigSchema = z.object({
  duration: z.number(),
  strikes: z.number(),
});

export const gossipConfigSchema = z.object({
  infect: z.number(),
  die: z.number(),
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
  gossip: gossipConfigSchema.optional(),
});
