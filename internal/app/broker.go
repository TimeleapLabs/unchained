package app

import (
	"github.com/KenshiTech/unchained/internal/constants"
	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/ethereum"
	"github.com/KenshiTech/unchained/internal/log"
	"github.com/KenshiTech/unchained/internal/pos"
	"github.com/KenshiTech/unchained/internal/transport/server"
	"github.com/KenshiTech/unchained/internal/transport/server/websocket"
)

// Broker starts the Unchained broker and contains its DI.
func Broker() {
	log.Logger.
		With("Version", constants.Version).
		With("Protocol", constants.ProtocolVersion).
		Info("Running Unchained | Broker")

	bls.InitClientIdentity()

	ethRPC := ethereum.New()
	pos.New(ethRPC)

	server.New(
		websocket.WithWebsocket(),
	)
}
