package correctness

import (
	"context"
	"fmt"
	"github.com/KenshiTech/unchained/internal/model"
	"github.com/KenshiTech/unchained/internal/utils/address"
	"math/big"
	"os"
	"sync"
	"time"

	"github.com/KenshiTech/unchained/internal/consts"

	"github.com/KenshiTech/unchained/internal/repository"

	"github.com/KenshiTech/unchained/internal/crypto/ethereum"
	"github.com/KenshiTech/unchained/internal/pos"
	"github.com/KenshiTech/unchained/internal/service/evmlog"
	"github.com/puzpuzpuz/xsync/v3"

	"github.com/KenshiTech/unchained/internal/config"
	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/crypto/shake"
	"github.com/KenshiTech/unchained/internal/ent"
	"github.com/KenshiTech/unchained/internal/utils"
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
	Info      model.Correctness
	Hash      bls12381.G1Affine
	Consensus bool
	Voted     *big.Int
}

type Service struct {
	ethRPC          *ethereum.Repository
	pos             *pos.Repository
	signerRepo      repository.Signer
	correctnessRepo repository.CorrectnessReport

	signatureCache *lru.Cache[bls12381.G1Affine, []model.Signature]
	consensus      *lru.Cache[Key, xsync.MapOf[bls12381.G1Affine, big.Int]]

	DebouncedSaveSignatures func(key bls12381.G1Affine, arg SaveSignatureArgs)
	signatureMutex          *sync.Mutex
	supportedTopics         map[[64]byte]bool
}

func (s *Service) IsNewSigner(signature model.Signature, records []*ent.CorrectnessReport) bool {
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
	signer model.Signer,
	hash bls12381.G1Affine,
	info model.Correctness,
	debounce bool) error {
	if supported := s.supportedTopics[info.Topic]; !supported {
		return consts.ErrTokenNotSupported
	}

	s.signatureMutex.Lock()
	defer s.signatureMutex.Unlock()

	signatures, ok := s.signatureCache.Get(hash)

	if !ok {
		signatures = make([]model.Signature, 0)
	}

	// Check for duplicates
	for _, sig := range signatures {
		if sig.Signer.PublicKey == signer.PublicKey {
			return consts.ErrDuplicateSignature
		}
	}

	packed := model.Signature{
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
		utils.Logger.
			With("Address", address.Calculate(signer.PublicKey[:])).
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
		return nil
	}

	err = s.SaveSignatures(saveArgs)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) SaveSignatures(args SaveSignatureArgs) error {
	signatures, ok := s.signatureCache.Get(args.Hash)
	if !ok {
		return consts.ErrSignatureNotfound
	}

	ctx := context.Background()

	var newSigners []model.Signer
	var newSignatures []bls12381.G1Affine
	var keys [][]byte

	for i := range signatures {
		signature := signatures[i]
		keys = append(keys, signature.Signer.PublicKey[:])
	}

	currentRecords, err := s.correctnessRepo.Find(ctx, args.Info.Hash, args.Info.Topic[:], args.Info.Timestamp)
	if err != nil && !ent.IsNotFound(err) {
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

	signerIDs, err := s.signerRepo.GetSingerIDsByKeys(context.TODO(), keys)
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

func (s *Service) init() {
	var err error

	s.DebouncedSaveSignatures = utils.Debounce[bls12381.G1Affine, SaveSignatureArgs](5*time.Second, s.SaveSignatures)
	s.signatureMutex = new(sync.Mutex)
	s.supportedTopics = make(map[[64]byte]bool)
	s.signatureCache, err = lru.New[bls12381.G1Affine, []model.Signature](LruSize)

	if err != nil {
		panic(err)
	}
}

func New(
	ethRPC *ethereum.Repository,
	pos *pos.Repository,
	signerRepo repository.Signer,
	correctnessRepo repository.CorrectnessReport,
) *Service {
	c := Service{
		ethRPC:          ethRPC,
		pos:             pos,
		signerRepo:      signerRepo,
		correctnessRepo: correctnessRepo,
	}
	c.init()

	for _, conf := range config.App.Plugins.Correctness {
		c.supportedTopics[[64]byte(shake.Shake([]byte(conf)))] = true
	}

	var err error
	c.consensus, err = lru.New[Key, xsync.MapOf[bls12381.G1Affine, big.Int]](evmlog.LruSize)
	if err != nil {
		utils.Logger.
			Error("Failed to create correctness consensus cache.")
		os.Exit(1)
	}

	return &c
}
