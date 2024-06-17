package handler

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/service/correctness"
	"github.com/TimeleapLabs/unchained/internal/service/evmlog"
	"github.com/TimeleapLabs/unchained/internal/service/frost"
	"github.com/TimeleapLabs/unchained/internal/service/uniswap"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
)

type Handler interface {
	Challenge(message []byte) []byte
	CorrectnessReport(ctx context.Context, message []byte)
	EventLog(ctx context.Context, message []byte)
	PriceReport(ctx context.Context, message []byte)

	ConfirmFrostHandshake(ctx context.Context, message []byte)
	InitFrostIdentity(ctx context.Context, message []byte)
	RequestToSign(ctx context.Context, message []byte)
}

type worker struct {
	frostService frost.Service
}

type consumer struct {
	correctness correctness.Service
	uniswap     uniswap.Service
	evmlog      evmlog.Service
}

func NewWorkerHandler(frostService frost.Service) Handler {
	conn.Send(consts.OpCodeRegisterConsumer, []byte(config.App.Network.SubscribedChannel))

	return &worker{
		frostService: frostService,
	}
}

func NewConsumerHandler(
	correctness correctness.Service,
	uniswap uniswap.Service,
	evmlog evmlog.Service,
) Handler {
	conn.Send(consts.OpCodeRegisterConsumer, []byte(config.App.Network.SubscribedChannel))

	return &consumer{
		correctness: correctness,
		uniswap:     uniswap,
		evmlog:      evmlog,
	}
}
