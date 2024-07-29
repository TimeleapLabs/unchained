package handler

import "github.com/TimeleapLabs/unchained/internal/service/rpc"

type worker struct {
	rpc *rpc.Worker
}

func NewWorkerHandler(rpc *rpc.Worker) Handler {
	return &worker{
		rpc: rpc,
	}
}
