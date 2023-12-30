// TODO: Add retries and proper error handling
export const debounce = (fn: Function, wait: number) => {
  const timeouts = new Map<any, NodeJS.Timeout>();
  const executions = new Map<any, Promise<void>>();

  return ({ key, args }: { key: any; args: any[] }) => {
    clearTimeout(timeouts.get(key));
    timeouts.set(
      key,
      setTimeout(async () => {
        timeouts.delete(key);
        if (executions.has(key)) {
          await executions.get(key);
        }
        executions.set(
          key,
          new Promise(async (resolve) => {
            try {
              await fn.apply(null, args);
            } catch (error) {}
            executions.delete(key);
            resolve();
          })
        );
      }, wait)
    );
  };
};
