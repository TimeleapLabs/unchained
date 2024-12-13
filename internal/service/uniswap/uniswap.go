package uniswap

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/repository"
	"github.com/TimeleapLabs/unchained/internal/service/pos"
	"github.com/TimeleapLabs/unchained/internal/utils/address"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"
	"github.com/TimeleapLabs/unchained/internal/ent"
	"github.com/TimeleapLabs/unchained/internal/service/evmlog"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
	"github.com/TimeleapLabs/unchained/internal/utils"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/puzpuzpuz/xsync/v3"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
)

const (
	MaxBlockNumberDelta = 96
	SizeOfPriceCacheLru = 128
)

var DebouncedSaveSignatures func(key model.AssetKey, arg SaveSignatureArgs)

type Service interface {
	checkAndCacheSignature(
		reportedValues *xsync.MapOf[bls12381.G1Affine, big.Int], signature bls12381.G1Affine, signer model.Signer,
		hash bls12381.G1Affine, totalVoted *big.Int,
	) error
	saveSignatures(ctx context.Context, args SaveSignatureArgs) error
	GetBlockNumber(ctx context.Context, network string) (*uint64, error)
	GetPriceAtBlockFromPair(network string, blockNumber uint64, pairAddr string, decimalDif int64, inverse bool) (*big.Int, error)
	SyncBlocks(ctx context.Context, token model.Token, key model.TokenKey, latest uint64) error
	TokenKey(token model.Token) *model.TokenKey
	RecordSignature(
		ctx context.Context, signature bls12381.G1Affine, signer model.Signer, hash bls12381.G1Affine, info model.PriceInfo, debounce bool, historical bool,
	) error
	ProcessBlocks(ctx context.Context, chain string) error
}

type service struct {
	ethRPC         ethereum.RPC
	pos            pos.Service
	signerRepo     repository.Signer
	assetPriceRepo repository.AssetPrice

	consensus       *lru.Cache[model.AssetKey, xsync.MapOf[bls12381.G1Affine, big.Int]]
	signatureCache  *lru.Cache[bls12381.G1Affine, []model.Signature]
	SupportedTokens map[model.TokenKey]bool
	signatureMutex  sync.Mutex

	twoOneNineTwo big.Int
	tenEighteen   big.Int
	tenEighteenF  big.Float
	LastBlock     xsync.MapOf[model.TokenKey, uint64]
	PriceCache    map[string]*lru.Cache[uint64, big.Int]
	crossPrices   map[string]big.Int
	crossTokens   map[string]model.TokenKey
	LastPrice     big.Int
}

func (s *service) checkAndCacheSignature(
	reportedValues *xsync.MapOf[bls12381.G1Affine, big.Int], signature bls12381.G1Affine, signer model.Signer,
	hash bls12381.G1Affine, totalVoted *big.Int,
) error {
	s.signatureMutex.Lock()
	defer s.signatureMutex.Unlock()

	cached, _ := s.signatureCache.Get(hash)

	packed := model.Signature{
		Signature: signature,
		Signer:    signer,
	}

	for _, item := range cached {
		if item.Signer.PublicKey == signer.PublicKey {
			utils.Logger.
				With("Address", address.Calculate(signer.PublicKey[:])).
				Debug("Duplicated signature")
			return fmt.Errorf("duplicated signature")
		}
	}

	reportedValues.Store(hash, *totalVoted)
	cached = append(cached, packed)
	s.signatureCache.Add(hash, cached)

	return nil
}

type SaveSignatureArgs struct {
	Info      model.PriceInfo
	Hash      bls12381.G1Affine
	Consensus bool
	Voted     *big.Int
}

