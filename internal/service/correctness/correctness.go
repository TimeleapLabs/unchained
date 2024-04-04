package correctness

import (
	"context"
	"sync"
	"time"

	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/crypto/bls"
	"github.com/KenshiTech/unchained/crypto/shake"
	"github.com/KenshiTech/unchained/datasets"
	"github.com/KenshiTech/unchained/db"
	"github.com/KenshiTech/unchained/ent"
	"github.com/KenshiTech/unchained/ent/correctnessreport"
	"github.com/KenshiTech/unchained/ent/signer"
	"github.com/KenshiTech/unchained/ethereum"
	"github.com/KenshiTech/unchained/utils"
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
	Info datasets.Correctness
	Hash bls12381.G1Affine
}
type Service struct {
	ethRPC *ethereum.Repository

	signatureCache          *lru.Cache[bls12381.G1Affine, []datasets.Signature]
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
		Processed: false,
	}

	signatures = append(signatures, packed)
	s.signatureCache.Add(hash, signatures)

	if debounce {
		s.DebouncedSaveSignatures(hash, SaveSignatureArgs{Hash: hash, Info: info})
	} else {
		s.SaveSignatures(SaveSignatureArgs{Hash: hash, Info: info})
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

	currentRecord, err := dbClient.CorrectnessReport.
		Query().
		Where(correctnessreport.And(
			correctnessreport.Hash(args.Info.Hash[:]),
			correctnessreport.Topic(args.Info.Topic[:]),
			correctnessreport.Timestamp(args.Info.Timestamp),
			correctnessreport.Correct(args.Info.Correct),
		)).
		Only(ctx)

	if err != nil && !ent.IsNotFound(err) {
		panic(err)
	}

	// Select the new signers and signatures
	for i := range signatures {
		signature := signatures[i]

		if currentRecord != nil {
			for _, signer := range currentRecord.Edges.Signers {
				if signature.Signer.PublicKey == [96]byte(signer.Key) {
					continue
				}
			}
		}

		newSigners = append(newSigners, signature.Signer)
		newSignatures = append(newSignatures, signature.Signature)
	}

	// TODO: This part can be a shared library
	err = dbClient.Signer.MapCreateBulk(newSigners, func(sc *ent.SignerCreate, i int) {
		signer := newSigners[i]
		sc.SetName(signer.Name).
			SetEvm(signer.EvmWallet).
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

	signerIds, err := dbClient.Signer.
		Query().
		Where(signer.KeyIn(keys...)).
		IDs(ctx)

	if err != nil {
		return
	}

	var aggregate bls12381.G1Affine

	if currentRecord != nil {
		currentSignature, err := bls.RecoverSignature([48]byte(currentRecord.Signature))

		if err != nil {
			panic(err)
		}

		newSignatures = append(newSignatures, currentSignature)
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
		AddSignerIDs(signerIds...).
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

func New(ethRPC *ethereum.Repository) *Service {
	c := Service{
		ethRPC: ethRPC,
	}
	c.init()

	for _, conf := range config.App.Plugins.Correctness {
		c.supportedTopics[[64]byte(shake.Shake([]byte(conf)))] = true
	}

	return &c
}
