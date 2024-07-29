package handler

import (
	"context"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/service/rpc"
	"github.com/TimeleapLabs/unchained/internal/service/rpc/dto"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/gorilla/websocket"
)

// unchainedRpc is a global variable that holds the rpc coordinator
var unchainedRpc = rpc.NewCoordinator()

// RegisterRpcFunction is a handler of network that registers a new worker
func RegisterRpcFunction(_ context.Context, conn *websocket.Conn, payload []byte) {
	request := new(dto.RegisterFunction).
		FromSiaBytes(payload[1:])

	utils.Logger.
		With("IP", conn.RemoteAddr().String()).
		With("Function", request.Function).
		Info("New Worker registered")

	unchainedRpc.RegisterWorker(request.Function, conn)
}

// CallFunction is a handler of network that calls a registered function
func CallFunction(_ context.Context, conn *websocket.Conn, payload []byte) {
	request := new(dto.RpcRequest).
		FromSiaBytes(payload[1:])

	utils.Logger.
		With("IP", conn.RemoteAddr().String()).
		With("ID", request.ID).
		With("Function", request.Method).
		Info("RPC Request")

	unchainedRpc.RegisterTask(request.ID, conn)
	worker := unchainedRpc.GetRandomWorker(request.Method)

	if worker != nil {
		utils.Logger.
			With("IP", conn.RemoteAddr().String()).
			With("Function", request.Method).
			Info("RPC Request Sent to Worker")

		Send(worker, consts.OpCodeRpcRequest, payload[1:])
	}
}

// ResponseFunction is a handler of network that sends a response to requester
func ResponseFunction(_ context.Context, conn *websocket.Conn, payload []byte) {
	response := new(dto.RpcResponse).
		FromSiaBytes(payload[1:])

	task := unchainedRpc.GetTask(response.ID)
	if task != nil {
		utils.Logger.
			With("IP", conn.RemoteAddr().String()).
			With("ID", response.ID).
			Info("RPC Response")

		Send(task, consts.OpCodeRpcResponse, payload[1:])
	}
}
