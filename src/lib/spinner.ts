import ora from "ora";
import { colorizer } from "./logger/format.js";
import { sockets } from "./constants.js";

import { Ora } from "ora";

interface OraExtended extends Ora {
  isEnabled: boolean;
}

export const makeSpinner = (message: string) => {
  const instance = ora(message);

  const formattedSpace = colorizer.colorize("level_info", " ");
  const formattedTimestamp = colorizer.colorize(
    "timestamp",
    ` ${new Date().toISOString()} `
  );
  const formattedLevel = colorizer.colorize("level_info", ` INFO ---> `);

  const peers = sockets.size.toString().padStart(4, "·");
  const formattedPeers = colorizer.colorize("level_info", `[${peers}] `);

  instance.prefixText =
    formattedSpace + formattedPeers + formattedTimestamp + formattedLevel;

  return instance as OraExtended;
};
