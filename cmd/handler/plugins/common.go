package plugins

import "github.com/gorilla/websocket"

var conn *websocket.Conn
var closed = false

// Read is a function that reads messages from the websocket connection.
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

// CloseSocket is a function that closes the websocket connection.
func CloseSocket() {
	if conn != nil {
		closed = true
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
