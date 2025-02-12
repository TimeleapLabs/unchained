package handler

import (
	"crypto/ed25519"

	"github.com/TimeleapLabs/timeleap/internal/model"
	"github.com/TimeleapLabs/timeleap/internal/transport/server/packet"
)

func ToBroadcastPacket(payload []byte, signature [64]byte, signer ed25519.PublicKey) ([]byte, error) {
	attestation := new(model.Message).FromBytes(payload)

	broadcastPacket := packet.BroadcastMessagePacket{
		Info:      *attestation,
		Signature: signature,
		Signer:    signer,
	}

	return broadcastPacket.Sia().Bytes(), nil
}