func (s *service) saveSignatures(ctx context.Context, args SaveSignatureArgs) error {
	utils.Logger.
		With("Block", args.Info.Asset.Block).
		With("Hash", fmt.Sprintf("%x", args.Hash.Bytes())[:8]).
		Debug("Saving into DB")

	signatures, ok := s.signatureCache.Get(args.Hash)
	if !ok {
		utils.Logger.
			With("Block", args.Info.Asset.Block).
			With("Hash", fmt.Sprintf("%x", args.Hash.Bytes())[:8]).
			Debug("Cache not found")
		return consts.ErrSignatureNotfound
	}

	var newSigners []model.Signer
	var newSignatures []bls12381.G1Affine
	var keys [][]byte

	currentRecords, err := s.assetPriceRepo.Find(
		ctx,
		args.Info.Asset.Block, args.Info.Asset.Token.Chain, args.Info.Asset.Token.Name, args.Info.Asset.Token.Pair,
	)

	if err != nil && !ent.IsNotFound(err) {
		return err
	}

	for i := range signatures {
		signature := signatures[i]
		keys = append(keys, signature.Signer.PublicKey[:])

		if !IsNewSigner(signature, currentRecords) {
			continue
		}

		newSignatures = append(newSignatures, signature.Signature)
		newSigners = append(newSigners, signature.Signer)
	}

	err = s.signerRepo.CreateSigners(ctx, newSigners)
	if err != nil {
		utils.Logger.
			With("Block", args.Info.Asset.Block).
			With("Hash", fmt.Sprintf("%x", args.Hash.Bytes())[:8]).
			Debug("Failed to upsert token signers.")
		return err
	}

	signerIDs, err := s.signerRepo.GetSingerIDsByKeys(ctx, keys)

	if err != nil {
		utils.Logger.
			With("Block", args.Info.Asset.Block).
			With("Hash", fmt.Sprintf("%x", args.Hash.Bytes())[:8]).
			Debug("Filed to upsert signers")
		return err
	}

	var aggregate bls12381.G1Affine

	for _, record := range currentRecords {
		if record.Price.Cmp(&args.Info.Price) == 0 {
			currentAggregate, err := bls.RecoverSignature([48]byte(record.Signature))

			if err != nil {
				utils.Logger.
					With("Block", args.Info.Asset.Block).
					With("Hash", fmt.Sprintf("%x", args.Hash.Bytes())[:8]).
					With("Error", err).
					Debug("Failed to recover signature")
				return consts.ErrCantRecoverSignature
			}

			newSignatures = append(newSignatures, currentAggregate)
			break
		}
	}

	aggregate, err = bls.AggregateSignatures(newSignatures)

	if err != nil {
		utils.Logger.
			With("Block", args.Info.Asset.Block).
			With("Hash", fmt.Sprintf("%x", args.Hash.Bytes())[:8]).
			Debug("Filed to aggregate signatures")
		return consts.ErrCantAggregateSignatures
	}

	signatureBytes := aggregate.Bytes()

	// TODO: Handle cases where signerIDs need to be removed
	err = s.assetPriceRepo.Upsert(ctx, model.AssetPrice{
		Pair:         strings.ToLower(args.Info.Asset.Token.Pair),
		Name:         args.Info.Asset.Token.Name,
		Chain:        args.Info.Asset.Token.Chain,
		Block:        args.Info.Asset.Block,
		Price:        args.Info.Price,
		SignersCount: uint64(len(signatures)),
		Signature:    signatureBytes[:],
		Consensus:    args.Consensus,
		Voted:        *args.Voted,
		SignerIDs:    signerIDs,
	})

	if err != nil {
		utils.Logger.
			With("Block", args.Info.Asset.Block).
			With("Hash", fmt.Sprintf("%x", args.Hash.Bytes())[:8]).
			Debug("Failed to upsert asset price")
		return err
	}

	return nil
}

func (s *service) GetBlockNumber(ctx context.Context, network string) (*uint64, error) {
	blockNumber, err := s.ethRPC.GetBlockNumber(ctx, network)

	if err != nil {
		s.ethRPC.RefreshRPC(network)
		return nil, err
	}

	return &blockNumber, nil
}

func (s *service) GetPriceAtBlockFromPair(
	network string, blockNumber uint64, pairAddr string, decimalDif int64, inverse bool,
) (*big.Int, error) {
	pair, err := s.ethRPC.GetNewUniV3Contract(network, pairAddr, false)

	if err != nil {
		s.ethRPC.RefreshRPC(network)
		return nil, err
	}

	data, err := pair.Slot0(
		&bind.CallOpts{
			//nolint:gosec // refactoring needed
			BlockNumber: big.NewInt(int64(blockNumber)),
		})

	if err != nil {
		s.ethRPC.RefreshRPC(network)
		return nil, err
	}

	s.LastPrice = *s.priceFromSqrtX96(data.SqrtPriceX96, decimalDif, inverse)
	lruCache := s.PriceCache[strings.ToLower(pairAddr)]
	lruCache.Add(blockNumber, s.LastPrice)

	return &s.LastPrice, nil
}

