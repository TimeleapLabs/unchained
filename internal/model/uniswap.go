package model

import (
	"math/big"

	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/utils"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"

	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
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

type BroadcastPricePacket struct {
	Info      PriceInfo
	Signature [48]byte
	Signer    Signer
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

func (t *TokenKey) FromBytes(payload []byte) *TokenKey {
	siaMessage := sia.NewFromBytes(payload)
	t.Name = siaMessage.ReadString8()
	t.Pair = siaMessage.ReadString8()
	t.Chain = siaMessage.ReadString8()
	t.Delta = siaMessage.ReadInt64()
	t.Invert = siaMessage.ReadBool()
	t.Cross = siaMessage.ReadString8()

	return t
}

func (a *AssetKey) Sia() sia.Sia {
	return sia.New().
		EmbedBytes(a.Token.Sia().Bytes()).
		AddUInt64(a.Block)
}

func (a *AssetKey) FromBytes(payload []byte) *AssetKey {
	a.Token.FromBytes(payload)
	a.Block = sia.NewFromBytes(payload).ReadUInt64()

	return a
}

func (p *PriceInfo) Sia() sia.Sia {
	return sia.New().
		EmbedBytes(p.Asset.Sia().Bytes()).
		AddBigInt(&p.Price)
}

func (p *PriceInfo) FromBytes(payload []byte) *PriceInfo {
	p.Asset.FromBytes(payload)
	p.Price = *sia.NewFromBytes(payload).ReadBigInt()

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

func (b *BroadcastPricePacket) Sia() sia.Sia {
	return sia.New().
		EmbedBytes(b.Info.Sia().Bytes()).
		AddByteArray8(b.Signature[:]).
		EmbedBytes(b.Signer.Sia().Bytes())
}

func (b *BroadcastPricePacket) FromBytes(payload []byte) *BroadcastPricePacket {
	b.Info.FromBytes(payload)
	copy(b.Signature[:], sia.NewFromBytes(payload).ReadByteArray8())
	b.Signer.FromBytes(payload)

	return b
}
