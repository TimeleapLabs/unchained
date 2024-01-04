// hyperswarm.d.ts
declare module "hyperswarm" {
  import { EventEmitter } from "events";
  import { Duplex } from "stream";
  import { PeerInfo } from "../lib/types.ts";

  interface Discovery {
    flushed(): Promise<void>;
  }

  class HyperSwarm extends EventEmitter {
    constructor(opts?: { maxPeers?: number; maxParallel?: number });
    join(topic: Buffer | string): Discovery;
    on(
      event: "connection",
      listener: (socket: Duplex, info: PeerInfo) => void
    ): this;
    keyPair: {
      publicKey: Buffer;
    };
  }

  export default HyperSwarm;
}
