package handler

import (
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/middleware"
	"github.com/gorilla/websocket"
)

func (h Handler) EventLog(conn *websocket.Conn, payload []byte) ([]byte, error) {
	err := middleware.IsConnectionAuthenticated(conn)
	if err != nil {
		return []byte{}, err
	}

	priceReport := new(model.EventLogReportPacket).FromBytes(payload)

	signer, err := h.middleware.IsMessageValid(conn, priceReport.EventLog.Sia().Bytes(), priceReport.Signature)
	if err != nil {
		return []byte{}, err
	}

	broadcastPacket := model.BroadcastEventPacket{
		Info:      priceReport.EventLog,
		Signature: priceReport.Signature,
		Signer:    signer,
	}

	return broadcastPacket.Sia().Bytes(), nil
}
