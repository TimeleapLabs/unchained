package conn

import (
	"fmt"
	"time"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/queue"
	"github.com/TimeleapLabs/unchained/internal/utils"

	"github.com/TimeleapLabs/unchained/internal/config"

	"github.com/gorilla/websocket"
)

var conn *websocket.Conn
var writer *queue.WebSocketWriter
var IsClosed bool

// Start function create a new websocket connection to the broker.
func Start() {
	var err error

	var brokerURI = config.App.Network.Broker.URI

	utils.Logger.
		With("URL", fmt.Sprintf("%s/%s", brokerURI, consts.ProtocolVersion)).
		Info("Connecting to the broker")

	conn, _, err = websocket.DefaultDialer.Dial(
		fmt.Sprintf("%s/%s", brokerURI, consts.ProtocolVersion), nil,
	)
	if err != nil {
		utils.Logger.
			With("URI", fmt.Sprintf("%s/%s", brokerURI, consts.ProtocolVersion)).
			With("Error", err).
			Error("can't connect to broker")
		panic(err)
	}

	writer = queue.NewWebSocketWriter(conn, 1024)
}

func Reconnect(err error) {
	if websocket.IsUnexpectedCloseError(err) {
		Close()

		for i := 1; i < 6; i++ {
			time.Sleep(time.Duration(i) * 3 * time.Second)

			utils.Logger.
				With("URL", fmt.Sprintf("%s/%s", config.App.Network.Broker.URI, consts.ProtocolVersion)).
				With("Retry", i).
				Info("Reconnecting to broker")

			conn, _, err = websocket.DefaultDialer.Dial(
				fmt.Sprintf("%s/%s", config.App.Network.Broker.URI, consts.ProtocolVersion),
				nil,
			)
			if err != nil {
				utils.Logger.
					With("URI", fmt.Sprintf("%s/%s", config.App.Network.Broker.URI, consts.ProtocolVersion)).
					With("Error", err).
					Error("Cannot reconnect to broker")
			} else {
				IsClosed = false
				utils.Logger.Info("Connection with broker recovered")
				return
			}
		}

		panic("Cannot Connect to broker")
	}
}

// Close function gracefully disconnect from the broker.
func Close() {
	if conn != nil && config.App.Network.Broker.URI != "" {
		err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			utils.Logger.
				With("Error", err).
				Error("Cannot sent close packet")
		}

		IsClosed = true
		err = conn.Close()
		if err != nil {
			utils.Logger.
				With("Error", err).
				Error("Connection closed")
		}
	}
}

// Read function consume from the broker's messages and push them into a channel.
func Read() <-chan []byte {
	out := make(chan []byte)

	go func() {
		for {
			_, payload, err := conn.ReadMessage()
			if err != nil {
				utils.Logger.
					With("Error", err).
					Error("Read error")

				Reconnect(err)
				if IsClosed {
					break
				}

				continue
			}

			if payload[0] == byte(consts.OpCodeError) {
				utils.Logger.
					With("Error", string(payload[1:])).
					Error("Incoming error")

				Reconnect(err)
				if IsClosed {
					break
				}

				continue
			}

			out <- payload
		}
	}()

	return out
}

func SendSigned(opCode consts.OpCode, payload []byte) {
	writer.SendSigned(opCode, payload)
}

func Get() *websocket.Conn {
	return conn
}
