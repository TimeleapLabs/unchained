package app

import (
	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/constants"
	"github.com/KenshiTech/unchained/crypto/bls"
	"github.com/KenshiTech/unchained/db"
	"github.com/KenshiTech/unchained/ethereum"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/pos"
	correctnessService "github.com/KenshiTech/unchained/service/correctness"
	evmlogService "github.com/KenshiTech/unchained/service/evmlog"
	uniswapService "github.com/KenshiTech/unchained/service/uniswap"
	"github.com/KenshiTech/unchained/transport/client"
	"github.com/KenshiTech/unchained/transport/client/handler"
	"github.com/KenshiTech/unchained/transport/server"
	"github.com/KenshiTech/unchained/transport/server/gql"
)

func Consumer() {
	log.Start(config.App.System.Log)
	log.Logger.
		With("Version", constants.Version).
		With("Protocol", constants.ProtocolVersion).
		Info("Running Unchained | Consumer")

	err := config.Load(configPath, secretsPath)
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
