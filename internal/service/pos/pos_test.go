package pos

import (
	"context"
	"math/big"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PosTestSuite struct {
	suite.Suite
	service *MockService
}

func (s *PosTestSuite) SetupTest() {
	s.service = NewMock()
}

func (s *PosTestSuite) TestGetTotalVotingPower() {
	s.service.On("GetTotalVotingPower").Return(0, nil)

	_, err := s.service.GetTotalVotingPower()
	s.NoError(err)
}

func (s *PosTestSuite) TestGetVotingPower() {
	s.service.On("GetVotingPower", [20]byte{}, big.NewInt(0)).Return(0, nil)

	_, err := s.service.GetVotingPower([20]byte{}, big.NewInt(0))
	s.NoError(err)
}

func (s *PosTestSuite) TestGetVotingPowerFromContract() {
	s.service.On("GetVotingPowerFromContract", [20]byte{}, big.NewInt(0)).Return(0, nil)

	_, err := s.service.GetVotingPowerFromContract([20]byte{}, big.NewInt(0))
	s.NoError(err)
}

func (s *PosTestSuite) TestGetVotingPowerOfPublicKey() {
	s.service.On("GetVotingPowerOfPublicKey", [96]byte{}).Return(0, nil)

	_, err := s.service.GetVotingPowerOfPublicKey(context.TODO(), [96]byte{})
	s.NoError(err)
}

func TestPosSuite(t *testing.T) {
	suite.Run(t, new(PosTestSuite))
}
