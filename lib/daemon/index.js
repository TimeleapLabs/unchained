import * as uniswap from "../plugins/uniswap/uniswap.js";
import { gossip } from "../gossip/index.js";

const store = {};

// TODO: this needs to be improved
export const interval = setInterval(async () => {
  try {
    const ethPrice = await uniswap.work(
      store.config,
      "0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640",
      [18, 6],
      true
    );

    if (ethPrice) {
      await gossip({
        method: "attest",
        args: ["ethPrice", ethPrice],
      });
    }
  } catch (error) {}
}, 5000);

export const setup = (config) => {
  store.config = config;
};
