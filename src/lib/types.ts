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

export interface Config {
  name: string;
  log: string;
  store: string;
  rpc: RPCList;
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
