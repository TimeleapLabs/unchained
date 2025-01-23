package handler

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/transport/server/packet"
	"github.com/TimeleapLabs/unchained/internal/utils/hash"
)

// Attestation is a method that handles attestation report packets.
func (h *consumer) Attestation(ctx context.Context, message []byte) {
	packet := new(packet.BroadcastAttestationPacket).FromBytes(message)

	err := h.attestation.RecordSignature(
		ctx,
		packet.Signature,
		packet.Signer,
		hash.Hash(&packet.Info),
		packet.Info,
		true,
	)
	if err != nil {
		return
	}
}

// Attestation is not defined for worker nodes.
func (w worker) Attestation(_ context.Context, _ []byte) {}
