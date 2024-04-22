package pos

import (
	"testing"

	"github.com/KenshiTech/unchained/internal/crypto"
	"github.com/KenshiTech/unchained/internal/crypto/ethereum"
	"github.com/KenshiTech/unchained/internal/utils"
	"github.com/stretchr/testify/suite"
)

type PosTestSuite struct {
	suite.Suite
	service Service
}

func (s *PosTestSuite) SetupTest() {
	utils.SetupLogger("info")
	crypto.InitMachineIdentity(
		crypto.WithBlsIdentity(),
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
	_, err := s.service.GetVotingPower([20]byte{}, nil)
	s.NoError(err)
}

func (s *PosTestSuite) TestGetVotingPowerFromContract() {
	_, err := s.service.GetVotingPowerFromContract([20]byte{}, nil)
	s.NoError(err)
}

func (s *PosTestSuite) TestGetVotingPowerOfPublicKey() {
	_, err := s.service.GetVotingPowerOfPublicKey([96]byte{})
	s.NoError(err)
}

func (s *PosTestSuite) TestVotingPowerToFloat() {
	// := s.service.VotingPowerToFloat(nil)
	// s.NoError(err)
}

func TestPosTestSuite(t *testing.T) {
	suite.Run(t, new(PosTestSuite))
}
