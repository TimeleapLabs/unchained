package handler

import (
	"github.com/KenshiTech/unchained/internal/consts"
	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/model"
	"github.com/KenshiTech/unchained/internal/transport/server/websocket/middleware"
	"github.com/KenshiTech/unchained/internal/transport/server/websocket/store"
	"github.com/gorilla/websocket"
)

func Kosk(conn *websocket.Conn, payload []byte) error {
	challenge := new(model.ChallengePacket).FromBytes(payload)

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
