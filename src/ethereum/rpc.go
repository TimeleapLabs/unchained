package ethereum

import (
	"context"
	"fmt"
	"reflect"
	"sync"

	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/ethereum/contracts"
	"github.com/KenshiTech/unchained/log"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

var rpcList map[string][]string
var rpcIndex map[string]int
var Clients map[string]*ethclient.Client
var rpcMutex *sync.Mutex

func Start() {
	rpcConf := config.Config.Sub("rpc")
	networkNames := rpcConf.AllKeys()

	for _, name := range networkNames {
		conf := rpcConf.Get(name)
		rpcIndex[name] = 0

		switch reflect.TypeOf(conf).Kind() {
		case reflect.String:
			rpcList[name] = append(rpcList[name], conf.(string))

		case reflect.Slice:
			for _, rpc := range conf.([]interface{}) {
				rpcList[name] = append(rpcList[name], rpc.(string))
			}
		default:
			panic("RPC List Is Invalid")
		}

		RefreshRPC(name)
	}
}

func refreshRPCWithRetries(network string, retries int) bool {
	if retries == 0 {
		panic("Cannot connect to any of the provided RPCs")
	}

	if rpcIndex[network] == len(rpcList[network])-1 {
		rpcIndex[network] = 0
	} else {
		rpcIndex[network]++
	}

	var err error

	index := rpcIndex[network]
	Clients[network], err = ethclient.Dial(rpcList[network][index])

	if err != nil {
		return refreshRPCWithRetries(network, retries-1)
	}

	return true
}

func RefreshRPC(network string) {
	rpcMutex.Lock()
	defer rpcMutex.Unlock()
	refreshRPCWithRetries(network, len(rpcList))
}

func GetNewStakingContract(
	network string,
	address string,
	refresh bool) (*contracts.UnchainedStaking, error) {
	if refresh {
		RefreshRPC(network)
	}

	return contracts.NewUnchainedStaking(common.HexToAddress(address), Clients[network])
}

func GetNewUniV3Contract(network string, address string, refresh bool) (*contracts.UniV3, error) {
	if refresh {
		RefreshRPC(network)
	}

	return contracts.NewUniV3(common.HexToAddress(address), Clients[network])
}

func GetBlockNumber(network string) (uint64, error) {
	client, ok := Clients[network]

	if !ok {
		log.Logger.With("Network", network).Error("Client not found")
		return 0, fmt.Errorf("client not found")
	}

	return client.BlockNumber(context.Background())
}

func init() {
	rpcList = make(map[string][]string)
	rpcIndex = make(map[string]int)
	Clients = make(map[string]*ethclient.Client)
	rpcMutex = new(sync.Mutex)
}
