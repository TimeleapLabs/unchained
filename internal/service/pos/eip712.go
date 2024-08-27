package pos

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (s *service) Slash(_ context.Context, _ [20]byte, _ common.Address, _ *big.Int, _ []*big.Int) error {
	return nil
}
