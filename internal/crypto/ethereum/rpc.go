package ethereum

import (
	"context"
	"fmt"
	"sync"

	"github.com/KenshiTech/unchained/internal/crypto/ethereum/contracts"

	"github.com/KenshiTech/unchained/internal/config"

	"github.com/KenshiTech/unchained/internal/log"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Repository struct {
	rpcList  map[string][]string
	rpcIndex map[string]int
	Clients  map[string]*ethclient.Client
	rpcMutex *sync.Mutex
}

func (r *Repository) refreshRPCWithRetries(network string, retries int) bool {
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

func (r *Repository) RefreshRPC(network string) {
	r.rpcMutex.Lock()
	defer r.rpcMutex.Unlock()
	r.refreshRPCWithRetries(network, len(r.rpcList))
}

func (r *Repository) GetNewStakingContract(
	network string,
	address string,
	refresh bool) (*contracts.UnchainedStaking, error) {
	if refresh {
		r.RefreshRPC(network)
	}

	return contracts.NewUnchainedStaking(common.HexToAddress(address), r.Clients[network])
}

func (r *Repository) GetNewUniV3Contract(network string, address string, refresh bool) (*contracts.UniV3, error) {
	if refresh {
		r.RefreshRPC(network)
	}

	return contracts.NewUniV3(common.HexToAddress(address), r.Clients[network])
}

func (r *Repository) GetBlockNumber(network string) (uint64, error) {
	client, ok := r.Clients[network]

	if !ok {
		log.Logger.With("Network", network).Error("Client not found")
		return 0, fmt.Errorf("client not found")
	}

	return client.BlockNumber(context.Background())
}

func (r *Repository) init() {
	r.rpcList = make(map[string][]string)
	r.rpcIndex = make(map[string]int)
	r.Clients = make(map[string]*ethclient.Client)
	r.rpcMutex = new(sync.Mutex)
}

func New() *Repository {
	r := &Repository{}
	r.init()

	for _, rpc := range config.App.RPC {
		r.rpcIndex[rpc.Name] = 0
		r.rpcList[rpc.Name] = append(r.rpcList[rpc.Name], rpc.Nodes...)
		r.RefreshRPC(rpc.Name)
	}

	return r
}
