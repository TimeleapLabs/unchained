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

func (m mockRPC) GetClient(_ string) *ethclient.Client {
	// TODO implement me
	panic("implement me")
}

func (m mockRPC) RefreshRPC(_ string) {}

func (m mockRPC) GetNewStakingContract(_ string, address string, _ bool) (*contracts.UnchainedStaking, error) {
	return contracts.NewUnchainedStaking(
		common.HexToAddress(address),
		m.backend,
	)
}

func (m mockRPC) GetNewUniV3Contract(_ string, address string, _ bool) (*contracts.UniV3, error) {
	return contracts.NewUniV3(
		common.HexToAddress(address),
		m.backend,
	)
}

func (m mockRPC) GetBlockNumber(_ context.Context, _ string) (uint64, error) {
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
