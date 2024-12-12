package model

import (
	"math/big"

	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/utils"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"

	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
)

type TokenKey struct {
	Name   string
	Pair   string
	Chain  string
	Delta  int64
	Invert bool
	Cross  string
}

func (t *TokenKey) Sia() sia.Sia {
	return sia.New().
		AddString8(t.Name).
		AddString8(t.Pair).
		AddString8(t.Chain).
		AddInt64(t.Delta).
		AddBool(t.Invert).
		AddString8(t.Cross)
}

func (t *TokenKey) FromSia(sia sia.Sia) *TokenKey {
	t.Name = sia.ReadString8()
	t.Pair = sia.ReadString8()
	t.Chain = sia.ReadString8()
	t.Delta = sia.ReadInt64()
	t.Invert = sia.ReadBool()
	t.Cross = sia.ReadString8()

	return t
}

type AssetKey struct {
	Token TokenKey
	Block uint64
}

func (a *AssetKey) Sia() sia.Sia {
	return sia.New().
		EmbedBytes(a.Token.Sia().Bytes()).
		AddUInt64(a.Block)
}

func (a *AssetKey) FromSia(sia sia.Sia) *AssetKey {
	a.Token.FromSia(sia)
	a.Block = sia.ReadUInt64()

	return a
}

type PriceInfo struct {
	Asset AssetKey
	Price big.Int
}

func (p *PriceInfo) Sia() sia.Sia {
	return sia.New().
		EmbedBytes(p.Asset.Sia().Bytes()).
		AddBigInt(&p.Price)
}

func (p *PriceInfo) FromBytes(payload []byte) *PriceInfo {
	siaMessage := sia.NewFromBytes(payload)
	return p.FromSia(siaMessage)
}

func (p *PriceInfo) FromSia(sia sia.Sia) *PriceInfo {
	p.Asset.FromSia(sia)
	p.Price = *sia.ReadBigInt()

	return p
}

func (p *PriceInfo) Bls() (bls12381.G1Affine, error) {
	hash, err := bls.Hash(p.Sia().Bytes())
	if err != nil {
		utils.Logger.With("err", err).Error("Can't hash bls")
		return bls12381.G1Affine{}, err
	}

	return hash, err
}

type BroadcastPricePacket struct {
	Info      PriceInfo
	Signature [48]byte
	Signer    Signer
}

func (b *BroadcastPricePacket) Sia() sia.Sia {
	return sia.New().
		EmbedBytes(b.Info.Sia().Bytes()).
		AddByteArray8(b.Signature[:]).
		EmbedBytes(b.Signer.Sia().Bytes())
}

func (b *BroadcastPricePacket) FromBytes(payload []byte) *BroadcastPricePacket {
	siaMessage := sia.NewFromBytes(payload)
	b.Info.FromSia(siaMessage)
	copy(b.Signature[:], siaMessage.ReadByteArray8())
	b.Signer.FromSia(siaMessage)

	return b
}
