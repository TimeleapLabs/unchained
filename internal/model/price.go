package model

import (
	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
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
	siaMessage := sia.NewFromBytes(payload)
	return p.FromSia(siaMessage)
}

func (p *PriceReportPacket) FromSia(sia sia.Sia) *PriceReportPacket {
	p.PriceInfo.FromSia(sia)
	copy(p.Signature[:], sia.ReadByteArray8())

	return p
}
