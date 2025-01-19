package packet

import (
	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
	"github.com/TimeleapLabs/unchained/internal/model"
)

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
