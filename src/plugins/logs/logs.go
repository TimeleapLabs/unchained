package logs

import (
	"context"
	"math/big"
	"os"
	"sort"
	"time"

	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/ethereum"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/net/client"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	goEthereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-co-op/gocron/v2"
)

// var lastSynced map[string]uint64
var abiMap map[string]abi.ABI
var lastBlock map[string]uint64

type LogConf struct {
	Name    string `mapstructure:"name"`
	Chain   string `mapstructure:"chain"`
	Abi     string `mapstructure:"abi"`
	Event   string `mapstructure:"event"`
	Address string `mapstructure:"address"`
	From    uint64 `mapstructure:"from"`
}

func GetBlockNumber(network string) (*uint64, error) {
	blockNumber, err := ethereum.GetBlockNumber(network)

	if err != nil {
		ethereum.RefreshRPC(network)
		return nil, err
	}

	return &blockNumber, nil
}

type Event struct {
	Key   [32]byte
	Value [32]byte
}

func Start() {

	scheduler, err := gocron.NewScheduler()

	if err != nil {
		panic(err)
	}

	var configs []LogConf
	if err := config.Config.UnmarshalKey("plugins.logs", &configs); err != nil {
		panic(err)
	}

	for _, conf := range configs {
		file, err := os.Open(conf.Abi)
		if err != nil {
			panic(err)
		}

		contractAbi, err := abi.JSON(file)
		if err != nil {
			panic(err)
		}

		abiMap[conf.Abi] = contractAbi
		file.Close()
	}

	caser := cases.Title(language.English, cases.NoLower)

	_, err = scheduler.NewJob(
		gocron.DurationJob(5*time.Second),
		gocron.NewTask(
			func() {

				if client.IsClientSocketClosed {
					return
				}

				for _, conf := range configs {

					blockNumber, err := GetBlockNumber(conf.Chain)

					if err != nil {
						return
					}

					if lastBlock[conf.Chain] == *blockNumber {
						return
					}

					lastBlock[conf.Chain] = *blockNumber

					contractAddress := common.HexToAddress(conf.Address)

					query := goEthereum.FilterQuery{
						FromBlock: big.NewInt(int64(lastBlock[conf.Chain])),
						ToBlock:   big.NewInt(int64(lastBlock[conf.Chain])),
						Addresses: []common.Address{contractAddress},
					}

					client := ethereum.Clients[conf.Chain]
					logs, err := client.FilterLogs(context.Background(), query)

					if err != nil {
						panic(err)
					}

					contractAbi := abiMap[conf.Abi]

					for _, vLog := range logs {
						eventSignature := vLog.Topics[0]
						eventAbi, err := contractAbi.EventByID(eventSignature)

						if eventAbi.Name != conf.Event {
							continue
						}

						if err != nil {
							panic(err)
						}

						// Unpack the log's data using the event's name
						// This gives you a map of the event's arguments
						eventData := make(map[string]interface{})
						err = contractAbi.UnpackIntoMap(eventData, eventAbi.Name, vLog.Data)
						if err != nil {
							panic(err)
						}

						for i, arg := range eventAbi.Inputs {
							if arg.Indexed {
								switch arg.Type.String() {
								case "address":
									if len(vLog.Topics) > i {
										eventData[arg.Name] = common.BytesToAddress(vLog.Topics[i+1].Bytes()).Hex()
									}
								case "uint256", "uint8", "uint16", "uint32", "uint64":
									if len(vLog.Topics) > i {
										num := new(big.Int).SetBytes(vLog.Topics[i+1][:])
										eventData[arg.Name] = num
									}
								}
								// TODO: Add support for more types
							}
						}

						var keys []string
						for k := range eventData {
							keys = append(keys, k)
						}

						message := log.Logger.
							With("Event", conf.Event).
							With("Block", lastBlock[conf.Chain])

						sort.Strings(keys)
						for _, key := range keys {
							message = message.
								With(caser.String(key), eventData[key])
						}

						message.Info(conf.Name)
					}
				}
			},
		),
	)

	if err != nil {
		panic(err)
	}

	scheduler.Start()
}

func init() {
	abiMap = make(map[string]abi.ABI)
	lastBlock = make(map[string]uint64)
}
