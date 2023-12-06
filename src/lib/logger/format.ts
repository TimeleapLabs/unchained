import { inspect } from "util";
import { format } from "winston";
import { sockets } from "../constants.js";

export const colorizer = format.colorize();

colorizer.addColors({
  level_error: "inverse red",
  level_warn: "inverse yellow",
  level_info: "inverse green",
  level_http: "inverse green",
  level_verbose: "inverse cyan",
  level_debug: "inverse blue",
  level_silly: "inverse magenta",

  msg_error: "red",
  msg_warn: "yellow",
  msg_info: "green",
  msg_http: "green",
  msg_verbose: "cyan",
  msg_debug: "blue",
  msg_silly: "magenta",

  timestamp: "inverse",
});

export const customFormat = (): any => {
  return format.combine(
    format.splat(),
    format.timestamp(),
    format.printf(({ timestamp, level, message, ...rest }) => {
      const formattedSpace = colorizer.colorize(`level_${level}`, " ");
      const formattedTimestamp = colorizer.colorize(
        "timestamp",
        ` ${timestamp} `
      );
      const formattedLevel = colorizer.colorize(
        `level_${level}`,
        ` ${level
          .toUpperCase()
          .padEnd(7, "-")
          .replace(/([A-Z])-/, "$1 ")}-> `
      );

      const peers = sockets.size.toString().padStart(4, "·");
      const formattedPeers = colorizer.colorize(
        `level_${level}`,
        `[${peers}] `
      );

      let formattedMessage = "";
      if (typeof message === "string") {
        formattedMessage = colorizer.colorize(`msg_${level}`, " " + message);
      } else {
        formattedMessage = " " + inspect(message, false, null, true);
      }

      const restStr = JSON.stringify(rest);
      if (restStr !== "{}") formattedMessage += "\t" + restStr;

      return (
        formattedSpace +
        formattedPeers +
        formattedTimestamp +
        formattedLevel +
        formattedMessage
      );
    })
  );
};
