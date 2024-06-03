package pos

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (s *service) Slash(ctx context.Context, address [20]byte, to common.Address, amount *big.Int, nftIDs []*big.Int) error {
	return nil
}
