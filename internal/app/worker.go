package app

import (
	"github.com/KenshiTech/unchained/internal/config"
	"github.com/KenshiTech/unchained/internal/constants"
	"github.com/KenshiTech/unchained/internal/crypto"
	"github.com/KenshiTech/unchained/internal/crypto/ethereum"
	"github.com/KenshiTech/unchained/internal/log"
	"github.com/KenshiTech/unchained/internal/persistence"
	"github.com/KenshiTech/unchained/internal/pos"
	"github.com/KenshiTech/unchained/internal/scheduler"
	evmlogService "github.com/KenshiTech/unchained/internal/service/evmlog"
	uniswapService "github.com/KenshiTech/unchained/internal/service/uniswap"
	"github.com/KenshiTech/unchained/internal/transport/client"
	"github.com/KenshiTech/unchained/internal/transport/client/conn"
	"github.com/KenshiTech/unchained/internal/transport/client/handler"
)

// Worker starts the Unchained worker and contains its DI.
func Worker() {
	log.Logger.
		With("Mode", "Worker").
		With("Version", constants.Version).
		With("Protocol", constants.ProtocolVersion).
		Info("Running Unchained")

	crypto.InitMachineIdentity(
		crypto.WithEvmSigner(),
		crypto.WithBlsIdentity(),
	)

	ethRPC := ethereum.New()
	pos := pos.New(ethRPC)
	badger := persistence.New(config.App.System.ContextPath)

	evmLogService := evmlogService.New(ethRPC, pos)
	uniswapService := uniswapService.New(ethRPC, pos)

	scheduler := scheduler.New(
		scheduler.WithEthLogs(evmLogService, ethRPC, badger),
		scheduler.WithUniswapEvents(uniswapService, ethRPC),
	)

	conn.Start()

	handler := handler.NewWorkerHandler()
	client.NewRPC(handler)

	scheduler.Start()
}
