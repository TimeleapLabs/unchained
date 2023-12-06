import winston from "winston";
import { customFormat } from "./format.js";
export const logger = winston.createLogger({
    level: "info",
    transports: [
        new winston.transports.Console({
            handleExceptions: true,
            format: customFormat(),
        }),
    ],
});
