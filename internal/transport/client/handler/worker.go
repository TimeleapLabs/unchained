package handler

import "github.com/TimeleapLabs/unchained/internal/service/rpc"

type worker struct {
	rpc *rpc.Worker
}

func NewWorkerHandler(rpc *rpc.Worker) Handler {
	// Register the worker functions with the broker
	rpc.RegisterWorker()

	return &worker{
		rpc: rpc,
	}
}
