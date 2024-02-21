package datasets

import "math/big"

type PriceInfo struct {
	Block uint64
	Price big.Int
	Asset string
	Chain string
	Pair  string
}

type PriceReport struct {
	PriceInfo PriceInfo
	Signature [48]byte
}

type BroadcastPacket struct {
	Info      PriceInfo
	Signature [48]byte
	Signers   [][]byte
}
