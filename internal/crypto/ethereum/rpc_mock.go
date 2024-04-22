package ethereum

import (
	"github.com/KenshiTech/unchained/internal/crypto/ethereum/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/ethclient"
)

type mockRpc struct {
	backend *backends.SimulatedBackend
}

func (m mockRpc) GetClient(network string) *ethclient.Client {
	//TODO implement me
	panic("implement me")
}

func (m mockRpc) RefreshRPC(_ string) {}

func (m mockRpc) GetNewStakingContract(_ string, address string, _ bool) (*contracts.UnchainedStaking, error) {
	return contracts.NewUnchainedStaking(
		common.HexToAddress(address),
		m.backend,
	)
}

func (m mockRpc) GetNewUniV3Contract(_ string, address string, _ bool) (*contracts.UniV3, error) {
	return contracts.NewUniV3(
		common.HexToAddress(address),
		m.backend,
	)
}

func (m mockRpc) GetBlockNumber(_ string) (uint64, error) {
	var blockNumber uint64 = 1000
	return blockNumber, nil
}

func NewMock() Rpc {
	return &mockRpc{
		backend: backends.NewSimulatedBackend(
			core.DefaultGenesisBlock().Alloc,
			9000000,
		),
	}
}
