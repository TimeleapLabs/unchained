package middleware

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/store"
	"github.com/gorilla/websocket"
)

func IsConnectionAuthenticated(conn *websocket.Conn) error {
	challenge, ok := store.Challenges.Load(conn)
	if !ok || !challenge.Passed {
		return consts.ErrMissingKosk
	}

	return nil
}
