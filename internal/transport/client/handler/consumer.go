package handler

import (
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/service/attestation"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
)

// consumer is a struct that holds the services required by the consumer handler.
type consumer struct {
	attestation attestation.Service
}

// NewConsumerHandler is a function that creates a new consumer handler.
func NewConsumerHandler(
	attestation attestation.Service,
) Handler {
	conn.Send(consts.OpCodeRegisterConsumer, []byte(config.App.Network.SubscribedChannel))

	return &consumer{
		attestation: attestation,
	}
}
