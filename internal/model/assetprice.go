package model

import "math/big"

type AssetPrice struct {
	Pair         string
	Name         string
	Chain        string
	Block        uint64
	Price        big.Int
	SignersCount uint64
	Signature    []byte
	Consensus    bool
	Voted        big.Int
	SignerIDs    []int
}
