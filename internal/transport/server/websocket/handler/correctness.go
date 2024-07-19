package handler

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/transport/server/pubsub"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/middleware"
	"github.com/gorilla/websocket"
)

func CorrectnessRecord(conn *websocket.Conn, payload []byte) {
	err := middleware.IsConnectionAuthenticated(conn)
	if err != nil {
		SendError(conn, consts.OpCodeError, err)
		return
	}

	correctness := new(model.CorrectnessReportPacket).FromBytes(payload)
	correctnessHash, err := correctness.Correctness.Bls()
	if err != nil {
		SendError(conn, consts.OpCodeError, consts.ErrInternalError)
		return
	}

	signer, err := middleware.IsMessageValid(conn, correctnessHash, correctness.Signature)
	if err != nil {
		SendError(conn, consts.OpCodeError, err)
		return
	}

	broadcastPacket := model.BroadcastCorrectnessPacket{
		Info:      correctness.Correctness,
		Signature: correctness.Signature,
		Signer:    signer,
	}

	pubsub.Publish(consts.ChannelCorrectnessReport, consts.OpCodeCorrectnessReportBroadcast, broadcastPacket.Sia().Bytes())
	SendMessage(conn, consts.OpCodeFeedback, "signature.accepted")
}
