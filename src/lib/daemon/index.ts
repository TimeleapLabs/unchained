import * as uniswap from "../plugins/uniswap/uniswap.js";
import * as score from "../score/index.js";
import { runWithRetries } from "../utils/retry.js";
import { Cron } from "croner";
import { printScores } from "../score/print.js";
import { config } from "../constants.js";
import { queryNetworkFor } from "../network/index.js";
import { toMurmurCached } from "../crypto/murmur/index.js";
import { hashObject } from "../utils/hash.js";
import { epoch, seconds } from "../utils/time.js";

interface Cache {
  want: string;
  dataset: string;
  calls: number;
  created: number;
  getHave: (want: string) => Promise<any>;
}

let waveCache: Cache[] = [];

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

// TODO: We need to expose this as an API
export const runTasks = (): void => {
  Cron("*/5 * * * * *", async () => {
    try {
      const result = await runWithRetries(uniswap.work, uniswapArgs);
      if (result && !(result instanceof Symbol)) {
        const want = await toMurmurCached(hashObject(result.metric));
        await queryNetworkFor(want, result.dataset, uniswap.getHave);
        const created = epoch();
        const { dataset } = result;
        const args: Cache = {
          want,
          dataset,
          calls: 0,
          created,
          getHave: uniswap.getHave,
        };
        waveCache.push(args);
      }
    } catch (error) {
      // Handle the error or log it
    }
  });

  Cron("0 */5 * * * *", async () => {
    try {
      const scores = score.resetAllScores();
      printScores(scores).catch(() => null);
      const result = await score.getScoresPayload(scores);
      const want = await toMurmurCached(hashObject(result.metric));
      await queryNetworkFor(want, result.dataset, score.getHave);
      const created = epoch();
      const { dataset } = result;
      const args: Cache = {
        want,
        dataset,
        calls: 0,
        created,
        getHave: score.getHave,
      };
      waveCache.push(args);
      // TODO: We need retries here
      if (!config.lite) {
        await score.storeSprintScores();
      }
    } catch (error) {
      console.trace(error);
      // Handle the error or log it
    }
  });

  Cron("*/1 * * * * *", async () => {
    try {
      waveCache = waveCache.filter((item) => item.calls <= config.waves.count);
      const now = epoch();
      for (const item of waveCache.toReversed()) {
        if (now - item.created >= seconds(item.calls ** 2)) {
          item.calls++;
          await queryNetworkFor(item.want, item.dataset, item.getHave);
        }
      }
    } catch (error) {
      // Handle the error or log it
    }
  });
};
