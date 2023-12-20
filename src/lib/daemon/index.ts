import * as uniswap from "../plugins/uniswap/uniswap.js";
import { gossip } from "../gossip/index.js";
import { runWithRetries } from "../utils/retry.js";
import { schedule } from "node-cron";
import { printScores } from "../score/print.js";

interface UniswapArgs {
  blockchain: string;
  asset: string;
  source: string;
}

const uniswapArgs: [UniswapArgs, string, [number, number], boolean] = [
  { blockchain: "ethereum", asset: "Ethereum", source: "uniswap" },
  "0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640",
  [18, 6],
  true,
];

export const runTasks = (): void => {
  schedule("*/5 * * * * *", async () => {
    try {
      const result = await runWithRetries(uniswap.work, uniswapArgs);
      if (result && !(result instanceof Symbol)) {
        await gossip(result, []);
      }
    } catch (error) {
      // Handle the error or log it
    }
  }).start();

  schedule("0 */5 * * * *", async () => {
    try {
      printScores();
    } catch (error) {
      // Handle the error or log it
    }
  }).start();
};
