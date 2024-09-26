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
	_posService := pos.New(ethRPC)

	var eventLogRepo repository.EventLog
	var proofRepo repository.Proof
	var assetPrice repository.AssetPrice
	var correctnessRepo repository.CorrectnessReport

	if config.App.Mongo.URL != "" {
		utils.Logger.Info("MongoDB configuration found, initializing...")
		db := mongo.New()

		eventLogRepo = mongoRepo.NewEventLog(db)
		proofRepo = mongoRepo.NewProof(db)
		assetPrice = mongoRepo.NewAssetPrice(db)
		correctnessRepo = mongoRepo.NewCorrectness(db)
	} else {
		utils.Logger.Info("Postgresql configuration found, initializing...")
		db := postgres.New()
		db.Migrate()

		eventLogRepo = postgresRepo.NewEventLog(db)
		proofRepo = postgresRepo.NewProof(db)
		assetPrice = postgresRepo.NewAssetPrice(db)
		correctnessRepo = postgresRepo.NewCorrectness(db)
	}

	_correctnessService := correctnessService.New(_posService, proofRepo, correctnessRepo)
	_evmLogService := evmlogService.New(ethRPC, _posService, eventLogRepo, proofRepo, nil)
	_uniswapService := uniswapService.New(ethRPC, _posService, proofRepo, assetPrice)

	conn.Start()

	consumerHandler := handler.NewConsumerHandler(_correctnessService, _uniswapService, _evmLogService)
	client.NewRPC(consumerHandler)

	select {}
}
