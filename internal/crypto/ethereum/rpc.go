package ethereum

import (
	"context"
	"sync"

	"github.com/KenshiTech/unchained/internal/utils"

	"github.com/KenshiTech/unchained/internal/consts"

	"github.com/KenshiTech/unchained/internal/crypto/ethereum/contracts"

	"github.com/KenshiTech/unchained/internal/config"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Rpc interface {
	RefreshRPC(network string)
	GetClient(network string) *ethclient.Client
	GetNewStakingContract(network string, address string, refresh bool) (*contracts.UnchainedStaking, error)
	GetNewUniV3Contract(network string, address string, refresh bool) (*contracts.UniV3, error)
	GetBlockNumber(network string) (uint64, error)
}

type repository struct {
	rpcList  map[string][]string
	rpcIndex map[string]int
	Clients  map[string]*ethclient.Client
	rpcMutex *sync.Mutex
}

func (r *repository) GetClient(chain string) *ethclient.Client {
	return r.Clients[chain]
}

func (r *repository) refreshRPCWithRetries(network string, retries int) bool {
	if retries == 0 {
		panic("Cannot connect to any of the provided RPCs")
	}

	if r.rpcIndex[network] == len(r.rpcList[network])-1 {
		r.rpcIndex[network] = 0
	} else {
		r.rpcIndex[network]++
	}

	var err error

	index := r.rpcIndex[network]
	r.Clients[network], err = ethclient.Dial(r.rpcList[network][index])

	if err != nil {
		return r.refreshRPCWithRetries(network, retries-1)
	}

	return true
}

func (r *repository) RefreshRPC(network string) {
	r.rpcMutex.Lock()
	defer r.rpcMutex.Unlock()
	r.refreshRPCWithRetries(network, len(r.rpcList))
}

func (r *repository) GetNewStakingContract(network string, address string, refresh bool) (*contracts.UnchainedStaking, error) {
	if refresh {
		r.RefreshRPC(network)
	}

	return contracts.NewUnchainedStaking(common.HexToAddress(address), r.Clients[network])
}

func (r *repository) GetNewUniV3Contract(network string, address string, refresh bool) (*contracts.UniV3, error) {
	if refresh {
		r.RefreshRPC(network)
	}

	return contracts.NewUniV3(common.HexToAddress(address), r.Clients[network])
}

// GetBlockNumber returns the most recent block number.
func (r *repository) GetBlockNumber(network string) (uint64, error) {
	client, ok := r.Clients[network]

	if !ok {
		utils.Logger.With("Network", network).Error("Client not found")
		return 0, consts.ErrClientNotFound
	}

	return client.BlockNumber(context.Background())
}

func (r *repository) init() {
	r.rpcList = make(map[string][]string)
	r.rpcIndex = make(map[string]int)
	r.Clients = make(map[string]*ethclient.Client)
	r.rpcMutex = new(sync.Mutex)
}

func New() Rpc {
	r := &repository{}
	r.init()

	for _, rpc := range config.App.RPC {
		r.rpcIndex[rpc.Name] = 0
		r.rpcList[rpc.Name] = append(r.rpcList[rpc.Name], rpc.Nodes...)
		r.RefreshRPC(rpc.Name)
	}

	return r
}
