package handler

import (
	"context"

	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/model"
	"github.com/KenshiTech/unchained/internal/utils"
)

func (h *consumer) CorrectnessReport(ctx context.Context, message []byte) {
	packet := new(model.BroadcastCorrectnessPacket).FromBytes(message)

	correctnessHash, err := packet.Info.Bls()
	if err != nil {
		return
	}

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
		correctnessHash,
		packet.Info,
		true,
	)
	if err != nil {
		return
	}
}

func (w worker) CorrectnessReport(_ context.Context, _ []byte) {}
