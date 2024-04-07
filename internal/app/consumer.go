package app

import (
	"github.com/KenshiTech/unchained/internal/config"
	"github.com/KenshiTech/unchained/internal/constants"
	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/db"
	"github.com/KenshiTech/unchained/internal/ethereum"
	"github.com/KenshiTech/unchained/internal/log"
	"github.com/KenshiTech/unchained/internal/pos"
	correctnessService "github.com/KenshiTech/unchained/internal/service/correctness"
	evmlogService "github.com/KenshiTech/unchained/internal/service/evmlog"
	uniswapService "github.com/KenshiTech/unchained/internal/service/uniswap"
	"github.com/KenshiTech/unchained/internal/transport/client"
	"github.com/KenshiTech/unchained/internal/transport/client/handler"
	"github.com/KenshiTech/unchained/internal/transport/server"
	"github.com/KenshiTech/unchained/internal/transport/server/gql"
)

// Consumer starts the Unchained consumer and contains its DI.
func Consumer() {
	log.Logger.
		With("Version", constants.Version).
		With("Protocol", constants.ProtocolVersion).
		Info("Running Unchained | Consumer")

	err := config.Load(config.App.System.ConfigPath, config.App.System.SecretsPath)
	if err != nil {
		panic(err)
	}

	bls.InitClientIdentity()

	ethRPC := ethereum.New()
	pos := pos.New(ethRPC)
	db.Start()

	server.New(
		gql.WithGraphQL(),
	)

	correctnessService := correctnessService.New(ethRPC)
	evmLogService := evmlogService.New(ethRPC, pos)
	uniswapService := uniswapService.New(ethRPC, pos)

	handler := handler.New(correctnessService, uniswapService, evmLogService)
	client.Consume(handler)
}
