package app

import (
	"github.com/TimeleapLabs/timeleap/internal/config"
	"github.com/TimeleapLabs/timeleap/internal/consts"
	"github.com/TimeleapLabs/timeleap/internal/crypto"
	"github.com/TimeleapLabs/timeleap/internal/crypto/ethereum"
	"github.com/TimeleapLabs/timeleap/internal/repository"
	mongoRepo "github.com/TimeleapLabs/timeleap/internal/repository/mongo"
	attestationService "github.com/TimeleapLabs/timeleap/internal/service/attestation"
	"github.com/TimeleapLabs/timeleap/internal/service/pos"
	"github.com/TimeleapLabs/timeleap/internal/transport/client"
	"github.com/TimeleapLabs/timeleap/internal/transport/client/conn"
	"github.com/TimeleapLabs/timeleap/internal/transport/client/handler"
	"github.com/TimeleapLabs/timeleap/internal/transport/database/mongo"
	"github.com/TimeleapLabs/timeleap/internal/utils"
)

// Consumer starts the Timeleap consumer and contains its DI.
func Consumer() {
	utils.Logger.
		With("Mode", "Consumer").
		With("Version", consts.Version).
		With("Protocol", consts.ProtocolVersion).
		Info("Running Timeleap")

	crypto.InitMachineIdentity(
		crypto.WithEvmSigner(),
		crypto.WithEd25519Identity(),
	)

	ethRPC := ethereum.New()
	posService := pos.New(ethRPC)

	var proofRepo repository.Proof
	var attestationRepo repository.Attestation

	if config.App.Mongo.URL != "" {
		utils.Logger.Info("MongoDB configuration found, initializing...")
		db := mongo.New()

		proofRepo = mongoRepo.NewProof(db)
		attestationRepo = mongoRepo.NewAttestation(db)
	} else {
		utils.Logger.Error("MongoDB configuration not found, exiting...")
		return
	}

	_attestationService := attestationService.New(posService, proofRepo, attestationRepo)

	conn.Start()

	consumerHandler := handler.NewConsumerHandler(_attestationService)
	client.NewRPC(consumerHandler)

	select {}
}
