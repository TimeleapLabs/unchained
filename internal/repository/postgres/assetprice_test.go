package postgres

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/repository"
	"github.com/TimeleapLabs/unchained/internal/transport/database"
	"github.com/TimeleapLabs/unchained/internal/transport/database/postgres"
	"github.com/TimeleapLabs/unchained/internal/utils"
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
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

	db       *embeddedpostgres.EmbeddedPostgres
	ins      database.Database
	cacheDir string
	repo     repository.AssetPrice
}

func (s *AssetPriceRepositoryTestSuite) SetupTest() {
	utils.SetupLogger(config.App.System.Log)

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
	s.repo = NewAssetPrice(s.ins)
}

func (s *AssetPriceRepositoryTestSuite) TestUpsert() {
	err := s.repo.Upsert(context.TODO(), SampleAssetPrice)
	s.Require().NoError(err)
}

func (s *AssetPriceRepositoryTestSuite) TestFind() {
	err := s.repo.Upsert(context.TODO(), SampleAssetPrice)
	s.Require().NoError(err)

	assetPrices, err := s.repo.Find(context.TODO(), SampleAssetPrice.Block, SampleAssetPrice.Chain, SampleAssetPrice.Name, SampleAssetPrice.Pair)
	s.Require().NoError(err)
	s.Len(assetPrices, 1)
}

func (s *AssetPriceRepositoryTestSuite) TearDownSuite() {
	s.T().Log("Stopping the pg server")
	err := s.db.Stop()
	s.Require().NoError(err)
	os.RemoveAll(s.cacheDir)
}

func TestAssetPriceRepositoryTestSuite(t *testing.T) {
	suite.Run(t, &AssetPriceRepositoryTestSuite{})
}
