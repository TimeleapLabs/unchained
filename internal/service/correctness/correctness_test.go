package correctness

import (
	"encoding/hex"
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
			PublicKey: hex.EncodeToString([]byte{1, 2, 3}),
		},
	}
	SignerTwo = Signature{
		Signer: model.Signer{
			PublicKey: hex.EncodeToString([]byte{3, 2, 1}),
		},
	}
)

type CorrectnessTestSuite struct {
	suite.Suite
	service Service
}

func (s *CorrectnessTestSuite) SetupTest() {
	db := postgres.New()

	posService := new(pos.MockService)

	signerRepo := postgresRepo.NewProof(db)
	correctnessRepo := postgresRepo.NewCorrectness(db)

	s.service = New(posService, signerRepo, correctnessRepo)
}

func TestCorrectnessSuite(t *testing.T) {
	suite.Run(t, new(CorrectnessTestSuite))
}
