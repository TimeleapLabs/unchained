import * as uniswap from "../plugins/uniswap/uniswap.js";
import * as score from "../score/index.js";
import { runWithRetries } from "../utils/retry.js";
import { Cron } from "croner";
import { printScores } from "../score/print.js";
import { config } from "../constants.js";
import { queryNetworkFor } from "../network/index.js";
import { toMurmurCached } from "../crypto/murmur/index.js";
import { hashObject } from "../utils/hash.js";
import { epoch, jitter, seconds } from "../utils/time.js";

interface Task {
  running: boolean;
  want: string;
  dataset: string;
  calls: number;
  lastQuery: number;
  getHave: (want: string) => Promise<any>;
}

let waveCache: Task[] = [];

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
        const lastQuery = epoch();
        const { dataset } = result;
        const args: Task = {
          running: false,
          want,
          dataset,
          calls: 0,
          lastQuery,
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
      const lastQuery = epoch();
      const { dataset } = result;
      const args: Task = {
        running: false,
        want,
        dataset,
        calls: 0,
        lastQuery,
        getHave: score.getHave,
      };
      waveCache.push(args);
      // TODO: We need retries here
      if (!config.lite) {
        await score.storeSprintScores();
      }
    } catch (error) {
      // Handle the error or log it
    }
  });

  Cron("*/1 * * * * *", async () => {
    try {
      waveCache = waveCache.filter((item) => item.calls <= config.waves.count);
      const now = epoch();
      for (const item of waveCache.toReversed()) {
        if (item.running) {
          continue;
        }

        const nextCall = item.lastQuery + seconds(15 + 7.5 * item.calls);

        if (now <= nextCall) {
          continue;
        }

        item.running = true;
        item.lastQuery = epoch();

        const success = await queryNetworkFor(
          item.want,
          item.dataset,
          item.getHave
        );

        if (success) {
          item.calls++;
        }

        item.running = false;
      }
    } catch (error) {
      // Handle the error or log it
    }
  });
};
