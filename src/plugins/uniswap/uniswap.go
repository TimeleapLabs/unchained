package uniswap

import (
	"context"
	"fmt"
	"math/big"
	"strings"
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
	"github.com/KenshiTech/unchained/net/consumer"
	"github.com/KenshiTech/unchained/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/go-co-op/gocron/v2"
	"github.com/gorilla/websocket"
	"github.com/vmihailenco/msgpack/v5"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	lru "github.com/hashicorp/golang-lru/v2"
)

var DebouncedSaveSignatures func(key AssetKey, arg datasets.PriceInfo)
var signatureMutex *sync.Mutex

var priceCache map[string]*lru.Cache[uint64, big.Int]
var signatureCache *lru.Cache[AssetKey, []bls.Signature]
var aggregateCache *lru.Cache[AssetKey, bls12381.G1Affine]

var twoOneNineTwo big.Int
var tenEighteen big.Int
var tenEighteenF big.Float
var lastBlock map[string]uint64
var crossPrices map[string]big.Int
var lastPrice big.Int
var caser cases.Caser

type Token struct {
	Id     *string  `mapstructure:"id"`
	Chain  string   `mapstructure:"chain"`
	Name   string   `mapstructure:"name"`
	Pair   string   `mapstructure:"pair"`
	Unit   string   `mapstructure:"unit"`
	Symbol string   `mapstructure:"symbol"`
	Delta  int64    `mapstructure:"delta"`
	Invert bool     `mapstructure:"invert"`
	Cross  []string `mapstructure:"cross"`
}

type AssetKey struct {
	Asset string
	Pair  string
	Chain string
	Block uint64
}

// TODO: This needs to work with different datasets
func RecordSignature(
	signature bls12381.G1Affine,
	signer bls.Signer,
	info datasets.PriceInfo) {

	signatureMutex.Lock()
	defer signatureMutex.Unlock()

	// TODO: This needs to take `chain` into account
	lruCache := priceCache[strings.ToLower(info.Pair)]

	if lruCache == nil {
		return
	}

	// TODO: Needs optimization
	if !lruCache.Contains(info.Block) {

		var tokens []Token
		if err := config.Config.UnmarshalKey("plugins.uniswap.tokens", &tokens); err != nil {
			panic(err)
		}

		var found Token

		for _, token := range tokens {
			if strings.EqualFold(token.Pair, info.Pair) &&
				strings.EqualFold(token.Name, info.Asset) {
				found = token
				break
			}
		}

		if len(found.Pair) == 0 {
			return
		}

		blockNumber, _, err := GetPriceFromPair(
			info.Chain,
			found.Pair,
			found.Delta,
			found.Invert,
		)

		if err != nil {
			return
		}

		lastBlock[info.Chain] = *blockNumber
	}

	if lastBlock[info.Chain]-info.Block > 16 {
		return // Data too old
	}

	key := AssetKey{
		Block: info.Block,
		Asset: info.Asset,
		Chain: info.Chain,
		Pair:  info.Pair,
	}

	cachedPrice, ok := lruCache.Get(info.Block)

	if !ok {
		return
	}

	if cachedPrice.Cmp(&info.Price) != 0 {
		return
	}

	cached, ok := signatureCache.Get(key)
	packed := bls.Signature{Signature: signature, Signer: signer, Processed: false}

	if !ok {
		signatureCache.Add(key, []bls.Signature{packed})
		// TODO: This should not only write to DB,
		// TODO: but also report to "consumers"
		DebouncedSaveSignatures(key, info)
		return
	}

	for _, item := range cached {
		if item.Signer.PublicKey == signer.PublicKey {
			return
		}
	}

	cached = append(cached, packed)
	signatureCache.Add(key, cached)

	DebouncedSaveSignatures(key, info)
}

