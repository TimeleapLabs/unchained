package client

import (
	"github.com/KenshiTech/unchained/datasets"
	"github.com/KenshiTech/unchained/log"
	"github.com/gorilla/websocket"
	"github.com/vmihailenco/msgpack/v5"
)

func Consume(message []byte) {
	var packet datasets.BroadcastPacket
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

func StartConsumer() {
	Client.WriteMessage(websocket.BinaryMessage, []byte{6})
}
