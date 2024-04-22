package app

import (
	"github.com/KenshiTech/unchained/internal/consts"
	"github.com/KenshiTech/unchained/internal/crypto"
	"github.com/KenshiTech/unchained/internal/crypto/ethereum"
	postgresRepo "github.com/KenshiTech/unchained/internal/repository/postgres"
	correctnessService "github.com/KenshiTech/unchained/internal/service/correctness"
	evmlogService "github.com/KenshiTech/unchained/internal/service/evmlog"
	"github.com/KenshiTech/unchained/internal/service/pos"
	uniswapService "github.com/KenshiTech/unchained/internal/service/uniswap"
	"github.com/KenshiTech/unchained/internal/transport/client"
	"github.com/KenshiTech/unchained/internal/transport/client/conn"
	"github.com/KenshiTech/unchained/internal/transport/client/handler"
	"github.com/KenshiTech/unchained/internal/transport/database/postgres"
	"github.com/KenshiTech/unchained/internal/transport/server"
	"github.com/KenshiTech/unchained/internal/transport/server/gql"
	"github.com/KenshiTech/unchained/internal/utils"
)

// Consumer starts the Unchained consumer and contains its DI.
func Consumer() {
	utils.Logger.
		With("Mode", "Consumer").
		With("Version", consts.Version).
		With("Protocol", consts.ProtocolVersion).
		Info("Running Unchained")

	crypto.InitMachineIdentity(
		crypto.WithEvmSigner(),
		crypto.WithBlsIdentity(),
	)

	ethRPC := ethereum.New()
	pos := pos.New(ethRPC)
	db := postgres.New()

	eventLogRepo := postgresRepo.NewEventLog(db)
	signerRepo := postgresRepo.NewSigner(db)
	assetPrice := postgresRepo.NewAssetPrice(db)
	correctnessRepo := postgresRepo.NewCorrectness(db)

	correctnessService := correctnessService.New(pos, signerRepo, correctnessRepo)
	evmLogService := evmlogService.New(ethRPC, pos, eventLogRepo, signerRepo, nil)
	uniswapService := uniswapService.New(ethRPC, pos, signerRepo, assetPrice)

	conn.Start()

	handler := handler.NewConsumerHandler(correctnessService, uniswapService, evmLogService)
	client.NewRPC(handler)

	server.New(
		gql.WithGraphQL(db),
	)
}
