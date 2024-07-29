package rpc

import (
	"context"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/service/rpc/runtime"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

// Runtime is a type that holds the runtime of a function
type Runtime string

const (
	Mock    Runtime = "mock"
	Wasm    Runtime = "wasm"
	Python  Runtime = "python"
	Webhook Runtime = "webhook"
	Docker  Runtime = "docker"
)

// meta is a struct that holds the information of a function
type meta struct {
	runtime Runtime
	path    string
}

// Worker is a struct that holds the functions that the worker can run
type Worker struct {
	functions map[string]meta
}

// RunFunction runs a function with the given name and parameters
func (w *Worker) RunFunction(ctx context.Context, name string, params []byte) ([]byte, error) {
	switch w.functions[name].runtime {
	case Wasm:
		result, err := runtime.RunWasmFromFile(ctx, w.functions[name].path, params)
		if err != nil {
			utils.Logger.With("err", err).Error("Failed to run wasm")
			return nil, err
		}
		return result, nil
	case Python:
		result, err := runtime.RunPython(ctx, w.functions[name].path, params)
		if err != nil {
			utils.Logger.With("err", err).Error("Failed to run python")
			return nil, err
		}
		return result, nil
	case Webhook:
		result, err := runtime.RunWebhook(ctx, w.functions[name].path, params)
		if err != nil {
			utils.Logger.With("err", err).Error("Failed to run wasm")
			return nil, err
		}

		return result, nil
	case Mock:
		return runtime.RunMock(params)
	}

	return nil, consts.ErrInternalError
}

// NewWorker creates a new worker
func NewWorker(options ...func(s *Worker)) *Worker {
	worker := &Worker{
		functions: make(map[string]meta),
	}

	for _, o := range options {
		o(worker)
	}

	return worker
}
