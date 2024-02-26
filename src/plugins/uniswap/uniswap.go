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
	"github.com/KenshiTech/unchained/constants/opcodes"
	"github.com/KenshiTech/unchained/datasets"
	"github.com/KenshiTech/unchained/db"
	"github.com/KenshiTech/unchained/ent"
	"github.com/KenshiTech/unchained/ent/helpers"
	"github.com/KenshiTech/unchained/ent/signer"
	"github.com/KenshiTech/unchained/ethereum"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/net/client"
	"github.com/KenshiTech/unchained/net/consumer"
	"github.com/KenshiTech/unchained/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/go-co-op/gocron/v2"
	"github.com/puzpuzpuz/xsync/v3"
	"github.com/vmihailenco/msgpack/v5"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	lru "github.com/hashicorp/golang-lru/v2"
)

var DebouncedSaveSignatures func(key bls12381.G1Affine, arg SaveSignatureArgs)
var signatureMutex *sync.Mutex

type TokenKey struct {
	Pair  string
	Chain string
}

type AssetKey struct {
	Asset string
	Pair  string
	Chain string
	Block uint64
}

type Token struct {
	Id     *string  `mapstructure:"id"`
	Chain  string   `mapstructure:"chain"`
	Name   string   `mapstructure:"name"`
	Pair   string   `mapstructure:"pair"`
	Unit   string   `mapstructure:"unit"`
	Symbol string   `mapstructure:"symbol"`
	Delta  int64    `mapstructure:"delta"`
	Invert bool     `mapstructure:"invert"`
	Store  bool     `mapstructure:"store"`
	Send   bool     `mapstructure:"send"`
	Cross  []string `mapstructure:"cross"`
}

var priceCache map[string]*lru.Cache[uint64, big.Int]
var consensus *lru.Cache[AssetKey, map[bls12381.G1Affine]uint64]
var signatureCache *lru.Cache[bls12381.G1Affine, []bls.Signature]
var aggregateCache *lru.Cache[bls12381.G1Affine, bls12381.G1Affine]
var supportedTokens map[TokenKey]bool

var twoOneNineTwo big.Int
var tenEighteen big.Int
var tenEighteenF big.Float
var lastBlock *xsync.MapOf[TokenKey, uint64]
var crossPrices map[string]big.Int
var lastPrice big.Int

