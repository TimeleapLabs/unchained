package uniswap

import (
	"github.com/KenshiTech/unchained/internal/config"
	"github.com/KenshiTech/unchained/internal/service/uniswap"
)

type Uniswap struct {
	chain          string
	uniswapService uniswap.Service
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
	uniswapService uniswap.Service,
) *Uniswap {
	u := Uniswap{
		chain:          chanName,
		uniswapService: uniswapService,
	}

	return &u
}
