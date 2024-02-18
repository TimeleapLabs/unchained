package uniswap

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/KenshiTech/unchained/bls"
	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/datasets"
	"github.com/KenshiTech/unchained/db"
	"github.com/KenshiTech/unchained/ent"
	"github.com/KenshiTech/unchained/ent/signer"
	"github.com/KenshiTech/unchained/ethereum"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/net/client"
	"github.com/KenshiTech/unchained/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/go-co-op/gocron/v2"
	"github.com/gorilla/websocket"
	"github.com/vmihailenco/msgpack/v5"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	lru "github.com/hashicorp/golang-lru/v2"
)

var DebouncedSaveSignatures func(key uint64, arg uint64)
var signatureMutex *sync.Mutex

var priceCache *lru.Cache[uint64, big.Int]
var signatureCache *lru.Cache[uint64, []bls.Signature]
var aggregateCache *lru.Cache[uint64, bls12381.G1Affine]

var twoOneNinetyTwo big.Int
var tenEighteen big.Int
var tenEighteenF big.Float
var lastBlock uint64
var lastPrice big.Int

type Token struct {
	Name   string `mapstructure:"name"`
	Pair   string `mapstructure:"pair"`
	Delta  int64  `mapstructure:"delta"`
	Invert bool   `mapstructure:"invert"`
}

// TODO: This needs to work with different datasets
func RecordSignature(signature bls12381.G1Affine, signer bls.Signer, block uint64) {

	signatureMutex.Lock()
	defer signatureMutex.Unlock()

	// TODO: Needs optimization
	if !priceCache.Contains(block) {

		var tokens []Token
		if err := config.Config.UnmarshalKey("plugins.uniswap.tokens", &tokens); err != nil {
			panic(err)
		}

		eth := tokens[0]
		blockNumber, _, err := GetPriceFromPair(eth.Pair, eth.Delta, eth.Invert)

		if err != nil {
			return
		}

		lastBlock = *blockNumber
	}

	if lastBlock-block > 16 {
		return // Data too old
	}

	cached, ok := signatureCache.Get(block)
	packed := bls.Signature{Signature: signature, Signer: signer, Processed: false}

	if !ok {
		signatureCache.Add(block, []bls.Signature{packed})
		// TODO: This looks ugly
		DebouncedSaveSignatures(block, block)
		return
	}

	for _, item := range cached {
		if item.Signer.PublicKey == signer.PublicKey {
			return
		}
	}

	cached = append(cached, packed)
	signatureCache.Add(block, cached)

	DebouncedSaveSignatures(block, block)
}

func SaveSignatures(block uint64) {

	dbClient := db.GetClient()
	price, ok := priceCache.Get(block)

	if !ok {
		return
	}

	signatures, ok := signatureCache.Get(block)

	if !ok {
		return
	}

	ctx := context.Background()

	// TODO: Cache this
	datasetId, err := dbClient.DataSet.
		Create().
		SetName("uniswap::ethereum::ethereum").
		OnConflictColumns("name").
		UpdateName().
		ID(ctx)

	if err != nil {
		panic(err)
	}

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

	err = dbClient.Signer.MapCreateBulk(newSigners, func(sc *ent.SignerCreate, i int) {
		signer := newSigners[i]
		sc.SetName(signer.Name).
			SetKey(signer.PublicKey[:]).
			SetShortkey(signer.ShortPublicKey[:]).
			SetPoints(0)
	}).
		OnConflictColumns("shortkey").
		UpdateName().
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
	currentAggregate, ok := aggregateCache.Get(block)

	if ok {
		newSignatures = append(newSignatures, currentAggregate)
	}

	aggregate, err = bls.AggregateSignatures(newSignatures)

	if err != nil {
		return
	}

	signatureBytes := aggregate.Bytes()

	err = dbClient.AssetPrice.
		Create().
		SetBlock(block).
		SetPrice(&price).
		SetSignersCount(uint64(len(signatures))).
		SetSignature(signatureBytes[:]).
		AddDataSetIDs(datasetId).
		AddSignerIDs(signerIds...).
		OnConflictColumns("block").
		UpdateNewValues().
		Exec(ctx)

	if err != nil {
		panic(err)
	}

	for _, signature := range signatures {
		signature.Processed = true
	}

	aggregateCache.Add(block, aggregate)

}

// TODO: Each pair should have its own LRU-Cache
func GetPriceFromCache(block uint64) (big.Int, bool) {
	return priceCache.Get(block)
}

