package consumers

import (
	"github.com/KenshiTech/unchained/constants/opcodes"
	"github.com/KenshiTech/unchained/crypto/bls"
	"github.com/KenshiTech/unchained/datasets"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/net/shared"
	"github.com/KenshiTech/unchained/plugins/correctness"
	"github.com/KenshiTech/unchained/plugins/logs"
	"github.com/KenshiTech/unchained/plugins/uniswap"
	"github.com/gorilla/websocket"

	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

// TODO: These functions share a huge chunk of code
func ConsumePriceReport(message []byte) {
	packet := new(datasets.BroadcastPricePacket).DeSia(&sia.Sia{Content: message})
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

	uniswap.RecordSignature(
		signature,
		packet.Signer,
		hash,
		packet.Info,
		true,
		false,
	)
}

func ConsumeEventLog(message []byte) {
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

	logs.RecordSignature(
		signature,
		packet.Signer,
		hash,
		packet.Info,
		true,
		false,
	)
}

func ConsumeCorrectnessReport(message []byte) {
	packet := new(datasets.BroadcastCorrectnessPacket).DeSia(&sia.Sia{Content: message})
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

	correctness.RecordSignature(
		signature,
		packet.Signer,
		hash,
		packet.Info,
		true,
	)
}

func StartConsumer() {
	err := shared.Client.WriteMessage(websocket.BinaryMessage, []byte{byte(opcodes.RegisterConsumer)})
	if err != nil {
		panic(err)
	}
}
