package handler

import (
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/service/correctness"
	"github.com/TimeleapLabs/unchained/internal/service/evmlog"
	"github.com/TimeleapLabs/unchained/internal/service/uniswap"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
)

// consumer is a struct that holds the services required by the consumer handler.
type consumer struct {
	correctness correctness.Service
	uniswap     uniswap.Service
	evmlog      evmlog.Service
}

// NewConsumerHandler is a function that creates a new consumer handler.
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
