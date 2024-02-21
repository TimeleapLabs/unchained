package consumer

import (
	"github.com/KenshiTech/unchained/net/repository"
	"github.com/gorilla/websocket"
)

func Broadcast(message []byte) {
	repository.Consumers.Range(func(consumer *websocket.Conn, _ bool) bool {
		consumer.WriteMessage(websocket.BinaryMessage, message)
		return true
	})
}
