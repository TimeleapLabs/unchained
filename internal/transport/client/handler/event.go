package handler

import (
	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/model"
	"github.com/KenshiTech/unchained/internal/utils"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

func (h *consumer) EventLog(message []byte) {
	packet := new(model.BroadcastEventPacket).DeSia(&sia.Sia{Content: message})

	eventLogHash, err := packet.Info.Bls()
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

	err = h.evmlog.RecordSignature(
		signature,
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

func (w worker) EventLog(_ []byte) {}
