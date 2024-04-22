package model

import (
	"math/big"

	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/utils"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"

	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

type CorrectnessReportPacket struct {
	Correctness
	Signature [48]byte
}

func (c *CorrectnessReportPacket) Sia() *sia.Sia {
	return new(sia.Sia).
		EmbedSia(c.Correctness.Sia()).
		AddByteArray8(c.Signature[:])
}

func (c *CorrectnessReportPacket) DeSia(sia *sia.Sia) *CorrectnessReportPacket {
	c.Correctness.DeSia(sia)
	copy(c.Signature[:], sia.ReadByteArray8())

	return c
}

//////

type BroadcastCorrectnessPacket struct {
	Info      Correctness
	Signature [48]byte
	Signer    Signer
}

func (b *BroadcastCorrectnessPacket) Sia() *sia.Sia {
	return new(sia.Sia).
		EmbedSia(b.Info.Sia()).
		AddByteArray8(b.Signature[:]).
		EmbedSia(b.Signer.Sia())
}

func (b *BroadcastCorrectnessPacket) DeSia(sia *sia.Sia) *BroadcastCorrectnessPacket {
	b.Info.DeSia(sia)
	copy(b.Signature[:], sia.ReadByteArray8())
	b.Signer.DeSia(sia)

	return b
}

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
}

func (c *Correctness) Sia() *sia.Sia {
	return new(sia.Sia).
		AddUInt64(c.Timestamp).
		AddByteArray8(c.Hash).
		AddByteArray8(c.Topic[:]).
		AddBool(c.Correct)
}

func (c *Correctness) DeSia(sia *sia.Sia) *Correctness {
	c.Timestamp = sia.ReadUInt64()
	copy(c.Hash, sia.ReadByteArray8())
	copy(c.Topic[:], sia.ReadByteArray8())
	c.Correct = sia.ReadBool()

	return c
}

func (c *Correctness) Bls() (bls12381.G1Affine, error) {
	hash, err := bls.Hash(c.Sia().Content)
	if err != nil {
		utils.Logger.Error("Can't hash bls: %v", err)
		return bls12381.G1Affine{}, err
	}

	return hash, err
}
