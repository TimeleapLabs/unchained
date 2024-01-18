// TODO: Add retries and proper error handling
export const debounce = (fn: Function, wait: number) => {
  const timeouts = new Map<any, NodeJS.Timeout>();

  return ({ key, args }: { key: any; args: any[] }) => {
    // Clear existing timeout
    clearTimeout(timeouts.get(key));

    // Set a new timeout
    timeouts.set(
      key,
      setTimeout(() => {
        timeouts.delete(key);
        fn.apply(null, args);
      }, wait)
    );
  };
};

export const debounceAsync = <T>(
  fn: (...args: any) => Promise<T>,
  wait: number
) => {
  const timeouts = new Map<any, NodeJS.Timeout>();
  const executions = new Map<any, Promise<T>>();

  return async ({ key, args }: { key: any; args: any[] }) => {
    // Clear existing timeout
    clearTimeout(timeouts.get(key));

    // Set a new timeout
    return new Promise<T>((resolve, reject) => {
      timeouts.set(
        key,
        setTimeout(async () => {
          try {
            timeouts.delete(key);

            // Wait for the previous execution to finish
            await executions.get(key)?.catch(() => null);

            // Execute the function
            const execution = fn(...args);
            executions.set(key, execution);

            // Resolve or reject based on execution
            const result = await execution;
            return resolve(result);
          } catch (error) {
            reject(error);
          } finally {
            executions.delete(key);
          }
        }, wait)
      );
    });
  };
};
