package app

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/service/rpc"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/transport/client"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
	"github.com/TimeleapLabs/unchained/internal/transport/client/handler"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

// Worker starts the Unchained worker and contains its DI.
func Worker(_ context.Context) {
	utils.Logger.
		With("Mode", "Worker").
		With("Version", consts.Version).
		With("Protocol", consts.ProtocolVersion).
		Info("Running Unchained")

	crypto.InitMachineIdentity(
		crypto.WithEvmSigner(),
		crypto.WithBlsIdentity(),
	)

	rpcFunctions := []rpc.Option{}
	for _, fun := range config.App.Functions {
		switch fun.Type { //nolint: gocritic // This is a switch case for different types of rpc functions
		case "unix":
			rpcFunctions = append(rpcFunctions, rpc.WithUnixSocket(fun.Name, fun.Endpoint))
		}
	}
	rpcService := rpc.NewWorker(rpcFunctions...)

	conn.Start()

	workerHandler := handler.NewWorkerHandler(rpcService)
	client.NewRPC(workerHandler)
}