func GetPriceFromPair(pairAddr string, decimalDif int64, inverse bool) (*uint64, *big.Int, error) {
	blockNumber, err := ethereum.GetBlockNumber()

	if err != nil {
		ethereum.RefreshRPC()
		return nil, nil, err
	}

	if blockNumber == lastBlock {
		return &blockNumber, &lastPrice, nil
	}

	pair, err := ethereum.GetNewUniV3Contract(pairAddr, false)

	if err != nil {
		ethereum.RefreshRPC()
		return nil, nil, err
	}

	data, err := pair.Slot0(
		&bind.CallOpts{
			BlockNumber: big.NewInt(int64(blockNumber)),
		})

	if err != nil {
		ethereum.RefreshRPC()
		return nil, nil, err
	}

	lastPrice = *priceFromSqrtX96(data.SqrtPriceX96, 6, true)
	priceCache.Add(blockNumber, lastPrice)

	return &blockNumber, &lastPrice, nil
}

func priceFromSqrtX96(sqrtPriceX96 *big.Int, decimalDif int64, inverse bool) *big.Int {
	var priceX96 big.Int
	var raw big.Int
	var price big.Int
	var factor big.Int

	// const raw = (fetchedSqrtPriceX96**2 / 2**192) * 10**6;
	priceX96.Exp(sqrtPriceX96, big.NewInt(2), nil)
	raw.Div(&priceX96, &twoOneNinetyTwo)

	if inverse {
		factor.Exp(big.NewInt(10), big.NewInt(36-decimalDif), nil)
		price.Div(&factor, &raw)
	} else {
		// TODO: needs work
		factor.Exp(big.NewInt(10), big.NewInt(decimalDif), nil)
		price.Div(&raw, &factor)
	}
	return &price
}

func Start() {

	scheduler, err := gocron.NewScheduler()

	if err != nil {
		panic(err)
	}

	var tokens []Token
	if err := config.Config.UnmarshalKey("plugins.uniswap.tokens", &tokens); err != nil {
		panic(err)
	}

	eth := tokens[0]

	_, err = scheduler.NewJob(
		gocron.DurationJob(5*time.Second),
		gocron.NewTask(
			func() {

				if client.IsClientSocketClosed {
					return
				}

				blockNumber, price, err := GetPriceFromPair(eth.Pair, eth.Delta, eth.Invert)

				if err != nil {
					return
				}

				if lastBlock == *blockNumber {
					return
				}

				lastBlock = *blockNumber

				var priceF big.Float
				priceF.Quo(new(big.Float).SetInt(price), &tenEighteenF)

				priceStr := fmt.Sprintf("$%.18f", &priceF)

				log.Logger.
					With("Block", *blockNumber).
					With("Price", priceStr).
					Info("Ethereum")

				priceInfo := datasets.PriceInfo{Price: *price, Block: *blockNumber}
				toHash, err := msgpack.Marshal(&priceInfo)

				if err != nil {
					panic(err)
				}

				signature, _ := bls.Sign(*bls.ClientSecretKey, toHash)
				compressedSignature := signature.Bytes()

				priceReport := datasets.PriceReport{
					PriceInfo: priceInfo,
					Signature: compressedSignature,
				}

				payload, err := msgpack.Marshal(&priceReport)

				if err != nil {
					panic(err)
				}

				if !client.IsClientSocketClosed {
					client.Client.WriteMessage(websocket.BinaryMessage, append([]byte{1, 0}, payload...))
				}

			},
		),
	)

	if err != nil {
		panic(err)
	}

	scheduler.Start()
}

func init() {

	DebouncedSaveSignatures = utils.Debounce[uint64, uint64](5*time.Second, SaveSignatures)
	signatureMutex = new(sync.Mutex)

	twoOneNinetyTwo.Exp(big.NewInt(2), big.NewInt(192), nil)
	tenEighteen.Exp(big.NewInt(10), big.NewInt(18), nil)
	tenEighteenF.SetInt(&tenEighteen)

	var err error
	priceCache, err = lru.New[uint64, big.Int](24)

	if err != nil {
		panic(err)
	}

	signatureCache, err = lru.New[uint64, []bls.Signature](24)

	if err != nil {
		panic(err)
	}

	aggregateCache, err = lru.New[uint64, bls12381.G1Affine](24)

	if err != nil {
		panic(err)
	}
}
