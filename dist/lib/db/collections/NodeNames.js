export let nodeNames;
export const initCollection = async (db) => {
    nodeNames = db.collection("nodeNames");
    await nodeNames.createIndex({ address: 1, name: 1 });
};
