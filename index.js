#!/usr/bin/env node

import { program } from "commander";
import { startAction } from "./lib/cli/actions/start.js";

program
  .name("unchained")
  .description("Kenshi Unchained CLI Node")
  .version("0.1.0");

program
  .command("start")
  .description("start a node on the Kenshi Unchained HyperSwarm")
  .argument("<config>", "config file in YAML format")
  .option("--log <level>", "log level")
  .option("--store <path>", "path to persistent storage")
  .option("--ask", "ask for private key")
  .action(startAction);

program.parse();
