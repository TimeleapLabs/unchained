package ethereum

import (
	"context"
	"log"
	"reflect"

	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/ethereum/contracts"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

var rpcList []string
var rpcIndex int
var Client *ethclient.Client

func Start() {
	rpcConfig := config.Config.Get("rpc.ethereum")
	rpcIndex = 0

	switch reflect.TypeOf(rpcConfig).Kind() {
	case reflect.String:
		rpcList = append(rpcList, rpcConfig.(string))

	case reflect.Slice:
		for _, rpc := range rpcConfig.([]interface{}) {
			rpcList = append(rpcList, rpc.(string))
		}
	default:
		panic("RPC List Is Invalid")
	}

	RefreshRPC()
}

func refreshRPCWithRetries(retries int) bool {
	if retries == 0 {
		log.Fatal("Cannot connect to any of the provided RPCs")
	}

	if rpcIndex == len(rpcList)-1 {
		rpcIndex = 0
	} else {
		rpcIndex++
	}

	var err error

	Client, err = ethclient.Dial(rpcList[rpcIndex])

	if err != nil {
		return refreshRPCWithRetries(retries - 1)
	}

	return true
}

func RefreshRPC() {
	refreshRPCWithRetries(len(rpcList))
}

func GetNewUniV3Contract(address string, refresh bool) (*contracts.UniV3, error) {
	if refresh {
		RefreshRPC()
	}

	return contracts.NewUniV3(common.HexToAddress(address), Client)
}

func GetBlockNumber() (uint64, error) {
	return Client.BlockNumber(context.Background())
}
