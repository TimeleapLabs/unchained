package correctness

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/KenshiTech/unchained/address"
	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/crypto/bls"
	"github.com/KenshiTech/unchained/crypto/shake"
	"github.com/KenshiTech/unchained/datasets"
	"github.com/KenshiTech/unchained/db"
	"github.com/KenshiTech/unchained/ent"
	"github.com/KenshiTech/unchained/ent/signer"
	"github.com/KenshiTech/unchained/ethereum"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/pos"
	"github.com/KenshiTech/unchained/utils"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	lru "github.com/hashicorp/golang-lru/v2"
)

type CorrectnessKey struct {
	Topic   [64]byte
	Hash    [64]byte
	Correct bool
}

var consensus *lru.Cache[CorrectnessKey, map[bls12381.G1Affine]big.Int]
var signatureCache *lru.Cache[bls12381.G1Affine, []bls.Signature]
var aggregateCache *lru.Cache[bls12381.G1Affine, bls12381.G1Affine]
var DebouncedSaveSignatures func(key bls12381.G1Affine, arg SaveSignatureArgs)
var signatureMutex *sync.Mutex
var supportedTopics map[[64]byte]bool

type CorrectnessConf struct {
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
	signer bls.Signer,
	hash bls12381.G1Affine,
	info datasets.Correctness,
	debounce bool,
	historical bool) {

	if supported := supportedTopics[info.Topic]; !supported {
		return
	}

	signatureMutex.Lock()
	defer signatureMutex.Unlock()

	key := CorrectnessKey{
		Topic:   info.Topic,
		Hash:    info.Hash,
		Correct: info.Correct,
	}

	if !consensus.Contains(key) {
		consensus.Add(key, make(map[bls12381.G1Affine]big.Int))
	}

	consensusChain := config.Config.GetString("pos.chain")
	blockNumber, err := GetBlockNumber(consensusChain)

	if err != nil {
		log.Logger.
			With("Error", err).
			Error("Failed to get the latest block number")
		return
	}

	votingPower, err := pos.GetVotingPowerOfPublicKey(
		signer.PublicKey,
		big.NewInt(int64(*blockNumber)),
	)

	if err != nil {
		log.Logger.
			With("Address", address.Calculate(signer.PublicKey[:])).
			With("Error", err).
			Error("Failed to get voting power")
		return
	}

	reportedValues, _ := consensus.Get(key)
	voted := reportedValues[hash]
	totalVoted := new(big.Int).Add(votingPower, &voted)
	isMajority := true

	for _, reportCount := range reportedValues {
		if reportCount.Cmp(totalVoted) == 1 {
			isMajority = false
			break
		}
	}

	cached, _ := signatureCache.Get(hash)

	packed := bls.Signature{
		Signature: signature,
		Signer:    signer,
		Processed: false,
	}

	for _, item := range cached {
		if item.Signer.PublicKey == signer.PublicKey {
			return
		}
	}

	reportedValues[hash] = *totalVoted
	cached = append(cached, packed)
	signatureCache.Add(hash, cached)

	if isMajority {
		if debounce {
			DebouncedSaveSignatures(hash, SaveSignatureArgs{Hash: hash, Info: info})
		} else {
			SaveSignatures(SaveSignatureArgs{Hash: hash, Info: info})
		}
	}
}

func SaveSignatures(args SaveSignatureArgs) {

	dbClient := db.GetClient()
	signatures, ok := signatureCache.Get(args.Hash)

	if !ok {
		return
	}

	ctx := context.Background()

	var newSigners []bls.Signer
	var newSignatures []bls12381.G1Affine
	var keys [][]byte

	for i := range signatures {
		signature := signatures[i]
		keys = append(keys, signature.Signer.PublicKey[:])
		if !signature.Processed {
			newSignatures = append(newSignatures, signature.Signature)
			newSigners = append(newSigners, signature.Signer)
		}
	}

	// TODO: This part can be a shared library
	err := dbClient.Signer.MapCreateBulk(newSigners, func(sc *ent.SignerCreate, i int) {
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
	currentAggregate, ok := aggregateCache.Get(args.Hash)

	if ok {
		newSignatures = append(newSignatures, currentAggregate)
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

	for _, signature := range signatures {
		signature.Processed = true
	}

	aggregateCache.Add(args.Hash, aggregate)
}

func Setup() {
	if !config.Config.IsSet("plugins.correctness") {
		return
	}

	var configs []string
	if err := config.Config.UnmarshalKey("plugins.correctness", &configs); err != nil {
		panic(err)
	}

	for _, conf := range configs {
		supportedTopics[[64]byte(shake.Shake([]byte(conf)))] = true
	}
}

func init() {

	DebouncedSaveSignatures = utils.Debounce[bls12381.G1Affine, SaveSignatureArgs](5*time.Second, SaveSignatures)
	signatureMutex = new(sync.Mutex)
	supportedTopics = make(map[[64]byte]bool)

	var err error
	signatureCache, err = lru.New[bls12381.G1Affine, []bls.Signature](128)

	if err != nil {
		panic(err)
	}

	consensus, err = lru.New[CorrectnessKey, map[bls12381.G1Affine]big.Int](128)

	if err != nil {
		panic(err)
	}

	aggregateCache, err = lru.New[bls12381.G1Affine, bls12381.G1Affine](128)

	if err != nil {
		panic(err)
	}
}
