import { work } from "./uniswap.js";

export const plugin = {
  async uniSpot(...args) {
    return await getPrice(this.config, ...args);
  },
};
