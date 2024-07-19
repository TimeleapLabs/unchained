package handler

import (
	"context"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/service/rpc"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/gorilla/websocket"
)

var unchainedRpc = rpc.New()

func RegisterRpcFunction(ctx context.Context, conn *websocket.Conn, payload []byte) {
	request := new(rpc.RegisterFunction).
		FromSiaBytes(payload[1:])

	utils.Logger.
		With("IP", conn.RemoteAddr().String()).
		With("Function", request.Function).
		Info("New Worker registered")

	unchainedRpc.RegisterWorker(request.Function, conn)
}

func CallFunction(ctx context.Context, conn *websocket.Conn, payload []byte) {
	request := new(rpc.TextToImageRpcRequest).
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

func ResponseFunction(ctx context.Context, conn *websocket.Conn, payload []byte) {
	response := new(rpc.TextToImageRpcResponse).
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
