package conn

import (
	"fmt"
	"sync"
	"time"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/utils"

	"github.com/TimeleapLabs/unchained/internal/crypto"

	"github.com/TimeleapLabs/unchained/internal/config"

	"github.com/gorilla/websocket"
)

var conn *websocket.Conn
var IsClosed bool
var mu = new(sync.Mutex)

// Start function create a new websocket connection to the broker.
func Start() {
	var err error

	utils.Logger.
		With("URL", fmt.Sprintf("%s/%s", config.App.Network.BrokerURI, consts.ProtocolVersion)).
		Info("Connecting to the broker")

	conn, _, err = websocket.DefaultDialer.Dial(
		fmt.Sprintf("%s/%s", config.App.Network.BrokerURI, consts.ProtocolVersion), nil,
	)
	if err != nil {
		utils.Logger.
			With("URI", fmt.Sprintf("%s/%s", config.App.Network.BrokerURI, consts.ProtocolVersion)).
			With("Error", err).
			Error("can't connect to broker")
		panic(err)
	}

	Send(consts.OpCodeHello, crypto.Identity.ExportEvmSigner().Sia().Bytes())
}

func Reconnect(err error) {
	if websocket.IsUnexpectedCloseError(err) {
		Close()
		hello := crypto.Identity.ExportEvmSigner().Sia().Bytes()

		for i := 1; i < 6; i++ {
			time.Sleep(time.Duration(i) * 3 * time.Second)

			utils.Logger.
				With("URL", fmt.Sprintf("%s/%s", config.App.Network.BrokerURI, consts.ProtocolVersion)).
				With("Retry", i).
				Info("Reconnecting to broker")

			conn, _, err = websocket.DefaultDialer.Dial(
				fmt.Sprintf("%s/%s", config.App.Network.BrokerURI, consts.ProtocolVersion),
				nil,
			)
			if err != nil {
				utils.Logger.
					With("URI", fmt.Sprintf("%s/%s", config.App.Network.BrokerURI, consts.ProtocolVersion)).
					With("Error", err).
					Error("Can't reconnect to broker")
			} else {
				IsClosed = false
				Send(consts.OpCodeHello, hello)
				utils.Logger.Info("Connection with broker recovered")
				return
			}
		}

		panic("Cant Connect to broker")
	}
}

// Close function gracefully disconnect from the broker.
func Close() {
	if conn != nil && config.App.Network.BrokerURI != "" {
		err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			utils.Logger.
				With("Error", err).
				Error("Can't sent close packet")
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

// SendRaw function sends byte array data to the broker.
func SendRaw(data []byte) error {
	mu.Lock()
	defer mu.Unlock()
	return conn.WriteMessage(websocket.BinaryMessage, data)
}

// Send function sends a message with specific opCode to the broker.
func Send(opCode consts.OpCode, payload []byte) {
	err := SendRaw(
		append([]byte{byte(opCode)}, payload...),
	)
	if err != nil {
		utils.Logger.Error("Can't send packet: %v", err)
	}
}
