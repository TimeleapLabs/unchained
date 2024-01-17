import { sia, desia } from "sializer";

export const serialize = (input: any) => sia(input);

export const parse = (payload: Buffer) => {
  try {
    return desia(payload);
  } catch (error) {
    return error;
  }
};
