package websocket

import (
	"context"
	"fmt"
	"net/http"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/transport/server/pubsub"
	"github.com/TimeleapLabs/unchained/internal/utils"

	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/handler"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/store"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func WithWebsocket() func() {
	return func() {
		utils.Logger.Info("Starting a websocket server")

		versionedRoot := fmt.Sprintf("/%s", consts.ProtocolVersion)
		http.HandleFunc(versionedRoot, multiplexer)
	}
}

func multiplexer(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // remove this line in production

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.Logger.Error("Can't upgrade the HTTP connection: %v", err)
		return
	}

	ctx, cancel := context.WithCancel(context.TODO())

	defer store.Signers.Delete(conn)
	defer store.Challenges.Delete(conn)
	defer cancel()

	for {
		_, payload, err := conn.ReadMessage()
		if err != nil {
			utils.Logger.Error("Can't read message: %v", err)

			err := conn.Close()
			if err != nil {
				utils.Logger.Error("Can't close connection: %v", err)
			}

			break
		}

		if len(payload) == 0 {
			continue
		}

		switch consts.OpCode(payload[0]) {
		case consts.OpCodeHello:
			handler.Hello(conn, payload[1:])
		case consts.OpCodePriceReport:
			handler.PriceReport(conn, payload[1:])
		case consts.OpCodeEventLog:
			handler.EventLog(conn, payload[1:])
		case consts.OpCodeCorrectnessReport:
			handler.CorrectnessRecord(conn, payload[1:])
		case consts.OpCodeKoskResult:
			handler.Kosk(conn, payload[1:])
		case consts.OpCodeRegisterConsumer:
			utils.Logger.
				With("IP", conn.RemoteAddr().String()).
				With("Channel", string(payload[1:])).
				Info("New Consumer registered")

			go handler.BroadcastListener(ctx, conn, pubsub.Subscribe(string(payload[1:])))
		case consts.OpCodeRegisterRpcFunction:
			handler.RegisterRpcFunction(ctx, conn, payload[1:])
		case consts.OpCodeRpcRequest:
			handler.CallFunction(ctx, conn, payload[1:])
		case consts.OpCodeRpcResponse:
			handler.ResponseFunction(ctx, conn, payload[1:])
		default:
			handler.SendError(conn, consts.OpCodeError, consts.ErrNotSupportedInstruction)
		}
	}
}
