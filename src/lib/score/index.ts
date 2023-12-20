const peerScoreMap = new Map<string, number>();

export const addOnePoint = (peer: string) => {
  const current = peerScoreMap.get(peer) || 0;
  peerScoreMap.set(peer, current + 1);
};

export const resetScore = (peer: string) => peerScoreMap.set(peer, 0);
export const resetAllScores = () => peerScoreMap.clear();
export const getScoreOf = (peer: string): number => peerScoreMap.get(peer) || 0;
export const getAllScores = () => peerScoreMap.entries();
