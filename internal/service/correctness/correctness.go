package correctness

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"sync"
	"time"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/repository"
	"github.com/TimeleapLabs/unchained/internal/service/pos"
	"github.com/TimeleapLabs/unchained/internal/utils/address"

	"github.com/puzpuzpuz/xsync/v3"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/utils"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	lru "github.com/hashicorp/golang-lru/v2"
)

const (
	LruSize = 128
)

type SaveSignatureArgs struct {
	Info      model.Correctness
	Hash      bls12381.G1Affine
	Consensus bool
	Voted     *big.Int
}

type Service interface {
	RecordSignature(
		ctx context.Context, signature bls12381.G1Affine, signer model.Signer, hash bls12381.G1Affine, info model.Correctness, debounce bool,
	) error
	SaveSignatures(ctx context.Context, args SaveSignatureArgs) error
}

type service struct {
	pos             pos.Service
	proofRepo       repository.Proof
	correctnessRepo repository.CorrectnessReport

	signatureCache *lru.Cache[bls12381.G1Affine, []Signature]
	consensus      *lru.Cache[Key, xsync.MapOf[bls12381.G1Affine, big.Int]]

	DebouncedSaveSignatures func(key bls12381.G1Affine, arg SaveSignatureArgs)
	signatureMutex          *sync.Mutex
	supportedTopics         map[[64]byte]bool
}

// TODO: How should we handle older records?
// Possible Solution: Add a not after timestamp to the document.
func (s *service) RecordSignature(
	ctx context.Context, signature bls12381.G1Affine, signer model.Signer, hash bls12381.G1Affine, info model.Correctness, debounce bool,
) error {
	if supported := s.supportedTopics[[64]byte(info.Topic)]; !supported {
		utils.Logger.
			With("Topic", info.Topic).
			Debug("Token not supported")
		return consts.ErrTopicNotSupported
	}

	s.signatureMutex.Lock()
	defer s.signatureMutex.Unlock()

	signatures, ok := s.signatureCache.Get(hash)
	if !ok {
		signatures = make([]Signature, 0)
	}

	// Check for duplicates
	for _, sig := range signatures {
		if sig.Signer.PublicKey == signer.PublicKey {
			return consts.ErrDuplicateSignature
		}
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
	voted, ok := reportedValues.Load(hash)
	if !ok {
		voted = *big.NewInt(0)
	}

	votingPower, err := s.pos.GetVotingPowerOfEvm(ctx, signer.EvmAddress)
	if err != nil {
		publicKeyBytes, err := hex.DecodeString(signer.PublicKey)
		if err != nil {
			utils.Logger.Error("Can't decode public key: %v", err)
			return err
		}

		utils.Logger.
			With("Address", address.Calculate(publicKeyBytes)).
			With("Error", err).
			Error("Failed to get voting power")
		return err
	}

	totalVoted := new(big.Int).Add(votingPower, &voted)

	isMajority := true
	reportedValues.Range(func(_ bls12381.G1Affine, value big.Int) bool {
		if value.Cmp(totalVoted) == 1 {
			isMajority = false
		}
		return isMajority
	})

	reportedValues.Store(hash, *totalVoted)
	signatures = append(signatures, Signature{
		Signature: signature,
		Signer:    signer,
	})
	s.signatureCache.Add(hash, signatures)

	saveArgs := SaveSignatureArgs{
		Info:      info,
		Hash:      hash,
		Consensus: isMajority,
		Voted:     totalVoted,
	}

	if debounce {
		s.DebouncedSaveSignatures(hash, saveArgs)
		return nil
	}

	err = s.SaveSignatures(ctx, saveArgs)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) SaveSignatures(ctx context.Context, args SaveSignatureArgs) error {
	signatures, ok := s.signatureCache.Get(args.Hash)
	if !ok {
		return consts.ErrSignatureNotfound
	}

	currentRecords, err := s.correctnessRepo.Find(ctx, args.Info.Hash, args.Info.Topic, args.Info.Timestamp)
	if err != nil {
		return err
	}

	var newSigners []model.Signer
	var newSignatures []bls12381.G1Affine
	// Select the new signers and signatures
	for _, signature := range signatures {
		newSigners = append(newSigners, signature.Signer)
		newSignatures = append(newSignatures, signature.Signature)
	}

	// TODO: This part can be a shared library
	for _, record := range currentRecords {
		if record.Correct == args.Info.Correct {
			currentSignature, err := bls.RecoverSignature([48]byte(record.Signature))

			if err != nil {
				return err
			}

			newSignatures = append(newSignatures, currentSignature)
			break
		}
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

	err = s.correctnessRepo.Upsert(ctx, model.Correctness{
		SignersCount: uint64(len(signatures)),
		Signature:    signatureBytes[:],
		Consensus:    args.Consensus,
		Voted:        args.Voted.Int64(),
		Timestamp:    args.Info.Timestamp,
		Hash:         args.Info.Hash,
		Topic:        args.Info.Topic,
		Correct:      args.Info.Correct,
	})
	if err != nil {
		return err
	}

	s.signatureCache.Remove(args.Hash)

	return nil
}

func New(
	pos pos.Service, proofRepo repository.Proof, correctnessRepo repository.CorrectnessReport,
) Service {
	c := service{
		pos:             pos,
		proofRepo:       proofRepo,
		correctnessRepo: correctnessRepo,
	}

	var err error
	c.DebouncedSaveSignatures = utils.Debounce[bls12381.G1Affine, SaveSignatureArgs](5*time.Second, c.SaveSignatures)
	c.signatureMutex = new(sync.Mutex)
	c.supportedTopics = make(map[[64]byte]bool)
	c.signatureCache, err = lru.New[bls12381.G1Affine, []Signature](LruSize)
	if err != nil {
		panic(err)
	}

	for _, conf := range config.App.Plugins.Correctness {
		c.supportedTopics[[64]byte(utils.Shake([]byte(conf)))] = true
	}

	c.consensus, err = lru.New[Key, xsync.MapOf[bls12381.G1Affine, big.Int]](LruSize)
	if err != nil {
		utils.Logger.
			Error("Failed to create correctness consensus cache.")
		os.Exit(1)
	}

	return &c
}
