package evmlog

import (
	"context"
	"fmt"
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"
	"github.com/TimeleapLabs/unchained/internal/model"
	postgresRepo "github.com/TimeleapLabs/unchained/internal/repository/postgres"
	"github.com/TimeleapLabs/unchained/internal/service/pos"
	"github.com/TimeleapLabs/unchained/internal/transport/database"
	"github.com/TimeleapLabs/unchained/internal/transport/database/postgres"
	"github.com/TimeleapLabs/unchained/internal/utils"
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/google/uuid"
	mock2 "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"os"
	"path"
	"runtime"
	"testing"
)

var (
	sampleSigner = model.Signer{
		Name:           "",
		EvmAddress:     "0x12345",
		PublicKey:      "0x321",
		ShortPublicKey: "0x123",
	}

	sampleEventLog = model.EventLog{
		LogIndex:     1,
		Block:        1,
		Address:      "0x12345",
		Event:        "event",
		Chain:        "eth",
		TxHash:       make([]byte, 32),
		Args:         nil,
		Consensus:    false,
		SignersCount: 0,
		Signature:    nil,
		Voted:        0,
	}
)

type EvmLogTestSuite struct {
	suite.Suite

	db       *embeddedpostgres.EmbeddedPostgres
	ins      database.Database
	cacheDir string
	service  Service
}

func (s *EvmLogTestSuite) SetupTest() {
	utils.SetupLogger("info")
	cachePath := fmt.Sprintf("embedded-postgres-go-%s", uuid.NewString())
	cacheDir, err := os.MkdirTemp("", cachePath)
	cacheContextDir, err := os.MkdirTemp("", fmt.Sprintf("context-test-%s", uuid.NewString()))
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
	posService.On("GetBlockNumber", mock2.Anything, "eth").Return(10, nil)
	posService.On("GetVotingPowerOfEvm", mock2.Anything, "0x12345").Return(10, nil)

	ethRPC := ethereum.NewMock()
	badger := NewBadger(cacheContextDir)
	proofRepo := postgresRepo.NewProof(s.ins)
	evmlogRepo := postgresRepo.NewEventLog(s.ins)

	_, testFilePath, _, _ := runtime.Caller(0)

	config.App.Plugins.EthLog = &config.EthLog{
		Events: []config.Event{
			{
				Name:    "event",
				Chain:   "eth",
				Event:   "event",
				Address: "0x12345",
				Abi:     path.Join(path.Dir(testFilePath), "../../abi/UniV3.json"),
			},
		},
	}
	s.service = New(ethRPC, posService, evmlogRepo, proofRepo, badger)
}

func (s *EvmLogTestSuite) TestRecordSignatures() {
	_, _, shortPublicKey := bls.GenerateBlsKeyPair()

	s.Run("Should return event not supported", func() {
		signature, err := bls.RecoverSignature(shortPublicKey.Bytes())
		s.Require().NoError(err)

		sampleEventLog := sampleEventLog
		sampleEventLog.Event = "not_supported_event"
		err = s.service.RecordSignature(context.TODO(), signature, sampleSigner, sampleEventLog, false, false)
		s.ErrorIs(err, consts.ErrEventNotSupported)
	})

	s.Run("Should return data is too old", func() {
		signature, err := bls.RecoverSignature(shortPublicKey.Bytes())
		s.Require().NoError(err)

		sampleEventLog := sampleEventLog
		sampleEventLog.Block = 100
		err = s.service.RecordSignature(context.TODO(), signature, sampleSigner, sampleEventLog, false, false)
		s.ErrorIs(err, consts.ErrDataTooOld)
	})

	s.Run("Should successfully", func() {
		signature, err := bls.RecoverSignature(shortPublicKey.Bytes())
		s.Require().NoError(err)

		sampleEventLog := sampleEventLog
		sampleEventLog.Block = 950
		err = s.service.RecordSignature(context.TODO(), signature, sampleSigner, sampleEventLog, false, false)
		s.Require().NoError(err)
	})
}

func (s *EvmLogTestSuite) TearDownSuite() {
	s.T().Log("Stopping the pg server")
	err := s.db.Stop()
	s.Require().NoError(err)
	os.RemoveAll(s.cacheDir)
}

func TestEvmLogTestSuite(t *testing.T) {
	suite.Run(t, new(EvmLogTestSuite))
}
