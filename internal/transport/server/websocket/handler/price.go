package handler

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/transport/server/packet"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/middleware"
	"github.com/gorilla/websocket"
)

// PriceReport check signature of message and return price info.
func PriceReport(conn *websocket.Conn, payload []byte) ([]byte, error) {
	err := middleware.IsConnectionAuthenticated(conn)
	if err != nil {
		return []byte{}, err
	}

	priceReport := new(packet.PriceReportPacket).FromBytes(payload)
	priceInfoHash, err := priceReport.PriceInfo.Bls()
	if err != nil {
		return []byte{}, consts.ErrInternalError
	}

	signer, err := middleware.IsMessageValid(conn, priceInfoHash, priceReport.Signature)
	if err != nil {
		return []byte{}, err
	}

	priceInfo := packet.BroadcastPricePacket{
		Info:      priceReport.PriceInfo,
		Signature: priceReport.Signature,
		Signer:    signer,
	}

	return priceInfo.Sia().Bytes(), nil
}
