package handler

import (
	"github.com/TimeleapLabs/unchained/internal/transport/server/packet"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/middleware"
	"github.com/gorilla/websocket"
)

// AttestationRecord is a handler for attestation report.
func AttestationRecord(conn *websocket.Conn, payload []byte) ([]byte, error) {
	attestation := new(packet.AttestationPacket).FromBytes(payload)

	signer, err := middleware.IsMessageValid(conn, attestation.Attestation.Sia().Bytes(), attestation.Signature)
	if err != nil {
		return []byte{}, err
	}

	broadcastPacket := packet.BroadcastAttestationPacket{
		Info:      attestation.Attestation,
		Signature: attestation.Signature,
		Signer:    signer,
	}

	return broadcastPacket.Sia().Bytes(), nil
}
