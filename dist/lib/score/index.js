const peerScoreMap = new Map();
export const addOnePoint = (peer) => {
    const current = peerScoreMap.get(peer) || 0;
    peerScoreMap.set(peer, current + 1);
};
export const resetScore = (peer, map = peerScoreMap) => map.set(peer, 0);
export const resetAllScores = (map = peerScoreMap) => {
    const clone = new Map(map.entries());
    map.clear();
    return clone;
};
export const getScoreOf = (peer, map = peerScoreMap) => map.get(peer) || 0;
export const getAllScores = (map = peerScoreMap) => map.entries();
