package handler

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/middleware"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/store"
	"github.com/gorilla/websocket"
)

func Kosk(conn *websocket.Conn, payload []byte) error {
	challenge := new(model.ChallengePacket).FromBytes(payload)

	_, err := middleware.IsMessageValid(conn, challenge.Random[:], challenge.Signature)
	if err != nil {
		return consts.ErrInvalidKosk
	}

	challenge.Passed = true
	store.Challenges.Store(conn, *challenge)

	return nil
}
