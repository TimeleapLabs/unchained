package ethereum

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/ethclient"
)

type mockRPC struct {
	backend *backends.SimulatedBackend
}

func (m mockRPC) GetClient() *ethclient.Client {
	// TODO implement me
	panic("implement me")
}

func (m mockRPC) RefreshRPC() {}

func (m mockRPC) GetNewStakingContract(address string, _ bool) (*contracts.ProofOfStake, error) {
	return contracts.NewProofOfStake(
		common.HexToAddress(address),
		m.backend,
	)
}

func (m mockRPC) GetBlockNumber(_ context.Context) (uint64, error) {
	var blockNumber uint64 = 1000
	return blockNumber, nil
}

func NewMock() RPC {
	return &mockRPC{
		backend: backends.NewSimulatedBackend(
			core.DefaultGenesisBlock().Alloc,
			9000000,
		),
	}
}
