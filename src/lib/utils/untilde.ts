import { homedir } from "os";

export const untilde = (path: string): string =>
  path.replace(/^~(?=$|\/|\\)/, homedir());
