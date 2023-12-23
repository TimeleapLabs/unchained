// TODO: Add retries and proper error handling
export const debounce = (fn: Function, wait: number) => {
  const timeouts = new Map<any, NodeJS.Timeout>();

  return ({ key, args }: { key: any; args: any[] }) => {
    clearTimeout(timeouts.get(key));
    timeouts.set(
      key,
      setTimeout(() => {
        timeouts.delete(key);
        fn.apply(null, args);
      }, wait)
    );
  };
};
