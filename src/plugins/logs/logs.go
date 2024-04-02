package logs

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/KenshiTech/unchained/src/address"
	"github.com/KenshiTech/unchained/src/config"
	"github.com/KenshiTech/unchained/src/constants/opcodes"
	"github.com/KenshiTech/unchained/src/crypto/bls"
	"github.com/KenshiTech/unchained/src/datasets"
	"github.com/KenshiTech/unchained/src/db"
	"github.com/KenshiTech/unchained/src/ent"
	"github.com/KenshiTech/unchained/src/ent/signer"
	"github.com/KenshiTech/unchained/src/ethereum"
	"github.com/KenshiTech/unchained/src/log"
	"github.com/KenshiTech/unchained/src/net/shared"
	"github.com/KenshiTech/unchained/src/persistence"
	"github.com/KenshiTech/unchained/src/pos"
	"github.com/KenshiTech/unchained/src/utils"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	"github.com/dgraph-io/badger/v4"
	goEthereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-co-op/gocron/v2"
	lru "github.com/hashicorp/golang-lru/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type EventKey struct {
	Chain    string
	LogIndex uint64
	TxHash   [32]byte
}

type SupportKey struct {
	Chain   string
	Address string
	Event   string
}

var consensus *lru.Cache[EventKey, map[bls12381.G1Affine]big.Int]
var signatureCache *lru.Cache[bls12381.G1Affine, []datasets.Signature]
var aggregateCache *lru.Cache[bls12381.G1Affine, bls12381.G1Affine]
var DebouncedSaveSignatures func(key bls12381.G1Affine, arg SaveSignatureArgs)
var signatureMutex *sync.Mutex
var supportedEvents map[SupportKey]bool

const (
	BlockOutOfRange = 96
	LruSize         = 128
)

// var lastSynced map[string]uint64.
var abiMap map[string]abi.ABI
var lastSyncedBlock map[config.Event]uint64

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

type SaveSignatureArgs struct {
	Info datasets.EventLog
	Hash bls12381.G1Affine
}

func RecordSignature(
	signature bls12381.G1Affine,
	signer datasets.Signer,
	hash bls12381.G1Affine,
	info datasets.EventLog,
	debounce bool,
	historical bool) {
	supportKey := SupportKey{
		Chain:   info.Chain,
		Address: info.Address,
		Event:   info.Event,
	}

	if supported := supportedEvents[supportKey]; !supported {
		return
	}

	// TODO: Standalone mode shouldn't call this or check consensus
	blockNumber, err := GetBlockNumber(info.Chain)

	if err != nil {
		panic(err)
	}

	if !historical {
		// TODO: this won't work for Arbitrum
		// TODO: we disallow syncing historical events here
		if *blockNumber-info.Block > BlockOutOfRange {
			log.Logger.
				With("Packet", info.Block).
				With("Current", *blockNumber).
				Debug("Data too old")
			return // Data too old
		}
	}

	signatureMutex.Lock()
	defer signatureMutex.Unlock()

	key := EventKey{
		Chain:    info.Chain,
		TxHash:   info.TxHash,
		LogIndex: info.LogIndex,
	}

	if !consensus.Contains(key) {
		consensus.Add(key, make(map[bls12381.G1Affine]big.Int))
	}

	votingPower, err := pos.GetVotingPowerOfPublicKey(
		signer.PublicKey,
		big.NewInt(int64(*blockNumber)),
	)

	if err != nil {
		log.Logger.
			With("Address", address.Calculate(signer.PublicKey[:])).
			With("Error", err).
			Error("Failed to get voting power")
		return
	}

	reportedValues, _ := consensus.Get(key)
	voted := reportedValues[hash]
	totalVoted := new(big.Int).Add(votingPower, &voted)
	isMajority := true

	for _, reportCount := range reportedValues {
		if reportCount.Cmp(totalVoted) == 1 {
			isMajority = false
			break
		}
	}

	cached, _ := signatureCache.Get(hash)

	packed := datasets.Signature{
		Signature: signature,
		Signer:    signer,
		Processed: false,
	}

	for _, item := range cached {
		if item.Signer.PublicKey == signer.PublicKey {
			return
		}
	}

	reportedValues[hash] = *totalVoted
	cached = append(cached, packed)
	signatureCache.Add(hash, cached)

	if isMajority {
		if debounce {
			DebouncedSaveSignatures(hash, SaveSignatureArgs{Hash: hash, Info: info})
		} else {
			SaveSignatures(SaveSignatureArgs{Hash: hash, Info: info})
		}
	}
}

