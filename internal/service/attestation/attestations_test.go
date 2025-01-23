package attestation

import (
	"crypto/ed25519"
	"testing"

	"github.com/TimeleapLabs/unchained/internal/transport/database/postgres"

	"github.com/TimeleapLabs/unchained/internal/model"
	postgresRepo "github.com/TimeleapLabs/unchained/internal/repository/postgres"
	"github.com/TimeleapLabs/unchained/internal/service/pos"
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
	db := postgres.New()

	posService := new(pos.MockService)

	signerRepo := postgresRepo.NewProof(db)
	attestationRepo := postgresRepo.NewAttestation(db)

	s.service = New(posService, signerRepo, attestationRepo)
}

func TestAttestationSuite(t *testing.T) {
	suite.Run(t, new(AttestationTestSuite))
}
