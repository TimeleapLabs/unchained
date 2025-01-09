package worker

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/service/rpc/dto"
	"github.com/TimeleapLabs/unchained/internal/service/rpc/runtime"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/google/uuid"
	"github.com/puzpuzpuz/xsync/v3"
)

type Option func(s *Worker)

// meta is a struct that holds the information of a function.

type resourceUsage struct {
	CPU int
	GPU int
}

// Worker is a struct that holds the functions that the worker can run.
type Worker struct {
	Plugins      map[string]dto.Plugin
	CurrentTasks xsync.MapOf[uuid.UUID, resourceUsage]
	CPUUsage     int
	GPUUsage     int
	MaxCPU       int
	MaxGPU       int
	Overloaded   bool
}

// RunFunction runs a function with the given name and parameters.
func (w *Worker) RunFunction(ctx context.Context, pluginName string, params *dto.RPCRequest) error {
	// Check if plugin exists
	if _, ok := w.Plugins[pluginName]; !ok {
		utils.Logger.
			With("plugin", pluginName).
			Error("Plugin not found")
		return consts.ErrPluginNotFound
	}

	// Check if function exists
	if _, ok := w.Plugins[pluginName].Functions[params.Method]; !ok {
		utils.Logger.
			With("plugin", pluginName).
			With("function", params.Method).
			Error("Function not found")
		return consts.ErrFunctionNotFound
	}

	method := w.Plugins[pluginName].Functions[params.Method]

	// Make sure we're not overloading the worker
	if w.Overloaded || w.CPUUsage+method.CPU > w.MaxCPU || w.GPUUsage+method.GPU > w.MaxGPU {
		utils.Logger.
			With("cpu", w.CPUUsage).
			With("gpu", w.GPUUsage).
			With("method", params.Method).
			Error("Overloaded")
		// TODO: We should notify the broker that we're overloaded so it can stop sending us requests
		return consts.ErrOverloaded
	}

	// Record CPU and GPU units
	w.CPUUsage += method.CPU
	w.GPUUsage += method.GPU

	// Record the current task to release the resources when the task is done
	w.CurrentTasks.Store(params.ID, resourceUsage{
		CPU: method.CPU,
		GPU: method.GPU,
	})

	switch w.Plugins[pluginName].Runtime {
	case WebSocket:
		err := runtime.RunWebSocketCall(ctx, w.Plugins[pluginName].Writer, params)
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

// RegisterWorker registers the functions with the broker.
func (w *Worker) RegisterWorker() {
	// Register the functions
	payload := dto.RegisterWorker{
		Plugins: make([]dto.Plugin, 0, len(w.Plugins)),
		CPU:     w.MaxCPU,
		GPU:     w.MaxGPU,
	}

	for _, p := range w.Plugins {
		payload.Plugins = append(payload.Plugins, p)
	}

	conn.Send(consts.OpCodeRegisterWorker, payload.Sia().Bytes())
}

// NewWorker creates a new worker.
func NewWorker(options ...Option) *Worker {
	worker := &Worker{
		Plugins:      make(map[string]dto.Plugin),
		CurrentTasks: *xsync.NewMapOf[uuid.UUID, resourceUsage](),
		MaxCPU:       config.App.RPC.CPUs,
		MaxGPU:       config.App.RPC.GPUs,
	}

	for _, o := range options {
		o(worker)
	}

	return worker
}
