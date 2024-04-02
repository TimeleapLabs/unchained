package uniswap

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/KenshiTech/unchained/src/address"
	"github.com/KenshiTech/unchained/src/config"
	"github.com/KenshiTech/unchained/src/constants/opcodes"
	"github.com/KenshiTech/unchained/src/crypto/bls"
	"github.com/KenshiTech/unchained/src/crypto/shake"
	"github.com/KenshiTech/unchained/src/datasets"
	"github.com/KenshiTech/unchained/src/db"
	"github.com/KenshiTech/unchained/src/ent"
	"github.com/KenshiTech/unchained/src/ent/helpers"
	"github.com/KenshiTech/unchained/src/ent/signer"
	"github.com/KenshiTech/unchained/src/ethereum"
	"github.com/KenshiTech/unchained/src/log"
	"github.com/KenshiTech/unchained/src/net/shared"
	"github.com/KenshiTech/unchained/src/pos"
	"github.com/KenshiTech/unchained/src/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/go-co-op/gocron/v2"
	"github.com/puzpuzpuz/xsync/v3"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	lru "github.com/hashicorp/golang-lru/v2"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

var DebouncedSaveSignatures func(key datasets.AssetKey, arg SaveSignatureArgs)
var signatureMutex *sync.Mutex

const (
	OldBlockNumber      = 96
	LruSize             = 1024
	SizeOfPriceCacheLru = 128
)

