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
	"github.com/TimeleapLabs/unchained/internal/utils/hash"
	"github.com/stretchr/testify/suite"
	"github.com/tryvium-travels/memongo"
)

var sampleAttestation = model.Attestation{
	Timestamp: 999,
	Topic:     []byte{},
	Meta: map[string]interface{}{
		"Correct": true,
	},
}

type AttestationRepositoryTestSuite struct {
	suite.Suite
	dbServer *memongo.Server
	repo     repository.Attestation
}

func (s *AttestationRepositoryTestSuite) SetupTest() {
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
	s.repo = NewAttestation(db)
}

func (s *AttestationRepositoryTestSuite) TestUpsert() {
	err := s.repo.Upsert(context.TODO(), hash.Hash(&sampleAttestation), sampleAttestation)
	s.NoError(err)
}

func (s *AttestationRepositoryTestSuite) TestFind() {
	err := s.repo.Upsert(context.TODO(), hash.Hash(&sampleAttestation), sampleAttestation)
	s.NoError(err)

	hashed := hash.Hash(&sampleAttestation)
	_, err = s.repo.Find(context.TODO(), hashed)
	s.NoError(err)
}

func (s *AttestationRepositoryTestSuite) TearDownTest() {
	if s.dbServer != nil {
		s.dbServer.Stop()
	}
}

func TestAttestationRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AttestationRepositoryTestSuite))
}
