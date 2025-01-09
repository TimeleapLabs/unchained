package rpc

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/service/rpc/dto"
	"github.com/TimeleapLabs/unchained/internal/service/rpc/runtime"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/google/uuid"

	"github.com/gorilla/websocket"
)

type Option func(s *Worker)

// meta is a struct that holds the information of a function.

type plugin struct {
	name      string
	conn      *websocket.Conn
	runtime   Runtime
	functions map[string]config.Function
}

type resourceUsage struct {
	CPU int
	GPU int
}

// Worker is a struct that holds the functions that the worker can run.
type Worker struct {
	plugins      map[string]plugin
	currentTasks map[uuid.UUID]resourceUsage
	cpuUsage     int
	gpuUsage     int
	overloaded   bool
}

// RunFunction runs a function with the given name and parameters.
func (w *Worker) RunFunction(ctx context.Context, pluginName string, params *dto.RPCRequest) error {
	// Check if plugin exists
	if _, ok := w.plugins[pluginName]; !ok {
		utils.Logger.With("plugin", pluginName).Error("Plugin not found")
		return consts.ErrPluginNotFound
	}

	// Check if function exists
	if _, ok := w.plugins[pluginName].functions[params.Method]; !ok {
		utils.Logger.With("plugin", pluginName).With("function", params.Method).Error("Function not found")
		return consts.ErrFunctionNotFound
	}

	method := w.plugins[pluginName].functions[params.Method]

	// Make sure we're not overloading the worker
	if w.overloaded || w.cpuUsage+method.CPU > config.App.RPC.CPUs || w.gpuUsage+method.GPU > config.App.RPC.GPUs {
		utils.Logger.With("cpu", w.cpuUsage).With("gpu", w.gpuUsage).With("method", params.Method).Error("Overloaded")
		// TODO: We should notify the broker that we're overloaded so it can stop sending us requests
		return consts.ErrOverloaded
	}

	// Record CPU and GPU units
	w.cpuUsage += method.CPU
	w.gpuUsage += method.GPU

	// Record the current task to release the resources when the task is done
	w.currentTasks[params.ID] = resourceUsage{
		CPU: method.CPU,
		GPU: method.GPU,
	}

	switch w.plugins[pluginName].runtime {
	case WebSocket:
		err := runtime.RunWebSocketCall(ctx, w.plugins[pluginName].conn, params)
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
		functionNames := []string{}
		for name := range plugin.functions {
			functionNames = append(functionNames, name)
		}
		w.registerFunctions(plugin.name, functionNames, string(plugin.runtime))
	}
}

// NewWorker creates a new worker.
func NewWorker(options ...Option) *Worker {
	worker := &Worker{
		plugins:      make(map[string]plugin),
		currentTasks: make(map[uuid.UUID]resourceUsage),
	}

	for _, o := range options {
		o(worker)
	}

	return worker
}
