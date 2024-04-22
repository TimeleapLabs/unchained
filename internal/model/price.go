package model

import (
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

type PriceReportPacket struct {
	PriceInfo PriceInfo
	Signature [48]byte
}

func (p *PriceReportPacket) Sia() *sia.Sia {
	return new(sia.Sia).
		EmbedSia(p.PriceInfo.Sia()).
		AddByteArray8(p.Signature[:])
}

func (p *PriceReportPacket) DeSia(sia *sia.Sia) *PriceReportPacket {
	p.PriceInfo.DeSia(sia)
	copy(p.Signature[:], sia.ReadByteArray8())

	return p
}
