package shared

// TODO: This package needs a better name

import (
	"sync"

	"github.com/KenshiTech/unchained/src/constants/opcodes"
	"github.com/KenshiTech/unchained/src/log"

	"github.com/gorilla/websocket"
)

var Client *websocket.Conn
var IsClientSocketClosed = false
var mu *sync.Mutex

func SendRaw(data []byte) error {
	mu.Lock()
	defer mu.Unlock()
	return Client.WriteMessage(websocket.BinaryMessage, data)
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

func Close() {

}

func init() {
	mu = new(sync.Mutex)
}
