// hyperswarm.d.ts
declare module "hyperswarm" {
  import { EventEmitter } from "events";
  import { Duplex } from "stream";

  interface ConnectionInfo {
    publicKey: Buffer;
  }

  interface Discovery {
    flushed(): Promise<void>;
  }

  class HyperSwarm extends EventEmitter {
    constructor(opts?: { maxPeers?: number; maxParallel?: number });
    join(topic: Buffer | string): Discovery;
    on(
      event: "connection",
      listener: (socket: Duplex, info: ConnectionInfo) => void
    ): this;
    keyPair: {
      publicKey: Buffer;
    };
  }

  export default HyperSwarm;
}
