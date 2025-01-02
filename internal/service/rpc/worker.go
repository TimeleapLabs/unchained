package rpc

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/service/rpc/dto"
	"github.com/TimeleapLabs/unchained/internal/service/rpc/runtime"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
	"github.com/TimeleapLabs/unchained/internal/utils"

	"github.com/gorilla/websocket"
)

type Option func(s *Worker)

// meta is a struct that holds the information of a function.
type meta struct {
	runtime Runtime
	path    string
	conn    *websocket.Conn
}

// Worker is a struct that holds the functions that the worker can run.
type Worker struct {
	functions map[string]meta
}

// RunFunction runs a function with the given name and parameters.
func (w *Worker) RunFunction(ctx context.Context, name string, params *dto.RPCRequest) error {
	switch w.functions[name].runtime {
	case WebSocket:
		err := runtime.RunWebSocketCall(ctx, w.functions[name].conn, params)
		if err != nil {
			utils.Logger.With("err", err).Error("Failed to run function")
			return err
		}

		return nil
	case Mock:
		return runtime.RunMock(params.Sia().Bytes())
	}

	return consts.ErrInternalError
}

// registerFunction registers a function with the broker.
func (w *Worker) registerFunction(name string, runtime string) {
	payload := dto.RegisterFunction{Function: name, Runtime: runtime}
	conn.Send(consts.OpCodeRegisterRPCFunction, payload.Sia().Bytes())
}

// RegisterFunctions registers the functions with the broker.
func (w *Worker) RegisterFunctions() {
	// Register the functions
	for name, fun := range w.functions {
		w.registerFunction(name, string(fun.runtime))
	}
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
