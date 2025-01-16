package packet

import (
	"github.com/TimeleapLabs/unchained/internal/model"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

type AttestationPacket struct {
	model.Attestation
	Signature [64]byte
}

func (c *AttestationPacket) Sia() sia.Sia {
	return sia.New().
		EmbedBytes(c.Attestation.Sia().Bytes()).
		AddByteArray8(c.Signature[:])
}

func (c *AttestationPacket) FromBytes(payload []byte) *AttestationPacket {
	siaMessage := sia.NewFromBytes(payload)

	c.Attestation.FromSia(siaMessage)
	copy(c.Signature[:], siaMessage.ReadByteArray8())

	return c
}

type BroadcastAttestationPacket struct {
	Info      model.Attestation
	Signature [64]byte
	Signer    model.Signer
}

func (b *BroadcastAttestationPacket) Sia() sia.Sia {
	return sia.New().
		EmbedBytes(b.Info.Sia().Bytes()).
		AddByteArray8(b.Signature[:]).
		EmbedBytes(b.Signer.Sia().Bytes())
}

func (b *BroadcastAttestationPacket) FromBytes(payload []byte) *BroadcastAttestationPacket {
	siaMessage := sia.NewFromBytes(payload)

	b.Info.FromSia(siaMessage)
	copy(b.Signature[:], siaMessage.ReadByteArray8())
	b.Signer.FromSia(siaMessage)

	return b
}
