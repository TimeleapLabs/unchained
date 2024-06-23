package handler

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/transport/server/packet"

	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

func (h *consumer) PriceReport(ctx context.Context, message []byte) {
	packet := new(packet.BroadcastPricePacket).FromBytes(message)

	priceInfoHash, err := packet.Info.Bls()
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

	err = h.uniswap.RecordSignature(
		ctx,
		signature,
		packet.Signer,
		priceInfoHash,
		packet.Info,
		true,
		false,
	)
	if err != nil {
		return
	}
}

func (w worker) PriceReport(_ context.Context, _ []byte) {}
