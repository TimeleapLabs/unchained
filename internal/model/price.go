package model

import (
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

type PriceReportPacket struct {
	PriceInfo PriceInfo
	Signature [48]byte
}

func (p *PriceReportPacket) Sia() sia.Sia {
	return sia.New().
		EmbedBytes(p.PriceInfo.Sia().Bytes()).
		AddByteArray8(p.Signature[:])
}

func (p *PriceReportPacket) FromBytes(payload []byte) *PriceReportPacket {
	p.PriceInfo.FromBytes(payload)
	copy(p.Signature[:], sia.NewFromBytes(payload).ReadByteArray8())

	return p
}
