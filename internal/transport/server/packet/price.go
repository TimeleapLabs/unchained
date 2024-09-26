package packet

import (
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/service/uniswap/types"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

type PriceReportPacket struct {
	PriceInfo types.PriceInfo
	Signature [48]byte
}

func (p *PriceReportPacket) Sia() sia.Sia {
	return sia.New().
		EmbedBytes(p.PriceInfo.Sia().Bytes()).
		AddByteArray8(p.Signature[:])
}

func (p *PriceReportPacket) FromBytes(payload []byte) *PriceReportPacket {
	siaMessage := sia.NewFromBytes(payload)
	return p.FromSia(siaMessage)
}

func (p *PriceReportPacket) FromSia(sia sia.Sia) *PriceReportPacket {
	p.PriceInfo.FromSia(sia)
	copy(p.Signature[:], sia.ReadByteArray8())

	return p
}

type BroadcastPricePacket struct {
	Info      types.PriceInfo
	Signature [48]byte
	Signer    model.Signer
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
