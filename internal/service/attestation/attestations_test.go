package attestation

import (
	"crypto/ed25519"
	"testing"

	"github.com/TimeleapLabs/timeleap/internal/transport/database/mongo"

	"github.com/TimeleapLabs/timeleap/internal/model"
	mongoRepo "github.com/TimeleapLabs/timeleap/internal/repository/mongo"
	"github.com/TimeleapLabs/timeleap/internal/service/pos"
	"github.com/stretchr/testify/suite"
)

var (
	SignerOne = Signature{
		Signer: model.Signer{
			PublicKey: ed25519.PublicKey{1, 2, 3},
		},
	}
	SignerTwo = Signature{
		Signer: model.Signer{
			PublicKey: ed25519.PublicKey{3, 2, 1},
		},
	}
)

type AttestationTestSuite struct {
	suite.Suite
	service Service
}

func (s *AttestationTestSuite) SetupTest() {
	db := mongo.New()

	posService := new(pos.MockService)

	signerRepo := mongoRepo.NewProof(db)
	attestationRepo := mongoRepo.NewAttestation(db)

	s.service = New(posService, signerRepo, attestationRepo)
}

func TestAttestationSuite(t *testing.T) {
	suite.Run(t, new(AttestationTestSuite))
}
