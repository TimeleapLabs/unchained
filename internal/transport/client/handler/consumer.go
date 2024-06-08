package handler

import (
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/service/correctness"
	"github.com/TimeleapLabs/unchained/internal/service/evmlog"
	"github.com/TimeleapLabs/unchained/internal/service/uniswap"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
	"github.com/TimeleapLabs/unchained/internal/transport/client/store"
)

type postgresConsumer struct {
	correctness correctness.Service
	uniswap     uniswap.Service
	evmlog      evmlog.Service
}

type schnorrConsumer struct {
	signerRepository store.SignerRepository
}

func NewSchnorrConsumerHandler(
	signerRepository store.SignerRepository,
) Handler {
	return &schnorrConsumer{
		signerRepository: signerRepository,
	}
}

func NewPostgresConsumerHandler(
	correctness correctness.Service,
	uniswap uniswap.Service,
	evmlog evmlog.Service,
) Handler {
	conn.Send(consts.OpCodeRegisterConsumer, []byte(config.App.Network.SubscribedChannel))

	return &postgresConsumer{
		correctness: correctness,
		uniswap:     uniswap,
		evmlog:      evmlog,
	}
}