type Token struct {
	ID     *string  `mapstructure:"id"`
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

func NewTokensFromCfg(input []config.Token) []Token {
	result := []Token{}
	for _, t := range input {
		result = append(result, NewTokenFromCfg(t))
	}

	return result
}

func NewTokenFromCfg(input config.Token) Token {
	return Token{
		Chain:  input.Chain,
		Name:   input.Name,
		Pair:   input.Pair,
		Unit:   input.Unit,
		Delta:  input.Delta,
		Invert: input.Invert,
		Send:   input.Send,
	}
}

var priceCache map[string]*lru.Cache[uint64, big.Int]
var consensus *lru.Cache[datasets.AssetKey, xsync.MapOf[bls12381.G1Affine, big.Int]]
var signatureCache *lru.Cache[bls12381.G1Affine, []datasets.Signature]
var aggregateCache *lru.Cache[bls12381.G1Affine, bls12381.G1Affine]
var supportedTokens map[datasets.TokenKey]bool

var twoOneNineTwo big.Int
var tenEighteen big.Int
var tenEighteenF big.Float
var lastBlock *xsync.MapOf[datasets.TokenKey, uint64]
var crossPrices map[string]big.Int
var crossTokens map[string]datasets.TokenKey
var lastPrice big.Int

func CheckAndCacheSignature(
	reportedValues *xsync.MapOf[bls12381.G1Affine, big.Int],
	signature bls12381.G1Affine, signer datasets.Signer,
	hash bls12381.G1Affine,
	totalVoted *big.Int) error {
	signatureMutex.Lock()
	defer signatureMutex.Unlock()

	cached, _ := signatureCache.Get(hash)

	packed := datasets.Signature{
		Signature: signature,
		Signer:    signer,
		Processed: false,
	}

	for _, item := range cached {
		if item.Signer.PublicKey == signer.PublicKey {
			log.Logger.
				With("Address", address.Calculate(signer.PublicKey[:])).
				Debug("Duplicated signature")
			return fmt.Errorf("duplicated signature")
		}
	}

	reportedValues.Store(hash, *totalVoted)
	cached = append(cached, packed)
	signatureCache.Add(hash, cached)

	return nil
}

// TODO: This needs to work with different datasets
// TODO: Can we turn this into a library func?
func RecordSignature(
	signature bls12381.G1Affine,
	signer datasets.Signer,
	hash bls12381.G1Affine,
	info datasets.PriceInfo,
	debounce bool,
	historical bool) {
	if supported := supportedTokens[info.Asset.Token]; !supported {
		log.Logger.
			With("Name", info.Asset.Token.Name).
			With("Chain", info.Asset.Token.Chain).
			With("Pair", info.Asset.Token.Pair).
			Debug("Token not supported")
		return
	}

	// TODO: Standalone mode shouldn't call this or check consensus
	blockNumber, err := GetBlockNumber(info.Asset.Token.Chain)
	if err != nil {
		log.Logger.
			With("Network", info.Asset.Token.Chain).
			With("Error", err).
			Error("Failed to get the latest block number")
		ethereum.RefreshRPC(info.Asset.Token.Chain)
		// TODO: we should retry
		return
	}

	if !historical {
		// TODO: this won't work for Arbitrum
		if *blockNumber-info.Asset.Block > OldBlockNumber {
			log.Logger.
				With("Packet", info.Asset.Block).
				With("Current", *blockNumber).
				Debug("Data too old")
			return // Data too old
		}
	}

	if !consensus.Contains(info.Asset) {
		consensus.Add(info.Asset, *xsync.NewMapOf[bls12381.G1Affine, big.Int]())
	}

	reportedValues, _ := consensus.Get(info.Asset)
	isMajority := true
	voted, ok := reportedValues.Load(hash)
	if !ok {
		voted = *big.NewInt(0)
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

	totalVoted := new(big.Int).Add(votingPower, &voted)

	reportedValues.Range(func(_ bls12381.G1Affine, value big.Int) bool {
		if value.Cmp(totalVoted) == 1 {
			isMajority = false
		}
		return isMajority
	})

	err = CheckAndCacheSignature(&reportedValues, signature, signer, hash, totalVoted)
	if err != nil {
		return
	}

	if !isMajority {
		log.Logger.Debug("Not a majority")
		return
	}

	if !debounce {
		SaveSignatures(SaveSignatureArgs{Hash: hash, Info: info})
		return
	}

	if isMajority {
		reportLog := log.Logger.
			With("Block", info.Asset.Block).
			With("Price", info.Price.String()).
			With("Token", info.Asset.Token.Name)

		reportedValues.Range(func(hash bls12381.G1Affine, value big.Int) bool {
			reportLog = reportLog.With(
				fmt.Sprintf("%x", hash.Bytes())[:8],
				value.String(),
			)
			return true
		})
		reportedValues.Range(func(hash bls12381.G1Affine, value big.Int) bool {
			reportLog = reportLog.With(
				fmt.Sprintf("%x", hash.Bytes())[:8],
				value.String(),
			)
			return true
		})

		reportLog.
			With("Majority", fmt.Sprintf("%x", hash.Bytes())[:8]).
			Debug("Values")

		DebouncedSaveSignatures(
			info.Asset,
			SaveSignatureArgs{Hash: hash, Info: info},
		)

		if debounce {
			DebouncedSaveSignatures(
				info.Asset,
				SaveSignatureArgs{Hash: hash, Info: info},
			)
		} else {
			SaveSignatures(SaveSignatureArgs{Hash: hash, Info: info})
		}
	} else {
		log.Logger.Debug("Not a majority")
	}
}

type SaveSignatureArgs struct {
	Info datasets.PriceInfo
	Hash bls12381.G1Affine
}

func SaveSignatures(args SaveSignatureArgs) {
	dbClient := db.GetClient()
	log.Logger.
		With("Block", args.Info.Asset.Block).
		With("Hash", fmt.Sprintf("%x", args.Hash.Bytes())[:8]).
		Debug("Saving into DB")

	signatures, ok := signatureCache.Get(args.Hash)
	if !ok {
		log.Logger.
			With("Block", args.Info.Asset.Block).
			With("Hash", fmt.Sprintf("%x", args.Hash.Bytes())[:8]).
			Debug("Cache not found")
		return
	}

	ctx := context.Background()

	var newSigners []datasets.Signer
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
		newSigner := newSigners[i]
		sc.SetName(newSigner.Name).
			SetEvm(newSigner.EvmWallet).
			SetKey(newSigner.PublicKey[:]).
			SetShortkey(newSigner.ShortPublicKey[:]).
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
		log.Logger.
			With("Block", args.Info.Asset.Block).
			With("Hash", fmt.Sprintf("%x", args.Hash.Bytes())[:8]).
			Debug("Failed to upsert token signers.")
		panic(err)
	}

	signerIds, err := dbClient.Signer.
		Query().
		Where(signer.KeyIn(keys...)).
		IDs(ctx)

	if err != nil {
		log.Logger.
			With("Block", args.Info.Asset.Block).
			With("Hash", fmt.Sprintf("%x", args.Hash.Bytes())[:8]).
			Debug("Filed to upsert signers")
		return
	}

	var aggregate bls12381.G1Affine
	currentAggregate, ok := aggregateCache.Get(args.Hash)

	if ok {
		newSignatures = append(newSignatures, currentAggregate)
	}

	aggregate, err = bls.AggregateSignatures(newSignatures)

	if err != nil {
		log.Logger.
			With("Block", args.Info.Asset.Block).
			With("Hash", fmt.Sprintf("%x", args.Hash.Bytes())[:8]).
			Debug("Filed to aggregate signatures")
		return
	}

	signatureBytes := aggregate.Bytes()

	// TODO: Handle cases where signerIds need to be removed
	err = dbClient.AssetPrice.
		Create().
		SetPair(strings.ToLower(args.Info.Asset.Token.Pair)).
		SetAsset(args.Info.Asset.Token.Name).
		SetChain(args.Info.Asset.Token.Chain).
		SetBlock(args.Info.Asset.Block).
		SetPrice(&helpers.BigInt{Int: args.Info.Price}).
		SetSignersCount(uint64(len(signatures))).
		SetSignature(signatureBytes[:]).
		AddSignerIDs(signerIds...).
		OnConflictColumns("block", "chain", "asset", "pair").
		UpdateNewValues().
		Exec(ctx)

	if err != nil {
		log.Logger.
			With("Block", args.Info.Asset.Block).
			With("Hash", fmt.Sprintf("%x", args.Hash.Bytes())[:8]).
			Debug("Failed to upsert asset price")
		panic(err)
	}

	// TODO: We probably need a context-aware mutex here
	for inx := range signatures {
		signatures[inx].Processed = true
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

func New() {
	if config.App.Plugins.Uniswap != nil {
		return
	}

	for _, t := range config.App.Plugins.Uniswap.Tokens {
		token := NewTokenFromCfg(t)
		var err error
		priceCache[strings.ToLower(token.Pair)], err = lru.New[uint64, big.Int](SizeOfPriceCacheLru)

		if err != nil {
			log.Logger.Error("Failed to initialize token map.")
			os.Exit(1)
		}

		key := tokenKey(token)
		supportedTokens[*key] = true
	}
}

func syncBlock(token Token, caser cases.Caser, key *datasets.TokenKey, blockInx uint64) {
	lastSynced, ok := lastBlock.Load(*key)

	if ok && blockInx <= lastSynced {
		return
	}

	price, err := GetPriceAtBlockFromPair(
		token.Chain,
		blockInx,
		token.Pair,
		token.Delta,
		token.Invert,
	)

	if err != nil {
		log.Logger.Error(
			fmt.Sprintf("Failed to get token price from %s RPC.", token.Chain))
		ethereum.RefreshRPC(token.Chain)
		return
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

	if token.ID != nil {
		crossPrices[*token.ID] = *price
	}

	var priceF big.Float
	priceF.Quo(new(big.Float).SetInt(price), &tenEighteenF)
	priceStr := fmt.Sprintf("%.18f %s", &priceF, token.Unit)

	lastSynced, ok = lastBlock.Load(*key)

	if ok && blockInx <= lastSynced {
		return
	}

	log.Logger.
		With("Block", blockInx).
		With("Price", priceStr).
		Info(caser.String(token.Name))

	key = tokenKey(token)

	priceInfo := datasets.PriceInfo{
		Price: *price,
		Asset: datasets.AssetKey{
			Block: blockInx,
			Token: *key,
		},
	}

	toHash := priceInfo.Sia().Content
	signature, hash := bls.Sign(*bls.ClientSecretKey, toHash)

	if token.Send && !shared.IsClientSocketClosed {
		compressedSignature := signature.Bytes()

		priceReport := datasets.PriceReport{
			PriceInfo: priceInfo,
			Signature: compressedSignature,
		}

		payload := priceReport.Sia().Content
		shared.Send(opcodes.PriceReport, payload)
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

	lastBlock.Store(*key, blockInx)
}

func syncBlocks(token Token, key datasets.TokenKey, latest uint64) {
	block, ok := lastBlock.Load(key)
	if !ok {
		return
	}

	caser := cases.Title(language.English, cases.NoLower)

	for blockInx := block + 1; blockInx < latest; blockInx++ {
		syncBlock(token, caser, &key, blockInx)
	}
}

func tokenKey(token Token) *datasets.TokenKey {
	var cross []datasets.TokenKey

	for _, id := range token.Cross {
		cross = append(cross, crossTokens[id])
	}

	toHash := new(sia.ArraySia[datasets.TokenKey]).
		AddArray8(cross, func(s *sia.ArraySia[datasets.TokenKey], item datasets.TokenKey) {
			s.EmbedSia(item.Sia())
		}).Content

	hash := shake.Shake(toHash)

	key := datasets.TokenKey{
		Name:   strings.ToLower(token.Name),
		Pair:   strings.ToLower(token.Pair),
		Chain:  strings.ToLower(token.Chain),
		Delta:  token.Delta,
		Invert: token.Invert,
		Cross:  string(hash),
	}

	return &key
}

func createTask(tokens []Token, chain string) func() {
	return func() {
		currBlockNumber, err := GetBlockNumber(chain)

		if err != nil {
			log.Logger.Error(
				fmt.Sprintf("Couldn't get latest block from %s RPC.", chain))
			ethereum.RefreshRPC(chain)
			return
		}

		for _, token := range tokens {
			if token.Chain != chain {
				continue
			}

			// TODO: this can be cached
			key := tokenKey(token)
			tokenLastBlock, exists := lastBlock.Load(*key)

			if !exists {
				lastBlock.Store(*key, *currBlockNumber-1)
			} else if tokenLastBlock == *currBlockNumber {
				return
			}

			syncBlocks(token, *key, *currBlockNumber)
		}
	}
}

func Listen() {
	scheduler, err := gocron.NewScheduler()

	if err != nil {
		log.Logger.Error("Failed to create token scheduler.")
		os.Exit(1)
	}

	for _, token := range config.App.Plugins.Uniswap.Tokens {
		priceCache[strings.ToLower(token.Pair)], err = lru.New[uint64, big.Int](SizeOfPriceCacheLru)

		if err != nil {
			log.Logger.Error("Failed to create token price cache.")
			os.Exit(1)
		}
	}

	for name, duration := range config.App.Plugins.Uniswap.Schedule {
		task := createTask(NewTokensFromCfg(config.App.Plugins.Uniswap.Tokens), name)

		_, err = scheduler.NewJob(
			gocron.DurationJob(duration),
			gocron.NewTask(task),
		)

		if err != nil {
			log.Logger.Error("Failed to schedule token price task.")
			os.Exit(1)
		}
	}

	scheduler.Start()
}

func init() {
	DebouncedSaveSignatures = utils.Debounce[datasets.AssetKey, SaveSignatureArgs](5*time.Second, SaveSignatures)
	signatureMutex = new(sync.Mutex)

	twoOneNineTwo.Exp(big.NewInt(2), big.NewInt(192), nil)
	tenEighteen.Exp(big.NewInt(10), big.NewInt(18), nil)
	tenEighteenF.SetInt(&tenEighteen)

	// TODO: Should use AssetKey
	priceCache = make(map[string]*lru.Cache[uint64, big.Int])
	supportedTokens = make(map[datasets.TokenKey]bool)

	var err error
	signatureCache, err = lru.New[bls12381.G1Affine, []datasets.Signature](LruSize)

	if err != nil {
		log.Logger.Error("Failed to create token price signature cache.")
		os.Exit(1)
	}

	// TODO: This is vulnerable to flood attacks
	consensus, err = lru.New[datasets.AssetKey, xsync.MapOf[bls12381.G1Affine, big.Int]](LruSize)

	if err != nil {
		log.Logger.Error("Failed to create token price consensus cache.")
		os.Exit(1)
	}

	aggregateCache, err = lru.New[bls12381.G1Affine, bls12381.G1Affine](LruSize)

	if err != nil {
		log.Logger.Error("Failed to create token price aggregate cache.")
		os.Exit(1)
	}

	lastBlock = xsync.NewMapOf[datasets.TokenKey, uint64]()
	crossPrices = make(map[string]big.Int)
	crossTokens = make(map[string]datasets.TokenKey)
}
