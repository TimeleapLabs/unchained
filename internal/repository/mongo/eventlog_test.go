package mongo

import (
	"log"
	"runtime"
	"testing"

	"golang.org/x/net/context"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/repository"
	"github.com/TimeleapLabs/unchained/internal/transport/database/mongo"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/stretchr/testify/suite"
	"github.com/tryvium-travels/memongo"
)

var sampleEventLog = model.EventLog{
	LogIndex:     999,
	Block:        999,
	Address:      "123",
	Event:        "321",
	Chain:        "ETH",
	TxHash:       [32]byte{},
	Args:         nil,
	Consensus:    false,
	SignersCount: 100,
	Signature:    nil,
	Voted:        0,
}

type EventLogRepositoryTestSuite struct {
	suite.Suite
	dbServer *memongo.Server
	repo     repository.EventLog
}

func (s *EventLogRepositoryTestSuite) SetupTest() {
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
	s.repo = NewEventLog(db)
}

func (s *EventLogRepositoryTestSuite) TestUpsert() {
	err := s.repo.Upsert(context.TODO(), sampleEventLog)
	s.NoError(err)
}

func (s *EventLogRepositoryTestSuite) TestFind() {
	err := s.repo.Upsert(context.TODO(), sampleEventLog)
	s.NoError(err)

	result, err := s.repo.Find(context.TODO(), sampleEventLog.Block, sampleEventLog.TxHash[:], sampleEventLog.LogIndex)
	s.NoError(err)
	s.Len(result, 1)
}

func (s *EventLogRepositoryTestSuite) TearDownTest() {
	if s.dbServer != nil {
		s.dbServer.Stop()
	}
}

func TestEventLogRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(EventLogRepositoryTestSuite))
}
