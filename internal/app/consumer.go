package app

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"
	postgresRepo "github.com/TimeleapLabs/unchained/internal/repository/postgres"
	correctnessService "github.com/TimeleapLabs/unchained/internal/service/correctness"
	evmlogService "github.com/TimeleapLabs/unchained/internal/service/evmlog"
	"github.com/TimeleapLabs/unchained/internal/service/pos"
	uniswapService "github.com/TimeleapLabs/unchained/internal/service/uniswap"
	"github.com/TimeleapLabs/unchained/internal/transport/client"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
	"github.com/TimeleapLabs/unchained/internal/transport/client/handler"
	"github.com/TimeleapLabs/unchained/internal/transport/client/store"
	"github.com/TimeleapLabs/unchained/internal/transport/database/postgres"
	"github.com/TimeleapLabs/unchained/internal/transport/server"
	"github.com/TimeleapLabs/unchained/internal/transport/server/gql"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

type JobType string

const (
	PostgresConsumer JobType = "postgres"
	SchnorrConsumer  JobType = "schnorr"
)

// Consumer starts the Unchained consumer and contains its DI.
func Consumer(jobType JobType) {
	utils.Logger.
		With("Mode", "Consumer").
		With("Version", consts.Version).
		With("Protocol", consts.ProtocolVersion).
		Info("Running Unchained")

	crypto.InitMachineIdentity(
		crypto.WithEvmSigner(),
		crypto.WithBlsIdentity(),
	)

	if jobType == PostgresConsumer {
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

		handler := handler.NewPostgresConsumerHandler(correctnessService, uniswapService, evmLogService)
		client.NewRPC(handler)

		server.New(
			gql.WithGraphQL(db),
		)
	} else {
		signerRepo := store.New()

		conn.Start()

		handler := handler.NewSchnorrConsumerHandler(signerRepo)
		client.NewRPC(handler)

		select {}
	}
}
