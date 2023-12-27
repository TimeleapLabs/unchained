import { logger } from "../logger/index.js";
import { state } from "../constants.js";
// Poor man's NTP
export const runAtNextInterval = (callback, interval) => {
    const now = new Date();
    const wait = (interval - (now.getSeconds() % interval)) * 1000 - now.getMilliseconds();
    setTimeout(() => {
        runAtNextInterval(callback, interval);
        if (state.connected) {
            logger.debug("Running daemon's work");
        }
        callback();
    }, wait);
};
