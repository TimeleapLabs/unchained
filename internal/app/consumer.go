package app

import (
	"github.com/TimeleapLabs/unchained/internal/constants"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"
	"github.com/TimeleapLabs/unchained/internal/db"
	"github.com/TimeleapLabs/unchained/internal/log"
	"github.com/TimeleapLabs/unchained/internal/pos"
	correctnessService "github.com/TimeleapLabs/unchained/internal/service/correctness"
	evmlogService "github.com/TimeleapLabs/unchained/internal/service/evmlog"
	uniswapService "github.com/TimeleapLabs/unchained/internal/service/uniswap"
	"github.com/TimeleapLabs/unchained/internal/transport/client"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
	"github.com/TimeleapLabs/unchained/internal/transport/client/handler"
	"github.com/TimeleapLabs/unchained/internal/transport/server"
	"github.com/TimeleapLabs/unchained/internal/transport/server/gql"
)

// Consumer starts the Unchained consumer and contains its DI.
func Consumer() {
	log.Logger.
		With("Mode", "Consumer").
		With("Version", constants.Version).
		With("Protocol", constants.ProtocolVersion).
		Info("Running Unchained")

	crypto.InitMachineIdentity(
		crypto.WithEvmSigner(),
		crypto.WithBlsIdentity(),
	)

	ethRPC := ethereum.New()
	pos := pos.New(ethRPC)
	db.Start()

	correctnessService := correctnessService.New(ethRPC, pos)
	evmLogService := evmlogService.New(ethRPC, pos)
	uniswapService := uniswapService.New(ethRPC, pos)

	conn.Start()

	handler := handler.NewConsumerHandler(correctnessService, uniswapService, evmLogService)
	client.NewRPC(handler)

	server.New(
		gql.WithGraphQL(),
	)
}
