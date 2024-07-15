package uniswap

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
	"github.com/TimeleapLabs/unchained/internal/service/uniswap/types"
	"github.com/TimeleapLabs/unchained/internal/transport/database"
	"github.com/TimeleapLabs/unchained/internal/transport/database/postgres"
	"github.com/TimeleapLabs/unchained/internal/utils"
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/google/uuid"
	mock2 "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"math/big"
	"os"
	"testing"
)

var (
	sampleSigner = model.Signer{
		Name:           "",
		EvmAddress:     "0x12345",
		PublicKey:      "0x321",
		ShortPublicKey: "0x123",
	}

	samplePriceInfo = types.PriceInfo{
		Asset: types.AssetKey{
			Token: types.TokenKey{
				Name:   "ethereum",
				Pair:   "0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640",
				Chain:  "ethereum",
				Delta:  12,
				Invert: true,
				Cross:  string(utils.Shake(types.TokenKeys([]types.TokenKey{}).Sia().Bytes())),
			},
			Block: 100,
		},
		Price: *big.NewInt(1000),
	}
)

type UniswapTestSuite struct {
	suite.Suite

	db       *embeddedpostgres.EmbeddedPostgres
	ins      database.Database
	cacheDir string
	service  Service
}

func (s *UniswapTestSuite) SetupTest() {
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
	posService.On("GetBlockNumber", mock2.Anything, "eth").Return(10, nil)
	posService.On("GetVotingPowerOfEvm", mock2.Anything, "0x12345").Return(10, nil)

	ethRPC := ethereum.NewMock()
	assetPriceRepo := postgresRepo.NewAssetPrice(s.ins)
	proofRepo := postgresRepo.NewProof(s.ins)

	config.App.Plugins.Uniswap = &config.Uniswap{
		Tokens: []config.Token{
			{
				Name:   "ethereum",
				Pair:   "0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640",
				Chain:  "ethereum",
				Delta:  12,
				Invert: true,
				Unit:   "USDT",
				Send:   true,
			},
		},
	}
	s.service = New(ethRPC, posService, proofRepo, assetPriceRepo)
}

func (s *UniswapTestSuite) TestRecordSignature() {
	_, _, shortPublicKey := bls.GenerateBlsKeyPair()

	s.Run("Should return token not supported", func() {
		signature, err := bls.RecoverSignature(shortPublicKey.Bytes())
		s.Require().NoError(err)

		samplePriceInfo := samplePriceInfo
		samplePriceInfo.Asset.Token.Name += "invalid"

		err = s.service.RecordSignature(context.TODO(), signature, sampleSigner, samplePriceInfo, false, false)
		s.ErrorIs(err, consts.ErrTokenNotSupported)
	})

	s.Run("Should return data is too old", func() {
		signature, err := bls.RecoverSignature(shortPublicKey.Bytes())
		s.Require().NoError(err)

		samplePriceInfo := samplePriceInfo
		samplePriceInfo.Asset.Block = 100

		err = s.service.RecordSignature(context.TODO(), signature, sampleSigner, samplePriceInfo, false, false)
		s.ErrorIs(err, consts.ErrDataTooOld)
	})

	s.Run("Should run successfully", func() {
		signature, err := bls.RecoverSignature(shortPublicKey.Bytes())
		s.Require().NoError(err)

		samplePriceInfo := samplePriceInfo
		samplePriceInfo.Asset.Block = 950

		err = s.service.RecordSignature(context.TODO(), signature, sampleSigner, samplePriceInfo, false, false)
		s.Require().NoError(err)
	})
}

func (s *UniswapTestSuite) TearDownSuite() {
	s.T().Log("Stopping the pg server")
	err := s.db.Stop()
	s.Require().NoError(err)
	os.RemoveAll(s.cacheDir)
}

func TestEvmLogTestSuite(t *testing.T) {
	suite.Run(t, new(UniswapTestSuite))
}
