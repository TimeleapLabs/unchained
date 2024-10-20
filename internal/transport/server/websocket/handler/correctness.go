package handler

import (
	"github.com/TimeleapLabs/unchained/internal/transport/server/packet"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/middleware"
	"github.com/gorilla/websocket"
)

// AttestationRecord is a handler for attestation report.
func AttestationRecord(conn *websocket.Conn, payload []byte) ([]byte, error) {
	err := middleware.IsConnectionAuthenticated(conn)
	if err != nil {
		return []byte{}, err
	}

	attestation := new(packet.AttestationPacket).FromBytes(payload)

	signer, err := middleware.IsMessageValid(conn, *attestation.Attestation.Bls(), attestation.Signature)
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
