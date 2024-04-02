package consumer

import (
	"github.com/KenshiTech/unchained/src/log"
	"github.com/KenshiTech/unchained/src/net/repository"
	"github.com/gorilla/websocket"
)

func Broadcast(message []byte) {
	repository.Consumers.Range(func(consumer *websocket.Conn, _ bool) bool {
		mu, ok := repository.BroadcastMutex.Load(consumer)
		if ok {
			mu.Lock()
			defer mu.Unlock()
			err := consumer.WriteMessage(websocket.BinaryMessage, message)
			if err != nil {
				log.Logger.Error(err.Error())
			}
		}
		return true
	})
}
