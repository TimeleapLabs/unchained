const peerScoreMap = new Map<string, number>();

export const addOnePoint = (peer: string) => {
  const current = peerScoreMap.get(peer) || 0;
  peerScoreMap.set(peer, current + 1);
};

export const resetScore = (
  peer: string,
  map: Map<string, number> = peerScoreMap
) => map.set(peer, 0);

export const resetAllScores = (
  map: Map<string, number> = peerScoreMap
): Map<string, number> => {
  const clone = new Map(map.entries());
  map.clear();
  return clone;
};

export const getScoreOf = (
  peer: string,
  map: Map<string, number> = peerScoreMap
): number => map.get(peer) || 0;

export const getAllScores = (map: Map<string, number> = peerScoreMap) =>
  map.entries();
