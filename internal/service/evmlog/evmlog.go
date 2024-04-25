package evmlog

import (
	"context"
	"fmt"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/repository"
	"github.com/TimeleapLabs/unchained/internal/utils/address"
	"math/big"
	"slices"
	"sync"
	"time"

	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/ent"
	"github.com/TimeleapLabs/unchained/internal/ent/helpers"
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
	Info      model.EventLog
	Hash      bls12381.G1Affine
	Consensus bool
	Voted     *big.Int
}

type Service struct {
	ethRPC       *ethereum.Repository
	pos          *pos.Repository
	eventLogRepo repository.EventLog
	signerRepo   repository.Signer

	consensus               *lru.Cache[EventKey, map[bls12381.G1Affine]big.Int]
	signatureCache          *lru.Cache[bls12381.G1Affine, []model.Signature]
	DebouncedSaveSignatures func(key bls12381.G1Affine, arg SaveSignatureArgs)
	signatureMutex          *sync.Mutex
	supportedEvents         map[SupportKey]bool
	abiMap                  map[string]abi.ABI
	lastSyncedBlock         map[config.Event]uint64
}

func (s *Service) GetBlockNumber(network string) (*uint64, error) {
	blockNumber, err := s.ethRPC.GetBlockNumber(network)

	if err != nil {
		s.ethRPC.RefreshRPC(network)
		return nil, err
	}

	return &blockNumber, nil
}

func (s *Service) RecordSignature(
	signature bls12381.G1Affine, signer model.Signer, hash bls12381.G1Affine, info model.EventLog, debounce bool, historical bool,
) error {
	supportKey := SupportKey{
		Chain:   info.Chain,
		Address: info.Address,
		Event:   info.Event,
	}

	if supported := s.supportedEvents[supportKey]; !supported {
		return consts.ErrTokenNotSupported
	}

	// TODO: Standalone mode shouldn't call this or check consensus
	blockNumber, err := s.GetBlockNumber(info.Chain)

	if err != nil {
		return err
	}

	if !historical {
		// TODO: this won't work for Arbitrum
		// TODO: we disallow syncing historical events here
		if *blockNumber-info.Block > BlockOutOfRange {
			utils.Logger.
				With("Packet", info.Block).
				With("Current", *blockNumber).
				Debug("Data too old")
			return consts.ErrDataTooOld
		}
	}

	s.signatureMutex.Lock()
	defer s.signatureMutex.Unlock()

	key := EventKey{
		Chain:    info.Chain,
		TxHash:   info.TxHash,
		LogIndex: info.LogIndex,
	}

	if !s.consensus.Contains(key) {
		s.consensus.Add(key, make(map[bls12381.G1Affine]big.Int))
	}

	votingPower, err := s.pos.GetVotingPowerOfPublicKey(signer.PublicKey)
	if err != nil {
		utils.Logger.
			With("Address", address.Calculate(signer.PublicKey[:])).
			With("Error", err).
			Error("Failed to get voting power")
		return err
	}

	reportedValues, _ := s.consensus.Get(key)
	voted := reportedValues[hash]
	totalVoted := new(big.Int).Add(votingPower, &voted)
	isMajority := true

	for _, reportCount := range reportedValues {
		if reportCount.Cmp(totalVoted) == 1 {
			isMajority = false
			break
		}
	}

	cached, _ := s.signatureCache.Get(hash)

	packed := model.Signature{
		Signature: signature,
		Signer:    signer,
	}

	for _, item := range cached {
		if item.Signer.PublicKey == signer.PublicKey {
			return consts.ErrDuplicateSignature
		}
	}

	reportedValues[hash] = *totalVoted
	cached = append(cached, packed)
	s.signatureCache.Add(hash, cached)

	saveArgs := SaveSignatureArgs{
		Hash:      hash,
		Info:      info,
		Consensus: isMajority,
		Voted:     totalVoted,
	}

	if debounce {
		s.DebouncedSaveSignatures(hash, saveArgs)
		return nil
	}

	err = s.SaveSignatures(saveArgs)
	if err != nil {
		return err
	}

	return nil
}

