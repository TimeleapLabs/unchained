package handler

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/transport/server/packet"

	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

// CorrectnessReport is a method that handles correctness report packets.
func (h *consumer) CorrectnessReport(ctx context.Context, message []byte) {
	packet := new(packet.BroadcastCorrectnessPacket).FromBytes(message)

	signature, err := bls.RecoverSignature(packet.Signature)
	if err != nil {
		utils.Logger.
			With("Error", err).
			Error("Failed to recover packet signature")

		return
	}

	err = h.correctness.RecordSignature(
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

// CorrectnessReport is not defined for worker nodes.
func (w worker) CorrectnessReport(_ context.Context, _ []byte) {}
