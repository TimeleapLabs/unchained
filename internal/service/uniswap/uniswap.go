package uniswap

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"
	"sync"

	"github.com/KenshiTech/unchained/config"

	"github.com/KenshiTech/unchained/address"
	"github.com/KenshiTech/unchained/constants/opcodes"
	"github.com/KenshiTech/unchained/crypto/bls"
	"github.com/KenshiTech/unchained/crypto/shake"
	"github.com/KenshiTech/unchained/datasets"
	"github.com/KenshiTech/unchained/db"
	"github.com/KenshiTech/unchained/ent"
	"github.com/KenshiTech/unchained/ent/helpers"
	"github.com/KenshiTech/unchained/ent/signer"
	"github.com/KenshiTech/unchained/ethereum"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/pos"
	"github.com/KenshiTech/unchained/service/evmlog"
	"github.com/KenshiTech/unchained/transport/client/conn"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	lru "github.com/hashicorp/golang-lru/v2"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
	"github.com/puzpuzpuz/xsync/v3"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	OldBlockNumber = 96
)

var DebouncedSaveSignatures func(key datasets.AssetKey, arg SaveSignatureArgs)

type Service struct {
	ethRPC *ethereum.Repository
	pos    *pos.Repository

	consensus       *lru.Cache[datasets.AssetKey, xsync.MapOf[bls12381.G1Affine, big.Int]]
	signatureCache  *lru.Cache[bls12381.G1Affine, []datasets.Signature]
	aggregateCache  *lru.Cache[bls12381.G1Affine, bls12381.G1Affine]
	SupportedTokens map[datasets.TokenKey]bool
	signatureMutex  sync.Mutex

	twoOneNineTwo big.Int
	tenEighteen   big.Int
	tenEighteenF  big.Float
	LastBlock     xsync.MapOf[datasets.TokenKey, uint64]
	PriceCache    map[string]*lru.Cache[uint64, big.Int]
	crossPrices   map[string]big.Int
	crossTokens   map[string]datasets.TokenKey
	LastPrice     big.Int
}

func (u *Service) CheckAndCacheSignature(
	reportedValues *xsync.MapOf[bls12381.G1Affine, big.Int], signature bls12381.G1Affine, signer datasets.Signer,
	hash bls12381.G1Affine, totalVoted *big.Int,
) error {
	u.signatureMutex.Lock()
	defer u.signatureMutex.Unlock()

	cached, _ := u.signatureCache.Get(hash)

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
	u.signatureCache.Add(hash, cached)

	return nil
}

