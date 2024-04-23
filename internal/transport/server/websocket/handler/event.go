package handler

import (
	"github.com/KenshiTech/unchained/internal/consts"
	"github.com/KenshiTech/unchained/internal/model"
	"github.com/KenshiTech/unchained/internal/transport/server/websocket/middleware"
	"github.com/gorilla/websocket"
)

func EventLog(conn *websocket.Conn, payload []byte) ([]byte, error) {
	err := middleware.IsConnectionAuthenticated(conn)
	if err != nil {
		return []byte{}, err
	}

	priceReport := new(model.EventLogReportPacket).FromBytes(payload)
	priceInfoHash, err := priceReport.EventLog.Bls()
	if err != nil {
		return []byte{}, consts.ErrInternalError
	}

	signer, err := middleware.IsMessageValid(conn, priceInfoHash, priceReport.Signature)
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
