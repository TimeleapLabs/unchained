import crypto from "node:crypto";

const randomIndex = (length: number): number => {
  if (length <= 0) {
    throw new Error("Array length must be greater than 0.");
  }

  let index: number, randomByte: number;

  do {
    randomByte = crypto.randomBytes(1)[0];
    index = randomByte % length;
  } while (randomByte - index >= 256 - (256 % length));

  return index;
};

export const randomDistinct = (length: number, count: number): number[] => {
  const set = new Set<number>();
  while (set.size < count) {
    const random = randomIndex(length);
    set.add(random);
  }
  return [...set];
};
