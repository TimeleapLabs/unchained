package correctness

import (
	"testing"

	"github.com/TimeleapLabs/unchained/internal/model"
	postgresRepo "github.com/TimeleapLabs/unchained/internal/repository/postgres"
	"github.com/TimeleapLabs/unchained/internal/service/pos"
	"github.com/TimeleapLabs/unchained/internal/transport/database/mock"
	"github.com/stretchr/testify/suite"
)

var (
	SignerOne = Signature{
		Signer: model.Signer{
			PublicKey: [96]byte{1, 2, 3},
		},
	}
	SignerTwo = Signature{
		Signer: model.Signer{
			PublicKey: [96]byte{3, 2, 1},
		},
	}
)

type CorrectnessTestSuite struct {
	suite.Suite
	service Service
}

func (s *CorrectnessTestSuite) SetupTest() {
	db := mock.New(s.T())

	posService := new(pos.MockService)

	signerRepo := postgresRepo.NewProof(db)
	correctnessRepo := postgresRepo.NewCorrectness(db)

	s.service = New(posService, signerRepo, correctnessRepo)
}

func TestCorrectnessSuite(t *testing.T) {
	suite.Run(t, new(CorrectnessTestSuite))
}
