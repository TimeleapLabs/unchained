package app

import (
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"
	"github.com/TimeleapLabs/unchained/internal/repository"
	mongoRepo "github.com/TimeleapLabs/unchained/internal/repository/mongo"
	postgresRepo "github.com/TimeleapLabs/unchained/internal/repository/postgres"
	correctnessService "github.com/TimeleapLabs/unchained/internal/service/correctness"
	evmlogService "github.com/TimeleapLabs/unchained/internal/service/evmlog"
	"github.com/TimeleapLabs/unchained/internal/service/pos"
	uniswapService "github.com/TimeleapLabs/unchained/internal/service/uniswap"
	"github.com/TimeleapLabs/unchained/internal/transport/client"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
	"github.com/TimeleapLabs/unchained/internal/transport/client/handler"
	"github.com/TimeleapLabs/unchained/internal/transport/database/mongo"
	"github.com/TimeleapLabs/unchained/internal/transport/database/postgres"
	"github.com/TimeleapLabs/unchained/internal/utils"
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

	var eventLogRepo repository.EventLog
	var proofRepo repository.Proof
	var assetPrice repository.AssetPrice
	var correctnessRepo repository.CorrectnessReport

	if config.App.Mongo.URL != "" {
		db := mongo.New()

		eventLogRepo = mongoRepo.NewEventLog(db)
		proofRepo = mongoRepo.NewProof(db)
		assetPrice = mongoRepo.NewAssetPrice(db)
		correctnessRepo = mongoRepo.NewCorrectness(db)
	} else {
		db := postgres.New()
		db.Migrate()

		eventLogRepo = postgresRepo.NewEventLog(db)
		proofRepo = postgresRepo.NewProof(db)
		assetPrice = postgresRepo.NewAssetPrice(db)
		correctnessRepo = postgresRepo.NewCorrectness(db)
	}

	correctnessService := correctnessService.New(pos, proofRepo, correctnessRepo)
	evmLogService := evmlogService.New(ethRPC, pos, eventLogRepo, proofRepo, nil)
	uniswapService := uniswapService.New(ethRPC, pos, proofRepo, assetPrice)

	conn.Start()

	handler := handler.NewConsumerHandler(correctnessService, uniswapService, evmLogService)
	client.NewRPC(handler)

	select {}
}