func SaveSignatures(info datasets.PriceInfo) {

	dbClient := db.GetClient()
	lruCache := priceCache[strings.ToLower(info.Pair)]
	price, ok := lruCache.Get(info.Block)

	if !ok {
		return
	}

	key := AssetKey{
		Block: info.Block,
		Asset: info.Asset,
		Chain: info.Chain,
		Pair:  info.Pair,
	}

	signatures, ok := signatureCache.Get(key)

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

	err := dbClient.Signer.MapCreateBulk(newSigners, func(sc *ent.SignerCreate, i int) {
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
	currentAggregate, ok := aggregateCache.Get(key)

	if ok {
		newSignatures = append(newSignatures, currentAggregate)
	}

	aggregate, err = bls.AggregateSignatures(newSignatures)

	if err != nil {
		return
	}

	signatureBytes := aggregate.Bytes()

	packet := datasets.BroadcastPacket{
		Info:      info,
		Signers:   keys,
		Signature: signatureBytes,
	}

	payload, err := msgpack.Marshal(&packet)

	if err == nil {
		// TODO: Handle errors in a proper way
		consumer.Broadcast(append([]byte{7, 0}, payload...))
	}

	err = dbClient.AssetPrice.
		Create().
		SetPair(info.Pair).
		SetAsset(info.Asset).
		SetChain(info.Chain).
		SetBlock(info.Block).
		SetPrice(&price).
		SetSignersCount(uint64(len(signatures))).
		SetSignature(signatureBytes[:]).
		AddSignerIDs(signerIds...).
		OnConflictColumns("block", "chain", "asset", "pair").
		UpdateNewValues().
		Exec(ctx)

	if err != nil {
		panic(err)
	}

	for _, signature := range signatures {
		signature.Processed = true
	}

	aggregateCache.Add(key, aggregate)
}

func GetPriceFromCache(block uint64, pair string) (big.Int, bool) {
	lruCache := priceCache[strings.ToLower(pair)]
	return lruCache.Get(block)
}

func GetBlockNumber(network string) (*uint64, error) {
	blockNumber, err := ethereum.GetBlockNumber(network)

	if err != nil {
		ethereum.RefreshRPC(network)
		return nil, err
	}

	return &blockNumber, nil
}

func GetPriceAtBlockFromPair(
	network string,
	blockNumber uint64,
	pairAddr string,
	decimalDif int64,
	inverse bool) (*big.Int, error) {

	pair, err := ethereum.GetNewUniV3Contract(network, pairAddr, false)

	if err != nil {
		ethereum.RefreshRPC(network)
		return nil, err
	}

	data, err := pair.Slot0(
		&bind.CallOpts{
			BlockNumber: big.NewInt(int64(blockNumber)),
		})

	if err != nil {
		ethereum.RefreshRPC(network)
		return nil, err
	}

	lastPrice = *priceFromSqrtX96(data.SqrtPriceX96, decimalDif, inverse)
	lruCache := priceCache[strings.ToLower(pairAddr)]
	lruCache.Add(blockNumber, lastPrice)

	return &lastPrice, nil
}

func GetPriceFromPair(
	network string,
	pairAddr string,
	decimalDif int64,
	inverse bool) (*uint64, *big.Int, error) {

	blockNumber, err := ethereum.GetBlockNumber(network)

	if err != nil {
		ethereum.RefreshRPC(network)
		return nil, nil, err
	}

	lastPrice, err := GetPriceAtBlockFromPair(
		network,
		blockNumber,
		pairAddr,
		decimalDif,
		inverse)

	return &blockNumber, lastPrice, err
}

func priceFromSqrtX96(sqrtPriceX96 *big.Int, decimalDif int64, inverse bool) *big.Int {
	var decimalFix big.Int
	var powerUp big.Int
	var rawPrice big.Int
	var price big.Int
	var factor big.Int

	decimalFix.Mul(sqrtPriceX96, &tenEighteen)
	powerUp.Exp(&decimalFix, big.NewInt(2), nil)
	rawPrice.Div(&powerUp, &twoOneNineTwo)

	if inverse {
		factor.Exp(big.NewInt(10), big.NewInt(54+decimalDif), nil)
		price.Div(&factor, &rawPrice)
	} else {
		// TODO: needs work
		factor.Exp(big.NewInt(10), big.NewInt(18-decimalDif), nil)
		price.Div(&rawPrice, &factor)
	}

	return &price
}

func Setup() {
	var tokens []Token

	err := config.Config.UnmarshalKey("plugins.uniswap.tokens", &tokens)

	if err != nil {
		panic(err)
	}

	for _, token := range tokens {
		priceCache[strings.ToLower(token.Pair)], err = lru.New[uint64, big.Int](24)

		if err != nil {
			panic(err)
		}
	}
}

func createTask(tokens []Token, chain string) func() {

	return func() {

		if client.IsClientSocketClosed {
			return
		}

		currBlockNumber, err := GetBlockNumber(chain)

		if err != nil {
			panic(err)
		}

		if lastBlock[chain] == *currBlockNumber {
			return
		}

		if lastBlock[chain] == 0 {
			lastBlock[chain] = *currBlockNumber - 1
		}

	TOKENLOOP:
		for _, token := range tokens {

			for blockNumber := lastBlock[chain]; blockNumber < *currBlockNumber; blockNumber++ {
				if token.Chain != chain {
					continue
				}

				price, err := GetPriceAtBlockFromPair(
					token.Chain,
					blockNumber,
					token.Pair,
					token.Delta,
					token.Invert,
				)

				if err != nil {
					panic(err)
				}

				for _, cross := range token.Cross {
					stored := crossPrices[cross]

					if stored.Cmp(big.NewInt(0)) == 0 {
						continue TOKENLOOP
					}

					price.Mul(price, &stored)
					price.Div(price, &tenEighteen)
				}

				if token.Id != nil {
					crossPrices[*token.Id] = *price
				}

				var priceF big.Float
				priceF.Quo(new(big.Float).SetInt(price), &tenEighteenF)
				priceStr := fmt.Sprintf("%.18f %s", &priceF, token.Unit)

				log.Logger.
					With("Block", blockNumber).
					With("Price", priceStr).
					Info(caser.String(token.Name))

				priceInfo := datasets.PriceInfo{
					Price: *price,
					Block: blockNumber,
					Chain: token.Chain,
					Pair:  strings.ToLower(token.Pair),
					Asset: strings.ToLower(token.Name),
				}

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
			}

		}

		// TODO: block numbers should be tracked separately for each token
		lastBlock[chain] = *currBlockNumber
	}
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

	for _, token := range tokens {
		priceCache[strings.ToLower(token.Pair)], err = lru.New[uint64, big.Int](24)

		if err != nil {
			panic(err)
		}
	}

	scheduleConfs := config.Config.Sub("plugins.uniswap.schedule")
	scheduleNames := scheduleConfs.AllKeys()

	for index := range scheduleNames {
		name := scheduleNames[index]
		duration := scheduleConfs.GetDuration(name) * time.Millisecond
		task := createTask(tokens, name)

		_, err = scheduler.NewJob(
			gocron.DurationJob(duration),
			gocron.NewTask(task),
			gocron.WithSingletonMode(gocron.LimitModeReschedule),
		)

		if err != nil {
			panic(err)
		}
	}

	scheduler.Start()
}

func init() {

	DebouncedSaveSignatures = utils.Debounce[AssetKey, datasets.PriceInfo](5*time.Second, SaveSignatures)
	signatureMutex = new(sync.Mutex)

	twoOneNineTwo.Exp(big.NewInt(2), big.NewInt(192), nil)
	tenEighteen.Exp(big.NewInt(10), big.NewInt(18), nil)
	tenEighteenF.SetInt(&tenEighteen)

	// TODO: Should use AssetKey
	priceCache = make(map[string]*lru.Cache[uint64, big.Int])

	var err error
	signatureCache, err = lru.New[AssetKey, []bls.Signature](24)

	if err != nil {
		panic(err)
	}

	aggregateCache, err = lru.New[AssetKey, bls12381.G1Affine](24)

	if err != nil {
		panic(err)
	}

	lastBlock = make(map[string]uint64)
	crossPrices = make(map[string]big.Int)

	caser = cases.Title(language.English, cases.NoLower)

}
