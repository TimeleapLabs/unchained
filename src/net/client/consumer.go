package client

import (
	"github.com/KenshiTech/unchained/datasets"
	"github.com/KenshiTech/unchained/log"
	"github.com/gorilla/websocket"
	"github.com/vmihailenco/msgpack/v5"
)

func ConsumePriceReport(message []byte) {
	var packet datasets.BroadcastPricePacket
	err := msgpack.Unmarshal(message[1:], &packet)
	if err != nil {
		panic(err)
	}
	log.Logger.
		With("Validators", len(packet.Signers)).
		With("Asset", packet.Info.Asset).
		With("Block", packet.Info.Block).
		With("Price", packet.Info.Price.Text(10)).
		Info("Attestation")
}

func ConsumeEventLog(message []byte) {
	var packet datasets.BroadcastEventPacket
	err := msgpack.Unmarshal(message[1:], &packet)
	if err != nil {
		panic(err)
	}
	log.Logger.
		With("Validators", len(packet.Signers)).
		With("Chain", packet.Info.Chain).
		With("Address", packet.Info.Address).
		With("Event", packet.Info.Event).
		Info("Attestation")
}

func StartConsumer() {
	Client.WriteMessage(websocket.BinaryMessage, []byte{6})
}
