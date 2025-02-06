package app

import (
	"github.com/TimeleapLabs/timeleap/internal/consts"
	"github.com/TimeleapLabs/timeleap/internal/crypto"
	"github.com/TimeleapLabs/timeleap/internal/crypto/ethereum"
	"github.com/TimeleapLabs/timeleap/internal/service/pos"
	"github.com/TimeleapLabs/timeleap/internal/transport/server"
	"github.com/TimeleapLabs/timeleap/internal/transport/server/websocket"
	"github.com/TimeleapLabs/timeleap/internal/utils"
)

// Broker starts the Timeleap broker and contains its DI.
func Broker() {
	utils.Logger.
		With("Mode", "Broker").
		With("Version", consts.Version).
		With("Protocol", consts.ProtocolVersion).
		Info("Running Timeleap")

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
