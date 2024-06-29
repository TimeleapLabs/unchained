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
	IsNewSigner(signature Signature, records []model.Correctness) bool
	RecordSignature(
		ctx context.Context, signature bls12381.G1Affine, signer model.Signer, info model.Correctness, debounce bool,
	) error
}

type service struct {
	pos             pos.Service
	signerRepo      repository.Signer
	correctnessRepo repository.CorrectnessReport

	signatureCache *lru.Cache[bls12381.G1Affine, []Signature]
	consensus      *lru.Cache[Key, xsync.MapOf[bls12381.G1Affine, big.Int]]

	DebouncedSaveSignatures func(key bls12381.G1Affine, arg SaveSignatureArgs)
	signatureMutex          *sync.Mutex
	supportedTopics         map[[64]byte]bool
}

// IsNewSigner checks if the signer's pub key is in the records signers or not.
func (s *service) IsNewSigner(signature Signature, records []model.Correctness) bool {
	// TODO: This isn't efficient, we should use a map
	for _, record := range records {
		for _, signer := range record.Signers {
			if signature.Signer.PublicKey == signer.PublicKey {
				return false
			}
		}
	}

	return true
}

// TODO: How should we handle older records?
// Possible Solution: Add a not after timestamp to the document.
func (s *service) RecordSignature(
	ctx context.Context, signature bls12381.G1Affine, signer model.Signer, info model.Correctness, debounce bool,
) error {
	if supported := s.supportedTopics[info.Topic]; !supported {
		utils.Logger.
			With("Topic", info.Topic).
			Debug("Token not supported")
		return consts.ErrTopicNotSupported
	}

	s.signatureMutex.Lock()
	defer s.signatureMutex.Unlock()

	signatures, ok := s.signatureCache.Get(info.Bls())
	if !ok {
		signatures = make([]Signature, 0)
	}

	// Check for duplicates
	for _, sig := range signatures {
		if sig.Signer.PublicKey == signer.PublicKey {
			return consts.ErrDuplicateSignature
		}
	}

	packed := Signature{
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
	voted, ok := reportedValues.Load(info.Bls())
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

	reportedValues.Range(func(_ bls12381.G1Affine, value big.Int) bool {
		if value.Cmp(totalVoted) == 1 {
			isMajority = false
		}
		return isMajority
	})

	reportedValues.Store(info.Bls(), *totalVoted)
	signatures = append(signatures, packed)
	s.signatureCache.Add(info.Bls(), signatures)

	saveArgs := SaveSignatureArgs{
		Info:      info,
		Hash:      info.Bls(),
		Consensus: isMajority,
		Voted:     totalVoted,
	}

	if debounce {
		s.DebouncedSaveSignatures(info.Bls(), saveArgs)
		return nil
	}

	err = s.saveSignatures(ctx, saveArgs)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) saveSignatures(ctx context.Context, args SaveSignatureArgs) error {
	signatures, ok := s.signatureCache.Get(args.Hash)
	if !ok {
		return consts.ErrSignatureNotfound
	}

	var newSigners []model.Signer
	var newSignatures []bls12381.G1Affine
	var keys [][]byte

	for i := range signatures {
		signature := signatures[i]
		keys = append(keys, signature.Signer.PublicKey[:])
	}

	currentRecords, err := s.correctnessRepo.Find(ctx, args.Info.Hash, args.Info.Topic[:], args.Info.Timestamp)
	if err != nil {
		return err
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

	err = s.signerRepo.CreateSigners(ctx, newSigners)
	if err != nil {
		return err
	}

	signerIDs, err := s.signerRepo.GetSingerIDsByKeys(ctx, keys)
	if err != nil {
		return err
	}

	var aggregate bls12381.G1Affine

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

	aggregate, err = bls.AggregateSignatures(newSignatures)
	if err != nil {
		return consts.ErrCantAggregateSignatures
	}

	signatureBytes := aggregate.Bytes()

	err = s.correctnessRepo.Upsert(ctx, model.Correctness{
		SignersCount: uint64(len(signatures)),
		Signature:    signatureBytes[:],
		Consensus:    args.Consensus,
		Voted:        *args.Voted,
		SignerIDs:    signerIDs,
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

func (s *service) init() {
	var err error

	s.DebouncedSaveSignatures = utils.Debounce[bls12381.G1Affine, SaveSignatureArgs](5*time.Second, s.saveSignatures)
	s.signatureCache, err = lru.New[bls12381.G1Affine, []Signature](LruSize)

	if err != nil {
		panic(err)
	}
}

func New(
	pos pos.Service,
	signerRepo repository.Signer,
	correctnessRepo repository.CorrectnessReport,
) Service {
	c := service{
		pos:             pos,
		signerRepo:      signerRepo,
		correctnessRepo: correctnessRepo,
		signatureMutex:  new(sync.Mutex),
		supportedTopics: make(map[[64]byte]bool),
	}
	c.init()

	for _, conf := range config.App.Plugins.Correctness {
		c.supportedTopics[[64]byte(utils.Shake([]byte(conf)))] = true
	}

	var err error
	c.consensus, err = lru.New[Key, xsync.MapOf[bls12381.G1Affine, big.Int]](LruSize)
	if err != nil {
		utils.Logger.
			Error("Failed to create correctness consensus cache.")
		os.Exit(1)
	}

	return &c
}
