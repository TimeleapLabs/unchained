package uniswap

import (
	"math/big"
	"os"
	"strings"

	"github.com/KenshiTech/unchained/internal/model"
	"github.com/KenshiTech/unchained/internal/utils"

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
}

func (u *Uniswap) Run() {
	if config.App.Plugins.Uniswap == nil {
		return
	}

	err := u.uniswapService.ProcessBlocks(u.chain)
	if err != nil {
		panic(err)
	}
}

func New(
	chanName string,
	uniswapService *uniswap.Service,
) *Uniswap {
	u := Uniswap{
		chain:          chanName,
		uniswapService: uniswapService,
	}

	for _, t := range config.App.Plugins.Uniswap.Tokens {
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