// TODO: This needs to work with different datasets
// TODO: Can we turn this into a library func?
func (u *Service) RecordSignature(
	signature bls12381.G1Affine, signer datasets.Signer, hash bls12381.G1Affine, info datasets.PriceInfo, debounce bool, historical bool,
) {
	if supported := u.SupportedTokens[info.Asset.Token]; !supported {
		log.Logger.
			With("Name", info.Asset.Token.Name).
			With("Chain", info.Asset.Token.Chain).
			With("Pair", info.Asset.Token.Pair).
			Debug("Token not supported")
		return
	}

	// TODO: Standalone mode shouldn't call this or check consensus
	blockNumber, err := u.GetBlockNumber(info.Asset.Token.Chain)
	if err != nil {
		log.Logger.
			With("Network", info.Asset.Token.Chain).
			With("Error", err).
			Error("Failed to get the latest block number")
		u.ethRPC.RefreshRPC(info.Asset.Token.Chain)
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

	if !u.consensus.Contains(info.Asset) {
		u.consensus.Add(info.Asset, *xsync.NewMapOf[bls12381.G1Affine, big.Int]())
	}

	reportedValues, _ := u.consensus.Get(info.Asset)
	isMajority := true
	voted, ok := reportedValues.Load(hash)
	if !ok {
		voted = *big.NewInt(0)
	}

	votingPower, err := u.pos.GetVotingPowerOfPublicKey(
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

	err = u.CheckAndCacheSignature(&reportedValues, signature, signer, hash, totalVoted)
	if err != nil {
		return
	}

	if !isMajority {
		log.Logger.Debug("Not a majority")
		return
	}

	if !debounce {
		u.saveSignatures(SaveSignatureArgs{Hash: hash, Info: info})
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
			u.saveSignatures(SaveSignatureArgs{Hash: hash, Info: info})
		}
	} else {
		log.Logger.Debug("Not a majority")
	}
}

type SaveSignatureArgs struct {
	Info datasets.PriceInfo
	Hash bls12381.G1Affine
}

func (u *Service) saveSignatures(args SaveSignatureArgs) {
	dbClient := db.GetClient()
	log.Logger.
		With("Block", args.Info.Asset.Block).
		With("Hash", fmt.Sprintf("%x", args.Hash.Bytes())[:8]).
		Debug("Saving into DB")

	signatures, ok := u.signatureCache.Get(args.Hash)
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
	currentAggregate, ok := u.aggregateCache.Get(args.Hash)

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

	u.aggregateCache.Add(args.Hash, aggregate)
}

//
// func (u *Service) setPriceFromCache(block uint64, pair string) (big.Int, bool) {
//	lruCache := u.PriceCache[strings.ToLower(pair)]
//	return lruCache.Get(block)
//}

func (u *Service) GetBlockNumber(network string) (*uint64, error) {
	blockNumber, err := u.ethRPC.GetBlockNumber(network)

	if err != nil {
		u.ethRPC.RefreshRPC(network)
		return nil, err
	}

	return &blockNumber, nil
}

func (u *Service) GetPriceAtBlockFromPair(
	network string, blockNumber uint64, pairAddr string, decimalDif int64, inverse bool,
) (*big.Int, error) {
	pair, err := u.ethRPC.GetNewUniV3Contract(network, pairAddr, false)

	if err != nil {
		u.ethRPC.RefreshRPC(network)
		return nil, err
	}

	data, err := pair.Slot0(
		&bind.CallOpts{
			BlockNumber: big.NewInt(int64(blockNumber)),
		})

	if err != nil {
		u.ethRPC.RefreshRPC(network)
		return nil, err
	}

	u.LastPrice = *u.priceFromSqrtX96(data.SqrtPriceX96, decimalDif, inverse)
	lruCache := u.PriceCache[strings.ToLower(pairAddr)]
	lruCache.Add(blockNumber, u.LastPrice)

	return &u.LastPrice, nil
}

//
// func (u *Service) getPriceFromPair(
//	network string, pairAddr string, decimalDif int64, inverse bool,
// ) (*uint64, *big.Int, error) {
//	blockNumber, err := ethereum.GetBlockNumber(network)
//
//	if err != nil {
//		ethereum.RefreshRPC(network)
//		return nil, nil, err
//	}
//
//	lastPrice, err := u.GetPriceAtBlockFromPair(
//		network,
//		blockNumber,
//		pairAddr,
//		decimalDif,
//		inverse)
//
//	return &blockNumber, lastPrice, err
//}

func (u *Service) priceFromSqrtX96(sqrtPriceX96 *big.Int, decimalDif int64, inverse bool) *big.Int {
	var decimalFix big.Int
	var powerUp big.Int
	var rawPrice big.Int
	var price big.Int
	var factor big.Int

	decimalFix.Mul(sqrtPriceX96, &u.tenEighteen)
	powerUp.Exp(&decimalFix, big.NewInt(2), nil)
	rawPrice.Div(&powerUp, &u.twoOneNineTwo)

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

func (u *Service) syncBlock(token datasets.Token, caser cases.Caser, key *datasets.TokenKey, blockInx uint64) {
	lastSynced, ok := u.LastBlock.Load(*key)

	if ok && blockInx <= lastSynced {
		return
	}

	price, err := u.GetPriceAtBlockFromPair(
		token.Chain,
		blockInx,
		token.Pair,
		token.Delta,
		token.Invert,
	)

	if err != nil {
		log.Logger.Error(
			fmt.Sprintf("Failed to get token price from %s RPC.", token.Chain))
		u.ethRPC.RefreshRPC(token.Chain)
		return
	}

	for _, cross := range token.Cross {
		stored := u.crossPrices[cross]

		if stored.Cmp(big.NewInt(0)) == 0 {
			return
		}

		price.Mul(price, &stored)
	}

	for range token.Cross {
		price.Div(price, &u.tenEighteen)
	}

	if token.ID != nil {
		u.crossPrices[*token.ID] = *price
	}

	var priceF big.Float
	priceF.Quo(new(big.Float).SetInt(price), &u.tenEighteenF)
	priceStr := fmt.Sprintf("%.18f %s", &priceF, token.Unit)

	lastSynced, ok = u.LastBlock.Load(*key)

	if ok && blockInx <= lastSynced {
		return
	}

	log.Logger.
		With("Block", blockInx).
		With("Price", priceStr).
		Info(caser.String(token.Name))

	key = u.TokenKey(token)

	priceInfo := datasets.PriceInfo{
		Price: *price,
		Asset: datasets.AssetKey{
			Block: blockInx,
			Token: *key,
		},
	}

	toHash := priceInfo.Sia().Content
	signature, hash := bls.Sign(*bls.ClientSecretKey, toHash)

	if token.Send && !conn.IsClosed {
		compressedSignature := signature.Bytes()

		priceReport := datasets.PriceReport{
			PriceInfo: priceInfo,
			Signature: compressedSignature,
		}

		payload := priceReport.Sia().Content
		conn.Send(opcodes.PriceReport, payload)
	}

	if token.Store {
		u.RecordSignature(
			signature,
			bls.ClientSigner,
			hash,
			priceInfo,
			false,
			true,
		)
	}

	u.LastBlock.Store(*key, blockInx)
}

func (u *Service) SyncBlocks(token datasets.Token, key datasets.TokenKey, latest uint64) {
	block, ok := u.LastBlock.Load(key)
	if !ok {
		return
	}

	caser := cases.Title(language.English, cases.NoLower)

	for blockInx := block + 1; blockInx < latest; blockInx++ {
		u.syncBlock(token, caser, &key, blockInx)
	}
}

func (u *Service) TokenKey(token datasets.Token) *datasets.TokenKey {
	var cross []datasets.TokenKey

	for _, id := range token.Cross {
		cross = append(cross, u.crossTokens[id])
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

func New(ethRPC *ethereum.Repository, pos *pos.Repository) *Service {
	u := Service{
		ethRPC: ethRPC,
		pos:    pos,

		crossPrices:     map[string]big.Int{},
		crossTokens:     map[string]datasets.TokenKey{},
		SupportedTokens: map[datasets.TokenKey]bool{},
		PriceCache:      map[string]*lru.Cache[uint64, big.Int]{},
	}

	if config.App.Plugins.Uniswap != nil {
		for _, t := range config.App.Plugins.Uniswap.Tokens {
			token := datasets.NewTokenFromCfg(t)

			key := u.TokenKey(token)
			u.SupportedTokens[*key] = true
		}
	}

	u.twoOneNineTwo.Exp(big.NewInt(2), big.NewInt(192), nil)
	u.tenEighteen.Exp(big.NewInt(10), big.NewInt(18), nil)
	u.tenEighteenF.SetInt(&u.tenEighteen)

	var err error
	u.signatureCache, err = lru.New[bls12381.G1Affine, []datasets.Signature](evmlog.LruSize)
	if err != nil {
		log.Logger.Error("Failed to create token price signature cache.")
		os.Exit(1)
	}

	// TODO: This is vulnerable to flood attacks
	u.consensus, err = lru.New[datasets.AssetKey, xsync.MapOf[bls12381.G1Affine, big.Int]](evmlog.LruSize)
	if err != nil {
		log.Logger.Error("Failed to create token price consensus cache.")
		os.Exit(1)
	}

	u.aggregateCache, err = lru.New[bls12381.G1Affine, bls12381.G1Affine](evmlog.LruSize)
	if err != nil {
		log.Logger.Error("Failed to create token price aggregate cache.")
		os.Exit(1)
	}

	return &u
}
