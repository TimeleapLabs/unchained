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

var SampleAssetPrice = model.AssetPrice{
	Pair:         "USDT/ETH",
	Name:         "USDT",
	Chain:        "ETH",
	Block:        999,
	Price:        1000,
	SignersCount: 10,
	Signature:    nil,
	Consensus:    false,
	Voted:        1000,
}

type AssetPriceRepositoryTestSuite struct {
	suite.Suite
	dbServer *memongo.Server
	repo     repository.AssetPrice
}

func (s *AssetPriceRepositoryTestSuite) SetupTest() {
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
	s.repo = NewAssetPrice(db)
}

func (s *AssetPriceRepositoryTestSuite) TestUpsert() {
	s.Run("Upsert asset price", func() {
		err := s.repo.Upsert(context.TODO(), SampleAssetPrice)
		s.NoError(err)
	})
}

func (s *AssetPriceRepositoryTestSuite) TestFind() {
	s.Run("Find asset price", func() {
		err := s.repo.Upsert(context.TODO(), SampleAssetPrice)
		s.NoError(err)

		assetPrices, err := s.repo.Find(context.TODO(), SampleAssetPrice.Block, SampleAssetPrice.Chain, SampleAssetPrice.Name, SampleAssetPrice.Pair)
		s.NoError(err)
		s.Len(assetPrices, 1)
	})
}

func (s *AssetPriceRepositoryTestSuite) TearDownSuite() {
	s.T().Log("Stopping the mongo server")
	s.dbServer.Stop()
}

func TestAssetPriceRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AssetPriceRepositoryTestSuite))
}
