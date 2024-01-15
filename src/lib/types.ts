import { Duplex } from "stream";
import { SecretKey, PublicKey } from "@chainsafe/bls/types";

export interface State {
  connected: boolean;
}

export interface KeyPair {
  publicKey: PublicKey;
  secretKey: SecretKey;
}

interface RPCList {
  ethereum: string | string[];
}

interface DatabaseConfig {
  url: string;
}

interface PeerConfig {
  max: number;
  parallel: number;
}

interface JailConfig {
  duration: number;
  strikes: number;
}

export interface Config {
  name: string;
  log: string;
  rpc: RPCList;
  lite: boolean;
  database?: DatabaseConfig;
  secretKey: string;
  publicKey: string;
  peers: PeerConfig;
  jail: JailConfig;
  waves: number;
}

export interface MetaData {
  socket: Duplex;
  peer: string;
  peerAddr: string;
  murmurAddr?: string;
  name: string;
  publicKey?: string;
  needsDrain?: boolean;
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

export interface ObjectType<V> {
  [key: string]: V;
}

export type StringAnyObject = ObjectType<any>;

export interface AssetPriceMetric {
  block: number;
}

export interface AssetPriceValue {
  price: number;
}

export interface SignatureInput<MT, VT> {
  metric: MT;
  value: VT;
}

export interface WaveRequest<T, V> {
  method: string;
  dataset: string;
  metric: T;
  signature: string;
  signer: string;
  payload?: SignatureInput<T, V>;
}

export interface SignatureItem {
  signer: string;
  signature: string;
}

export interface PeerInfo {
  publicKey: Buffer;
  priority: number;
  ban(flag: boolean): void;
}

export interface Murmur {
  address: string;
}
