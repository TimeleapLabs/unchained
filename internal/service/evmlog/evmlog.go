package evmlog

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"slices"
	"time"

	"github.com/KenshiTech/unchained/internal/service/pos"

	"github.com/KenshiTech/unchained/internal/scheduler/persistence"
	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/KenshiTech/unchained/internal/consts"
	"github.com/KenshiTech/unchained/internal/model"

	"github.com/KenshiTech/unchained/internal/repository"

	"github.com/KenshiTech/unchained/internal/crypto/ethereum"

	"github.com/KenshiTech/unchained/internal/config"
	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/ent"
	"github.com/KenshiTech/unchained/internal/ent/helpers"
	"github.com/KenshiTech/unchained/internal/transport/client/conn"
	"github.com/KenshiTech/unchained/internal/utils"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
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
	ethRPC       ethereum.Rpc
	pos          pos.Service
	eventLogRepo repository.EventLog
	signerRepo   repository.Signer
	persistence  *persistence.BadgerRepository

	consensus               *lru.Cache[EventKey, map[bls12381.G1Affine]big.Int]
	signatureCache          *lru.Cache[bls12381.G1Affine, []model.Signature]
	DebouncedSaveSignatures func(key bls12381.G1Affine, arg SaveSignatureArgs)
	supportedEvents         map[SupportKey]bool
	lastSyncedBlock         map[config.Event]uint64
	abiMap                  map[string]abi.ABI
}

func (s *Service) GetBlockNumber(network string) (*uint64, error) {
	blockNumber, err := s.ethRPC.GetBlockNumber(network)

	if err != nil {
		s.ethRPC.RefreshRPC(network)
		return nil, err
	}

	return &blockNumber, nil
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

		if !isNewSigner(signature, currentRecords) {
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

	priceReport := model.EventLogReportPacket{
		EventLog:  event,
		Signature: compressedSignature,
	}

	payload := priceReport.Sia().Content
	conn.Send(consts.OpCodeEventLog, payload)
}

func New(
	ethRPC ethereum.Rpc,
	pos pos.Service,
	eventLogRepo repository.EventLog,
	signerRepo repository.Signer,
	persistence *persistence.BadgerRepository,
) *Service {
	s := Service{
		ethRPC:       ethRPC,
		pos:          pos,
		eventLogRepo: eventLogRepo,
		signerRepo:   signerRepo,
		persistence:  persistence,

		lastSyncedBlock: map[config.Event]uint64{},
		supportedEvents: make(map[SupportKey]bool),
		abiMap:          map[string]abi.ABI{},
	}

	s.DebouncedSaveSignatures = utils.Debounce[bls12381.G1Affine, SaveSignatureArgs](5*time.Second, s.SaveSignatures)

	for _, conf := range config.App.Plugins.EthLog.Events {
		key := SupportKey{
			Chain:   conf.Chain,
			Address: conf.Address,
			Event:   conf.Event,
		}
		s.supportedEvents[key] = true

		if _, exists := s.abiMap[conf.Abi]; exists {
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

		s.abiMap[conf.Abi] = contractAbi
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}

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
