package handler

import (
	"context"
	"fmt"

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
		FromSiaBytes(payload[1:])

	utils.Logger.
		With("IP", conn.RemoteAddr().String()).
		With("Function", request.Function).
		Info("New Worker registered")

	unchainedRPC.RegisterWorker(request.Function, conn)
}

// CallFunction is a handler of network that calls a registered function.
func CallFunction(_ context.Context, conn *websocket.Conn, payload []byte) {
	request := new(dto.RPCRequest).
		FromSiaBytes(payload)

	utils.Logger.
		With("IP", conn.RemoteAddr().String()).
		With("ID", request.ID).
		With("Function", request.Method).
		Info("RPC Request")

	unchainedRPC.RegisterTask(request.ID, conn)
	worker := unchainedRPC.GetRandomWorker(request.Method)

	fmt.Println(worker)
	fmt.Println(request.Method)

	if worker != nil {
		utils.Logger.
			With("IP", conn.RemoteAddr().String()).
			With("Function", request.Method).
			Info("RPC Request Sent to Worker")

		Send(worker, consts.OpCodeRPCRequest, payload[1:])
	}
}

// ResponseFunction is a handler of network that sends a response to requester.
func ResponseFunction(_ context.Context, conn *websocket.Conn, payload []byte) {
	response := new(dto.RPCResponse).
		FromSiaBytes(payload[1:])

	task := unchainedRPC.GetTask(response.ID)
	if task != nil {
		utils.Logger.
			With("IP", conn.RemoteAddr().String()).
			With("ID", response.ID).
			Info("RPC Response")

		Send(task, consts.OpCodeRPCResponse, payload[1:])
	}
}
