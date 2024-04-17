package model

import (
	"math/big"

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

type PriceReport struct {
	PriceInfo PriceInfo
	Signature [48]byte
}

type BroadcastPricePacket struct {
	Info      PriceInfo
	Signature [48]byte
	Signer    Signer
}

func (t *TokenKey) Sia() *sia.Sia {
	return new(sia.Sia).
		AddString8(t.Name).
		AddString8(t.Pair).
		AddString8(t.Chain).
		AddInt64(t.Delta).
		AddBool(t.Invert).
		AddString8(t.Cross)
}

func (t *TokenKey) DeSia(sia *sia.Sia) *TokenKey {
	t.Name = sia.ReadString8()
	t.Pair = sia.ReadString8()
	t.Chain = sia.ReadString8()
	t.Delta = sia.ReadInt64()
	t.Invert = sia.ReadBool()
	t.Cross = sia.ReadString8()

	return t
}

func (a *AssetKey) Sia() *sia.Sia {
	return new(sia.Sia).
		EmbedSia(a.Token.Sia()).
		AddUInt64(a.Block)
}

func (a *AssetKey) DeSia(sia *sia.Sia) *AssetKey {
	a.Token.DeSia(sia)
	a.Block = sia.ReadUInt64()

	return a
}

func (p *PriceInfo) Sia() *sia.Sia {
	return new(sia.Sia).
		EmbedSia(p.Asset.Sia()).
		AddBigInt(&p.Price)
}

func (p *PriceInfo) DeSia(sia *sia.Sia) *PriceInfo {
	p.Asset.DeSia(sia)
	p.Price = *sia.ReadBigInt()

	return p
}

func (p *PriceReport) Sia() *sia.Sia {
	return new(sia.Sia).
		EmbedSia(p.PriceInfo.Sia()).
		AddByteArray8(p.Signature[:])
}

func (p *PriceReport) DeSia(sia *sia.Sia) *PriceReport {
	p.PriceInfo.DeSia(sia)
	copy(p.Signature[:], sia.ReadByteArray8())

	return p
}

func (b *BroadcastPricePacket) Sia() *sia.Sia {
	return new(sia.Sia).
		EmbedSia(b.Info.Sia()).
		AddByteArray8(b.Signature[:]).
		EmbedSia(b.Signer.Sia())
}

func (b *BroadcastPricePacket) DeSia(sia *sia.Sia) *BroadcastPricePacket {
	b.Info.DeSia(sia)
	copy(b.Signature[:], sia.ReadByteArray8())
	b.Signer.DeSia(sia)

	return b
}
