package app

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/service/rpc"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"
	"github.com/TimeleapLabs/unchained/internal/repository/postgres"
	"github.com/TimeleapLabs/unchained/internal/scheduler"
	evmlogService "github.com/TimeleapLabs/unchained/internal/service/evmlog"
	"github.com/TimeleapLabs/unchained/internal/service/pos"
	uniswapService "github.com/TimeleapLabs/unchained/internal/service/uniswap"
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

	ethRPC := ethereum.New()

	eventLogRepo := postgres.NewEventLog(nil)
	proofRepo := postgres.NewProof(nil)
	assetPrice := postgres.NewAssetPrice(nil)

	badger := evmlogService.NewBadger(config.App.System.ContextPath)
	evmLogService := evmlogService.New(ethRPC, pos, eventLogRepo, proofRepo, badger)
	uniswapService := uniswapService.New(ethRPC, pos, proofRepo, assetPrice)
	_posService := pos.New(ethRPC)
	_badgerService := evmlogService.NewBadger(config.App.System.ContextPath)
	_evmLogService := evmlogService.New(ethRPC, _posService, eventLogRepo, signerRepo, _badgerService)
	_uniswapService := uniswapService.New(ethRPC, _posService, signerRepo, assetPrice)

	rpcFunctions := []rpc.Option{}
	for _, fun := range config.App.Functions {
		switch fun.Type { //nolint: gocritic // This is a switch case for different types of rpc functions
		case "unix":
			rpcFunctions = append(rpcFunctions, rpc.WithUnixSocket(fun.Name, fun.Endpoint))
		}
	}
	rpcService := rpc.NewWorker(rpcFunctions...)

	taskScheduler := scheduler.New(
		scheduler.WithEthLogs(_evmLogService),
		scheduler.WithUniswapEvents(_uniswapService),
	)

	conn.Start()

	workerHandler := handler.NewWorkerHandler(rpcService)
	client.NewRPC(workerHandler)

	taskScheduler.Start()
}
