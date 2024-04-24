package scheduler

import (
	"context"

	"github.com/KenshiTech/unchained/internal/service/uniswap"
)

// Uniswap is a scheduler for Uniswap and keep task's dependencies.
type Uniswap struct {
	chain          string
	uniswapService uniswap.Service
}

// Run will trigger by the scheduler and process the Uniswap blocks.
func (u *Uniswap) Run() {
	err := u.uniswapService.ProcessBlocks(context.TODO(), u.chain)
	if err != nil {
		panic(err)
	}
}

// NewUniswap will create a new Uniswap task.
func NewUniswap(
	chanName string,
	uniswapService uniswap.Service,
) *Uniswap {
	u := Uniswap{
		chain:          chanName,
		uniswapService: uniswapService,
	}

	return &u
}
