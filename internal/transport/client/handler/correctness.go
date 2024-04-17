package handler

import (
	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/model"
	"github.com/KenshiTech/unchained/internal/utils"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

func (h *consumer) CorrectnessReport(message []byte) {
	packet := new(model.BroadcastCorrectnessPacket).DeSia(&sia.Sia{Content: message})
	toHash := packet.Info.Sia().Content
	hash, err := bls.Hash(toHash)

	if err != nil {
		utils.Logger.
			With("Error", err).
			Error("Hash error")

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
		hash,
		packet.Info,
		true,
	)

	if err != nil {
		utils.Logger.
			With("Error", err).
			Error("Failed to record signature")
	}
}

func (w worker) CorrectnessReport(_ []byte) {}
