package pos

import (
	"context"
	"math/big"
	"testing"

	"github.com/TimeleapLabs/timeleap/internal/config"
	"github.com/TimeleapLabs/timeleap/internal/crypto"
	"github.com/TimeleapLabs/timeleap/internal/crypto/ethereum"
	"github.com/TimeleapLabs/timeleap/internal/utils"
	"github.com/stretchr/testify/suite"
)

type PosTestSuite struct {
	suite.Suite
	service Service
}

func (s *PosTestSuite) SetupTest() {
	config.App.System.AllowGenerateSecrets = true

	utils.SetupLogger("info")
	crypto.InitMachineIdentity(
		crypto.WithEd25519Identity(),
		crypto.WithEvmSigner(),
	)

	ethRPC := ethereum.NewMock()
	s.service = New(ethRPC)
}

func (s *PosTestSuite) TestGetTotalVotingPower() {
	_, err := s.service.GetTotalVotingPower()
	s.NoError(err)
}

func (s *PosTestSuite) TestGetVotingPower() {
	_, err := s.service.GetVotingPower([20]byte{}, big.NewInt(1000))
	s.NoError(err)
}

func (s *PosTestSuite) TestGetVotingPowerFromContract() {
	_, err := s.service.GetVotingPowerFromContract([20]byte{}, big.NewInt(1000))
	s.Error(err)
}

func (s *PosTestSuite) TestGetVotingPowerOfPublicKey() {
	_, err := s.service.GetVotingPowerOfPublicKey(context.TODO(), [96]byte{})
	s.NoError(err)
}

func TestPosSuite(t *testing.T) {
	suite.Run(t, new(PosTestSuite))
}
