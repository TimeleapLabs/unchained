package handler

import (
	"github.com/KenshiTech/unchained/internal/config"
	"github.com/KenshiTech/unchained/internal/consts"
	"github.com/KenshiTech/unchained/internal/service/correctness"
	"github.com/KenshiTech/unchained/internal/service/evmlog"
	"github.com/KenshiTech/unchained/internal/service/uniswap"
	"github.com/KenshiTech/unchained/internal/transport/client/conn"
)

type consumer struct {
	correctness correctness.Service
	uniswap     uniswap.Service
	evmlog      evmlog.Service
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
