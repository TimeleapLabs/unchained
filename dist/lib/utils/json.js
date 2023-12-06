export const parse = (payload) => {
    try {
        return JSON.parse(payload);
    }
    catch (error) {
        return error;
    }
};
