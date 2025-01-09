package handler

import (
	"context"
	"strings"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/service/rpc"
	"github.com/TimeleapLabs/unchained/internal/service/rpc/dto"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/gorilla/websocket"
)

// unchainedRPC is a global variable that holds the rpc coordinator.
var unchainedRPC = rpc.NewCoordinator()

// RegisterWorker is a handler of network that registers a new worker.
func RegisterWorker(_ context.Context, conn *websocket.Conn, payload []byte) {
	request := new(dto.RegisterWorker).
		FromSiaBytes(payload)

	pluginNames := make([]string, 0, len(request.Plugins))
	for _, plugin := range request.Plugins {
		pluginNames = append(pluginNames, plugin.Name)
	}

	plugins := strings.Join(pluginNames, ", ")

	utils.Logger.
		With("IP", conn.RemoteAddr().String()).
		With("Plugins", plugins).
		With("CPU", request.CPU).
		With("GPU", request.GPU).
		Info("New Worker registered")

	unchainedRPC.RegisterWorker(request, conn)
}

// CallFunction is a handler of network that calls a registered function.
func CallFunction(_ context.Context, conn *websocket.Conn, payload []byte) {
	request := new(dto.RPCRequest).
		FromSiaBytes(payload)

	utils.Logger.
		With("IP", conn.RemoteAddr().String()).
		With("ID", request.ID).
		With("Plugin", request.Plugin).
		With("Function", request.Method).
		Info("RPC Request")

	worker, function := unchainedRPC.GetRandomWorker(request.Plugin, request.Method)

	if worker != nil && function != nil {
		unchainedRPC.RegisterTask(request.ID, worker.Conn, conn, function.CPU, function.GPU)

		utils.Logger.
			With("IP", conn.RemoteAddr().String()).
			With("ID", request.ID).
			With("Worker", worker.Conn.RemoteAddr().String()).
			With("Plugin", request.Plugin).
			With("Function", request.Method).
			With("WorkerCPU", worker.CPUUsage).
			With("WorkerGPU", worker.GPUUsage).
			Info("RPC Request Sent to Worker")

		Send(worker.Conn, consts.OpCodeRPCRequest, payload)
	} // TODO: Handle if worker is nil
}

// ResponseFunction is a handler of network that sends a response to requester.
func ResponseFunction(_ context.Context, conn *websocket.Conn, payload []byte) {
	response := new(dto.RPCResponse).
		FromSiaBytes(payload)

	task, ok := unchainedRPC.GetTask(response.ID)
	if ok {
		utils.Logger.
			With("IP", conn.RemoteAddr().String()).
			With("ID", response.ID).
			Info("RPC Response")

		Send(task.Client, consts.OpCodeRPCResponse, payload)
		unchainedRPC.UnregisterTask(response.ID)
	} else {
		utils.Logger.
			With("IP", conn.RemoteAddr().String()).
			With("ID", response.ID).
			Error("Task not found")
	}
}
