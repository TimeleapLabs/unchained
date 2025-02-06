package handler

import (
	"context"
	"strings"

	"github.com/TimeleapLabs/timeleap/internal/consts"
	"github.com/TimeleapLabs/timeleap/internal/service/rpc"
	"github.com/TimeleapLabs/timeleap/internal/service/rpc/dto"
	"github.com/TimeleapLabs/timeleap/internal/transport/server/websocket/queue"
	"github.com/TimeleapLabs/timeleap/internal/utils"
	"github.com/gorilla/websocket"
)

// timeleapRPC is a global variable that holds the rpc coordinator.
var timeleapRPC = rpc.NewCoordinator()

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
		With("RAM", request.RAM).
		Info("New Worker registered")

	timeleapRPC.RegisterWorker(request, conn)
}

func UnregisterWorker(conn *websocket.Conn) {
	utils.Logger.
		With("IP", conn.RemoteAddr().String()).
		Info("Worker unregistered")

	timeleapRPC.UnregisterWorker(conn)
}

func WorkerOverload(_ context.Context, conn *websocket.Conn, payload []byte) {
	overload := new(dto.WorkerOverload).
		FromSiaBytes(payload)

	utils.Logger.
		With("IP", conn.RemoteAddr().String()).
		With("FailedTaskID", overload.FailedTaskID).
		With("CPU", overload.CPU).
		With("GPU", overload.GPU).
		With("RAM", overload.RAM).
		Error("Worker Overload")

	task, ok := timeleapRPC.GetTask(overload.FailedTaskID)
	if ok {
		task.Client.SendError(consts.OpCodeError, consts.ErrOverloaded)
		timeleapRPC.UnregisterTask(overload.FailedTaskID)
	}
}

// CallFunction is a handler of network that calls a registered function.
func CallFunction(_ context.Context, wsQueue *queue.WebSocketWriter, payload []byte) {
	request := new(dto.RPCRequest).
		FromSiaBytes(payload)

	utils.Logger.
		With("IP", wsQueue.Conn.RemoteAddr().String()).
		With("ID", request.ID).
		With("Plugin", request.Plugin).
		With("Function", request.Method).
		Info("RPC Request")

	worker, function := timeleapRPC.GetRandomWorker(request.Plugin, request.Method, request.Timeout)

	if worker != nil && function != nil {
		timeleapRPC.RegisterTask(
			request.ID,
			worker.Conn,
			wsQueue,
			function.CPU,
			function.GPU,
			function.RAM,
			request.Timeout,
		)

		utils.Logger.
			With("IP", wsQueue.Conn.RemoteAddr().String()).
			With("ID", request.ID).
			With("Worker", worker.Conn.RemoteAddr().String()).
			With("Plugin", request.Plugin).
			With("Function", request.Method).
			With("WorkerCPU", worker.CPUUsage).
			With("WorkerGPU", worker.GPUUsage).
			With("WorkerRAM", worker.RAMUsage).
			Info("RPC Request Sent to Worker")

		worker.Writer.SendSigned(consts.OpCodeRPCRequest, payload)
	} else {
		utils.Logger.
			With("IP", wsQueue.Conn.RemoteAddr().String()).
			With("ID", request.ID).
			With("Plugin", request.Plugin).
			With("Function", request.Method).
			Error("Worker not found")

		wsQueue.SendErrorSigned(consts.OpCodeError, consts.ErrNoWorker)
	}
}

// ResponseFunction is a handler of network that sends a response to requester.
func ResponseFunction(_ context.Context, wsQueue *queue.WebSocketWriter, payload []byte) {
	response := new(dto.RPCResponse).
		FromSiaBytes(payload)

	task, ok := timeleapRPC.GetTask(response.ID)
	if ok {
		utils.Logger.
			With("IP", wsQueue.Conn.RemoteAddr().String()).
			With("ID", response.ID).
			Info("RPC Response")

		task.Client.SendSigned(consts.OpCodeRPCResponse, payload)
		timeleapRPC.UnregisterTask(response.ID)
	} else {
		utils.Logger.
			With("IP", wsQueue.Conn.RemoteAddr().String()).
			With("ID", response.ID).
			Error("Task not found")
	}
}
