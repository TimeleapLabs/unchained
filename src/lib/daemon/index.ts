import * as uniswap from "../plugins/uniswap/uniswap.js";
import { gossip } from "../gossip/index.js";
import { runWithRetries } from "../utils/retry.js";
import { Cron } from "croner";
import {
  getScoresPayload,
  resetAllScores,
  storeSprintScores,
} from "../score/index.js";
import { printScores } from "../score/print.js";
import { murmur } from "../constants.js";

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
  Cron("*/5 * * * * *", async () => {
    try {
      const result = await runWithRetries(uniswap.work, uniswapArgs);
      if (result && !(result instanceof Symbol)) {
        await gossip(result, [murmur.address]);
      }
    } catch (error) {
      // Handle the error or log it
    }
  });

  Cron("0 */5 * * * *", async () => {
    try {
      const scores = resetAllScores();
      printScores(scores);
      const payload = getScoresPayload(scores);
      await gossip(payload, [murmur.address]);
      // TODO: We need retries here
      await storeSprintScores();
    } catch (error) {
      // Handle the error or log it
    }
  });
};
