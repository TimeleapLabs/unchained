package app

import (
	"context"

	"github.com/TimeleapLabs/timeleap/internal/service/rpc"
	"github.com/TimeleapLabs/timeleap/internal/service/rpc/worker"

	"github.com/TimeleapLabs/timeleap/internal/config"
	"github.com/TimeleapLabs/timeleap/internal/consts"
	"github.com/TimeleapLabs/timeleap/internal/crypto"
	"github.com/TimeleapLabs/timeleap/internal/transport/client"
	"github.com/TimeleapLabs/timeleap/internal/transport/client/conn"
	"github.com/TimeleapLabs/timeleap/internal/transport/client/handler"
	"github.com/TimeleapLabs/timeleap/internal/utils"
)

// Worker starts the Timeleap worker and contains its DI.
func Worker(_ context.Context) {
	utils.Logger.
		With("Mode", "Worker").
		With("Version", consts.Version).
		With("Protocol", consts.ProtocolVersion).
		Info("Running Timeleap")

	crypto.InitMachineIdentity(
		crypto.WithEvmSigner(),
		crypto.WithEd25519Identity(),
	)

	rpcFunctions := []worker.Option{}
	for _, plugin := range config.App.Plugins {
		switch plugin.Type { //nolint: gocritic // This is a switch case for different types of rpc functions
		case "websocket":
			rpcFunctions = append(rpcFunctions,
				rpc.WithWebSocket(
					plugin.Name,
					plugin.Functions,
					plugin.Endpoint,
					plugin.PublicKey,
				),
			)
		}
	}
	rpcService := worker.NewWorker(rpcFunctions...)

	conn.Start()

	workerHandler := handler.NewWorkerHandler(rpcService)
	client.NewRPC(workerHandler)

	select {}
}
