package mongo

import (
	"context"
	"log"
	"runtime"
	"testing"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/repository"
	"github.com/TimeleapLabs/unchained/internal/transport/database/mongo"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/stretchr/testify/suite"
	"github.com/tryvium-travels/memongo"
)

var sampleSigner = model.Signer{}

type SignerRepositoryTestSuite struct {
	suite.Suite
	dbServer *memongo.Server
	repo     repository.Proof
}

func (s *SignerRepositoryTestSuite) SetupTest() {
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
	s.repo = NewProof(db)
}

func (s *SignerRepositoryTestSuite) TestUpsert() {
	err := s.repo.CreateProof(context.TODO(), [48]byte{}, []model.Signer{sampleSigner})
	s.Require().NoError(err)
}

func (s *SignerRepositoryTestSuite) TearDownTest() {
	if s.dbServer != nil {
		s.dbServer.Stop()
	}
}

func TestSignerRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(SignerRepositoryTestSuite))
}
