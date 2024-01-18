import { AsyncResource } from "node:async_hooks";
import { EventEmitter } from "node:events";
import { Worker } from "node:worker_threads";
import os from "node:os";

const kTaskInfo = Symbol("kTaskInfo");
const kWorkerFreedEvent = Symbol("kWorkerFreedEvent");

type AsyncResourceCallback = (this: null, ...args: any[]) => unknown;

class CustomWorker extends Worker {
  [kTaskInfo]: WorkerPoolTaskInfo | null = null;
}

class WorkerPoolTaskInfo extends AsyncResource {
  resolve: AsyncResourceCallback;
  reject: AsyncResourceCallback;

  constructor(resolve: AsyncResourceCallback, reject: AsyncResourceCallback) {
    super("WorkerPoolTaskInfo");
    this.resolve = resolve;
    this.reject = reject;
  }

  done(err: Error | null, result: any) {
    if (err) {
      this.runInAsyncScope(this.reject, null, result);
    } else {
      this.runInAsyncScope(this.resolve, null, result);
    }
    this.emitDestroy();
  }
}

interface Task {
  task: any;
  reject: AsyncResourceCallback;
  resolve: AsyncResourceCallback;
}

export default class WorkerPool extends EventEmitter {
  numThreads: number;
  tasks: Task[];
  workers: CustomWorker[];
  freeWorkers: CustomWorker[];

  constructor(numThreads: number) {
    super();
    this.numThreads = numThreads;
    this.workers = [];
    this.freeWorkers = [];
    this.tasks = [];

    for (let i = 0; i < numThreads; i++) this.addNewWorker();

    // Any time the kWorkerFreedEvent is emitted, dispatch
    // the next task pending in the queue, if any.
    this.on(kWorkerFreedEvent, () => {
      if (this.tasks.length > 0) {
        const { task, resolve, reject } = this.tasks.shift() as Task;
        this.retryTask(task, resolve, reject);
      }
    });
  }

  addNewWorker() {
    const worker = new CustomWorker(new URL("worker.js", import.meta.url));
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
      } else {
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

  runTask<T>(task: any): Promise<T> {
    if (this.freeWorkers.length === 0) {
      return new Promise((resolve, reject) => {
        this.tasks.push({ task, resolve, reject });
      });
    }

    return new Promise((resolve, reject) => {
      const worker = this.freeWorkers.pop() as CustomWorker;
      worker[kTaskInfo] = new WorkerPoolTaskInfo(resolve, reject);
      worker.postMessage(task);
    });
  }

  retryTask(
    task: any,
    resolve: AsyncResourceCallback,
    reject: AsyncResourceCallback
  ): void {
    if (this.freeWorkers.length === 0) {
      this.tasks.push({ task, resolve, reject });
    } else {
      const worker = this.freeWorkers.pop() as CustomWorker;
      worker[kTaskInfo] = new WorkerPoolTaskInfo(resolve, reject);
      worker.postMessage(task);
    }
  }

  close() {
    for (const worker of this.workers) worker.terminate();
  }
}

export const pool = new WorkerPool(os.availableParallelism());

export const verifyAggregate = async (
  signers: Uint8Array[],
  signature: Uint8Array,
  data: any
): Promise<boolean> => {
  return await pool.runTask<boolean>({
    method: "verifyAggregate",
    args: [signers, signature, data],
  });
};

export const aggregate = async (
  signatures: Uint8Array[]
): Promise<Uint8Array> => {
  return await pool.runTask<Uint8Array>({
    method: "aggregate",
    args: [signatures],
  });
};
