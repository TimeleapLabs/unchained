import { PeerInfo } from "../types.js";
import { logger } from "../logger/index.js";
import { epoch } from "../utils/time.js";

const jail = new Map<string, number>();
const strikes = new Map<string, number>();

const JAIL_TIME = 15 * 60 * 1000; // 15 minutes
const STRIKES_TO_JAIL = 3;

const peerKey = (info: PeerInfo) => info.publicKey.toString("base64");

export const isJailed = (name: string, info: PeerInfo) => {
  const key = peerKey(info);
  const jailedAt = jail.get(key);
  if (!jailedAt) {
    return false;
  }
  const isFree = jailedAt > epoch() + JAIL_TIME;
  if (isFree) {
    jail.delete(key);
    logger.info(`Peer ${name} is freed from jail.`);
  }
  return isFree;
};

export const strike = (name: string, info: PeerInfo) => {
  const key = peerKey(info);
  const score = 1 + (strikes.get(key) || 0);
  if (score >= STRIKES_TO_JAIL) {
    strikes.delete(key);
    jail.set(key, epoch());
    logger.info(`Jailed peer ${name} for too many connection errors.`);
    return true;
  } else {
    strikes.set(key, score);
    return false;
  }
};
