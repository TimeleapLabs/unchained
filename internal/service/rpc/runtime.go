package rpc

import (
	"github.com/gorilla/websocket"
)

// Runtime is a type that holds the runtime of a function.
type Runtime string

const (
	Mock      Runtime = "Mock"
	WebSocket Runtime = "WebSocket"
)

func WithMockTask(name string) func(s *Worker) {
	return func(s *Worker) {
		s.functions[name] = meta{
			runtime: Mock,
		}
	}
}

func WithWebSocket(name string, url string) func(s *Worker) {
	return func(s *Worker) {
		meta := meta{
			runtime: WebSocket,
			path:    url,
		}

		// TODO: NEED A HANDLER TO HANDLE THE CONNECTION
		wsConn, httpResp, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			panic(err)
		}

		if httpResp.StatusCode != 101 {
			panic("Failed to establish websocket connection")
		}

		meta.conn = wsConn
		s.functions[name] = meta
	}
}
