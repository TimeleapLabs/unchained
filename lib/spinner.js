import ora from "ora";
import { colorizer } from "./logger/format.js";

export const makeSpinner = (message) => {
  const instance = ora(message);

  const formattedSpace = colorizer.colorize("level_info", " ");
  const formattedTimestamp = colorizer.colorize(
    "timestamp",
    ` ${new Date().toISOString()} `
  );
  const formattedLevel = colorizer.colorize(
    "level_info",
    ` ${"INFO".padEnd(7, " ")} | `
  );

  instance.prefixText = formattedSpace + formattedTimestamp + formattedLevel;

  return instance;
};
