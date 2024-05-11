package evmlog

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"strings"

	"github.com/TimeleapLabs/unchained/internal/crypto/bls"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/utils"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/dgraph-io/badger/v4"
	goEthereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (s *service) ProcessBlocks(ctx context.Context, chain string) error {
	for _, conf := range config.App.Plugins.EthLog.Events {
		if conf.Chain != chain {
			continue
		}

		// check if processing is needed
		blockNumber, err := s.ethRPC.GetBlockNumber(ctx, chain)
		if err != nil {
			s.ethRPC.RefreshRPC(chain)
			return err
		}

		allowedBlock := blockNumber - conf.Confirmations

		if s.lastSyncedBlock[conf]+1 >= allowedBlock {
			return consts.ErrAlreadySynced
		}

		contractAddress := common.HexToAddress(conf.Address)
		contextKey := fmt.Sprintf("plugins.logs.events.%s", conf.Name)
		fromBlock := s.lastSyncedBlock[conf] + 1

		if s.lastSyncedBlock[conf] == 0 {
			contextBlock, err := s.persistence.ReadUInt64(contextKey)

			if err != nil && !errors.Is(err, badger.ErrKeyNotFound) {
				return err
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

		rpcClient := s.ethRPC.GetClient(conf.Chain)
		logs, err := rpcClient.FilterLogs(ctx, query)

		if err != nil {
			return err
		}

		contractAbi := s.abiMap[conf.Abi]
		caser := cases.Title(language.English, cases.NoLower)

		for _, vLog := range logs {
			eventSignature := vLog.Topics[0]
			eventAbi, err := contractAbi.EventByID(eventSignature)

			if eventAbi.Name != conf.Event {
				continue
			}

			if err != nil {
				return err
			}

			eventData := make(map[string]interface{})
			err = contractAbi.UnpackIntoMap(eventData, eventAbi.Name, vLog.Data)
			if err != nil {
				return err
			}

			indexedParams := make([]abi.Argument, 0)
			for _, input := range eventAbi.Inputs {
				if input.Indexed {
					indexedParams = append(indexedParams, input)
				}
			}

			err = abi.ParseTopicsIntoMap(eventData, indexedParams, vLog.Topics[1:])
			if err != nil {
				return err
			}

			var keys []string
			for k := range eventData {
				keys = append(keys, k)
			}

			message := utils.Logger.
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

			args := []model.EventLogArg{}
			for _, key := range keys {
				value := eventData[key]

				if strings.HasPrefix(argTypes[key], "uint") || strings.HasPrefix(argTypes[key], "int") {
					value = value.(*big.Int).String()
				}

				args = append(
					args,
					model.EventLogArg{
						Name:  key,
						Value: value,
						Type:  argTypes[key],
					},
				)
			}

			event := model.EventLog{
				LogIndex: uint64(vLog.Index),
				Block:    vLog.BlockNumber,
				Address:  vLog.Address.Hex(),
				Event:    conf.Event,
				Chain:    conf.Chain,
				TxHash:   vLog.TxHash,
				Args:     args,
			}

			toHash := event.Sia().Bytes()
			signature, err := s.signer.Sign(toHash)
			if err != nil {
				panic(err)
			}

			hashedMessage, err := bls.Hash(signature)
			if err != nil {
				panic(err)
			}

			if conf.Send {
				s.SendPriceReport(signature, event)
			}

			if conf.Store {
				err = s.RecordSignature(
					ctx,
					signature,
					*crypto.Identity.ExportEvmSigner(),
					hashedMessage,
					event,
					false,
					true,
				)
				if err != nil {
					return err
				}
			}
		}

		s.lastSyncedBlock[conf] = toBlock
		err = s.persistence.WriteUint64(contextKey, toBlock)
		if err != nil {
			return err
		}
	}

	return nil
}
