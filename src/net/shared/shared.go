package shared

// TODO: This package needs a better name

import (
	"sync"

	"github.com/gorilla/websocket"
)

var Client *websocket.Conn
var IsClientSocketClosed = false
var mu *sync.Mutex

func Send(data []byte) error {
	mu.Lock()
	defer mu.Unlock()
	return Client.WriteMessage(websocket.BinaryMessage, data)
}

func init() {
	mu = new(sync.Mutex)
}
