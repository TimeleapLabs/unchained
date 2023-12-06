import * as uniswap from "../plugins/uniswap/uniswap.js";
import { gossip } from "../gossip/index.js";
import { runAtNextInterval } from "../utils/time.js";
import { runWithRetries, CALLERROR, TIMEOUT } from "../utils/retry.js";
const uniswapArgs = [
    { blockchain: "ethereum", asset: "Ethereum", source: "uniswap" },
    "0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640",
    [18, 6],
    true,
];
export const run = () => {
    runAtNextInterval(async () => {
        try {
            const payload = await runWithRetries(uniswap.work, uniswapArgs);
            if (payload && payload !== CALLERROR && payload !== TIMEOUT) {
                const { request, signer, seen = [], signature, } = payload;
                await gossip({ request, signer, signature, seen });
            }
        }
        catch (error) {
            // Handle the error or log it
        }
    }, 5);
};
