package ethereum

import (
	"context"
	"math/big"
	"sync"

	goEthereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/utils"

	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum/contracts"

	"github.com/TimeleapLabs/unchained/internal/config"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

type RPC interface {
	RefreshRPC(network string)
	GetNewStakingContract(network string, address string, refresh bool) (*contracts.ProofOfStake, error)
	GetNewUniV3Contract(network string, address string, refresh bool) (*contracts.UniV3, error)

	GetBlockNumber(ctx context.Context, network string) (uint64, error)
	GetLogs(ctx context.Context, chain string, from, to *big.Int, addresses []common.Address) ([]types.Log, error)
}

type repository struct {
	list    map[string][]string
	index   map[string]int
	clients map[string]*ethclient.Client
	mutex   *sync.Mutex
}

func (r *repository) GetLogs(ctx context.Context, chain string, from, to *big.Int, addresses []common.Address) ([]types.Log, error) {
	client, isFound := r.clients[chain]
	if !isFound {
		utils.Logger.With("Network", chain).Error("Client not found")
		return nil, consts.ErrClientNotFound
	}

	logs, err := client.FilterLogs(ctx, goEthereum.FilterQuery{
		FromBlock: from,
		ToBlock:   to,
		Addresses: addresses,
	})
	if err != nil {
		return nil, err
	}

	return logs, nil
}

func (r *repository) refreshRPCWithRetries(network string, retries int) bool {
	if retries == 0 {
		panic("Cannot connect to any of the provided RPCs")
	}

	if r.index[network] == len(r.list[network])-1 {
		r.index[network] = 0
	} else {
		r.index[network]++
	}

	var err error

	index := r.index[network]
	r.clients[network], err = ethclient.Dial(r.list[network][index])

	if err != nil {
		return r.refreshRPCWithRetries(network, retries-1)
	}

	return true
}

func (r *repository) RefreshRPC(network string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	utils.Logger.With("Network", network).Info("Connecting to RPC")
	r.refreshRPCWithRetries(network, len(r.list))
}

func (r *repository) GetNewStakingContract(network string, address string, refresh bool) (*contracts.ProofOfStake, error) {
	if refresh {
		r.RefreshRPC(network)
	}

	client, isFound := r.clients[network]
	if !isFound {
		utils.Logger.With("Network", network).Error("Client not found")
		return nil, consts.ErrClientNotFound
	}

	return contracts.NewProofOfStake(common.HexToAddress(address), client)
}

func (r *repository) GetNewUniV3Contract(network string, address string, refresh bool) (*contracts.UniV3, error) {
	if refresh {
		r.RefreshRPC(network)
	}

	client, isFound := r.clients[network]
	if !isFound {
		utils.Logger.With("Network", network).Error("Client not found")
		return nil, consts.ErrClientNotFound
	}

	return contracts.NewUniV3(common.HexToAddress(address), client)
}

// GetBlockNumber returns the most recent block number.
func (r *repository) GetBlockNumber(ctx context.Context, network string) (uint64, error) {
	client, isFound := r.clients[network]
	if !isFound {
		utils.Logger.With("Network", network).Error("Client not found")
		return 0, consts.ErrClientNotFound
	}

	return client.BlockNumber(ctx)
}

func New() RPC {
	r := &repository{
		list:    map[string][]string{},
		index:   map[string]int{},
		clients: make(map[string]*ethclient.Client),
		mutex:   new(sync.Mutex),
	}

	for _, rpc := range config.App.RPC {
		r.index[rpc.Name] = 0
		r.list[rpc.Name] = append(r.list[rpc.Name], rpc.Nodes...)
		r.RefreshRPC(rpc.Name)
	}

	return r
}
