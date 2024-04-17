package uniswap

import (
	"fmt"
	"github.com/KenshiTech/unchained/internal/model"
	"github.com/KenshiTech/unchained/internal/utils"
	"math/big"
	"os"
	"strings"

	"github.com/KenshiTech/unchained/internal/crypto/ethereum"

	"github.com/KenshiTech/unchained/internal/config"
	"github.com/KenshiTech/unchained/internal/service/uniswap"
	lru "github.com/hashicorp/golang-lru/v2"
)

const (
	SizeOfPriceCacheLru = 128
)

type Uniswap struct {
	chain          string
	uniswapService *uniswap.Service
	ethRPC         *ethereum.Repository
}

func (u *Uniswap) Run() {
	if config.App.Plugins.Uniswap == nil {
		return
	}

	currBlockNumber, err := u.uniswapService.GetBlockNumber(u.chain)
	if err != nil {
		utils.Logger.Error(
			fmt.Sprintf("Couldn't get latest block from %s RPC.", u.chain))
		u.ethRPC.RefreshRPC(u.chain)
		return
	}

	for _, token := range model.NewTokensFromCfg(config.App.Plugins.Uniswap.Tokens) {
		if token.Chain != u.chain {
			continue
		}

		// TODO: this can be cached
		key := u.uniswapService.TokenKey(token)
		tokenLastBlock, exists := u.uniswapService.LastBlock.Load(*key)

		if !exists {
			u.uniswapService.LastBlock.Store(*key, *currBlockNumber-1)
		} else if tokenLastBlock == *currBlockNumber {
			return
		}

		err = u.uniswapService.SyncBlocks(token, *key, *currBlockNumber)
		if err != nil {
			return
		}
	}
}

func New(
	chanName string, tokens []config.Token,
	uniswapService *uniswap.Service,
	ethRPC *ethereum.Repository,
) *Uniswap {
	u := Uniswap{
		chain:          chanName,
		uniswapService: uniswapService,
		ethRPC:         ethRPC,
	}

	for _, t := range tokens {
		token := model.NewTokenFromCfg(t)
		var err error
		u.uniswapService.PriceCache[strings.ToLower(token.Pair)], err = lru.New[uint64, big.Int](SizeOfPriceCacheLru)

		if err != nil {
			utils.Logger.Error("Failed to initialize token map.")
			os.Exit(1)
		}
	}

	return &u
}
