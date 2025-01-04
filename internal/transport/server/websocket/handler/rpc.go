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

// RegisterRPCFunction is a handler of network that registers a new worker.
func RegisterRPCFunction(_ context.Context, conn *websocket.Conn, payload []byte) {
	request := new(dto.RegisterFunction).
		FromSiaBytes(payload)

	utils.Logger.
		With("IP", conn.RemoteAddr().String()).
		With("Plugin", request.Plugin).
		With("Functions", strings.Join(request.Functions, ",")).
		Info("New Worker registered")

	unchainedRPC.RegisterWorker(request.Plugin, conn)
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

	unchainedRPC.RegisterTask(request.ID, conn)
	worker := unchainedRPC.GetRandomWorker(request.Plugin)

	if worker != nil {
		utils.Logger.
			With("IP", conn.RemoteAddr().String()).
			With("Plugin", request.Plugin).
			With("Function", request.Method).
			Info("RPC Request Sent to Worker")

		Send(worker, consts.OpCodeRPCRequest, payload)
	}
}

// ResponseFunction is a handler of network that sends a response to requester.
func ResponseFunction(_ context.Context, conn *websocket.Conn, payload []byte) {
	response := new(dto.RPCResponse).
		FromSiaBytes(payload)

	task := unchainedRPC.GetTask(response.ID)
	if task != nil {
		utils.Logger.
			With("IP", conn.RemoteAddr().String()).
			With("ID", response.ID).
			Info("RPC Response")

		Send(task, consts.OpCodeRPCResponse, payload)
	} else {
		utils.Logger.
			With("IP", conn.RemoteAddr().String()).
			With("ID", response.ID).
			Error("Task not found")
	}
}
