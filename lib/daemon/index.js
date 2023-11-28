import * as uniswap from "../plugins/uniswap/uniswap.js";
import { gossip } from "../gossip/index.js";
import { runAtNextInterval } from "../utils/time.js";

export const run = () => {
  runAtNextInterval(async () => {
    try {
      const payload = await uniswap.work(
        {
          blockchain: "ethereum",
          asset: "Ethereum",
          source: "uniswap",
        },
        "0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640",
        [18, 6],
        true
      );

      if (payload) {
        await gossip(payload);
      }
    } catch (error) {}
  }, 5);
};
