package handler

import (
	"github.com/KenshiTech/unchained/internal/consts"
	"github.com/KenshiTech/unchained/internal/model"
	"github.com/KenshiTech/unchained/internal/transport/server/websocket/middleware"
	"github.com/gorilla/websocket"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

func CorrectnessRecord(conn *websocket.Conn, payload []byte) ([]byte, error) {
	err := middleware.IsConnectionAuthenticated(conn)
	if err != nil {
		return []byte{}, err
	}

	correctness := new(model.CorrectnessReportPacket).DeSia(&sia.Sia{Content: payload})
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

	return broadcastPacket.Sia().Content, nil
}
