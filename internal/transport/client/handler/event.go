package handler

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/transport/server/packet"

	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

// EventLog is a method that handles event log packets.
func (h *consumer) EventLog(ctx context.Context, message []byte) {
	packet := new(packet.BroadcastEventPacket).FromBytes(message)

	signature, err := bls.RecoverSignature(packet.Signature)
	if err != nil {
		utils.Logger.
			With("Error", err).
			Error("Failed to recover packet signature")

		return
	}

	err = h.evmlog.RecordSignature(
		ctx,
		signature,
		packet.Signer,
		*packet.Info.Bls(),
		packet.Info,
		true,
		false,
	)
	if err != nil {
		return
	}
}

// EventLog is not defined for worker nodes.
func (w worker) EventLog(_ context.Context, _ []byte) {}
