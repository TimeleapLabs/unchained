import * as uniswap from "../plugins/uniswap/uniswap.js";
import { gossip } from "../gossip/index.js";
import { runAtNextInterval } from "../utils/time.js";
import { runWithRetries, CALLERROR, TIMEOUT } from "../utils/retry.js";
import { GossipRequest } from "../types.js";

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
  runAtNextInterval(async () => {
    try {
      const result = await runWithRetries(uniswap.work, uniswapArgs);
      if (result && !(result instanceof Symbol)) {
        await gossip(result, []);
      }
    } catch (error) {
      // Handle the error or log it
    }
  }, 5);
};
