package correctness

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"sync"
	"time"

	"github.com/TimeleapLabs/unchained/internal/address"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"
	"github.com/TimeleapLabs/unchained/internal/log"
	"github.com/TimeleapLabs/unchained/internal/pos"
	"github.com/TimeleapLabs/unchained/internal/service/evmlog"
	"github.com/puzpuzpuz/xsync/v3"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/crypto/shake"
	"github.com/TimeleapLabs/unchained/internal/datasets"
	"github.com/TimeleapLabs/unchained/internal/db"
	"github.com/TimeleapLabs/unchained/internal/ent"
	"github.com/TimeleapLabs/unchained/internal/ent/correctnessreport"
	"github.com/TimeleapLabs/unchained/internal/ent/helpers"
	"github.com/TimeleapLabs/unchained/internal/ent/signer"
	"github.com/TimeleapLabs/unchained/internal/utils"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	lru "github.com/hashicorp/golang-lru/v2"
)

const (
	LruSize = 128
)

type Conf struct {
	Topic string `mapstructure:"name"`
}

type SaveSignatureArgs struct {
	Info      datasets.Correctness
	Hash      bls12381.G1Affine
	Consensus bool
	Voted     *big.Int
}
type Service struct {
	ethRPC *ethereum.Repository
	pos    *pos.Repository

	signatureCache *lru.Cache[bls12381.G1Affine, []datasets.Signature]
	consensus      *lru.Cache[Key, xsync.MapOf[bls12381.G1Affine, big.Int]]

	DebouncedSaveSignatures func(key bls12381.G1Affine, arg SaveSignatureArgs)
	signatureMutex          *sync.Mutex
	supportedTopics         map[[64]byte]bool
}

// TODO: This code should be moved to a shared library
func (s *Service) GetBlockNumber(network string) (*uint64, error) {
	blockNumber, err := s.ethRPC.GetBlockNumber(network)

	if err != nil {
		s.ethRPC.RefreshRPC(network)
		return nil, err
	}

	return &blockNumber, nil
}

func (s *Service) IsNewSigner(signature datasets.Signature, records []*ent.CorrectnessReport) bool {
	// TODO: This isn't efficient, we should use a map
	for _, record := range records {
		for _, signer := range record.Edges.Signers {
			if signature.Signer.PublicKey == [96]byte(signer.Key) {
				return false
			}
		}
	}

	return true
}

// TODO: How should we handle older records?
// Possible Solution: Add a not after timestamp to the document.
func (s *Service) RecordSignature(
	signature bls12381.G1Affine,
	signer datasets.Signer,
	hash bls12381.G1Affine,
	info datasets.Correctness,
	debounce bool) {
	if supported := s.supportedTopics[info.Topic]; !supported {
		return
	}

	s.signatureMutex.Lock()
	defer s.signatureMutex.Unlock()

	signatures, ok := s.signatureCache.Get(hash)

	if !ok {
		signatures = make([]datasets.Signature, 0)
	}

	// Check for duplicates
	for _, sig := range signatures {
		if sig.Signer.PublicKey == signer.PublicKey {
			return
		}
	}

	packed := datasets.Signature{
		Signature: signature,
		Signer:    signer,
	}

	key := Key{
		Hash:    fmt.Sprintf("%x", info.Hash),
		Topic:   fmt.Sprintf("%x", info.Topic),
		Correct: info.Correct,
	}

	if !s.consensus.Contains(key) {
		s.consensus.Add(key, *xsync.NewMapOf[bls12381.G1Affine, big.Int]())
	}

	reportedValues, _ := s.consensus.Get(key)
	isMajority := true
	voted, ok := reportedValues.Load(hash)
	if !ok {
		voted = *big.NewInt(0)
	}

	votingPower, err := s.pos.GetVotingPowerOfPublicKey(signer.PublicKey)
	if err != nil {
		log.Logger.
			With("Address", address.Calculate(signer.PublicKey[:])).
			With("Error", err).
			Error("Failed to get voting power")
		return
	}

	totalVoted := new(big.Int).Add(votingPower, &voted)

	reportedValues.Range(func(_ bls12381.G1Affine, value big.Int) bool {
		if value.Cmp(totalVoted) == 1 {
			isMajority = false
		}
		return isMajority
	})

	reportedValues.Store(hash, *totalVoted)
	signatures = append(signatures, packed)
	s.signatureCache.Add(hash, signatures)

	saveArgs := SaveSignatureArgs{
		Info:      info,
		Hash:      hash,
		Consensus: isMajority,
		Voted:     totalVoted,
	}

	if debounce {
		s.DebouncedSaveSignatures(hash, saveArgs)
	} else {
		s.SaveSignatures(saveArgs)
	}
}

