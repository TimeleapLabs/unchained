package model

import (
	"math/big"

	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/utils"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"

	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

type Correctness struct {
	SignersCount uint64
	Signature    []byte
	Consensus    bool
	Voted        big.Int
	SignerIDs    []int
	Timestamp    uint64
	Hash         []byte
	Topic        [64]byte
	Correct      bool

	Signers []Signer
}

func (c *Correctness) Sia() sia.Sia {
	return sia.New().
		AddUInt64(c.Timestamp).
		AddByteArray8(c.Hash).
		AddByteArray8(c.Topic[:]).
		AddBool(c.Correct)
}

func (c *Correctness) FromBytes(payload []byte) *Correctness {
	siaMessage := sia.NewFromBytes(payload)
	return c.FromSia(siaMessage)
}

func (c *Correctness) FromSia(sia sia.Sia) *Correctness {
	c.Timestamp = sia.ReadUInt64()
	copy(c.Hash, sia.ReadByteArray8())
	copy(c.Topic[:], sia.ReadByteArray8())
	c.Correct = sia.ReadBool()

	return c
}

func (c *Correctness) Bls() bls12381.G1Affine {
	hash, err := bls.Hash(c.Sia().Bytes())
	if err != nil {
		utils.Logger.Error("Can't hash bls: %v", err)
		return bls12381.G1Affine{}
	}

	return hash
}