func SaveSignatures(args SaveSignatureArgs) {
	dbClient := db.GetClient()
	signatures, ok := signatureCache.Get(args.Hash)

	if !ok {
		return
	}

	ctx := context.Background()

	var newSigners []datasets.Signer
	var newSignatures []bls12381.G1Affine
	var keys [][]byte

	for i := range signatures {
		signature := signatures[i]
		keys = append(keys, signature.Signer.PublicKey[:])
		if !signature.Processed {
			newSignatures = append(newSignatures, signature.Signature)
			newSigners = append(newSigners, signature.Signer)
		}
	}

	// TODO: This part can be a shared library
	err := dbClient.Signer.MapCreateBulk(newSigners, func(sc *ent.SignerCreate, i int) {
		signer := newSigners[i]
		sc.SetName(signer.Name).
			SetEvm(signer.EvmWallet).
			SetKey(signer.PublicKey[:]).
			SetShortkey(signer.ShortPublicKey[:]).
			SetPoints(0)
	}).
		OnConflictColumns("shortkey").
		UpdateName().
		UpdateEvm().
		UpdateKey().
		Update(func(su *ent.SignerUpsert) {
			su.AddPoints(1)
		}).
		Exec(ctx)

	if err != nil {
		panic(err)
	}

	signerIds, err := dbClient.Signer.
		Query().
		Where(signer.KeyIn(keys...)).
		IDs(ctx)

	if err != nil {
		return
	}

	var aggregate bls12381.G1Affine
	currentAggregate, ok := aggregateCache.Get(args.Hash)

	if ok {
		newSignatures = append(newSignatures, currentAggregate)
	}

	aggregate, err = bls.AggregateSignatures(newSignatures)

	if err != nil {
		return
	}

	signatureBytes := aggregate.Bytes()

	err = dbClient.EventLog.
		Create().
		SetBlock(args.Info.Block).
		SetChain(args.Info.Chain).
		SetAddress(args.Info.Address).
		SetEvent(args.Info.Event).
		SetIndex(args.Info.LogIndex).
		SetTransaction(args.Info.TxHash[:]).
		SetSignersCount(uint64(len(signatures))).
		SetSignature(signatureBytes[:]).
		SetArgs(args.Info.Args).
		AddSignerIDs(signerIds...).
		OnConflictColumns("block", "transaction", "index").
		UpdateNewValues().
		Exec(ctx)

	if err != nil {
		panic(err)
	}

	for inx := range signatures {
		signatures[inx].Processed = true
	}

	aggregateCache.Add(args.Hash, aggregate)
}

func createTask(configs []config.Event, chain string) func() {
	return func() {
		if shared.IsClientSocketClosed {
			return
		}

		for _, conf := range configs {
			if conf.Chain != chain {
				continue
			}

			blockNumber, err := GetBlockNumber(chain)
			allowedBlock := *blockNumber - conf.Confirmations

			if err != nil {
				return
			}

			if lastSyncedBlock[conf]+1 >= allowedBlock {
				return
			}

			contractAddress := common.HexToAddress(conf.Address)
			contextKey := fmt.Sprintf("plugins.logs.events.%s", conf.Name)
			fromBlock := lastSyncedBlock[conf] + 1

			if lastSyncedBlock[conf] == 0 {
				contextBlock, err := persistence.ReadUInt64(contextKey)

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

			rpcClient := ethereum.Clients[conf.Chain]
			logs, err := rpcClient.FilterLogs(context.Background(), query)

			if err != nil {
				panic(err)
			}

			contractAbi := abiMap[conf.Abi]
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
					compressedSignature := signature.Bytes()

					priceReport := datasets.EventLogReport{
						EventLog:  event,
						Signature: compressedSignature,
					}

					payload := priceReport.Sia().Content
					shared.Send(opcodes.EventLog, payload)
				}

				if conf.Store {
					RecordSignature(
						signature,
						bls.ClientSigner,
						hash,
						event,
						false,
						true,
					)
				}
			}

			lastSyncedBlock[conf] = toBlock
			err = persistence.WriteUint64(contextKey, toBlock)
			if err != nil {
				panic(err)
			}
		}
	}
}

func New() {
	if config.App.Plugins.EthLog != nil {
		return
	}

	for _, conf := range config.App.Plugins.EthLog.Events {
		key := SupportKey{
			Chain:   conf.Chain,
			Address: conf.Address,
			Event:   conf.Event,
		}
		supportedEvents[key] = true
	}
}

func Listen() {
	if config.App.Plugins.EthLog != nil {
		return
	}

	scheduler, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	for _, conf := range config.App.Plugins.EthLog.Events {
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
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}

	for name, duration := range config.App.Plugins.EthLog.Schedule {
		task := createTask(config.App.Plugins.EthLog.Events, name)

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
	DebouncedSaveSignatures = utils.Debounce[bls12381.G1Affine, SaveSignatureArgs](5*time.Second, SaveSignatures)
	signatureMutex = new(sync.Mutex)

	abiMap = make(map[string]abi.ABI)
	lastSyncedBlock = make(map[config.Event]uint64)
	supportedEvents = make(map[SupportKey]bool)

	var err error
	signatureCache, err = lru.New[bls12381.G1Affine, []datasets.Signature](LruSize)

	if err != nil {
		panic(err)
	}

	consensus, err = lru.New[EventKey, map[bls12381.G1Affine]big.Int](LruSize)

	if err != nil {
		panic(err)
	}

	aggregateCache, err = lru.New[bls12381.G1Affine, bls12381.G1Affine](LruSize)

	if err != nil {
		panic(err)
	}
}
