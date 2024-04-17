package conn

import (
	"fmt"
	"sync"
	"time"

	"github.com/KenshiTech/unchained/internal/utils"

	"github.com/KenshiTech/unchained/internal/consts"

	"github.com/KenshiTech/unchained/internal/crypto"

	"github.com/KenshiTech/unchained/internal/config"
	"github.com/gorilla/websocket"
)

var conn *websocket.Conn
var IsClosed bool
var mu = new(sync.Mutex)

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

	Send(consts.OpCodeHello, crypto.Identity.ExportBlsSigner().Sia().Content)
}

func Reconnect(err error) {
	IsClosed = true
	hello := crypto.Identity.ExportBlsSigner().Sia().Content

	if websocket.IsUnexpectedCloseError(err) {
		for i := 1; i < 6; i++ {
			time.Sleep(time.Duration(i) * 3 * time.Second)

			utils.Logger.
				With("URL", fmt.Sprintf("%s/%s", config.App.Network.BrokerURI, consts.ProtocolVersion)).
				With("Retry", i).
				Info("Reconnecting to broker")

			conn, _, err = websocket.DefaultDialer.Dial(config.App.Network.BrokerURI, nil)
			if err != nil {
				utils.Logger.
					With("URI", fmt.Sprintf("%s/%s", config.App.Network.BrokerURI, consts.ProtocolVersion)).
					With("Error", err).
					Error("can't connect to broker")
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

func Close() {
	if conn != nil && config.App.Network.BrokerURI != "" {
		err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			utils.Logger.
				With("Error", err).
				Error("Can't sent close packet")
		}

		IsClosed = false
		err = conn.Close()
		if err != nil {
			utils.Logger.
				With("Error", err).
				Error("Connection closed")
		}
	}
}

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
					return
				}

				continue
			}

			if payload[0] == byte(consts.OpCodeError) {
				utils.Logger.
					With("Error", string(payload[1:])).
					Error("Incoming error")

				Reconnect(err)
				if IsClosed {
					return
				}

				continue
			}

			out <- payload
		}
	}()

	return out
}

func SendRaw(data []byte) error {
	mu.Lock()
	defer mu.Unlock()
	return conn.WriteMessage(websocket.BinaryMessage, data)
}

func Send(opCode consts.OpCode, payload []byte) {
	err := SendRaw(
		append([]byte{byte(opCode)}, payload...),
	)
	if err != nil {
		utils.Logger.Error("Can't send packet: %v", err)
	}
}

func SendMessage(opCode consts.OpCode, message string) {
	Send(opCode, []byte(message))
}
