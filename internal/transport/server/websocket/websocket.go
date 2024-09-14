package websocket

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/transport/server/pubsub"
	"github.com/TimeleapLabs/unchained/internal/utils"

	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/handler"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/store"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// WithWebsocket is a function that starts a websocket server.
func WithWebsocket() func() {
	return func() {
		utils.Logger.Info("Starting a websocket server")

		versionedRoot := fmt.Sprintf("/%s", consts.ProtocolVersion)
		http.HandleFunc(versionedRoot, multiplexer)
	}
}

// multiplexer is a function that routes incoming messages to the appropriate handler.
func multiplexer(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(_ *http.Request) bool { return true } // remove this line in production

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.Logger.ErrorContext(context.TODO(), "Can't upgrade the HTTP connection", slog.Any("error", err))
		return
	}

	ctx, cancel := context.WithCancel(context.TODO())

	defer store.Signers.Delete(conn)
	defer store.Challenges.Delete(conn)
	defer cancel()

	for {
		messageType, payload, err := conn.ReadMessage()
		if err != nil {
			utils.Logger.ErrorContext(ctx, "Can't read message", slog.Any("error", err))

			err := conn.Close()
			if err != nil {
				utils.Logger.ErrorContext(ctx, "Can't close connection", slog.Any("error", err))
			}

			break
		}

		if len(payload) == 0 {
			continue
		}

		switch consts.OpCode(payload[0]) {
		case consts.OpCodeHello:
			utils.Logger.With("IP", conn.RemoteAddr().String()).Info("New Client Registered")
			result, err := handler.Hello(conn, payload[1:])
			if err != nil {
				handler.SendError(conn, messageType, consts.OpCodeError, err)
				continue
			}

			handler.SendMessage(conn, messageType, consts.OpCodeFeedback, "conf.ok")
			handler.Send(conn, messageType, consts.OpCodeKoskChallenge, result)
		case consts.OpCodePriceReport:
			result, err := handler.PriceReport(conn, payload[1:])
			if err != nil {
				handler.SendError(conn, messageType, consts.OpCodeError, err)
				continue
			}

			pubsub.Publish(consts.ChannelPriceReport, consts.OpCodePriceReportBroadcast, result)
			handler.SendMessage(conn, messageType, consts.OpCodeFeedback, "signature.accepted")
		case consts.OpCodeEventLog:
			result, err := handler.EventLog(conn, payload[1:])
			if err != nil {
				handler.SendError(conn, messageType, consts.OpCodeError, err)
				continue
			}

			pubsub.Publish(consts.ChannelEventLog, consts.OpCodeEventLogBroadcast, result)
			handler.SendMessage(conn, messageType, consts.OpCodeFeedback, "signature.accepted")
		case consts.OpCodeCorrectnessReport:
			result, err := handler.CorrectnessRecord(conn, payload[1:])
			if err != nil {
				handler.SendError(conn, messageType, consts.OpCodeError, err)
				continue
			}

			pubsub.Publish(consts.ChannelCorrectnessReport, consts.OpCodeCorrectnessReportBroadcast, result)
			handler.SendMessage(conn, messageType, consts.OpCodeFeedback, "signature.accepted")
		case consts.OpCodeKoskResult:
			err := handler.Kosk(conn, payload[1:])
			if err != nil {
				handler.SendError(conn, messageType, consts.OpCodeError, err)
				continue
			}

			handler.SendMessage(conn, messageType, consts.OpCodeFeedback, "kosk.ok")
		case consts.OpCodeRegisterConsumer:
			utils.Logger.
				With("IP", conn.RemoteAddr().String()).
				With("Channel", string(payload[1:])).
				Info("New Consumer registered")

			topic := string(payload[1:])
			go handler.BroadcastListener(ctx, conn, topic, pubsub.Subscribe(topic))
		default:
			handler.SendError(conn, messageType, consts.OpCodeError, consts.ErrNotSupportedInstruction)
		}
	}
}
