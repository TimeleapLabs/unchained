package ethereum

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
)

type mockRPC struct {
	backend *backends.SimulatedBackend
}

func (m mockRPC) GetLogs(_ context.Context, _ string, _, _ *big.Int, _ []common.Address) ([]types.Log, error) {
	return []types.Log{}, nil
}

func (m mockRPC) RefreshRPC(_ string) {}

func (m mockRPC) GetNewStakingContract(_ string, address string, _ bool) (*contracts.ProofOfStake, error) {
	return contracts.NewProofOfStake(
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
