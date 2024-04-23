package handler

import (
	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/datasets"
	"github.com/TimeleapLabs/unchained/internal/log"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

func (h *consumer) EventLog(message []byte) {
	packet := new(datasets.BroadcastEventPacket).DeSia(&sia.Sia{Content: message})
	toHash := packet.Info.Sia().Content
	hash, err := bls.Hash(toHash)

	if err != nil {
		log.Logger.
			With("Error", err).
			Error("Hash error")

		return
	}

	signature, err := bls.RecoverSignature(packet.Signature)

	if err != nil {
		log.Logger.
			With("Error", err).
			Error("Failed to recover packet signature")

		return
	}

	h.evmlog.RecordSignature(
		signature,
		packet.Signer,
		hash,
		packet.Info,
		true,
		false,
	)
}

func (w worker) EventLog(_ []byte) {}