func (s *service) priceFromSqrtX96(sqrtPriceX96 *big.Int, decimalDif int64, inverse bool) *big.Int {
	var decimalFix big.Int
	var powerUp big.Int
	var rawPrice big.Int
	var price big.Int
	var factor big.Int

	decimalFix.Mul(sqrtPriceX96, &s.tenEighteen)
	powerUp.Exp(&decimalFix, big.NewInt(2), nil)
	rawPrice.Div(&powerUp, &s.twoOneNineTwo)

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

func (s *service) syncBlock(ctx context.Context, token model.Token, caser cases.Caser, key *model.TokenKey, blockInx uint64) error {
	lastSynced, ok := s.LastBlock.Load(*key)

	if ok && blockInx <= lastSynced {
		return consts.ErrDataTooOld
	}

	price, err := s.GetPriceAtBlockFromPair(
		token.Chain,
		blockInx,
		token.Pair,
		token.Delta,
		token.Invert,
	)

	if err != nil {
		utils.Logger.Error(
			fmt.Sprintf("Failed to get token price from %s RPC.", token.Chain))
		s.ethRPC.RefreshRPC(token.Chain)
		return err
	}

	for _, cross := range token.Cross {
		stored := s.crossPrices[cross]

		if stored.Cmp(big.NewInt(0)) == 0 {
			return consts.ErrCrossPriceIsNotZero
		}

		price.Mul(price, &stored)
	}

	for range token.Cross {
		price.Div(price, &s.tenEighteen)
	}

	if token.ID != nil {
		s.crossPrices[*token.ID] = *price
	}

	var priceF big.Float
	priceF.Quo(new(big.Float).SetInt(price), &s.tenEighteenF)
	priceStr := fmt.Sprintf("%.18f %s", &priceF, token.Unit)

	lastSynced, ok = s.LastBlock.Load(*key)

	if ok && blockInx <= lastSynced {
		return consts.ErrDataTooOld
	}

	utils.Logger.
		With("Block", blockInx).
		With("Price", priceStr).
		Info(caser.String(token.Name))

	key = s.TokenKey(token)

	priceInfo := model.PriceInfo{
		Price: *price,
		Asset: model.AssetKey{
			Block: blockInx,
			Token: *key,
		},
	}

	signature, hash := crypto.Identity.Bls.Sign(priceInfo.Sia().Bytes())

	if token.Send && !conn.IsClosed {
		compressedSignature := signature.Bytes()
		priceReport := model.PriceReportPacket{
			PriceInfo: priceInfo,
			Signature: compressedSignature,
		}

		conn.Send(consts.OpCodePriceReport, priceReport.Sia().Bytes())
	}

	if token.Store {
		err = s.RecordSignature(
			ctx,
			signature,
			*crypto.Identity.ExportEvmSigner(),
			hash,
			priceInfo,
			false,
			true,
		)

		if err != nil {
			return err
		}
	}

	s.LastBlock.Store(*key, blockInx)

	return nil
}

func (s *service) SyncBlocks(ctx context.Context, token model.Token, key model.TokenKey, latest uint64) error {
	block, ok := s.LastBlock.Load(key)
	if !ok {
		return consts.ErrCantLoadLastBlock
	}

	caser := cases.Title(language.English, cases.NoLower)

	for blockInx := block + 1; blockInx < latest; blockInx++ {
		err := s.syncBlock(ctx, token, caser, &key, blockInx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *service) TokenKey(token model.Token) *model.TokenKey {
	var cross []model.TokenKey

	for _, id := range token.Cross {
		cross = append(cross, s.crossTokens[id])
	}

	toHash := new(sia.ArraySia[model.TokenKey]).
		AddArray8(cross, func(s *sia.ArraySia[model.TokenKey], item model.TokenKey) {
			s.EmbedBytes(item.Sia().Bytes())
		}).Bytes()

	hash := utils.Shake(toHash)

	key := model.TokenKey{
		Name:   strings.ToLower(token.Name),
		Pair:   strings.ToLower(token.Pair),
		Chain:  strings.ToLower(token.Chain),
		Delta:  token.Delta,
		Invert: token.Invert,
		Cross:  string(hash),
	}

	return &key
}

func New(
	ethRPC ethereum.RPC,
	pos pos.Service,
	signerRepo repository.Signer,
	assetPriceRepo repository.AssetPrice,
) Service {
	s := service{
		ethRPC:         ethRPC,
		pos:            pos,
		signerRepo:     signerRepo,
		assetPriceRepo: assetPriceRepo,

		consensus:       nil,
		signatureCache:  nil,
		SupportedTokens: map[model.TokenKey]bool{},
		signatureMutex:  sync.Mutex{},
		LastBlock:       *xsync.NewMapOf[model.TokenKey, uint64](),
		PriceCache:      map[string]*lru.Cache[uint64, big.Int]{},
		crossPrices:     map[string]big.Int{},
		crossTokens:     map[string]model.TokenKey{},
	}

	DebouncedSaveSignatures = utils.Debounce[model.AssetKey, SaveSignatureArgs](5*time.Second, s.saveSignatures)

	s.twoOneNineTwo.Exp(big.NewInt(2), big.NewInt(192), nil)
	s.tenEighteen.Exp(big.NewInt(10), big.NewInt(18), nil)
	s.tenEighteenF.SetInt(&s.tenEighteen)

	if config.App.Plugins.Uniswap != nil {
		for _, t := range config.App.Plugins.Uniswap.Tokens {
			token := model.NewTokenFromCfg(t)

			key := s.TokenKey(token)
			s.SupportedTokens[*key] = true
		}
	}

	for _, t := range config.App.Plugins.Uniswap.Tokens {
		token := model.NewTokenFromCfg(t)
		var err error
		s.PriceCache[strings.ToLower(token.Pair)], err = lru.New[uint64, big.Int](SizeOfPriceCacheLru)

		if err != nil {
			utils.Logger.Error("Failed to initialize token map.")
			os.Exit(1)
		}
	}

	var err error
	s.signatureCache, err = lru.New[bls12381.G1Affine, []model.Signature](evmlog.LruSize)
	if err != nil {
		utils.Logger.Error("Failed to create token price signature cache.")
		os.Exit(1)
	}

	// TODO: This is vulnerable to flood attacks
	s.consensus, err = lru.New[model.AssetKey, xsync.MapOf[bls12381.G1Affine, big.Int]](evmlog.LruSize)
	if err != nil {
		utils.Logger.Error("Failed to create token price consensus cache.")
		os.Exit(1)
	}

	return &s
}
