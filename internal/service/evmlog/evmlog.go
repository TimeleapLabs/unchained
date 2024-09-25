package evmlog

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"slices"
	"sync"
	"time"

	"github.com/TimeleapLabs/unchained/internal/service/correctness"
	"github.com/TimeleapLabs/unchained/internal/transport/server/packet"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/repository"
	"github.com/TimeleapLabs/unchained/internal/service/pos"
	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
	"github.com/TimeleapLabs/unchained/internal/utils"
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

type Service interface {
	GetBlockNumber(ctx context.Context, network string) (*uint64, error)
	SaveSignatures(ctx context.Context, args SaveSignatureArgs) error
	SendPriceReport(signature bls12381.G1Affine, event model.EventLog)
	ProcessBlocks(ctx context.Context, chain string) error
	RecordSignature(
		ctx context.Context, signature bls12381.G1Affine, signer model.Signer, hash bls12381.G1Affine, info model.EventLog, debounce bool, historical bool,
	) error
}

type service struct {
	ethRPC       ethereum.RPC
	pos          pos.Service
	eventLogRepo repository.EventLog
	proofRepo    repository.Proof
	persistence  *Badger

	consensus               *lru.Cache[EventKey, map[bls12381.G1Affine]big.Int]
	signatureCache          *lru.Cache[bls12381.G1Affine, []correctness.Signature]
	DebouncedSaveSignatures func(key bls12381.G1Affine, arg SaveSignatureArgs)
	signatureMutex          *sync.Mutex
	supportedEvents         map[SupportKey]bool
	lastSyncedBlock         map[config.Event]uint64
	abiMap                  map[string]abi.ABI
}

func (s *service) GetBlockNumber(ctx context.Context, network string) (*uint64, error) {
	blockNumber, err := s.ethRPC.GetBlockNumber(ctx, network)

	if err != nil {
		s.ethRPC.RefreshRPC(network)
		return nil, err
	}

	return &blockNumber, nil
}

func (s *service) SaveSignatures(ctx context.Context, args SaveSignatureArgs) error {
	signatures, ok := s.signatureCache.Get(args.Hash)
	if !ok {
		return consts.ErrInvalidSignature
	}

	var newSigners []model.Signer
	var newSignatures []bls12381.G1Affine

	currentRecords, err := s.eventLogRepo.Find(ctx, args.Info.Block, args.Info.TxHash[:], args.Info.LogIndex)
	if err != nil {
		return err
	}

	for i := range signatures {
		signature := signatures[i]

		newSignatures = append(newSignatures, signature.Signature)
		newSigners = append(newSigners, signature.Signer)
	}

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

	var aggregate bls12381.G1Affine
	aggregate, err = bls.AggregateSignatures(newSignatures)
	if err != nil {
		return consts.ErrCantAggregateSignatures
	}

	signatureBytes := aggregate.Bytes()

	err = s.proofRepo.CreateProof(ctx, signatureBytes, newSigners)
	if err != nil {
		return err
	}

	args.Info.SignersCount = uint64(len(signatures))
	args.Info.Consensus = args.Consensus
	args.Info.Signature = signatureBytes[:]
	args.Info.Voted = args.Voted.Int64()

	err = s.eventLogRepo.Upsert(ctx, args.Info)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) SendPriceReport(signature bls12381.G1Affine, event model.EventLog) {
	priceReport := packet.EventLogReportPacket{
		EventLog:  event,
		Signature: signature.Bytes(),
	}

	conn.Send(consts.OpCodeEventLog, priceReport.Sia().Bytes())
}

func New(
	ethRPC ethereum.RPC, pos pos.Service, eventLogRepo repository.EventLog, proofRepo repository.Proof, persistence *Badger,
) Service {
	s := service{
		ethRPC:       ethRPC,
		pos:          pos,
		eventLogRepo: eventLogRepo,
		proofRepo:    proofRepo,
		persistence:  persistence,

		signatureMutex:  new(sync.Mutex),
		lastSyncedBlock: map[config.Event]uint64{},
		supportedEvents: make(map[SupportKey]bool),
		abiMap:          map[string]abi.ABI{},
	}

	s.DebouncedSaveSignatures = utils.Debounce[bls12381.G1Affine, SaveSignatureArgs](5*time.Second, s.SaveSignatures)

	var err error
	s.signatureCache, err = lru.New[bls12381.G1Affine, []correctness.Signature](LruSize)
	if err != nil {
		panic(err)
	}

	s.consensus, err = lru.New[EventKey, map[bls12381.G1Affine]big.Int](LruSize)
	if err != nil {
		panic(err)
	}

	if config.App.Plugins.EthLog != nil {
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
	}

	return &s
}
