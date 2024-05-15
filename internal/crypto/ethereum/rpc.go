package ethereum

import (
	"context"
	"sync"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/utils"

	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum/contracts"

	"github.com/TimeleapLabs/unchained/internal/config"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

type RPC interface {
	RefreshRPC(network string)
	GetClient(network string) *ethclient.Client
	GetNewStakingContract(network string, address string, refresh bool) (*contracts.UnchainedStaking, error)
	GetNewUniV3Contract(network string, address string, refresh bool) (*contracts.UniV3, error)
	GetBlockNumber(ctx context.Context, network string) (uint64, error)
}

type repository struct {
	list    map[string][]string
	index   map[string]int
	clients map[string]*ethclient.Client
	mutex   *sync.Mutex
}

func (r *repository) GetClient(chain string) *ethclient.Client {
	client, isExist := r.clients[chain]
	if !isExist {
		utils.Logger.With("Network", chain).Error("Client not found")
		return nil
	}

	return client
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

func (r *repository) GetNewStakingContract(network string, address string, refresh bool) (*contracts.UnchainedStaking, error) {
	if refresh {
		r.RefreshRPC(network)
	}

	client := r.GetClient(network)
	if client == nil {
		return nil, consts.ErrClientNotFound
	}

	return contracts.NewUnchainedStaking(common.HexToAddress(address), client)
}

func (r *repository) GetNewUniV3Contract(network string, address string, refresh bool) (*contracts.UniV3, error) {
	if refresh {
		r.RefreshRPC(network)
	}

	client := r.GetClient(network)
	if client == nil {
		return nil, consts.ErrClientNotFound
	}

	return contracts.NewUniV3(common.HexToAddress(address), client)
}

// GetBlockNumber returns the most recent block number.
func (r *repository) GetBlockNumber(ctx context.Context, network string) (uint64, error) {
	client := r.GetClient(network)
	if client == nil {
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
