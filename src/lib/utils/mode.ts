export const getMode = (arr: number[]): number => {
  const frequencyMap = new Map<number, number>();

  for (const number of arr) {
    frequencyMap.set(number, 1 + (frequencyMap.get(number) || 0));
  }

  return [...frequencyMap.entries()].reduce((a, b) => (a[0] > b[0] ? a : b))[1];
};
