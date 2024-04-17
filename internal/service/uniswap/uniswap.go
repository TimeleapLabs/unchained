package uniswap

import (
	"context"
	"fmt"
	"github.com/KenshiTech/unchained/internal/model"
	"github.com/KenshiTech/unchained/internal/utils/address"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/KenshiTech/unchained/internal/consts"

	"github.com/KenshiTech/unchained/internal/repository"

	"github.com/KenshiTech/unchained/internal/crypto"
	"github.com/KenshiTech/unchained/internal/crypto/ethereum"

	"github.com/KenshiTech/unchained/internal/utils"

	"github.com/KenshiTech/unchained/internal/config"

	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/crypto/shake"
	"github.com/KenshiTech/unchained/internal/ent"
	"github.com/KenshiTech/unchained/internal/pos"
	"github.com/KenshiTech/unchained/internal/service/evmlog"
	"github.com/KenshiTech/unchained/internal/transport/client/conn"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	lru "github.com/hashicorp/golang-lru/v2"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
	"github.com/puzpuzpuz/xsync/v3"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	MaxBlockNumberDelta = 96
)

var DebouncedSaveSignatures func(key model.AssetKey, arg SaveSignatureArgs)

type Service struct {
	ethRPC         *ethereum.Repository
	pos            *pos.Repository
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

func (u *Service) CheckAndCacheSignature(
	reportedValues *xsync.MapOf[bls12381.G1Affine, big.Int], signature bls12381.G1Affine, signer model.Signer,
	hash bls12381.G1Affine, totalVoted *big.Int,
) error {
	u.signatureMutex.Lock()
	defer u.signatureMutex.Unlock()

	cached, _ := u.signatureCache.Get(hash)

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
	u.signatureCache.Add(hash, cached)

	return nil
}

// TODO: This needs to work with different datasets
// TODO: Can we turn this into a library func?
func (u *Service) RecordSignature(
	signature bls12381.G1Affine, signer model.Signer, hash bls12381.G1Affine, info model.PriceInfo, debounce bool, historical bool,
) error {
	if supported := u.SupportedTokens[info.Asset.Token]; !supported {
		utils.Logger.
			With("Name", info.Asset.Token.Name).
			With("Chain", info.Asset.Token.Chain).
			With("Pair", info.Asset.Token.Pair).
			Debug("Token not supported")
		return consts.ErrTokenNotSupported
	}

	// TODO: Standalone mode shouldn't call this or check consensus
	blockNumber, err := u.GetBlockNumber(info.Asset.Token.Chain)
	if err != nil {
		utils.Logger.
			With("Network", info.Asset.Token.Chain).
			With("Error", err).
			Error("Failed to get the latest block number")
		u.ethRPC.RefreshRPC(info.Asset.Token.Chain)
		// TODO: we should retry
		return err
	}

	if !historical {
		// TODO: this won't work for Arbitrum
		if *blockNumber-info.Asset.Block > MaxBlockNumberDelta {
			utils.Logger.
				With("Packet", info.Asset.Block).
				With("Current", *blockNumber).
				Debug("Data too old")
			return consts.ErrDataTooOld // Data too old
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

	votingPower, err := u.pos.GetVotingPowerOfPublicKey(signer.PublicKey)
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

	err = u.CheckAndCacheSignature(&reportedValues, signature, signer, hash, totalVoted)
	if err != nil {
		return err
	}

	saveArgs := SaveSignatureArgs{
		Hash:      hash,
		Info:      info,
		Voted:     totalVoted,
		Consensus: isMajority,
	}

	if !debounce {
		err = u.saveSignatures(saveArgs)
		if err != nil {
			return err
		}
		return nil
	}

	reportLog := utils.Logger.
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

	DebouncedSaveSignatures(info.Asset, saveArgs)

	return nil
}

type SaveSignatureArgs struct {
	Info      model.PriceInfo
	Hash      bls12381.G1Affine
	Consensus bool
	Voted     *big.Int
}

func IsNewSigner(signature model.Signature, records []*ent.AssetPrice) bool {
	for _, record := range records {
		for _, signer := range record.Edges.Signers {
			if signature.Signer.PublicKey == [96]byte(signer.Key) {
				return false
			}
		}
	}

	return true
}

func (u *Service) saveSignatures(args SaveSignatureArgs) error {
	utils.Logger.
		With("Block", args.Info.Asset.Block).
		With("Hash", fmt.Sprintf("%x", args.Hash.Bytes())[:8]).
		Debug("Saving into DB")

	signatures, ok := u.signatureCache.Get(args.Hash)
	if !ok {
		utils.Logger.
			With("Block", args.Info.Asset.Block).
			With("Hash", fmt.Sprintf("%x", args.Hash.Bytes())[:8]).
			Debug("Cache not found")
		return consts.ErrSignatureNotfound
	}

	ctx := context.Background()

	var newSigners []model.Signer
	var newSignatures []bls12381.G1Affine
	var keys [][]byte

	currentRecords, err := u.assetPriceRepo.Find(
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

	err = u.signerRepo.CreateSigners(ctx, newSigners)
	if err != nil {
		utils.Logger.
			With("Block", args.Info.Asset.Block).
			With("Hash", fmt.Sprintf("%x", args.Hash.Bytes())[:8]).
			Debug("Failed to upsert token signers.")
		return err
	}

	signerIDs, err := u.signerRepo.GetSingerIDsByKeys(ctx, keys)

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
	err = u.assetPriceRepo.Upsert(ctx, model.AssetPrice{
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

func (u *Service) syncBlock(token model.Token, caser cases.Caser, key *model.TokenKey, blockInx uint64) error {
	lastSynced, ok := u.LastBlock.Load(*key)

	if ok && blockInx <= lastSynced {
		return consts.ErrDataTooOld
	}

	price, err := u.GetPriceAtBlockFromPair(
		token.Chain,
		blockInx,
		token.Pair,
		token.Delta,
		token.Invert,
	)

	if err != nil {
		utils.Logger.Error(
			fmt.Sprintf("Failed to get token price from %s RPC.", token.Chain))
		u.ethRPC.RefreshRPC(token.Chain)
		return err
	}

	for _, cross := range token.Cross {
		stored := u.crossPrices[cross]

		if stored.Cmp(big.NewInt(0)) == 0 {
			return consts.ErrCrossPriceIsNotZero
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
		return consts.ErrDataTooOld
	}

	utils.Logger.
		With("Block", blockInx).
		With("Price", priceStr).
		Info(caser.String(token.Name))

	key = u.TokenKey(token)

	priceInfo := model.PriceInfo{
		Price: *price,
		Asset: model.AssetKey{
			Block: blockInx,
			Token: *key,
		},
	}

	toHash := priceInfo.Sia().Content
	signature, hash := bls.Sign(*crypto.Identity.Bls.SecretKey, toHash)

	if token.Send && !conn.IsClosed {
		compressedSignature := signature.Bytes()

		priceReport := model.PriceReport{
			PriceInfo: priceInfo,
			Signature: compressedSignature,
		}

		payload := priceReport.Sia().Content
		conn.Send(consts.OpCodePriceReport, payload)
	}

	if token.Store {
		err = u.RecordSignature(
			signature,
			*crypto.Identity.ExportBlsSigner(),
			hash,
			priceInfo,
			false,
			true,
		)

		if err != nil {
			return err
		}
	}

	u.LastBlock.Store(*key, blockInx)

	return nil
}

func (u *Service) SyncBlocks(token model.Token, key model.TokenKey, latest uint64) error {
	block, ok := u.LastBlock.Load(key)
	if !ok {
		return consts.ErrCantLoadLastBlock
	}

	caser := cases.Title(language.English, cases.NoLower)

	for blockInx := block + 1; blockInx < latest; blockInx++ {
		err := u.syncBlock(token, caser, &key, blockInx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *Service) TokenKey(token model.Token) *model.TokenKey {
	var cross []model.TokenKey

	for _, id := range token.Cross {
		cross = append(cross, u.crossTokens[id])
	}

	toHash := new(sia.ArraySia[model.TokenKey]).
		AddArray8(cross, func(s *sia.ArraySia[model.TokenKey], item model.TokenKey) {
			s.EmbedSia(item.Sia())
		}).Content

	hash := shake.Shake(toHash)

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
	ethRPC *ethereum.Repository,
	pos *pos.Repository,
	signerRepo repository.Signer,
	assetPriceRepo repository.AssetPrice,
) *Service {
	u := Service{
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
	DebouncedSaveSignatures = utils.Debounce[model.AssetKey, SaveSignatureArgs](5*time.Second, u.saveSignatures)
	u.twoOneNineTwo.Exp(big.NewInt(2), big.NewInt(192), nil)
	u.tenEighteen.Exp(big.NewInt(10), big.NewInt(18), nil)
	u.tenEighteenF.SetInt(&u.tenEighteen)

	if config.App.Plugins.Uniswap != nil {
		for _, t := range config.App.Plugins.Uniswap.Tokens {
			token := model.NewTokenFromCfg(t)

			key := u.TokenKey(token)
			u.SupportedTokens[*key] = true
		}
	}

	var err error
	u.signatureCache, err = lru.New[bls12381.G1Affine, []model.Signature](evmlog.LruSize)
	if err != nil {
		utils.Logger.Error("Failed to create token price signature cache.")
		os.Exit(1)
	}

	// TODO: This is vulnerable to flood attacks
	u.consensus, err = lru.New[model.AssetKey, xsync.MapOf[bls12381.G1Affine, big.Int]](evmlog.LruSize)
	if err != nil {
		utils.Logger.Error("Failed to create token price consensus cache.")
		os.Exit(1)
	}

	return &u
}
