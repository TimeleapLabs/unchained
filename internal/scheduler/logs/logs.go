package logs

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strings"

	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/crypto/bls"
	"github.com/KenshiTech/unchained/datasets"
	"github.com/KenshiTech/unchained/ethereum"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/persistence"
	"github.com/KenshiTech/unchained/service/evmlog"
	"github.com/dgraph-io/badger/v4"
	goEthereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type EvmLog struct {
	chain         string
	evmLogService *evmlog.Service
	ethRPC        *ethereum.Repository
	persistence   *persistence.BadgerRepository

	supportedEvents map[evmlog.SupportKey]bool
	abiMap          map[string]abi.ABI
	lastSyncedBlock map[config.Event]uint64
}

func (e *EvmLog) Run() {
	log.Logger.With("Chain", e.chain).Info("Run evm log task")

	for _, conf := range config.App.Plugins.EthLog.Events {
		if conf.Chain != e.chain {
			continue
		}

		blockNumber, err := e.evmLogService.GetBlockNumber(e.chain)
		allowedBlock := *blockNumber - conf.Confirmations

		if err != nil {
			return
		}

		if e.lastSyncedBlock[conf]+1 >= allowedBlock {
			return
		}

		contractAddress := common.HexToAddress(conf.Address)
		contextKey := fmt.Sprintf("plugins.logs.events.%s", conf.Name)
		fromBlock := e.lastSyncedBlock[conf] + 1

		if e.lastSyncedBlock[conf] == 0 {
			contextBlock, err := e.persistence.ReadUInt64(contextKey)

			if err != nil && !errors.Is(err, badger.ErrKeyNotFound) {
				panic(err)
			}

			fromBlock = allowedBlock - conf.Step
			if !errors.Is(err, badger.ErrKeyNotFound) {
				fromBlock = contextBlock
			} else if conf.From != nil {
				fromBlock = *conf.From
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

		rpcClient := e.ethRPC.Clients[conf.Chain]
		logs, err := rpcClient.FilterLogs(context.Background(), query)

		if err != nil {
			panic(err)
		}

		contractAbi := e.abiMap[conf.Abi]
		caser := cases.Title(language.English, cases.NoLower)

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

			indexedParams := make([]abi.Argument, 0)
			for _, input := range eventAbi.Inputs {
				if input.Indexed {
					indexedParams = append(indexedParams, input)
				}
			}

			err = abi.ParseTopicsIntoMap(eventData, indexedParams, vLog.Topics[1:])
			if err != nil {
				panic(err)
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

			argTypes := make(map[string]string)
			for _, input := range eventAbi.Inputs {
				argTypes[input.Name] = input.Type.String()
			}

			args := []datasets.EventLogArg{}
			for _, key := range keys {
				value := eventData[key]

				if strings.HasPrefix(argTypes[key], "uint") || strings.HasPrefix(argTypes[key], "int") {
					value = value.(*big.Int).String()
				}

				args = append(
					args,
					datasets.EventLogArg{
						Name:  key,
						Value: value,
						Type:  argTypes[key],
					},
				)
			}

			event := datasets.EventLog{
				LogIndex: uint64(vLog.Index),
				Block:    vLog.BlockNumber,
				Address:  vLog.Address.Hex(),
				Event:    conf.Event,
				Chain:    conf.Chain,
				TxHash:   vLog.TxHash,
				Args:     args,
			}

			toHash := event.Sia().Content
			signature, hash := bls.Sign(*bls.ClientSecretKey, toHash)

			if conf.Send {
				e.evmLogService.SendPriceReport(signature, event)
			}

			if conf.Store {
				e.evmLogService.RecordSignature(
					signature,
					bls.ClientSigner,
					hash,
					event,
					false,
					true,
				)
			}
		}

		e.lastSyncedBlock[conf] = toBlock
		err = e.persistence.WriteUint64(contextKey, toBlock)
		if err != nil {
			panic(err)
		}
	}
}

func New(
	chanName string, events []config.Event,
	evmLogService *evmlog.Service,
	ethRPC *ethereum.Repository,
	persistence *persistence.BadgerRepository,
) *EvmLog {
	e := EvmLog{
		chain:         chanName,
		evmLogService: evmLogService,
		ethRPC:        ethRPC,
		persistence:   persistence,

		supportedEvents: map[evmlog.SupportKey]bool{},
		abiMap:          map[string]abi.ABI{},
		lastSyncedBlock: map[config.Event]uint64{},
	}

	for _, conf := range events {
		key := evmlog.SupportKey{
			Chain:   conf.Chain,
			Address: conf.Address,
			Event:   conf.Event,
		}
		e.supportedEvents[key] = true

		if _, exists := e.abiMap[conf.Abi]; exists {
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

		e.abiMap[conf.Abi] = contractAbi
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}

	return &e
}
