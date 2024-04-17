package app

import (
	"github.com/KenshiTech/unchained/internal/config"
	"github.com/KenshiTech/unchained/internal/consts"
	"github.com/KenshiTech/unchained/internal/crypto"
	"github.com/KenshiTech/unchained/internal/crypto/ethereum"
	"github.com/KenshiTech/unchained/internal/pos"
	"github.com/KenshiTech/unchained/internal/repository/postgres"
	"github.com/KenshiTech/unchained/internal/scheduler"
	"github.com/KenshiTech/unchained/internal/scheduler/persistence"
	evmlogService "github.com/KenshiTech/unchained/internal/service/evmlog"
	uniswapService "github.com/KenshiTech/unchained/internal/service/uniswap"
	"github.com/KenshiTech/unchained/internal/transport/client"
	"github.com/KenshiTech/unchained/internal/transport/client/conn"
	"github.com/KenshiTech/unchained/internal/transport/client/handler"
	"github.com/KenshiTech/unchained/internal/utils"
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
	badger := persistence.New(config.App.System.ContextPath)

	eventLogRepo := postgres.NewEventLog(nil)
	signerRepo := postgres.NewSigner(nil)
	assetPrice := postgres.NewAssetPrice(nil)

	evmLogService := evmlogService.New(ethRPC, pos, eventLogRepo, signerRepo)
	uniswapService := uniswapService.New(ethRPC, pos, signerRepo, assetPrice)

	scheduler := scheduler.New(
		scheduler.WithEthLogs(evmLogService, ethRPC, badger),
		scheduler.WithUniswapEvents(uniswapService, ethRPC),
	)

	conn.Start()

	handler := handler.NewWorkerHandler()
	client.NewRPC(handler)

	scheduler.Start()
}
