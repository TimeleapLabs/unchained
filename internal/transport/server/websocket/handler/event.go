package handler

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/transport/server/pubsub"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/middleware"
	"github.com/gorilla/websocket"
)

// EventLog handles the event log packet from the client.
func EventLog(conn *websocket.Conn, payload []byte) {
	err := middleware.IsConnectionAuthenticated(conn)
	if err != nil {
		SendError(conn, consts.OpCodeError, err)
		return
	}

	priceReport := new(model.EventLogReportPacket).FromBytes(payload)
	priceInfoHash, err := priceReport.EventLog.Bls()
	if err != nil {
		SendError(conn, consts.OpCodeError, consts.ErrInternalError)
		return
	}

	signer, err := middleware.IsMessageValid(conn, priceInfoHash, priceReport.Signature)
	if err != nil {
		SendError(conn, consts.OpCodeError, err)
		return
	}

	broadcastPacket := model.BroadcastEventPacket{
		Info:      priceReport.EventLog,
		Signature: priceReport.Signature,
		Signer:    signer,
	}

	pubsub.Publish(consts.ChannelEventLog, consts.OpCodeEventLogBroadcast, broadcastPacket.Sia().Bytes())
	SendMessage(conn, consts.OpCodeFeedback, "signature.accepted")
}
