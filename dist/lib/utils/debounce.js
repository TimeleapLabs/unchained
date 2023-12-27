// TODO: Add retries and proper error handling
export const debounce = (fn, wait) => {
    const timeouts = new Map();
    return ({ key, args }) => {
        clearTimeout(timeouts.get(key));
        timeouts.set(key, setTimeout(() => {
            timeouts.delete(key);
            fn.apply(null, args);
        }, wait));
    };
};
