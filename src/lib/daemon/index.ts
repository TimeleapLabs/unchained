import * as uniswap from "../plugins/uniswap/uniswap.js";
import { gossip } from "../gossip/index.js";
import { runWithRetries } from "../utils/retry.js";
import { Cron } from "croner";
import {
  getScoresPayload,
  resetAllScores,
  storeSprintScores,
} from "../score/index.js";

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
        gossip(result, []);
      }
    } catch (error) {
      // Handle the error or log it
    }
  });

  Cron("0 */5 * * * *", () => {
    try {
      const scores = resetAllScores();
      const payload = getScoresPayload(scores);
      gossip(payload, []);
      // TODO: We need retries here
      storeSprintScores().catch(() => null);
    } catch (error) {
      // Handle the error or log it
    }
  });
};
