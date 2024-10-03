package correctness

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/service/pos"
	"github.com/TimeleapLabs/unchained/internal/utils/address"
	"math/big"
	"os"
	"sync"

	"github.com/puzpuzpuz/xsync/v3"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/utils"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	lru "github.com/hashicorp/golang-lru/v2"
)

const (
	LruSize = 128
)

// Service represents the correctness service which confirm and store the correctness reports.
type Service interface {
	RecordSignature(
		ctx context.Context, signature bls12381.G1Affine, signer model.Signer, hash bls12381.G1Affine, info model.Correctness,
	) error
}

type service struct {
	pos pos.Service

	signatureCache *lru.Cache[bls12381.G1Affine, []Signature]
	consensus      *lru.Cache[Key, xsync.MapOf[bls12381.G1Affine, big.Int]]

	signatureMutex  *sync.Mutex
	supportedTopics map[[64]byte]bool
}

// TODO: How should we handle older records?
// Possible Solution: Add a not after timestamp to the document.
func (s *service) RecordSignature(
	ctx context.Context, signature bls12381.G1Affine, signer model.Signer, hash bls12381.G1Affine, info model.Correctness,
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
			utils.Logger.With("Err", err).ErrorContext(ctx, "Can't decode public key")
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

	//saveArgs := SaveSignatureArgs{
	//	Info:      info,
	//	Hash:      hash,
	//	Consensus: isMajority,
	//	Voted:     totalVoted,
	//}

	// Maybe send to network...

	return nil
}

func New(pos pos.Service) Service {
	c := service{
		pos: pos,
	}

	var err error
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
