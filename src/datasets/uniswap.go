package datasets

import (
	"math/big"

	"github.com/KenshiTech/unchained/crypto/bls"
)

type TokenKey struct {
	Name   string
	Pair   string
	Chain  string
	Delta  int64
	Invert bool
	Cross  string
}

type AssetKey struct {
	Token TokenKey
	Block uint64
}

type PriceInfo struct {
	Asset AssetKey
	Price big.Int
}

type PriceReport struct {
	PriceInfo PriceInfo
	Signature [48]byte
}

type BroadcastPricePacket struct {
	Info      PriceInfo
	Signature [48]byte
	Signer    bls.Signer
}
