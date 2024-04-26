package handler

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/middleware"
	"github.com/gorilla/websocket"
)

func CorrectnessRecord(conn *websocket.Conn, payload []byte) ([]byte, error) {
	err := middleware.IsConnectionAuthenticated(conn)
	if err != nil {
		return []byte{}, err
	}

	correctness := new(model.CorrectnessReportPacket).FromBytes(payload)
	correctnessHash, err := correctness.Correctness.Bls()
	if err != nil {
		return []byte{}, consts.ErrInternalError
	}

	signer, err := middleware.IsMessageValid(conn, correctnessHash, correctness.Signature)
	if err != nil {
		return []byte{}, err
	}

	broadcastPacket := model.BroadcastCorrectnessPacket{
		Info:      correctness.Correctness,
		Signature: correctness.Signature,
		Signer:    signer,
	}

	return broadcastPacket.Sia().Bytes(), nil
}
