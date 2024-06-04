package handler

import (
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/middleware"
	"github.com/gorilla/websocket"
)

// PriceReport check signature of message and return price info.
func (h Handler) PriceReport(conn *websocket.Conn, payload []byte) ([]byte, error) {
	err := middleware.IsConnectionAuthenticated(conn)
	if err != nil {
		return []byte{}, err
	}

	priceReport := new(model.PriceReportPacket).FromBytes(payload)

	signer, err := h.middleware.IsMessageValid(conn, priceReport.PriceInfo.Sia().Bytes(), priceReport.Signature)
	if err != nil {
		return []byte{}, err
	}

	priceInfo := model.BroadcastPricePacket{
		Info:      priceReport.PriceInfo,
		Signature: priceReport.Signature,
		Signer:    signer,
	}

	return priceInfo.Sia().Bytes(), nil
}