func IsNewSigner(signature model.Signature, records []*ent.EventLog) bool {
	for _, record := range records {
		for _, signer := range record.Edges.Signers {
			if signature.Signer.PublicKey == [96]byte(signer.Key) {
				return false
			}
		}
	}

	return true
}

func sortEventArgs(lhs model.EventLogArg, rhs model.EventLogArg) int {
	if lhs.Name < rhs.Name {
		return -1
	}
	return 1
}

func (s *Service) SaveSignatures(args SaveSignatureArgs) error {
	signatures, ok := s.signatureCache.Get(args.Hash)
	if !ok {
		return consts.ErrInvalidSignature
	}

	ctx := context.Background()

	var newSigners []model.Signer
	var newSignatures []bls12381.G1Affine
	var keys [][]byte

	currentRecords, err := s.eventLogRepo.Find(ctx, args.Info.Block, args.Info.TxHash[:], args.Info.LogIndex)
	if err != nil && !ent.IsNotFound(err) {
		return err
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

	err = s.signerRepo.CreateSigners(ctx, newSigners)

	if err != nil {
		return err
	}

	signerIDs, err := s.signerRepo.GetSingerIDsByKeys(context.TODO(), keys)

	if err != nil {
		return err
	}

	var aggregate bls12381.G1Affine

	sortedCurrentArgs := make([]model.EventLogArg, len(args.Info.Args))
	copy(sortedCurrentArgs, args.Info.Args)
	slices.SortFunc(sortedCurrentArgs, sortEventArgs)

	for _, record := range currentRecords {
		sortedRecordArgs := make([]model.EventLogArg, len(args.Info.Args))
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
			utils.Logger.
				With("Block", args.Info.Block).
				With("Transaction", fmt.Sprintf("%x", args.Info.TxHash)).
				With("Index", args.Info.LogIndex).
				With("Event", args.Info.Event).
				With("Hash", fmt.Sprintf("%x", args.Hash.Bytes())[:8]).
				With("Error", err).
				Debug("Failed to recover signature")
			return consts.ErrCantRecoverSignature
		}

		newSignatures = append(newSignatures, currentAggregate)
		break
	}

	aggregate, err = bls.AggregateSignatures(newSignatures)

	if err != nil {
		return consts.ErrCantAggregateSignatures
	}

	signatureBytes := aggregate.Bytes()

	args.Info.SignersCount = uint64(len(signatures))
	args.Info.SignerIDs = signerIDs
	args.Info.Consensus = args.Consensus
	args.Info.Signature = signatureBytes[:]
	args.Info.Voted = &helpers.BigInt{Int: *args.Voted}
	err = s.eventLogRepo.Upsert(ctx, args.Info)

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) SendPriceReport(signature bls12381.G1Affine, event model.EventLog) {
	compressedSignature := signature.Bytes()

	priceReport := model.EventLogReport{
		EventLog:  event,
		Signature: compressedSignature,
	}

	payload := priceReport.Sia().Content
	conn.Send(consts.OpCodeEventLog, payload)
}

func New(
	ethRPC *ethereum.Repository,
	pos *pos.Repository,
	eventLogRepo repository.EventLog,
	signerRepo repository.Signer,
) *Service {
	s := Service{
		ethRPC:       ethRPC,
		pos:          pos,
		eventLogRepo: eventLogRepo,
		signerRepo:   signerRepo,
	}

	s.DebouncedSaveSignatures = utils.Debounce[bls12381.G1Affine, SaveSignatureArgs](5*time.Second, s.SaveSignatures)
	s.signatureMutex = new(sync.Mutex)

	s.abiMap = make(map[string]abi.ABI)
	s.lastSyncedBlock = make(map[config.Event]uint64)
	s.supportedEvents = make(map[SupportKey]bool)

	var err error
	s.signatureCache, err = lru.New[bls12381.G1Affine, []model.Signature](LruSize)
	if err != nil {
		panic(err)
	}

	s.consensus, err = lru.New[EventKey, map[bls12381.G1Affine]big.Int](LruSize)
	if err != nil {
		panic(err)
	}

	return &s
}
