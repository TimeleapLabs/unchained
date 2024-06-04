package handler

import (
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/middleware"
	"github.com/gorilla/websocket"
)

func (h Handler) CorrectnessRecord(conn *websocket.Conn, payload []byte) ([]byte, error) {
	err := middleware.IsConnectionAuthenticated(conn)
	if err != nil {
		return []byte{}, err
	}

	correctness := new(model.CorrectnessReportPacket).FromBytes(payload)

	signer, err := h.middleware.IsMessageValid(conn, correctness.Correctness.Sia().Bytes(), correctness.Signature)
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
