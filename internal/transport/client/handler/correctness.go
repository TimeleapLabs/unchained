package handler

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/model"
)

func (h *postgresConsumer) CorrectnessReport(ctx context.Context, message []byte) {
	packet := new(model.BroadcastCorrectnessPacket).FromBytes(message)

	correctnessHash, err := packet.Info.Bls()
	if err != nil {
		return
	}

	err = h.correctness.RecordSignature(
		ctx,
		packet.Signature[:],
		packet.Signer,
		correctnessHash,
		packet.Info,
		true,
	)
	if err != nil {
		return
	}
}

func (h *schnorrConsumer) CorrectnessReport(_ context.Context, _ []byte) {}

func (w worker) CorrectnessReport(_ context.Context, _ []byte) {}
