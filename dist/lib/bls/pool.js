import { AsyncResource } from "node:async_hooks";
import { EventEmitter } from "node:events";
import { Worker } from "node:worker_threads";
import { availableParallelism } from "os";
import { config } from "../constants.js";
const kTaskInfo = Symbol("kTaskInfo");
const kWorkerFreedEvent = Symbol("kWorkerFreedEvent");
class WorkerPoolTaskInfo extends AsyncResource {
    resolve;
    reject;
    constructor(resolve, reject) {
        super("WorkerPoolTaskInfo");
        this.resolve = resolve;
        this.reject = reject;
    }
    done(err, result) {
        if (err) {
            this.runInAsyncScope(this.reject, null, err);
        }
        else {
            this.runInAsyncScope(this.resolve, null, result);
        }
        this.emitDestroy(); // `TaskInfo`s are used only once.
    }
}
class AsyncWorker extends Worker {
    [kTaskInfo];
}
export default class WorkerPool extends EventEmitter {
    numThreads;
    workers;
    freeWorkers;
    tasks; // FIXME
    constructor(numThreads) {
        super();
        this.numThreads = numThreads;
        this.workers = [];
        this.freeWorkers = [];
        this.tasks = [];
        for (let i = 0; i < numThreads; i++)
            this.addNewWorker();
        // Any time the kWorkerFreedEvent is emitted, dispatch
        // the next task pending in the queue, if any.
        this.on(kWorkerFreedEvent, () => {
            if (this.tasks.length > 0) {
                const { task, resolve, reject } = this.tasks.shift();
                this.runTask(task, resolve, reject);
            }
        });
    }
    addNewWorker() {
        const worker = new AsyncWorker(new URL("worker.js", import.meta.url));
        worker.on("message", (result) => {
            // In case of success: Call the callback that was passed to `runTask`,
            // remove the `TaskInfo` associated with the Worker, and mark it as free
            // again.
            worker[kTaskInfo]?.done(null, result);
            worker[kTaskInfo] = null;
            this.freeWorkers.push(worker);
            this.emit(kWorkerFreedEvent);
        });
        worker.on("error", (err) => {
            // In case of an uncaught exception: Call the callback that was passed to
            // `runTask` with the error.
            if (worker[kTaskInfo]) {
                worker[kTaskInfo].done(err, null);
            }
            else {
                this.emit("error", err);
            }
            // Remove the worker from the list and start a new Worker to replace the
            // current one.
            this.workers.splice(this.workers.indexOf(worker), 1);
            this.addNewWorker();
        });
        this.workers.push(worker);
        this.freeWorkers.push(worker);
        this.emit(kWorkerFreedEvent);
    }
    // FIXME
    runTask(task, resolve, reject) {
        const worker = this.freeWorkers.pop();
        if (worker && resolve && reject) {
            worker[kTaskInfo] = new WorkerPoolTaskInfo(resolve, reject);
            worker.postMessage(task);
            return;
        }
        if (!worker && resolve && reject) {
            this.tasks.push({ task, resolve, reject });
            return;
        }
        return new Promise((resolve, reject) => {
            if (!worker) {
                // No free threads, wait until a worker thread becomes free.
                this.tasks.push({ task, resolve, reject });
                return;
            }
            else {
                worker[kTaskInfo] = new WorkerPoolTaskInfo(resolve, reject);
                worker.postMessage(task);
            }
        });
    }
    close() {
        for (const worker of this.workers)
            worker.terminate();
    }
}
let pool;
export const verify = (task) => pool.runTask(task);
export const initSignatureWorkers = () => {
    pool = new WorkerPool(config.parallelism || availableParallelism());
};
