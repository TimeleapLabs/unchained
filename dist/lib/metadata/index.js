const query = (address, name) => ({
    updateOne: {
        filter: { address },
        update: { $set: { name }, $setOnInsert: { address } },
        upsert: true,
    },
});
export const syncNodeNames = async () => {
    // const bulkOps = [...sockets.entries()]
    //   .map(([_, { publicKey, name }]) =>
    //     publicKey ? query(publicKey, name) : null
    //   )
    //   .filter(Boolean);
    // const { publicKey } = encodeKeys(keys);
    // const { name } = config;
    // bulkOps.push(query(publicKey, name));
    // await nodeNames.bulkWrite(bulkOps as AnyBulkWriteOperation<NodeNames>[]);
};
