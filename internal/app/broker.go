package app

import (
	"github.com/KenshiTech/unchained/internal/consts"
	"github.com/KenshiTech/unchained/internal/crypto"
	"github.com/KenshiTech/unchained/internal/crypto/ethereum"
	"github.com/KenshiTech/unchained/internal/service/pos"
	"github.com/KenshiTech/unchained/internal/transport/server"
	"github.com/KenshiTech/unchained/internal/transport/server/websocket"
	"github.com/KenshiTech/unchained/internal/utils"
)

// Broker starts the Unchained broker and contains its DI.
func Broker() {
	utils.Logger.
		With("Mode", "Broker").
		With("Version", consts.Version).
		With("Protocol", consts.ProtocolVersion).
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
