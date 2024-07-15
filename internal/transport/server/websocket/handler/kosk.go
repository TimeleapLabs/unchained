package handler

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/transport/server/packet"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/middleware"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/store"
	"github.com/gorilla/websocket"
)

func Kosk(conn *websocket.Conn, payload []byte) error {
	challenge := new(packet.ChallengePacket).FromBytes(payload)

	hash, err := bls.Hash(challenge.Random[:])
	if err != nil {
		return err
	}

	_, err = middleware.IsMessageValid(conn, hash, challenge.Signature)
	if err != nil {
		return consts.ErrInvalidKosk
	}

	challenge.Passed = true
	store.Challenges.Store(conn, *challenge)

	return nil
}
