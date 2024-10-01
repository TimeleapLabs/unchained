package app

import (
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"
	correctnessService "github.com/TimeleapLabs/unchained/internal/service/correctness"
	"github.com/TimeleapLabs/unchained/internal/service/pos"
	"github.com/TimeleapLabs/unchained/internal/service/rpc"
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
		crypto.WithBlsIdentity(),
		crypto.WithEvmSigner(),
	)

	ethRPC := ethereum.New()
	_pos := pos.New(ethRPC)

	correctnessService := correctnessService.New(_pos)

	rpcFunctions := []rpc.Option{}
	for _, fun := range config.App.Functions {
		switch fun.Type { //nolint: gocritic // This is a switch case for different types of rpc functions
		case "unix":
			rpcFunctions = append(rpcFunctions, rpc.WithUnixSocket(fun.Name, fun.Endpoint))
		}
	}

	rpcService := rpc.NewWorker(rpcFunctions...)

	server.New(
		websocket.WithWebsocket(),
	)
}
