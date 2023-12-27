import * as uniswap from "../plugins/uniswap/uniswap.js";
import { gossip } from "../gossip/index.js";
import { runWithRetries } from "../utils/retry.js";
import { printScores } from "../score/print.js";
import { syncNodeNames } from "../metadata/index.js";
import { Cron } from "croner";
import { config } from "../constants.js";
const uniswapArgs = [
    { blockchain: "ethereum", asset: "Ethereum", source: "uniswap" },
    "0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640",
    [18, 6],
    true,
];
export const runTasks = () => {
    Cron("*/5 * * * * *", async () => {
        try {
            const result = await runWithRetries(uniswap.work, uniswapArgs);
            if (result && !(result instanceof Symbol)) {
                await gossip(result, []);
            }
        }
        catch (error) {
            // Handle the error or log it
        }
    });
    Cron("0 */5 * * * *", async () => {
        try {
            printScores();
        }
        catch (error) {
            // Handle the error or log it
        }
    });
    if (!config.lite) {
        Cron("0 */1 * * * *", async () => {
            try {
                syncNodeNames();
            }
            catch (error) {
                // Handle the error or log it
            }
        });
    }
};