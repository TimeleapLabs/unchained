package logs

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"sort"
	"time"

	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/ethereum"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/net/client"
	"github.com/KenshiTech/unchained/persistence"
	"github.com/dgraph-io/badger/v4"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	goEthereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-co-op/gocron/v2"
)

type LogConf struct {
	Name          string  `mapstructure:"name"`
	Chain         string  `mapstructure:"chain"`
	Abi           string  `mapstructure:"abi"`
	Event         string  `mapstructure:"event"`
	Address       string  `mapstructure:"address"`
	From          *uint64 `mapstructure:"from"`
	Step          uint64  `mapstructure:"step"`
	Confrimations uint64  `mapstructure:"confirmations"`
}

// var lastSynced map[string]uint64
var abiMap map[string]abi.ABI
var lastSyncedBlock map[LogConf]uint64
var caser cases.Caser

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

func createTask(configs []LogConf, chain string) func() {
	return func() {

		if client.IsClientSocketClosed {
			return
		}

		for _, conf := range configs {

			if conf.Chain != chain {
				continue
			}

			blockNumber, err := GetBlockNumber(chain)
			allowedBlock := *blockNumber - conf.Confrimations

			if err != nil {
				return
			}

			if lastSyncedBlock[conf] == allowedBlock {
				return
			}

			contractAddress := common.HexToAddress(conf.Address)
			contextKey := fmt.Sprintf("plugins.logs.events.%s", conf.Name)
			fromBlock := lastSyncedBlock[conf]

			if fromBlock == 0 {
				contextBlock, err := persistence.ReadUInt64(contextKey)

				if err != nil && err != badger.ErrKeyNotFound {
					panic(err)
				}

				if err != badger.ErrKeyNotFound {
					fromBlock = contextBlock
				} else if conf.From != nil {
					fromBlock = *conf.From
				} else {
					fromBlock = allowedBlock - conf.Step
				}
			}

			toBlock := allowedBlock

			if fromBlock-toBlock > conf.Step {
				toBlock = fromBlock + conf.Step
			}

			query := goEthereum.FilterQuery{
				FromBlock: big.NewInt(int64(fromBlock)),
				ToBlock:   big.NewInt(int64(toBlock)),
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
					With("Block", vLog.BlockNumber)

				sort.Strings(keys)
				for _, key := range keys {
					message = message.
						With(caser.String(key), eventData[key])
				}

				message.Info(conf.Name)
			}

			lastSyncedBlock[conf] = toBlock
			persistence.WriteUint64(contextKey, toBlock)
		}

	}
}

func Start() {

	scheduler, err := gocron.NewScheduler()

	if err != nil {
		panic(err)
	}

	var configs []LogConf
	if err := config.Config.UnmarshalKey("plugins.logs.events", &configs); err != nil {
		panic(err)
	}

	for _, conf := range configs {

		if _, exists := abiMap[conf.Abi]; exists {
			continue
		}

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

	scheduleConfs := config.Config.Sub("plugins.logs.schedule")
	scheduleNames := scheduleConfs.AllKeys()

	for index := range scheduleNames {
		name := scheduleNames[index]
		duration := scheduleConfs.GetDuration(name) * time.Millisecond
		task := createTask(configs, name)

		_, err = scheduler.NewJob(
			gocron.DurationJob(duration),
			gocron.NewTask(task),
			gocron.WithSingletonMode(gocron.LimitModeReschedule),
		)

		if err != nil {
			panic(err)
		}
	}

	scheduler.Start()
}

func init() {
	abiMap = make(map[string]abi.ABI)
	lastSyncedBlock = make(map[LogConf]uint64)
	caser = cases.Title(language.English, cases.NoLower)
}
