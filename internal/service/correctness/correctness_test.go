package correctness

import (
	"context"
	"fmt"
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/transport/database"
	"github.com/TimeleapLabs/unchained/internal/transport/database/postgres"
	"github.com/TimeleapLabs/unchained/internal/utils"
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/google/uuid"
	mock2 "github.com/stretchr/testify/mock"
	"os"
	"testing"
	"time"

	"github.com/TimeleapLabs/unchained/internal/model"
	postgresRepo "github.com/TimeleapLabs/unchained/internal/repository/postgres"
	"github.com/TimeleapLabs/unchained/internal/service/pos"
	"github.com/stretchr/testify/suite"
)

var (
	SignatureOne = Signature{
		Signer: model.Signer{

			PublicKey:      "0x123",
			ShortPublicKey: "0x123",
		},
	}
	SignatureTwo = Signature{
		Signer: model.Signer{
			PublicKey:      "0x321",
			ShortPublicKey: "0x123",
		},
	}
	sampleSigner = model.Signer{
		Name:           "",
		EvmAddress:     "0x12345",
		PublicKey:      "0x321",
		ShortPublicKey: "0x123",
	}
)

var SampleCorrectness = []model.Correctness{
	{
		SignersCount: 1,
		Timestamp:    uint64(time.Now().Unix()),
		Topic:        utils.Shake([]byte("123")),
	},
	{
		SignersCount: 1,
		Timestamp:    uint64(time.Now().Unix()),
		Topic:        utils.Shake([]byte("1234")),
	},
}

type CorrectnessTestSuite struct {
	suite.Suite

	db       *embeddedpostgres.EmbeddedPostgres
	ins      database.Database
	cacheDir string
	service  Service
}

func (s *CorrectnessTestSuite) SetupTest() {
	utils.SetupLogger("info")
	cachePath := fmt.Sprintf("embedded-postgres-go-%s", uuid.NewString())
	cacheDir, err := os.MkdirTemp("", cachePath)
	s.Require().NoError(err)
	s.cacheDir = cacheDir

	s.db = embeddedpostgres.NewDatabase(
		embeddedpostgres.
			DefaultConfig().
			CachePath(s.cacheDir),
	)
	err = s.db.Start()
	s.Require().NoError(err)

	config.App.Postgres.URL = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	s.ins = postgres.New()
	s.ins.Migrate()

	posService := new(pos.MockService)
	posService.On("GetVotingPowerOfEvm", mock2.Anything, "0x12345").Return(10, nil)

	signerRepo := postgresRepo.NewProof(s.ins)
	correctnessRepo := postgresRepo.NewCorrectness(s.ins)

	config.App.Plugins.Correctness = []string{"123"}
	s.service = New(posService, signerRepo, correctnessRepo)
}

func (s *CorrectnessTestSuite) TestRecordSignatures() {
	s.Run("Should return topic not supported", func() {
		_, _, shortPublicKey := bls.GenerateBlsKeyPair()

		signature, err := bls.RecoverSignature(shortPublicKey.Bytes())
		s.Require().NoError(err)

		err = s.service.RecordSignature(
			context.TODO(),
			signature,
			sampleSigner,
			SampleCorrectness[1],
			false,
		)
		s.ErrorIs(err, consts.ErrTopicNotSupported)
	})

	s.Run("Should run without error", func() {
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
	})
}

func (s *CorrectnessTestSuite) TearDownSuite() {
	s.T().Log("Stopping the pg server")
	err := s.db.Stop()
	s.Require().NoError(err)
	os.RemoveAll(s.cacheDir)
}

func TestCorrectnessSuite(t *testing.T) {
	suite.Run(t, new(CorrectnessTestSuite))
}
