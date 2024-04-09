package conn

import (
	"fmt"
	"sync"
	"time"

	"github.com/KenshiTech/unchained/internal/config"
	"github.com/KenshiTech/unchained/internal/constants"
	"github.com/KenshiTech/unchained/internal/constants/opcodes"
	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/log"

	"github.com/gorilla/websocket"
)

var conn *websocket.Conn
var IsClosed bool
var mu = new(sync.Mutex)

func Start() {
	var err error

	log.Logger.
		With("URL", fmt.Sprintf("%s/%s", config.App.Network.BrokerURI, constants.ProtocolVersion)).
		Info("Connecting to the broker")

	conn, _, err = websocket.DefaultDialer.Dial(
		fmt.Sprintf("%s/%s", config.App.Network.BrokerURI, constants.ProtocolVersion), nil,
	)
	if err != nil {
		log.Logger.
			With("URI", fmt.Sprintf("%s/%s", config.App.Network.BrokerURI, constants.ProtocolVersion)).
			With("Error", err).
			Error("can't connect to broker")
		panic(err)
	}

	Send(opcodes.Hello, bls.ClientSigner.Sia().Content)
}

func Reconnect(err error) {
	IsClosed = true
	hello := bls.ClientSigner.Sia().Content

	if websocket.IsUnexpectedCloseError(err) {
		for i := 1; i < 6; i++ {
			time.Sleep(time.Duration(i) * 3 * time.Second)

			log.Logger.
				With("URL", fmt.Sprintf("%s/%s", config.App.Network.BrokerURI, constants.ProtocolVersion)).
				With("Retry", i).
				Info("Reconnecting to broker")

			conn, _, err = websocket.DefaultDialer.Dial(config.App.Network.BrokerURI, nil)
			if err != nil {
				log.Logger.
					With("URI", fmt.Sprintf("%s/%s", config.App.Network.BrokerURI, constants.ProtocolVersion)).
					With("Error", err).
					Error("can't connect to broker")
			} else {
				IsClosed = false
				Send(opcodes.Hello, hello)
				log.Logger.Info("Connection with broker recovered")
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
			log.Logger.
				With("Error", err).
				Error("Can't sent close packet")
		}

		IsClosed = false
		err = conn.Close()
		if err != nil {
			log.Logger.
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
				log.Logger.
					With("Error", err).
					Error("Read error")

				Reconnect(err)
				if IsClosed {
					return
				}

				continue
			}

			if payload[0] == byte(opcodes.Error) {
				log.Logger.
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

func Send(opCode opcodes.OpCode, payload []byte) {
	err := SendRaw(
		append([]byte{byte(opCode)}, payload...),
	)
	if err != nil {
		log.Logger.Error("Can't send packet: %v", err)
	}
}

func SendMessage(opCode opcodes.OpCode, message string) {
	Send(opCode, []byte(message))
}
