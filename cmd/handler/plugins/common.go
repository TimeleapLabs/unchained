package plugins

import "github.com/gorilla/websocket"

var conn *websocket.Conn
var closed = false

func Read() <-chan []byte {
	out := make(chan []byte)

	go func() {
		for {
			_, payload, err := conn.ReadMessage()
			if err != nil {
				if !closed {
					panic(err)
				}
			}

			out <- payload
		}
	}()

	return out
}

func CloseSocket() {
	if conn != nil {
		closed = true
		conn.WriteMessage(
			websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		conn.Close()
	}
}
