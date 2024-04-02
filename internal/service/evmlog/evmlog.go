package evmlog

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/KenshiTech/unchained/address"
	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/constants/opcodes"
	"github.com/KenshiTech/unchained/crypto/bls"
	"github.com/KenshiTech/unchained/datasets"
	"github.com/KenshiTech/unchained/db"
	"github.com/KenshiTech/unchained/ent"
	"github.com/KenshiTech/unchained/ent/signer"
	"github.com/KenshiTech/unchained/ethereum"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/pos"
	"github.com/KenshiTech/unchained/transport/client/conn"
	"github.com/KenshiTech/unchained/utils"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	"github.com/ethereum/go-ethereum/accounts/abi"
	lru "github.com/hashicorp/golang-lru/v2"
)

const (
	BlockOutOfRange = 96
	LruSize         = 128
)

type SaveSignatureArgs struct {
	Info datasets.EventLog
	Hash bls12381.G1Affine
}

type Service struct {
	consensus               *lru.Cache[EventKey, map[bls12381.G1Affine]big.Int]
	signatureCache          *lru.Cache[bls12381.G1Affine, []datasets.Signature]
	aggregateCache          *lru.Cache[bls12381.G1Affine, bls12381.G1Affine]
	DebouncedSaveSignatures func(key bls12381.G1Affine, arg SaveSignatureArgs)
	signatureMutex          *sync.Mutex
	supportedEvents         map[SupportKey]bool
	abiMap                  map[string]abi.ABI
	lastSyncedBlock         map[config.Event]uint64
}

func (e *Service) GetBlockNumber(network string) (*uint64, error) {
	blockNumber, err := ethereum.GetBlockNumber(network)

	if err != nil {
		ethereum.RefreshRPC(network)
		return nil, err
	}

	return &blockNumber, nil
}

func (e *Service) RecordSignature(
	signature bls12381.G1Affine, signer datasets.Signer, hash bls12381.G1Affine, info datasets.EventLog, debounce bool, historical bool,
) {
	supportKey := SupportKey{
		Chain:   info.Chain,
		Address: info.Address,
		Event:   info.Event,
	}

	if supported := e.supportedEvents[supportKey]; !supported {
		return
	}

	// TODO: Standalone mode shouldn't call this or check consensus
	blockNumber, err := e.GetBlockNumber(info.Chain)

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

	e.signatureMutex.Lock()
	defer e.signatureMutex.Unlock()

	key := EventKey{
		Chain:    info.Chain,
		TxHash:   info.TxHash,
		LogIndex: info.LogIndex,
	}

	if !e.consensus.Contains(key) {
		e.consensus.Add(key, make(map[bls12381.G1Affine]big.Int))
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

	reportedValues, _ := e.consensus.Get(key)
	voted := reportedValues[hash]
	totalVoted := new(big.Int).Add(votingPower, &voted)
	isMajority := true

	for _, reportCount := range reportedValues {
		if reportCount.Cmp(totalVoted) == 1 {
			isMajority = false
			break
		}
	}

	cached, _ := e.signatureCache.Get(hash)

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
	e.signatureCache.Add(hash, cached)

	if isMajority {
		if debounce {
			e.DebouncedSaveSignatures(hash, SaveSignatureArgs{Hash: hash, Info: info})
		} else {
			e.SaveSignatures(SaveSignatureArgs{Hash: hash, Info: info})
		}
	}
}

func (e *Service) SaveSignatures(args SaveSignatureArgs) {
	dbClient := db.GetClient()
	signatures, ok := e.signatureCache.Get(args.Hash)

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
	currentAggregate, ok := e.aggregateCache.Get(args.Hash)

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

	e.aggregateCache.Add(args.Hash, aggregate)
}

func (e *Service) SendPriceReport(signature bls12381.G1Affine, event datasets.EventLog) {
	compressedSignature := signature.Bytes()

	priceReport := datasets.EventLogReport{
		EventLog:  event,
		Signature: compressedSignature,
	}

	payload := priceReport.Sia().Content
	conn.Send(opcodes.EventLog, payload)
}

func New() *Service {
	s := Service{}

	s.DebouncedSaveSignatures = utils.Debounce[bls12381.G1Affine, SaveSignatureArgs](5*time.Second, s.SaveSignatures)
	s.signatureMutex = new(sync.Mutex)

	s.abiMap = make(map[string]abi.ABI)
	s.lastSyncedBlock = make(map[config.Event]uint64)
	s.supportedEvents = make(map[SupportKey]bool)

	var err error
	s.signatureCache, err = lru.New[bls12381.G1Affine, []datasets.Signature](LruSize)
	if err != nil {
		panic(err)
	}

	s.consensus, err = lru.New[EventKey, map[bls12381.G1Affine]big.Int](LruSize)
	if err != nil {
		panic(err)
	}

	s.aggregateCache, err = lru.New[bls12381.G1Affine, bls12381.G1Affine](LruSize)
	if err != nil {
		panic(err)
	}

	return &s
}