// TODO: This needs to work with different datasets
// TODO: Can we turn this into a library func?
func RecordSignature(
	signature bls12381.G1Affine,
	signer bls.Signer,
	hash bls12381.G1Affine,
	info datasets.PriceInfo,
	debounce bool,
	historical bool) {

	// TODO: Invert makes a difference here
	tokenKey := TokenKey{Chain: info.Chain, Pair: strings.ToLower(info.Pair)}
	if supported := supportedTokens[tokenKey]; !supported {
		return
	}

	if !historical {
		blockNumber, err := GetBlockNumber(info.Chain)

		if err != nil {
			panic(err)
		}

		// TODO: this won't work for Arbitrum
		if *blockNumber-info.Block > 96 {
			return // Data too old
		}
	}

	key := AssetKey{
		Asset: info.Asset,
		Chain: info.Chain,
		Pair:  strings.ToLower(info.Pair),
		Block: info.Block,
	}

	signatureMutex.Lock()
	defer signatureMutex.Unlock()

	if !consensus.Contains(key) {
		consensus.Add(key, make(map[bls12381.G1Affine]uint64))
	}

	reportedValues, _ := consensus.Get(key)
	reportedValues[hash]++
	isMajority := true
	count := reportedValues[hash]

	for _, reportCount := range reportedValues {
		if reportCount > count {
			isMajority = false
			break
		}
	}

	cached, ok := signatureCache.Get(hash)

	packed := bls.Signature{
		Signature: signature,
		Signer:    signer,
		Processed: false,
	}

	if !ok {
		signatureCache.Add(hash, []bls.Signature{packed})
		// TODO: This should not only write to DB,
		// TODO: but also report to "consumers"
		if isMajority {
			if debounce {
				DebouncedSaveSignatures(hash, SaveSignatureArgs{Hash: hash, Info: info})
			} else {
				SaveSignatures(SaveSignatureArgs{Hash: hash, Info: info})
			}
		}
		return
	}

	for _, item := range cached {
		if item.Signer.PublicKey == signer.PublicKey {
			return
		}
	}

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

type SaveSignatureArgs struct {
	Info datasets.PriceInfo
	Hash bls12381.G1Affine
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
	currentAggregate, ok := aggregateCache.Get(args.Hash)

	if ok {
		newSignatures = append(newSignatures, currentAggregate)
	}

	aggregate, err = bls.AggregateSignatures(newSignatures)

	if err != nil {
		return
	}

	signatureBytes := aggregate.Bytes()

	packet := datasets.BroadcastPricePacket{
		Info:      args.Info,
		Signers:   keys,
		Signature: signatureBytes,
	}

	payload, err := msgpack.Marshal(&packet)

	if err == nil {
		// TODO: Handle errors in a proper way
		consumer.Broadcast(
			append(
				[]byte{opcodes.PriceReportBroadcast, 0},
				payload...),
		)
	}

	err = dbClient.AssetPrice.
		Create().
		SetPair(strings.ToLower(args.Info.Pair)).
		SetAsset(args.Info.Asset).
		SetChain(args.Info.Chain).
		SetBlock(args.Info.Block).
		SetPrice(&helpers.BigInt{Int: args.Info.Price}).
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

	aggregateCache.Add(args.Hash, aggregate)
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
	if !config.Config.IsSet("plugins.uniswap") {
		return
	}

	var tokens []Token

	err := config.Config.UnmarshalKey("plugins.uniswap.tokens", &tokens)

	if err != nil {
		panic(err)
	}

	for _, token := range tokens {
		priceCache[strings.ToLower(token.Pair)], err = lru.New[uint64, big.Int](128)

		key := TokenKey{Chain: token.Chain, Pair: strings.ToLower(token.Pair)}
		supportedTokens[key] = true

		if err != nil {
			panic(err)
		}
	}
}

func syncBlocks(token Token, key TokenKey, latest uint64) {
	block, ok := lastBlock.Load(key)

	if !ok {
		return
	}

	caser := cases.Title(language.English, cases.NoLower)

	for currBlock := block + 1; currBlock < latest; currBlock++ {

		lastSycned, ok := lastBlock.Load(key)

		if ok && currBlock <= lastSycned {
			return
		}

		price, err := GetPriceAtBlockFromPair(
			token.Chain,
			currBlock,
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
				return
			}

			price.Mul(price, &stored)
		}

		for range token.Cross {
			price.Div(price, &tenEighteen)
		}

		if token.Id != nil {
			crossPrices[*token.Id] = *price
		}

		var priceF big.Float
		priceF.Quo(new(big.Float).SetInt(price), &tenEighteenF)
		priceStr := fmt.Sprintf("%.18f %s", &priceF, token.Unit)

		lastSycned, ok = lastBlock.Load(key)

		if ok && currBlock <= lastSycned {
			return
		}

		log.Logger.
			With("Block", currBlock).
			With("Price", priceStr).
			Info(caser.String(token.Name))

		priceInfo := datasets.PriceInfo{
			Price: *price,
			Block: currBlock,
			Chain: token.Chain,
			Pair:  strings.ToLower(token.Pair),
			Asset: strings.ToLower(token.Name),
		}

		toHash, err := msgpack.Marshal(&priceInfo)

		if err != nil {
			panic(err)
		}

		signature, hash := bls.Sign(*bls.ClientSecretKey, toHash)
		compressedSignature := signature.Bytes()

		priceReport := datasets.PriceReport{
			PriceInfo: priceInfo,
			Signature: compressedSignature,
		}

		payload, err := msgpack.Marshal(&priceReport)

		if err != nil {
			panic(err)
		}

		if token.Send && !client.IsClientSocketClosed {
			client.Send(
				append([]byte{opcodes.PriceReport, 0}, payload...),
			)
		}

		if token.Store {
			RecordSignature(
				signature,
				bls.ClientSigner,
				hash,
				priceInfo,
				false,
				true,
			)
		}

		lastBlock.Store(key, currBlock)
	}
}

func createTask(tokens []Token, chain string) func() {

	return func() {

		currBlockNumber, err := GetBlockNumber(chain)

		if err != nil {
			panic(err)
		}

		for _, token := range tokens {

			if token.Chain != chain {
				continue
			}

			key := TokenKey{Chain: token.Chain, Pair: token.Pair}
			tokenLastBlock, exists := lastBlock.Load(key)

			if !exists {
				lastBlock.Store(key, *currBlockNumber-1)
			} else if tokenLastBlock == *currBlockNumber {
				return
			}

			syncBlocks(token, key, *currBlockNumber)
		}
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
		priceCache[strings.ToLower(token.Pair)], err = lru.New[uint64, big.Int](128)

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
		)

		if err != nil {
			panic(err)
		}
	}

	scheduler.Start()
}

func init() {

	DebouncedSaveSignatures = utils.Debounce[bls12381.G1Affine, SaveSignatureArgs](5*time.Second, SaveSignatures)
	signatureMutex = new(sync.Mutex)

	twoOneNineTwo.Exp(big.NewInt(2), big.NewInt(192), nil)
	tenEighteen.Exp(big.NewInt(10), big.NewInt(18), nil)
	tenEighteenF.SetInt(&tenEighteen)

	// TODO: Should use AssetKey
	priceCache = make(map[string]*lru.Cache[uint64, big.Int])
	supportedTokens = make(map[TokenKey]bool)

	var err error
	signatureCache, err = lru.New[bls12381.G1Affine, []bls.Signature](128)

	if err != nil {
		panic(err)
	}

	consensus, err = lru.New[AssetKey, map[bls12381.G1Affine]uint64](128)

	if err != nil {
		panic(err)
	}

	aggregateCache, err = lru.New[bls12381.G1Affine, bls12381.G1Affine](128)

	if err != nil {
		panic(err)
	}

	lastBlock = xsync.NewMapOf[TokenKey, uint64]()
	crossPrices = make(map[string]big.Int)
}
