package handler

import (
	"github.com/KenshiTech/unchained/internal/constants/opcodes"
	"github.com/KenshiTech/unchained/internal/service/correctness"
	"github.com/KenshiTech/unchained/internal/service/evmlog"
	"github.com/KenshiTech/unchained/internal/service/uniswap"
	"github.com/KenshiTech/unchained/internal/transport/client/conn"
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
