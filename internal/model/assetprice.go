package model

import (
	"bytes"
	"encoding/gob"
	"math/big"
)

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
	Signers      []Signer
}

func (a AssetPrice) Hash() []byte {
	var b bytes.Buffer
	err := gob.NewEncoder(&b).Encode(a)
	if err != nil {
		return nil
	}
	return b.Bytes()
}
