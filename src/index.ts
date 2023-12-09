#!/usr/bin/env node

import { program } from "commander";
import { startAction } from "./lib/cli/actions/start.js";
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
  .action(startAction);

program.parse();
