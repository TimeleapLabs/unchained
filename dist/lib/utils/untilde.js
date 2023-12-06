import { homedir } from "os";
export const untilde = (path) => path.replace(/^~(?=$|\/|\\)/, homedir());
