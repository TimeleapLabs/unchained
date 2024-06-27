package correctness

import (
	"context"
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/utils"
	mock2 "github.com/stretchr/testify/mock"
	"testing"

	"github.com/TimeleapLabs/unchained/internal/model"
	postgresRepo "github.com/TimeleapLabs/unchained/internal/repository/postgres"
	"github.com/TimeleapLabs/unchained/internal/service/pos"
	"github.com/TimeleapLabs/unchained/internal/transport/database/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var (
	SignatureOne = Signature{
		Signer: model.Signer{
			PublicKey:      [96]byte{1, 2, 3},
			ShortPublicKey: [48]byte{1, 2, 3},
		},
	}
	SignatureTwo = Signature{
		Signer: model.Signer{
			PublicKey:      [96]byte{3, 2, 1},
			ShortPublicKey: [48]byte{1, 2, 3},
		},
	}
	sampleSigner = model.Signer{
		Name:           "",
		EvmAddress:     "12345",
		PublicKey:      [96]byte{3, 2, 1},
		ShortPublicKey: [48]byte{1, 2, 3},
	}
)

var SampleCorrectness = []model.Correctness{
	{
		SignersCount: 1,
		Topic:        [64]byte(utils.Shake([]byte("123"))),
		Signers: []model.Signer{
			{
				Name:           "test-1",
				EvmAddress:     "12345",
				PublicKey:      [96]byte{3, 2, 1},
				ShortPublicKey: [48]byte{3, 2, 1},
			},
		},
	},
}

type CorrectnessTestSuite struct {
	suite.Suite
	service Service
}

func (s *CorrectnessTestSuite) SetupTest() {
	utils.SetupLogger("info")
	db := mock.New(s.T())

	posService := new(pos.MockService)
	posService.On("GetVotingPowerOfEvm", mock2.Anything, "12345").Return(10, nil)

	signerRepo := postgresRepo.NewProof(db)
	correctnessRepo := postgresRepo.NewCorrectness(db)

	config.App.Plugins.Correctness = []string{"123"}
	s.service = New(posService, signerRepo, correctnessRepo)
}

func (s *CorrectnessTestSuite) TestIsNewSigner() {
	s.Run("Check if new signer with empty values", func() {
		isSigner := s.service.IsNewSigner(Signature{}, []model.Correctness{})
		assert.True(s.T(), isSigner)
	})

	s.Run("Check when sign is new signer", func() {
		isSigner := s.service.IsNewSigner(
			SignatureOne,
			SampleCorrectness,
		)
		assert.True(s.T(), isSigner)
	})

	s.Run("Check when sign is not new signer", func() {
		isSigner := s.service.IsNewSigner(
			SignatureTwo,
			SampleCorrectness,
		)

		assert.False(s.T(), isSigner)
	})
}

func (s *CorrectnessTestSuite) TestRecordSignatures() {
	_, _, shortPublicKey := bls.GenerateBlsKeyPair()

	signature, err := bls.RecoverSignature(shortPublicKey.Bytes())
	s.Require().NoError(err)

	err = s.service.RecordSignature(
		context.TODO(),
		signature,
		sampleSigner,
		SampleCorrectness[0],
		false,
	)
	s.Require().NoError(err)
}

func TestCorrectnessSuite(t *testing.T) {
	suite.Run(t, new(CorrectnessTestSuite))
}
