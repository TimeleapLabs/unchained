package handler

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/transport/server/packet"

	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

// Attestation is a method that handles attestation report packets.
func (h *consumer) Attestation(ctx context.Context, message []byte) {
	packet := new(packet.BroadcastAttestationPacket).FromBytes(message)

	signature, err := bls.RecoverSignature(packet.Signature)
	if err != nil {
		utils.Logger.
			With("Error", err).
			Error("Failed to recover packet signature")

		return
	}

	err = h.attestation.RecordSignature(
		ctx,
		signature,
		packet.Signer,
		*packet.Info.Bls(),
		packet.Info,
		true,
	)
	if err != nil {
		return
	}
}

// Attestation is not defined for worker nodes.
func (w worker) Attestation(_ context.Context, _ []byte) {}
