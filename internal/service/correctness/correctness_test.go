package correctness

import (
	"testing"

	"github.com/TimeleapLabs/unchained/internal/ent"
	"github.com/TimeleapLabs/unchained/internal/model"
	postgresRepo "github.com/TimeleapLabs/unchained/internal/repository/postgres"
	"github.com/TimeleapLabs/unchained/internal/service/pos"
	"github.com/TimeleapLabs/unchained/internal/transport/database/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var (
	SignerOne = model.Signature{
		Signer: model.Signer{
			PublicKey: [96]byte{1, 2, 3},
		},
	}
	SignerTwo = model.Signature{
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

	signerRepo := postgresRepo.NewSigner(db)
	correctnessRepo := postgresRepo.NewCorrectness(db)

	s.service = New(posService, signerRepo, correctnessRepo)
}

func (s *CorrectnessTestSuite) TestIsNewSigner() {
	s.Run("Check if new signer with empty values", func() {
		isSigner := s.service.IsNewSigner(model.Signature{}, []*ent.CorrectnessReport{})
		assert.False(s.T(), isSigner)
	})

	s.Run("Check when sign is new signer", func() {
		signers := make([]byte, 96)
		for i := 1; i < 4; i++ {
			signers[i] = byte(i)
		}

		isSigner := s.service.IsNewSigner(
			SignerOne,
			[]*ent.CorrectnessReport{
				{
					Edges: ent.CorrectnessReportEdges{
						Signers: []*ent.Signer{
							{
								Key: signers,
							},
						},
					},
				},
			},
		)
		assert.True(s.T(), isSigner)
	})

	s.Run("Check when sign is not new signer", func() {
		signers := make([]byte, 96)
		for i := 2; i < 4; i++ {
			signers[i] = byte(i)
		}

		isSigner := s.service.IsNewSigner(
			SignerTwo,
			[]*ent.CorrectnessReport{
				{
					Edges: ent.CorrectnessReportEdges{
						Signers: []*ent.Signer{
							{
								Key: signers,
							},
						},
					},
				},
			},
		)
		assert.True(s.T(), isSigner)
	})
}

func TestCorrectnessSuite(t *testing.T) {
	suite.Run(t, new(CorrectnessTestSuite))
}
