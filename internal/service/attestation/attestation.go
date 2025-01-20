package attestation

import (
	"bytes"
	"context"
	"errors"
	"time"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/repository"
	"github.com/TimeleapLabs/unchained/internal/service/pos"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"golang.org/x/crypto/ed25519"

	lru "github.com/hashicorp/golang-lru/v2"
)

const (
	LruSize = 128
)

// Service represents the attestation service which confirm and store the attestation reports.
type Service interface {
	RecordSignature(
		ctx context.Context,
		signature [64]byte,
		signer ed25519.PublicKey,
		hash [32]byte,
		info model.Attestation,
		debounce bool,
	) error
	SaveSignatures(ctx context.Context, args SaveSignatureArgs) error
}

type service struct {
	pos             pos.Service
	proofRepo       repository.Proof
	attestationRepo repository.Attestation
	signatureCache  *lru.Cache[[32]byte, []model.Signature]

	DebouncedSaveSignatures func(key [32]byte, arg SaveSignatureArgs)
	supportedTopics         map[[64]byte]bool
}

func (s *service) RecordSignature(
	ctx context.Context, signature [64]byte, signer ed25519.PublicKey, hash [32]byte, info model.Attestation, debounce bool,
) error {
	if supported := s.supportedTopics[[64]byte(info.Topic)]; !supported {
		utils.Logger.
			With("Topic", info.Topic).
			Debug("Topic not supported")
		return consts.ErrTopicNotSupported
	}

	signatures, ok := s.signatureCache.Get(hash)
	if !ok {
		signatures = []model.Signature{}
	}

	// Check for duplicates
	for _, sig := range signatures {
		if bytes.Equal(sig.PublicKey[:], signer[:]) {
			return consts.ErrDuplicateSignature
		}
	}

	signatures = append(signatures, model.Signature{
		Signature: signature,
		PublicKey: signer,
	})

	s.signatureCache.Add(hash, signatures)

	saveArgs := SaveSignatureArgs{
		Info: info,
		Hash: hash,
	}

	if debounce {
		s.DebouncedSaveSignatures(hash, saveArgs)
		return nil
	}

	err := s.SaveSignatures(ctx, saveArgs)
	if err != nil {
		return err
	}

	return nil
}

func alreadySigned(signers []model.Signature, pk []byte) bool {
	for _, s := range signers {
		if bytes.Equal(s.PublicKey[:], pk) {
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

	currentProof, err := s.proofRepo.Find(ctx, args.Hash)
	if err != nil && errors.Is(err, consts.ErrRecordNotfound) {
		return err
	}

	var newSignatures []model.Signature
	// Select the new signers and signatures
	for _, signature := range signatures {
		if !alreadySigned(currentProof.Signatures, signature.PublicKey[:]) {
			newSignatures = append(newSignatures, signature)
		}
	}

	if len(newSignatures) == 0 {
		return consts.ErrNoNewSigners
	}

	err = s.proofRepo.CreateProof(ctx, args.Hash, newSignatures)
	if err != nil {
		return err
	}

	if _, err := s.attestationRepo.Find(ctx, args.Hash); err != nil && errors.Is(err, consts.ErrRecordNotfound) {
		err = s.attestationRepo.Upsert(ctx, args.Hash, model.Attestation{
			Timestamp: args.Info.Timestamp,
			Topic:     args.Info.Topic,
			Meta:      args.Info.Meta,
		})

		if err != nil {
			return err
		}
	}

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

	c.DebouncedSaveSignatures = utils.Debounce[[32]byte, SaveSignatureArgs](5*time.Second, c.SaveSignatures)
	c.supportedTopics = make(map[[64]byte]bool)

	var err error
	c.signatureCache, err = lru.New[[32]byte, []model.Signature](LruSize)
	if err != nil {
		panic(err)
	}

	return &c
}
