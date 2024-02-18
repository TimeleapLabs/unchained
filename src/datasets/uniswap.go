package datasets

import "math/big"

type PriceInfo struct {
	Block uint64
	Price big.Int
}

type PriceReport struct {
	PriceInfo PriceInfo
	Signature [48]byte
}
