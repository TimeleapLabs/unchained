package app

import (
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"
	"github.com/TimeleapLabs/unchained/internal/repository/postgres"
	"github.com/TimeleapLabs/unchained/internal/scheduler"
	evmlogService "github.com/TimeleapLabs/unchained/internal/service/evmlog"
	"github.com/TimeleapLabs/unchained/internal/service/frost"
	"github.com/TimeleapLabs/unchained/internal/service/pos"
	uniswapService "github.com/TimeleapLabs/unchained/internal/service/uniswap"
	"github.com/TimeleapLabs/unchained/internal/transport/client"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
	"github.com/TimeleapLabs/unchained/internal/transport/client/handler"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

// Worker starts the Unchained worker and contains its DI.
func Worker() {
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

	pos := pos.New(ethRPC)

	eventLogRepo := postgres.NewEventLog(nil)
	signerRepo := postgres.NewSigner(nil)
	assetPrice := postgres.NewAssetPrice(nil)

	badger := evmlogService.NewBadger(config.App.System.ContextPath)
	evmLogService := evmlogService.New(ethRPC, pos, eventLogRepo, signerRepo, badger)
	uniswapService := uniswapService.New(ethRPC, pos, signerRepo, assetPrice)
	frostService := frost.New(pos)

	scheduler := scheduler.New(
		scheduler.WithEthLogs(evmLogService),
		scheduler.WithUniswapEvents(uniswapService),
		scheduler.WithFrostEvents(frostService),
	)

	conn.Start()

	handler := handler.NewWorkerHandler(frostService)
	client.NewRPC(handler)

	scheduler.Start()
}
