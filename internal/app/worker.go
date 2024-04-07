package app

import (
	"github.com/KenshiTech/unchained/internal/config"
	"github.com/KenshiTech/unchained/internal/constants"
	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/ethereum"
	"github.com/KenshiTech/unchained/internal/log"
	"github.com/KenshiTech/unchained/internal/persistence"
	"github.com/KenshiTech/unchained/internal/pos"
	"github.com/KenshiTech/unchained/internal/scheduler"
	correctnessService "github.com/KenshiTech/unchained/internal/service/correctness"
	evmlogService "github.com/KenshiTech/unchained/internal/service/evmlog"
	uniswapService "github.com/KenshiTech/unchained/internal/service/uniswap"
	"github.com/KenshiTech/unchained/internal/transport/client"
	"github.com/KenshiTech/unchained/internal/transport/client/handler"
)

// Worker starts the Unchained worker and contains its DI.
func Worker() {
	log.Logger.
		With("Version", constants.Version).
		With("Protocol", constants.ProtocolVersion).
		Info("Running Unchained | Worker")

	err := config.Load(config.App.System.ConfigPath, config.App.System.SecretsPath)
	if err != nil {
		panic(err)
	}

	bls.InitClientIdentity()

	ethRPC := ethereum.New()
	pos := pos.New(ethRPC)
	badger := persistence.New(config.App.System.ContextPath)

	correctnessService := correctnessService.New(ethRPC)
	evmLogService := evmlogService.New(ethRPC, pos)
	uniswapService := uniswapService.New(ethRPC, pos)

	scheduler.New(
		scheduler.WithEthLogs(evmLogService, ethRPC, badger),
		scheduler.WithUniswapEvents(uniswapService, ethRPC),
	)

	handler := handler.New(correctnessService, uniswapService, evmLogService)
	client.Consume(handler)
}
