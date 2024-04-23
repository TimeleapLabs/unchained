package app

import (
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/constants"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"
	"github.com/TimeleapLabs/unchained/internal/log"
	"github.com/TimeleapLabs/unchained/internal/persistence"
	"github.com/TimeleapLabs/unchained/internal/pos"
	"github.com/TimeleapLabs/unchained/internal/scheduler"
	evmlogService "github.com/TimeleapLabs/unchained/internal/service/evmlog"
	uniswapService "github.com/TimeleapLabs/unchained/internal/service/uniswap"
	"github.com/TimeleapLabs/unchained/internal/transport/client"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
	"github.com/TimeleapLabs/unchained/internal/transport/client/handler"
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
