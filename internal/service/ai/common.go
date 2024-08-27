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
		err := conn.WriteMessage(
			websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			return
		}
		err = conn.Close()
		if err != nil {
			return
		}
	}
}
