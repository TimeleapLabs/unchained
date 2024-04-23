package handler

import (
	"github.com/TimeleapLabs/unchained/internal/constants/opcodes"
	"github.com/TimeleapLabs/unchained/internal/service/correctness"
	"github.com/TimeleapLabs/unchained/internal/service/evmlog"
	"github.com/TimeleapLabs/unchained/internal/service/uniswap"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
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
