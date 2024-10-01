package handler

import (
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/service/correctness"
	"github.com/TimeleapLabs/unchained/internal/service/rpc"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
)

type handler struct {
	rpc         *rpc.Worker
	correctness correctness.Service
}

func NewHandler(
	rpc *rpc.Worker,
	correctness correctness.Service,
) Handler {
	rpc.RegisterFunctions()

	conn.Send(consts.OpCodeRegisterConsumer, []byte(config.App.Network.SubscribedChannel))

	return &handler{
		rpc:         rpc,
		correctness: correctness,
	}
}
