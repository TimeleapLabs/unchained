package handler

import (
	"github.com/TimeleapLabs/timeleap/internal/config"
	"github.com/TimeleapLabs/timeleap/internal/consts"
	"github.com/TimeleapLabs/timeleap/internal/service/attestation"
	"github.com/TimeleapLabs/timeleap/internal/transport/client/conn"
)

// consumer is a struct that holds the services required by the consumer handler.
type consumer struct {
	attestation attestation.Service
}

// NewConsumerHandler is a function that creates a new consumer handler.
func NewConsumerHandler(
	attestation attestation.Service,
) Handler {
	conn.SendSigned(
		consts.OpCodeRegisterConsumer,
		[]byte(config.App.Network.SubscribedChannel),
	)

	return &consumer{
		attestation: attestation,
	}
}
