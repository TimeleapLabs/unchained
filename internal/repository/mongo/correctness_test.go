package mongo

import (
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/repository"
	"github.com/TimeleapLabs/unchained/internal/transport/database/mongo"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/stretchr/testify/suite"
	"github.com/tryvium-travels/memongo"
	"log"
	"math/big"
	"runtime"
	"testing"
)

var sampleCorrectness = model.Correctness{
	SignersCount: 100,
	Signature:    nil,
	Consensus:    false,
	Voted:        *big.NewInt(1000),
	SignerIDs:    nil,
	Timestamp:    999,
	Hash:         nil,
	Topic:        [64]byte{},
	Correct:      false,
}

type CorrectnessRepositoryTestSuite struct {
	suite.Suite
	dbServer *memongo.Server
	repo     repository.CorrectnessReport
}

func (s *CorrectnessRepositoryTestSuite) SetupTest() {
	utils.SetupLogger(config.App.System.Log)

	var err error
	opts := &memongo.Options{
		MongoVersion: "5.0.0",
	}
	if runtime.GOARCH == "arm64" {
		if runtime.GOOS == "darwin" {
			opts.DownloadURL = "https://fastdl.mongodb.org/osx/mongodb-macos-x86_64-5.0.0.tgz"
		}
	}

	s.dbServer, err = memongo.StartWithOptions(opts)
	if err != nil {
		log.Fatal(err)
	}

	config.App.Mongo.URL = s.dbServer.URI()
	config.App.Mongo.Database = memongo.RandomDatabase()
	db := mongo.New()
	s.repo = NewCorrectness(db)
}

func (s *CorrectnessRepositoryTestSuite) TestUpsert() {
	err := s.repo.Upsert(nil, sampleCorrectness)
	s.NoError(err)
}

func (s *CorrectnessRepositoryTestSuite) TestFind() {
	err := s.repo.Upsert(nil, sampleCorrectness)
	s.NoError(err)

	result, err := s.repo.Find(nil, sampleCorrectness.Hash, sampleCorrectness.Topic[:], sampleCorrectness.Timestamp)
	s.NoError(err)
	s.Len(result, 1)
}

func (s *CorrectnessRepositoryTestSuite) TearDownTest() {
	if s.dbServer != nil {
		s.dbServer.Stop()
	}
}

func TestCorrectnessRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CorrectnessRepositoryTestSuite))
}
