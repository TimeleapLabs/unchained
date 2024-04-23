package uniswap

import (
	"context"
	"github.com/KenshiTech/unchained/internal/service/uniswap"
)

type Uniswap struct {
	chain          string
	uniswapService uniswap.Service
}

func (u *Uniswap) Run() {
	err := u.uniswapService.ProcessBlocks(context.TODO(), u.chain)
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