func (s *Service) SaveSignatures(args SaveSignatureArgs) {
	dbClient := db.GetClient()
	signatures, ok := s.signatureCache.Get(args.Hash)

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
	}

	currentRecords, err := dbClient.CorrectnessReport.
		Query().
		Where(correctnessreport.And(
			correctnessreport.Hash(args.Info.Hash[:]),
			correctnessreport.Topic(args.Info.Topic[:]),
			correctnessreport.Timestamp(args.Info.Timestamp),
		)).
		All(ctx)

	if err != nil && !ent.IsNotFound(err) {
		panic(err)
	}

	// Select the new signers and signatures

	for i := range signatures {
		signature := signatures[i]

		if !s.IsNewSigner(signature, currentRecords) {
			continue
		}

		newSigners = append(newSigners, signature.Signer)
		newSignatures = append(newSignatures, signature.Signature)
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

	for _, record := range currentRecords {
		if record.Correct == args.Info.Correct {
			currentSignature, err := bls.RecoverSignature([48]byte(record.Signature))

			if err != nil {
				panic(err)
			}

			newSignatures = append(newSignatures, currentSignature)
			break
		}
	}

	aggregate, err = bls.AggregateSignatures(newSignatures)

	if err != nil {
		return
	}

	signatureBytes := aggregate.Bytes()

	err = dbClient.CorrectnessReport.
		Create().
		SetCorrect(args.Info.Correct).
		SetSignersCount(uint64(len(signatures))).
		SetSignature(signatureBytes[:]).
		SetHash(args.Info.Hash[:]).
		SetTimestamp(args.Info.Timestamp).
		SetTopic(args.Info.Topic[:]).
		SetConsensus(args.Consensus).
		SetVoted(&helpers.BigInt{Int: *args.Voted}).
		AddSignerIDs(signerIDs...).
		OnConflictColumns("topic", "hash").
		UpdateNewValues().
		Exec(ctx)

	if err != nil {
		panic(err)
	}

	s.signatureCache.Remove(args.Hash)
}

func (s *Service) init() {
	var err error

	s.DebouncedSaveSignatures = utils.Debounce[bls12381.G1Affine, SaveSignatureArgs](5*time.Second, s.SaveSignatures)
	s.signatureMutex = new(sync.Mutex)
	s.supportedTopics = make(map[[64]byte]bool)
	s.signatureCache, err = lru.New[bls12381.G1Affine, []datasets.Signature](LruSize)

	if err != nil {
		panic(err)
	}
}

func New(ethRPC *ethereum.Repository, pos *pos.Repository) *Service {
	c := Service{
		ethRPC: ethRPC,
		pos:    pos,
	}
	c.init()

	for _, conf := range config.App.Plugins.Correctness {
		c.supportedTopics[[64]byte(shake.Shake([]byte(conf)))] = true
	}

	var err error
	c.consensus, err = lru.New[Key, xsync.MapOf[bls12381.G1Affine, big.Int]](evmlog.LruSize)
	if err != nil {
		log.Logger.
			Error("Failed to create correctness consensus cache.")
		os.Exit(1)
	}

	return &c
}
