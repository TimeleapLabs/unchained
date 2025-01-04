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

type plugin struct {
	name      string
	conn      *websocket.Conn
	runtime   Runtime
	functions []string
}

// Worker is a struct that holds the functions that the worker can run.
type Worker struct {
	plugins map[string]plugin
}

// RunFunction runs a function with the given name and parameters.
func (w *Worker) RunFunction(ctx context.Context, plugin string, params *dto.RPCRequest) error {
	switch w.plugins[plugin].runtime {
	case WebSocket:
		err := runtime.RunWebSocketCall(ctx, w.plugins[plugin].conn, params)
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
func (w *Worker) registerFunctions(plugin string, functions []string, runtime string) {
	payload := dto.RegisterFunction{Plugin: plugin, Functions: functions, Runtime: runtime}
	conn.Send(consts.OpCodeRegisterRPCFunction, payload.Sia().Bytes())
}

// RegisterFunctions registers the functions with the broker.
func (w *Worker) RegisterFunctions() {
	// Register the functions
	for _, plugin := range w.plugins {
		w.registerFunctions(plugin.name, plugin.functions, string(plugin.runtime))
	}
}

// NewWorker creates a new worker.
func NewWorker(options ...Option) *Worker {
	worker := &Worker{
		plugins: make(map[string]plugin),
	}

	for _, o := range options {
		o(worker)
	}

	return worker
}
