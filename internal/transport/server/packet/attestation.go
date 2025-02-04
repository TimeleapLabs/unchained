package packet

import (
	"crypto/ed25519"

	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
	"github.com/TimeleapLabs/unchained/internal/model"
)

type BroadcastAttestationPacket struct {
	Info      model.Attestation
	Signature [64]byte
	Signer    ed25519.PublicKey
}

func (b *BroadcastAttestationPacket) Sia() sia.Sia {
	return sia.New().
		EmbedBytes(b.Info.Sia().Bytes()).
		EmbedBytes(b.Signature[:]).
		EmbedBytes(b.Signer)
}

func (b *BroadcastAttestationPacket) FromBytes(payload []byte) *BroadcastAttestationPacket {
	siaMessage := sia.NewFromBytes(payload)

	b.Info.FromSia(siaMessage)
	copy(b.Signature[:], siaMessage.ReadByteArrayN(64))
	b.Signer = ed25519.PublicKey(siaMessage.ReadByteArrayN(32))

	return b
}
