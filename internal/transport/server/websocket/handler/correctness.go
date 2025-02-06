package handler

import (
	"crypto/ed25519"

	"github.com/TimeleapLabs/timeleap/internal/model"
	"github.com/TimeleapLabs/timeleap/internal/transport/server/packet"
)

// AttestationRecord is a handler for attestation report.
func AttestationRecord(payload []byte, signature [64]byte, signer ed25519.PublicKey) ([]byte, error) {
	attestation := new(model.Attestation).FromBytes(payload)

	broadcastPacket := packet.BroadcastAttestationPacket{
		Info:      *attestation,
		Signature: signature,
		Signer:    signer,
	}

	return broadcastPacket.Sia().Bytes(), nil
}
