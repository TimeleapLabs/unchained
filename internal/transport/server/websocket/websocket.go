package websocket

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/transport/server/pubsub"
	"github.com/TimeleapLabs/unchained/internal/utils"

	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/handler"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/queue"
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
		utils.Logger.Error("Cannot upgrade the HTTP connection", slog.Any("error", err))
		return
	}

	ctx, cancel := context.WithCancel(context.TODO())

	defer store.Signers.Delete(conn)
	defer cancel()

	// register a close handler to stop the for loop
	conn.SetCloseHandler(func(code int, text string) error {
		utils.Logger.With("Code", code).With("Text", text).Info("Connection closed")
		cancel()
		return nil
	})

	writer := queue.NewWebSocketWriter(conn, 10)

	for {
		// stop the loop if the context is done
		select {
		case <-ctx.Done():
			return
		default:
		}

		_, payload, err := conn.ReadMessage()
		if err != nil {
			var closeError *websocket.CloseError
			errIsClosedError := errors.As(err, &closeError) &&
				(closeError.Code == websocket.CloseNoStatusReceived || closeError.Code == websocket.CloseNormalClosure)

			if !errIsClosedError {
				utils.Logger.With("Err", err).ErrorContext(ctx, "Cannot read message")
			}

			err := conn.Close()
			if err != nil {
				utils.Logger.With("Err", err).ErrorContext(ctx, "Cannot close connection")
			}

			break
		}

		if len(payload) == 0 {
			continue
		}

		switch consts.OpCode(payload[0]) {
		case consts.OpCodeHello:
			utils.Logger.With("IP", conn.RemoteAddr().String()).Info("New Client Registered")
			err := handler.Hello(conn, payload[1:])
			if err != nil {
				writer.SendError(consts.OpCodeError, err)
				continue
			}

			writer.SendMessage(consts.OpCodeFeedback, "conf.ok")
		case consts.OpCodeAttestation:
			result, err := handler.AttestationRecord(conn, payload[1:])
			if err != nil {
				writer.SendError(consts.OpCodeError, err)
				continue
			}

			pubsub.Publish(consts.ChannelAttestation, consts.OpCodeAttestationBroadcast, result)
			writer.SendMessage(consts.OpCodeFeedback, "signature.accepted")
		case consts.OpCodeRegisterConsumer:
			utils.Logger.
				With("IP", conn.RemoteAddr().String()).
				With("Channel", string(payload[1:])).
				Info("New Consumer registered")

			topic := string(payload[1:])
			go handler.BroadcastListener(ctx, conn, topic, pubsub.Subscribe(topic))
		case consts.OpCodeRegisterWorker:
			go handler.RegisterWorker(ctx, conn, payload[1:])
		case consts.OpCodeWorkerOverload:
			go handler.WorkerOverload(ctx, conn, payload[1:])
		case consts.OpCodeRPCRequest:
			go handler.CallFunction(ctx, writer, payload[1:])
		case consts.OpCodeRPCResponse:
			go handler.ResponseFunction(ctx, writer, payload[1:])
		default:
			utils.Logger.With("OpCode", payload[0]).Error("Unsupported OpCode")
			writer.SendError(consts.OpCodeError, consts.ErrNotSupportedInstruction)
		}
	}
}
