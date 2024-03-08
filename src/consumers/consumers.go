package consumers

import (
	"github.com/KenshiTech/unchained/constants/opcodes"
	"github.com/KenshiTech/unchained/crypto/bls"
	"github.com/KenshiTech/unchained/datasets"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/net/shared"
	"github.com/KenshiTech/unchained/plugins/logs"
	"github.com/KenshiTech/unchained/plugins/uniswap"
	"github.com/gorilla/websocket"
	"github.com/vmihailenco/msgpack/v5"
)

// TODO: These functions share a huge chunk of code
func ConsumePriceReport(message []byte) {
	var packet datasets.BroadcastPricePacket
	err := msgpack.Unmarshal(message, &packet)

	if err != nil {
		log.Logger.
			With("Error", err).
			Error("Unmarshal error")

		return
	}

	toHash, err := msgpack.Marshal(&packet.Info)

	if err != nil {
		log.Logger.
			With("Error", err).
			Error("Marshal error")

		return
	}

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
	var packet datasets.BroadcastEventPacket
	err := msgpack.Unmarshal(message, &packet)

	if err != nil {
		log.Logger.
			With("Error", err).
			Error("Unmarshal error")

		return
	}

	toHash, err := msgpack.Marshal(&packet.Info)

	if err != nil {
		log.Logger.
			With("Error", err).
			Error("Marshal error")

		return
	}

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

func StartConsumer() {
	shared.Client.WriteMessage(websocket.BinaryMessage, []byte{opcodes.RegisterConsumer})
}
