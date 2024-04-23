package evmlog

import (
	"context"
	"fmt"
	"math/big"
	"slices"
	"sync"
	"time"

	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"

	"github.com/TimeleapLabs/unchained/internal/address"
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/constants/opcodes"
	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/datasets"
	"github.com/TimeleapLabs/unchained/internal/db"
	"github.com/TimeleapLabs/unchained/internal/ent"
	"github.com/TimeleapLabs/unchained/internal/ent/eventlog"
	"github.com/TimeleapLabs/unchained/internal/ent/helpers"
	"github.com/TimeleapLabs/unchained/internal/ent/signer"
	"github.com/TimeleapLabs/unchained/internal/log"
	"github.com/TimeleapLabs/unchained/internal/pos"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
	"github.com/TimeleapLabs/unchained/internal/utils"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	"github.com/ethereum/go-ethereum/accounts/abi"
	lru "github.com/hashicorp/golang-lru/v2"
)

const (
	BlockOutOfRange = 96
	LruSize         = 128
)

type SaveSignatureArgs struct {
	Info      datasets.EventLog
	Hash      bls12381.G1Affine
	Consensus bool
	Voted     *big.Int
}

type Service struct {
	ethRPC *ethereum.Repository
	pos    *pos.Repository

	consensus               *lru.Cache[EventKey, map[bls12381.G1Affine]big.Int]
	signatureCache          *lru.Cache[bls12381.G1Affine, []datasets.Signature]
	DebouncedSaveSignatures func(key bls12381.G1Affine, arg SaveSignatureArgs)
	signatureMutex          *sync.Mutex
	supportedEvents         map[SupportKey]bool
	abiMap                  map[string]abi.ABI
	lastSyncedBlock         map[config.Event]uint64
}

func (e *Service) GetBlockNumber(network string) (*uint64, error) {
	blockNumber, err := e.ethRPC.GetBlockNumber(network)

	if err != nil {
		e.ethRPC.RefreshRPC(network)
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

	votingPower, err := e.pos.GetVotingPowerOfPublicKey(signer.PublicKey)
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
	}

	for _, item := range cached {
		if item.Signer.PublicKey == signer.PublicKey {
			return
		}
	}

	reportedValues[hash] = *totalVoted
	cached = append(cached, packed)
	e.signatureCache.Add(hash, cached)

	saveArgs := SaveSignatureArgs{
		Hash:      hash,
		Info:      info,
		Consensus: isMajority,
		Voted:     totalVoted,
	}

	if debounce {
		e.DebouncedSaveSignatures(hash, saveArgs)
	} else {
		e.SaveSignatures(saveArgs)
	}
}

func IsNewSigner(signature datasets.Signature, records []*ent.EventLog) bool {
	for _, record := range records {
		for _, signer := range record.Edges.Signers {
			if signature.Signer.PublicKey == [96]byte(signer.Key) {
				return false
			}
		}
	}

	return true
}

func sortEventArgs(lhs datasets.EventLogArg, rhs datasets.EventLogArg) int {
	if lhs.Name < rhs.Name {
		return -1
	}
	return 1
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

	currentRecords, err := dbClient.EventLog.
		Query().
		Where(
			eventlog.Block(args.Info.Block),
			eventlog.TransactionEQ(args.Info.TxHash[:]),
			eventlog.IndexEQ(args.Info.LogIndex),
		).
		WithSigners().
		All(ctx)

	if err != nil && !ent.IsNotFound(err) {
		panic(err)
	}

	for i := range signatures {
		signature := signatures[i]
		keys = append(keys, signature.Signer.PublicKey[:])

		if !IsNewSigner(signature, currentRecords) {
			continue
		}

		newSignatures = append(newSignatures, signature.Signature)
		newSigners = append(newSigners, signature.Signer)
	}

	// TODO: This part can be a shared library
	err = dbClient.Signer.MapCreateBulk(newSigners, func(sc *ent.SignerCreate, i int) {
		signer := newSigners[i]
		sc.SetName(signer.Name).
			SetEvm(signer.EvmAddress).
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

	signerIDs, err := dbClient.Signer.
		Query().
		Where(signer.KeyIn(keys...)).
		IDs(ctx)

	if err != nil {
		return
	}

	var aggregate bls12381.G1Affine

	sortedCurrentArgs := make([]datasets.EventLogArg, len(args.Info.Args))
	copy(sortedCurrentArgs, args.Info.Args)
	slices.SortFunc(sortedCurrentArgs, sortEventArgs)

	for _, record := range currentRecords {
		sortedRecordArgs := make([]datasets.EventLogArg, len(args.Info.Args))
		copy(sortedRecordArgs, record.Args)
		slices.SortFunc(sortedRecordArgs, sortEventArgs)

		// compare args
		if len(sortedCurrentArgs) != len(sortedRecordArgs) {
			continue
		}

		for i := range sortedCurrentArgs {
			if sortedCurrentArgs[i].Name != sortedRecordArgs[i].Name ||
				sortedCurrentArgs[i].Value != sortedRecordArgs[i].Value {
				continue
			}
		}

		// Check if record mataches the passed event
		if record.Chain != args.Info.Chain ||
			record.Address != args.Info.Address ||
			record.Event != args.Info.Event {
			continue
		}

		currentAggregate, err := bls.RecoverSignature([48]byte(record.Signature))

		if err != nil {
			log.Logger.
				With("Block", args.Info.Block).
				With("Transaction", fmt.Sprintf("%x", args.Info.TxHash)).
				With("Index", args.Info.LogIndex).
				With("Event", args.Info.Event).
				With("Hash", fmt.Sprintf("%x", args.Hash.Bytes())[:8]).
				With("Error", err).
				Debug("Failed to recover signature")
			return
		}

		newSignatures = append(newSignatures, currentAggregate)
		break
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
		SetConsensus(args.Consensus).
		SetVoted(&helpers.BigInt{Int: *args.Voted}).
		AddSignerIDs(signerIDs...).
		OnConflictColumns("block", "transaction", "index").
		UpdateNewValues().
		Exec(ctx)

	if err != nil {
		panic(err)
	}
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

func New(
	ethRPC *ethereum.Repository,
	pos *pos.Repository,
) *Service {
	s := Service{
		ethRPC: ethRPC,
		pos:    pos,
	}

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

	return &s
}
