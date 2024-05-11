package handler

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/model"
)

func (h *consumer) EventLog(ctx context.Context, message []byte) {
	packet := new(model.BroadcastEventPacket).FromBytes(message)

	eventLogHash, err := packet.Info.Bls()
	if err != nil {
		return
	}

	err = h.evmlog.RecordSignature(
		ctx,
		packet.Signature[:],
		packet.Signer,
		eventLogHash,
		packet.Info,
		true,
		false,
	)
	if err != nil {
		return
	}
}

func (w worker) EventLog(_ context.Context, _ []byte) {}
