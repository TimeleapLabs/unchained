package handler

import (
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/transport/server/packet"
	"golang.org/x/crypto/ed25519"
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
