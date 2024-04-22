package handler

import (
	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/model"
	"github.com/KenshiTech/unchained/internal/utils"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

func (h *consumer) CorrectnessReport(message []byte) {
	packet := new(model.BroadcastCorrectnessPacket).DeSia(&sia.Sia{Content: message})

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

func (w worker) CorrectnessReport(_ []byte) {}
