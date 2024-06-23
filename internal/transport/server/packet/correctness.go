package packet

import (
	"github.com/TimeleapLabs/unchained/internal/model"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

type CorrectnessReportPacket struct {
	model.Correctness
	Signature [48]byte
}

func (c *CorrectnessReportPacket) Sia() sia.Sia {
	return sia.New().
		EmbedBytes(c.Correctness.Sia().Bytes()).
		AddByteArray8(c.Signature[:])
}

func (c *CorrectnessReportPacket) FromBytes(payload []byte) *CorrectnessReportPacket {
	siaMessage := sia.NewFromBytes(payload)

	c.Correctness.FromSia(siaMessage)
	copy(c.Signature[:], siaMessage.ReadByteArray8())

	return c
}

type BroadcastCorrectnessPacket struct {
	Info      model.Correctness
	Signature [48]byte
	Signer    model.Signer
}

func (b *BroadcastCorrectnessPacket) Sia() sia.Sia {
	return sia.New().
		EmbedBytes(b.Info.Sia().Bytes()).
		AddByteArray8(b.Signature[:]).
		EmbedBytes(b.Signer.Sia().Bytes())
}

func (b *BroadcastCorrectnessPacket) FromBytes(payload []byte) *BroadcastCorrectnessPacket {
	siaMessage := sia.NewFromBytes(payload)

	b.Info.FromSia(siaMessage)
	copy(b.Signature[:], siaMessage.ReadByteArray8())
	b.Signer.FromSia(siaMessage)

	return b
}
