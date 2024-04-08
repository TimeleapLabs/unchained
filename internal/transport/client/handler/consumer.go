package handler

import (
	"github.com/KenshiTech/unchained/constants/opcodes"
	"github.com/KenshiTech/unchained/service/correctness"
	"github.com/KenshiTech/unchained/service/evmlog"
	"github.com/KenshiTech/unchained/service/uniswap"
	"github.com/KenshiTech/unchained/transport/client/conn"
)

type consumer struct {
	correctness *correctness.Service
	uniswap     *uniswap.Service
	evmlog      *evmlog.Service
}

func NewConsumerHandler(
	correctness *correctness.Service,
	uniswap *uniswap.Service,
	evmlog *evmlog.Service,
) Handler {
	conn.Send(opcodes.RegisterConsumer, nil)

	return &consumer{
		correctness: correctness,
		uniswap:     uniswap,
		evmlog:      evmlog,
	}
}
