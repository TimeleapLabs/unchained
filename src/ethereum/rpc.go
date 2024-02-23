package ethereum

import (
	"context"
	"reflect"

	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/ethereum/contracts"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

var rpcList map[string][]string
var rpcIndex map[string]int
var Clients map[string]*ethclient.Client

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
	refreshRPCWithRetries(network, len(rpcList))
}

func GetNewUniV3Contract(network string, address string, refresh bool) (*contracts.UniV3, error) {
	if refresh {
		RefreshRPC(network)
	}

	return contracts.NewUniV3(common.HexToAddress(address), Clients[network])
}

func GetBlockNumber(network string) (uint64, error) {
	return Clients[network].BlockNumber(context.Background())
}

func init() {
	rpcList = make(map[string][]string)
	rpcIndex = make(map[string]int)
	Clients = make(map[string]*ethclient.Client)
}
