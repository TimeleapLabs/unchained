import { parentPort } from "node:worker_threads";
import { verify } from "./index.js";
parentPort?.on("message", (task) => {
    const result = verify(task);
    parentPort?.postMessage(result);
});
