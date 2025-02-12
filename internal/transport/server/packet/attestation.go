package packet

import (
	"crypto/ed25519"

	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
	"github.com/TimeleapLabs/timeleap/internal/model"
)

type BroadcastMessagePacket struct {
	Info      model.Message
	Signature [64]byte
	Signer    ed25519.PublicKey
}

func (b *BroadcastMessagePacket) Sia() sia.Sia {
	return sia.New().
		EmbedBytes(b.Info.Sia().Bytes()).
		EmbedBytes(b.Signature[:]).
		EmbedBytes(b.Signer)
}

func (b *BroadcastMessagePacket) FromBytes(payload []byte) *BroadcastMessagePacket {
	siaMessage := sia.NewFromBytes(payload)

	b.Info.FromSia(siaMessage)
	copy(b.Signature[:], siaMessage.ReadByteArrayN(64))
	b.Signer = ed25519.PublicKey(siaMessage.ReadByteArrayN(32))

	return b
}
