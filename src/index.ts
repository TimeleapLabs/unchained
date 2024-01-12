#!/usr/bin/env node

import { program } from "commander";
import { startAction } from "./lib/cli/actions/start.js";
import { addressAction } from "./lib/cli/actions/address.js";
import { initDbAction } from "./lib/cli/actions/postgres/migrate.js";
import { generateDbAction } from "./lib/cli/actions/postgres/generate.js";
import { revertDbAction } from "./lib/cli/actions/postgres/revert.js";
import { diagnoseAction } from "./lib/cli/actions/diagnose.js";
import { version } from "./lib/constants.js";

program
  .name("unchained")
  .description("Kenshi Unchained CLI Node")
  .version(version);

// TODO: Expose more options from the config file here
program
  .command("start")
  .description("start a node on the Kenshi Unchained HyperSwarm")
  .argument("<config>", "config file in YAML format")
  .option("--log <level>", "log level")
  .option("--lite", "run in lite mode")
  .option("--generate", "generate a secret key")
  .option("--max-peers <max>", "set max allowed active peer connections")
  .option(
    "--parallel-peers <parallel>",
    "set max allowed parallel peer connections"
  )
  .action(startAction);

program
  .command("address")
  .description("print the public Unchained address of a config file")
  .argument("<config>", "config file in YAML format")
  .option("--generate", "generate a secret key")
  .option("--ci", "run in ci mode")
  .action(addressAction);

program
  .command("diagnose")
  .description("perform a system check and print the details")
  .argument("<config>", "config file in YAML format")
  .action(diagnoseAction);

const postgres = program
  .command("postgres")
  .description("commands to manage postgres instances for Unchained");

postgres
  .command("migrate")
  .description("runs the database migrations to this Unchained version")
  .argument("<config>", "config file in YAML format")
  .action(initDbAction);

postgres
  .command("generate")
  .description("generates the Postgres client for this Unchained version")
  .argument("<config>", "config file in YAML format")
  .action(generateDbAction);

postgres
  .command("revert")
  .description("revert a Postgres migration")
  .argument("<migration>", "migration name")
  .argument("<config>", "config file in YAML format")
  .action(revertDbAction);

program.parse();
