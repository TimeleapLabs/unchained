import { MetaData } from "../types.js";
import { logger } from "../logger/index.js";
import { epoch } from "../utils/time.js";
import { config } from "../constants.js";

const jail = new Map<string, number>();
const strikes = new Map<string, number>();

export const isJailed = (meta: MetaData) => {
  const jailedAt = jail.get(meta.peerAddr);
  if (!jailedAt) {
    return false;
  }
  const isFree = jailedAt > epoch() + config.jail.duration;
  if (isFree) {
    jail.delete(meta.peerAddr);
    logger.debug(`Peer ${meta.name} is freed from jail.`);
  }
  return isFree;
};

export const strike = (meta: MetaData) => {
  const score = 1 + (strikes.get(meta.peerAddr) || 0);
  if (score >= config.jail.strikes) {
    strikes.delete(meta.peerAddr);
    jail.set(meta.peerAddr, epoch());
    logger.debug(`Jailed peer ${meta.name} for too many connection errors.`);
    return true;
  } else {
    strikes.set(meta.peerAddr, score);
    return false;
  }
};
