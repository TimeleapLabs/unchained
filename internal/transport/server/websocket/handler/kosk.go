package handler

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/middleware"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/store"
	"github.com/gorilla/websocket"
)

// Kosk handler check the result of signer challenge and store it.
func Kosk(conn *websocket.Conn, payload []byte) {
	challenge := new(model.ChallengePacket).FromBytes(payload)

	hash, err := bls.Hash(challenge.Random[:])
	if err != nil {
		SendError(conn, consts.OpCodeError, err)
		return
	}

	_, err = middleware.IsMessageValid(conn, hash, challenge.Signature)
	if err != nil {
		SendError(conn, consts.OpCodeError, consts.ErrInvalidKosk)
		return
	}

	challenge.Passed = true
	store.Challenges.Store(conn, *challenge)
	SendMessage(conn, consts.OpCodeFeedback, "kosk.ok")
}
