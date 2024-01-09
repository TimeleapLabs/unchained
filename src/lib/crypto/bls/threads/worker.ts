import { parentPort } from "node:worker_threads";
import { verifyAggregate } from "../index.js";
import { aggregate } from "../index.js";
import { ObjectType } from "../../../types.js";

const methods: ObjectType<Function> = { aggregate, verifyAggregate };

if (parentPort) {
  parentPort.on("message", (task) => {
    if (methods.hasOwnProperty(task.method)) {
      return parentPort?.postMessage(methods[task.method](...task.args));
    }
    return parentPort?.postMessage(null);
  });
}
