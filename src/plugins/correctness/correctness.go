package correctness

import (
	"context"
	"sync"
	"time"

	"github.com/KenshiTech/unchained/src/config"
	"github.com/KenshiTech/unchained/src/crypto/bls"
	"github.com/KenshiTech/unchained/src/crypto/shake"
	"github.com/KenshiTech/unchained/src/datasets"
	"github.com/KenshiTech/unchained/src/db"
	"github.com/KenshiTech/unchained/src/ent"
	"github.com/KenshiTech/unchained/src/ent/correctnessreport"
	"github.com/KenshiTech/unchained/src/ent/signer"
	"github.com/KenshiTech/unchained/src/ethereum"
	"github.com/KenshiTech/unchained/src/utils"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	lru "github.com/hashicorp/golang-lru/v2"
)

var signatureCache *lru.Cache[bls12381.G1Affine, []datasets.Signature]
var DebouncedSaveSignatures func(key bls12381.G1Affine, arg SaveSignatureArgs)
var signatureMutex *sync.Mutex
var supportedTopics map[[64]byte]bool

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

// TODO: This code should be moved to a shared library
func GetBlockNumber(network string) (*uint64, error) {
	blockNumber, err := ethereum.GetBlockNumber(network)

	if err != nil {
		ethereum.RefreshRPC(network)
		return nil, err
	}

	return &blockNumber, nil
}

func RecordSignature(
	signature bls12381.G1Affine,
	signer datasets.Signer,
	hash bls12381.G1Affine,
	info datasets.Correctness,
	debounce bool) {
	if supported := supportedTopics[info.Topic]; !supported {
		return
	}

	signatureMutex.Lock()
	defer signatureMutex.Unlock()

	signatures, ok := signatureCache.Get(hash)

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
	signatureCache.Add(hash, signatures)

	if debounce {
		DebouncedSaveSignatures(hash, SaveSignatureArgs{Hash: hash, Info: info})
	} else {
		SaveSignatures(SaveSignatureArgs{Hash: hash, Info: info})
	}
}

func SaveSignatures(args SaveSignatureArgs) {
	dbClient := db.GetClient()
	signatures, ok := signatureCache.Get(args.Hash)

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

	signatureCache.Remove(args.Hash)
}

func New() {
	for _, conf := range config.App.Plugins.EthLog.Correctness {
		supportedTopics[[64]byte(shake.Shake([]byte(conf)))] = true
	}

	for _, conf := range config.App.Plugins.Uniswap.Correctness {
		supportedTopics[[64]byte(shake.Shake([]byte(conf)))] = true
	}
}

func init() {
	DebouncedSaveSignatures = utils.Debounce[bls12381.G1Affine, SaveSignatureArgs](5*time.Second, SaveSignatures)
	signatureMutex = new(sync.Mutex)
	supportedTopics = make(map[[64]byte]bool)

	var err error
	signatureCache, err = lru.New[bls12381.G1Affine, []datasets.Signature](LruSize)

	if err != nil {
		panic(err)
	}
}
