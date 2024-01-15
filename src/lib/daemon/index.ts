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
import { queryNetworkFor, stats } from "../network/index.js";
import { toMurmur } from "../crypto/murmur/index.js";
import { hashObject } from "../utils/hash.js";
import { cache } from "../utils/cache.js";
import { epoch, minutes, seconds } from "../utils/time.js";

interface Cache {
  want: string;
  dataset: string;
  calls: number;
  created: number;
}

let wantCache: Cache[] = [];

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
        const want = await toMurmur(hashObject(result.metric));
        queryNetworkFor(want, result.dataset, [murmur.address]);
        wantCache.push({
          want,
          dataset: result.dataset,
          calls: 0,
          created: epoch(),
        });
        //await gossip(result, [murmur.address]);
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

  Cron("0 */1 * * * *", async () => {
    console.log(stats);
    stats.have = 0;
    stats.want = 0;
  });

  Cron(
    "*/1 * * * * *",
    async () => {
      try {
        wantCache = wantCache.filter((item) => item.calls <= 7);
        const now = epoch();
        for (const item of wantCache.toReversed()) {
          if (now - item.created >= seconds(item.calls ** 2)) {
            item.calls++;
            const have = await uniswap.getHave(item.want);
            queryNetworkFor(item.want, item.dataset, have);
          }
        }
      } catch (error) {
        // Handle the error or log it
      }
    },
    { protect: true }
  );
};
