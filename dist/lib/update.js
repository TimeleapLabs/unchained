import latestVersion from "latest-version";
import { version } from "./constants.js";
import { logger } from "./logger/index.js";
import semver from "semver";
export const checkForUpdates = async () => {
    const latestUnchained = await latestVersion("@kenshi.io/unchained");
    if (semver.gt(latestUnchained, version)) {
        const sudo = process.platform === "win32" ? "" : "sudo ";
        logger.warn("Update available! To update, run the following command:");
        logger.warn(`--> ${sudo}npm i -g @kenshi.io/unchained`);
    }
};
