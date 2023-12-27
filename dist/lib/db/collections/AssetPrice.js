export let assetPrices;
export const initCollection = async (db) => {
    assetPrices = db.collection("assetPrices");
    await assetPrices.createIndex({ block: -1, asset: 1 });
};
