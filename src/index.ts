#!/usr/bin/env node

import { program } from "commander";
import { startAction } from "./lib/cli/actions/start.js";
import { addressAction } from "./lib/cli/actions/address.js";
import { version } from "./lib/constants.js";

program
  .name("unchained")
  .description("Kenshi Unchained CLI Node")
  .version(version);

program
  .command("start")
  .description("start a node on the Kenshi Unchained HyperSwarm")
  .argument("<config>", "config file in YAML format")
  .option("--log <level>", "log level")
  .option("--lite", "run in lite mode")
  .option("--generate", "generate a secret key")
  .option("--gossip <size>", "set gossip bucket size")
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

program.parse();
