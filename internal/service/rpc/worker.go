package rpc

import (
	"context"
	"net"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/service/rpc/dto"
	"github.com/TimeleapLabs/unchained/internal/service/rpc/runtime"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

type Option func(s *Worker)

// meta is a struct that holds the information of a function.
type meta struct {
	runtime Runtime
	path    string
	conn    net.Conn
}

// Worker is a struct that holds the functions that the worker can run.
type Worker struct {
	functions map[string]meta
}

// RunFunction runs a function with the given name and parameters.
func (w *Worker) RunFunction(ctx context.Context, name string, params *dto.RPCRequest) ([]byte, error) {
	switch w.functions[name].runtime {
	case Unix:
		result, err := runtime.RunUnixCall(ctx, w.functions[name].conn, params)
		if err != nil {
			utils.Logger.With("err", err).Error("Failed to run wasm")
			return nil, err
		}

		return result.Sia().Bytes(), nil
	case Mock:
		return runtime.RunMock(params.Sia().Bytes())
	}

	return nil, consts.ErrInternalError
}

// NewWorker creates a new worker.
func NewWorker(options ...Option) *Worker {
	worker := &Worker{
		functions: make(map[string]meta),
	}

	for _, o := range options {
		o(worker)
	}

	return worker
}
