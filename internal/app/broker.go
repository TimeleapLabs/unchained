package app

import (
	"github.com/TimeleapLabs/unchained/internal/constants"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"
	"github.com/TimeleapLabs/unchained/internal/log"
	"github.com/TimeleapLabs/unchained/internal/pos"
	"github.com/TimeleapLabs/unchained/internal/transport/server"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket"
)

// Broker starts the Unchained broker and contains its DI.
func Broker() {
	log.Logger.
		With("Mode", "Broker").
		With("Version", constants.Version).
		With("Protocol", constants.ProtocolVersion).
		Info("Running Unchained")

	crypto.InitMachineIdentity(
		crypto.WithBlsIdentity(),
		crypto.WithEvmSigner(),
	)

	ethRPC := ethereum.New()
	pos.New(ethRPC)

	server.New(
		websocket.WithWebsocket(),
	)
}
