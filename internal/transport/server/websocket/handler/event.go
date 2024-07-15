package handler

import (
	"github.com/TimeleapLabs/unchained/internal/transport/server/packet"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/middleware"
	"github.com/gorilla/websocket"
)

func EventLog(conn *websocket.Conn, payload []byte) ([]byte, error) {
	err := middleware.IsConnectionAuthenticated(conn)
	if err != nil {
		return []byte{}, err
	}

	priceReport := new(packet.EventLogReportPacket).FromBytes(payload)

	signer, err := middleware.IsMessageValid(conn, *priceReport.EventLog.Bls(), priceReport.Signature)
	if err != nil {
		return []byte{}, err
	}

	broadcastPacket := packet.BroadcastEventPacket{
		Info:      priceReport.EventLog,
		Signature: priceReport.Signature,
		Signer:    signer,
	}

	return broadcastPacket.Sia().Bytes(), nil
}
