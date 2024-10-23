package attestation

import (
	"bytes"
	"context"
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

// Service represents the attestation service which confirm and store the attestation reports.
type Service interface {
	RecordSignature(
		ctx context.Context, signature bls12381.G1Affine, signer model.Signer, hash bls12381.G1Affine, info model.Attestation, debounce bool,
	) error
	SaveSignatures(ctx context.Context, args SaveSignatureArgs) error
}

type service struct {
	pos             pos.Service
	proofRepo       repository.Proof
	attestationRepo repository.Attestation

	signatureCache *lru.Cache[bls12381.G1Affine, []Signature]
	consensus      *lru.Cache[Key, xsync.MapOf[bls12381.G1Affine, big.Int]] // TODO: This needs to go

	DebouncedSaveSignatures func(key bls12381.G1Affine, arg SaveSignatureArgs)
	signatureMutex          *sync.Mutex
	supportedTopics         map[[64]byte]bool
}

// TODO: How should we handle older records?
// Possible Solution: Add a not after timestamp to the document.
func (s *service) RecordSignature(
	ctx context.Context, signature bls12381.G1Affine, signer model.Signer, hash bls12381.G1Affine, info model.Attestation, debounce bool,
) error {
	if supported := s.supportedTopics[[64]byte(info.Topic)]; !supported {
		utils.Logger.
			With("Topic", info.Topic).
			Debug("Topic not supported")
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
		if bytes.Equal(sig.Signer.PublicKey, signer.PublicKey) {
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
		utils.Logger.
			With("Address", address.Calculate(signer.PublicKey)).
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

// TODO: This part can be a shared library
func alreadySigned(signers []model.Signer, signer model.Signer) bool {
	for _, s := range signers {
		if bytes.Equal(s.PublicKey, signer.PublicKey) {
			return true
		}
	}
	return false
}

func (s *service) SaveSignatures(ctx context.Context, args SaveSignatureArgs) error {
	signatures, ok := s.signatureCache.Get(args.Hash)
	if !ok {
		return consts.ErrSignatureNotfound
	}

	blsHash := args.Info.Bls().Bytes()
	currentRecord, err := s.attestationRepo.Find(ctx, blsHash[:])
	if err != nil && err != consts.ErrRecordNotfound {
		return err
	}

	currentProof, err := s.proofRepo.Find(ctx, blsHash)
	if err != nil && err != consts.ErrRecordNotfound {
		return err
	}

	var newSigners []model.Signer
	var newSignatures []bls12381.G1Affine
	// Select the new signers and signatures
	for _, signature := range signatures {
		if !alreadySigned(currentProof.Signers, signature.Signer) {
			newSigners = append(newSigners, signature.Signer)
			newSignatures = append(newSignatures, signature.Signature)
		}
	}

	if len(newSigners) == 0 {
		return consts.ErrNoNewSigners
	}

	// TODO: This part can be a shared library
	if currentRecord.Signature != nil {
		currentSignature, err := bls.RecoverSignature([48]byte(currentRecord.Signature))

		if err != nil {
			return err
		}

		newSignatures = append(newSignatures, currentSignature)
	}

	var aggregate bls12381.G1Affine
	aggregate, err = bls.AggregateSignatures(newSignatures)
	if err != nil {
		return consts.ErrCantAggregateSignatures
	}

	signatureBytes := aggregate.Bytes()

	err = s.proofRepo.CreateProof(ctx, blsHash, newSigners)
	if err != nil {
		return err
	}

	err = s.attestationRepo.Upsert(ctx, model.Attestation{
		SignersCount: uint64(len(newSigners) + len(currentProof.Signers)),
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
	pos pos.Service, proofRepo repository.Proof, attestationRepo repository.Attestation,
) Service {
	c := service{
		pos:             pos,
		proofRepo:       proofRepo,
		attestationRepo: attestationRepo,
	}

	var err error
	c.DebouncedSaveSignatures = utils.Debounce[bls12381.G1Affine, SaveSignatureArgs](5*time.Second, c.SaveSignatures)
	c.signatureMutex = new(sync.Mutex)
	c.supportedTopics = make(map[[64]byte]bool)
	c.signatureCache, err = lru.New[bls12381.G1Affine, []Signature](LruSize)
	if err != nil {
		panic(err)
	}

	for _, conf := range config.App.Plugins.Attestation {
		c.supportedTopics[[64]byte(utils.Shake([]byte(conf)))] = true
	}

	c.consensus, err = lru.New[Key, xsync.MapOf[bls12381.G1Affine, big.Int]](LruSize)
	if err != nil {
		utils.Logger.
			Error("Failed to create attestation consensus cache.")
		os.Exit(1)
	}

	return &c
}
