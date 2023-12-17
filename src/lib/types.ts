import { Duplex } from "stream";
import { SecretKey, PublicKey } from "@chainsafe/bls/types";

export interface State {
  connected: boolean;
}

export interface Keys {
  publicKey?: PublicKey;
  secretKey?: SecretKey;
}

interface RPCList {
  ethereum: string;
}

interface DatabaseConfig {
  url: string;
  name: string;
}

export interface Config {
  name: string;
  log: string;
  rpc: RPCList;
  lite: boolean;
  database?: DatabaseConfig;
  secretKey: string;
}

export interface MetaData {
  socket: Duplex;
  peer: string;
  peerAddr: string;
  name: string;
  publicKey?: string;
}

export interface NodeSystemError extends Error {
  address?: string;
  code: string;
  dest: string;
  errno: number;
  info?: Object;
  message: string;
  path?: string;
  port?: number;
  syscall: string;
}

export interface StringAnyObject {
  [key: string]: any;
}

export interface AssetPriceMetric {
  block: number;
}

export interface AssetPriceValue {
  price: number;
}

export interface GossipSignatureInput<MT, VT> {
  metric: MT;
  value: VT;
}

export interface GossipRequest<T> {
  method: string;
  dataset: string;
  metric: T;
  signature: string;
}

export interface Gossip<T> {
  type: "gossip";
  request: GossipRequest<T>;
  origin: string;
  seen: string[];
}
