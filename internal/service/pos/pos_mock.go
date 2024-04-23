package pos

import (
	"context"
	"github.com/stretchr/testify/mock"
	"math/big"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) GetTotalVotingPower() (*big.Int, error) {
	args := m.Called()
	return big.NewInt(int64(args.Int(0))), args.Error(1)
}

func (m *MockService) GetVotingPowerFromContract(address [20]byte, block *big.Int) (*big.Int, error) {
	args := m.Called(address, block)
	return big.NewInt(int64(args.Int(0))), args.Error(1)
}

func (m *MockService) GetVotingPower(address [20]byte, block *big.Int) (*big.Int, error) {
	args := m.Called(address, block)
	return big.NewInt(int64(args.Int(0))), args.Error(1)
}

func (m *MockService) GetVotingPowerOfPublicKey(_ context.Context, pkBytes [96]byte) (*big.Int, error) {
	args := m.Called(pkBytes)
	return big.NewInt(int64(args.Int(0))), args.Error(1)
}
