package app

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"
	"github.com/TimeleapLabs/unchained/internal/service/pos"
	"github.com/TimeleapLabs/unchained/internal/transport/server"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

// Broker starts the Unchained broker and contains its DI.
func Broker() {
	utils.Logger.
		With("Mode", "Broker").
		With("Version", consts.Version).
		With("Protocol", consts.ProtocolVersion).
		Info("Running Unchained")

	crypto.InitMachineIdentity(
		crypto.WithEd25519Identity(),
		crypto.WithEvmSigner(),
	)

	ethRPC := ethereum.New()
	pos.New(ethRPC)

	server.New(
		websocket.WithWebsocket(),
	)
}
