package rpc

import (
	"context"
	"github.com/TimeleapLabs/unchained/internal/service/rpc/runtime"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

type Runtime string

const (
	Wasm    Runtime = "wasm"
	Python  Runtime = "python"
	Webhook Runtime = "webhook"
)

type meta struct {
	runtime Runtime
	path    string
}

type Worker struct {
	functions map[string]meta
}

func (w *Worker) WithFunction(name string, runtime Runtime, path string) {
	w.functions[name] = meta{
		runtime: runtime,
		path:    path,
	}
}

func (w *Worker) RunFunction(ctx context.Context, name string, params interface{}) ([]byte, error) {
	switch w.functions[name].runtime {
	case Wasm:
		result, err := runtime.RunWasm(ctx, w.functions[name].path)
		if err != nil {
			utils.Logger.With("err", err).Error("Failed to run wasm")
			return nil, err
		}
		return result, nil
	case Python:

	case Webhook:
		result, err := runtime.RunWebhook(ctx, w.functions[name].path, params.([]byte))
		if err != nil {
			utils.Logger.With("err", err).Error("Failed to run wasm")
			return nil, err
		}

		return result, nil
	}
}

func NewWorker() *Worker {
	return &Worker{
		functions: make(map[string]meta),
	}
}
