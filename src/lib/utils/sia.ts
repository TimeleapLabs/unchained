import { Sia, DeSia, constructors } from "sializer";
import type { ConstructorFactory } from "sializer";

const uint8: ConstructorFactory<Uint8Array, number[]> = {
  constructor: Uint8Array,
  code: 2,
  args: (item: Uint8Array) => [...item],
  build(...members: number[]) {
    return new Uint8Array(members);
  },
};

const UnchainedConstructors = [...constructors, uint8];

const sia = new Sia({ constructors: UnchainedConstructors });
const desia = new DeSia({ constructors: UnchainedConstructors });

export const serialize = (input: any) => {
  return sia.serialize(input);
};

export const parse = (payload: Buffer) => {
  try {
    return desia.deserialize(payload);
  } catch (error) {
    return error;
  }
};
