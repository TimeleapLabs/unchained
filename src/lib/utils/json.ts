export const parse = (payload: string) => {
  try {
    return JSON.parse(payload);
  } catch (error) {
    return error;
  }
};
