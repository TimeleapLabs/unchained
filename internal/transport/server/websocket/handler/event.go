package handler

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/transport/server/packet"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/middleware"
	"github.com/gorilla/websocket"
)

// EventLog handles the event log packet from the client.
func EventLog(conn *websocket.Conn, payload []byte) ([]byte, error) {
	err := middleware.IsConnectionAuthenticated(conn)
	if err != nil {
		return []byte{}, err
	}

	priceReport := new(packet.EventLogReportPacket).FromBytes(payload)
	priceInfoHash, err := priceReport.EventLog.Bls()
	if err != nil {
		return []byte{}, consts.ErrInternalError
	}

	signer, err := middleware.IsMessageValid(conn, priceInfoHash, priceReport.Signature)
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
