// base-ex.d.ts
declare module "base-ex" {
  class Base58 {
    constructor(alphabet: string);
    encode(input: Buffer | Uint8Array): string;
    decode(input: string): Buffer;
  }

  export { Base58 };
}
