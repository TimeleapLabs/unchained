package handler

import (
	"github.com/TimeleapLabs/unchained/internal/transport/server/packet"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/middleware"
	"github.com/gorilla/websocket"
)

func CorrectnessRecord(conn *websocket.Conn, payload []byte) ([]byte, error) {
	err := middleware.IsConnectionAuthenticated(conn)
	if err != nil {
		return []byte{}, err
	}

	correctness := new(packet.CorrectnessReportPacket).FromBytes(payload)
	correctnessHash := correctness.Correctness.Bls()

	signer, err := middleware.IsMessageValid(conn, correctnessHash, correctness.Signature)
	if err != nil {
		return []byte{}, err
	}

	broadcastPacket := packet.BroadcastCorrectnessPacket{
		Info:      correctness.Correctness,
		Signature: correctness.Signature,
		Signer:    signer,
	}

	return broadcastPacket.Sia().Bytes(), nil
}
