package handler

import (
	rpcWorker "github.com/TimeleapLabs/timeleap/internal/service/rpc/worker"
)

type worker struct {
	rpc *rpcWorker.Worker
}

func NewWorkerHandler(rpc *rpcWorker.Worker) Handler {
	// Register the worker functions with the broker
	rpc.RegisterWorker()

	return &worker{
		rpc: rpc,
	}
}
