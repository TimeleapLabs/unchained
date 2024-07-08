package ai

import (
	"github.com/gorilla/websocket"
)

func Read(conn *websocket.Conn, closed *bool) <-chan []byte {
	out := make(chan []byte)

	go func() {
		defer close(out)
		for {
			_, payload, err := conn.ReadMessage()
			if err != nil {
				if !*closed {
					panic(err)
				}
				return
			}
			out <- payload
		}
	}()

	return out
}

func CloseSocket(conn *websocket.Conn, closed *bool) {
	if conn != nil {
		*closed = true
		conn.WriteMessage(
			websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		conn.Close()
	}
}
