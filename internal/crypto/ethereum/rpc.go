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
	RefreshRPC()
	GetClient() *ethclient.Client
	GetNewStakingContract(address string, refresh bool) (*contracts.ProofOfStake, error)
	GetBlockNumber(ctx context.Context) (uint64, error)
}

type repository struct {
	list   []string
	index  int
	client *ethclient.Client
	mutex  *sync.Mutex
}

func (r *repository) GetClient() *ethclient.Client {
	if r.client == nil {
		utils.Logger.Error("PoS client not found")
		return nil
	}

	return r.client
}

func (r *repository) refreshRPCWithRetries(retries int) bool {
	if retries == 0 {
		panic("Cannot connect to any of the provided RPCs")
	}

	if r.index == len(r.list)-1 {
		r.index = 0
	} else {
		r.index++
	}

	var err error

	r.client, err = ethclient.Dial(r.list[r.index])

	if err != nil {
		return r.refreshRPCWithRetries(retries - 1)
	}

	return true
}

func (r *repository) RefreshRPC() {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	utils.Logger.Info("Connecting to the PoS RPC")
	r.refreshRPCWithRetries(len(r.list))
}

func (r *repository) GetNewStakingContract(address string, refresh bool) (*contracts.ProofOfStake, error) {
	if refresh {
		r.RefreshRPC()
	}

	client := r.GetClient()
	if client == nil {
		return nil, consts.ErrClientNotFound
	}

	return contracts.NewProofOfStake(common.HexToAddress(address), client)
}

// GetBlockNumber returns the most recent block number.
func (r *repository) GetBlockNumber(ctx context.Context) (uint64, error) {
	client := r.GetClient()
	if client == nil {
		return 0, consts.ErrClientNotFound
	}

	return client.BlockNumber(ctx)
}

func New() RPC {
	r := &repository{
		list:  config.App.ProofOfStake.RPC,
		index: 0,
		mutex: &sync.Mutex{},
	}

	r.RefreshRPC()

	return r
}
