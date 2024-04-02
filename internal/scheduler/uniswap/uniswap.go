package uniswap

import (
	"fmt"
	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/ethereum"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/service/uniswap"
	lru "github.com/hashicorp/golang-lru/v2"
	"math/big"
	"os"
	"strings"
)

const (
	SizeOfPriceCacheLru = 128
)

type Uniswap struct {
	chain          string
	uniswapService uniswap.Service
}

func (u *Uniswap) Run() {
	log.Logger.With("Chain", u.chain).Info("Run Uniswap task")

	currBlockNumber, err := u.uniswapService.GetBlockNumber(u.chain)
	if err != nil {
		log.Logger.Error(
			fmt.Sprintf("Couldn't get latest block from %s RPC.", u.chain))
		ethereum.RefreshRPC(u.chain)
		return
	}

	for _, token := range uniswap.NewTokensFromCfg(config.App.Plugins.Uniswap.Tokens) {
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

		u.uniswapService.SyncBlocks(token, *key, *currBlockNumber)
	}
}

func New(chanName string, tokens []config.Token) *Uniswap {
	u := Uniswap{
		chain: chanName,
	}

	for _, t := range tokens {
		token := uniswap.NewTokenFromCfg(t)
		var err error
		u.uniswapService.PriceCache[strings.ToLower(token.Pair)], err = lru.New[uint64, big.Int](SizeOfPriceCacheLru)

		if err != nil {
			log.Logger.Error("Failed to initialize token map.")
			os.Exit(1)
		}

		key := u.uniswapService.TokenKey(token)
		u.uniswapService.SupportedTokens[*key] = true
	}

	return &u
}
