package app

import (
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"
	"github.com/TimeleapLabs/unchained/internal/repository"
	mongoRepo "github.com/TimeleapLabs/unchained/internal/repository/mongo"
	postgresRepo "github.com/TimeleapLabs/unchained/internal/repository/postgres"
	attestationService "github.com/TimeleapLabs/unchained/internal/service/attestation"
	"github.com/TimeleapLabs/unchained/internal/service/pos"
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

	var proofRepo repository.Proof
	var attestationRepo repository.Attestation

	if config.App.Mongo.URL != "" {
		utils.Logger.Info("MongoDB configuration found, initializing...")
		db := mongo.New()

		proofRepo = mongoRepo.NewProof(db)
		attestationRepo = mongoRepo.NewAttestation(db)
	} else {
		utils.Logger.Info("Postgresql configuration found, initializing...")
		db := postgres.New()
		db.Migrate()

		proofRepo = postgresRepo.NewProof(db)
		attestationRepo = postgresRepo.NewAttestation(db)
	}

	_attestationService := attestationService.New(_posService, proofRepo, attestationRepo)

	conn.Start()

	consumerHandler := handler.NewConsumerHandler(_attestationService)
	client.NewRPC(consumerHandler)

	select {}
}
