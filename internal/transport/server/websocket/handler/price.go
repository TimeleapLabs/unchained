package handler

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/transport/server/pubsub"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/middleware"
	"github.com/gorilla/websocket"
)

// PriceReport check signature of message and return price info.
func PriceReport(conn *websocket.Conn, payload []byte) {
	err := middleware.IsConnectionAuthenticated(conn)
	if err != nil {
		SendError(conn, consts.OpCodeError, err)
		return
	}

	priceReport := new(model.PriceReportPacket).FromBytes(payload)
	priceInfoHash, err := priceReport.PriceInfo.Bls()
	if err != nil {
		SendError(conn, consts.OpCodeError, err)
		return
	}

	signer, err := middleware.IsMessageValid(conn, priceInfoHash, priceReport.Signature)
	if err != nil {
		SendError(conn, consts.OpCodeError, err)
		return
	}

	priceInfo := model.BroadcastPricePacket{
		Info:      priceReport.PriceInfo,
		Signature: priceReport.Signature,
		Signer:    signer,
	}

	pubsub.Publish(consts.ChannelPriceReport, consts.OpCodePriceReportBroadcast, priceInfo.Sia().Bytes())
	SendMessage(conn, consts.OpCodeFeedback, "signature.accepted")
}
