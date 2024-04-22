package handler

import (
	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/model"
	"github.com/KenshiTech/unchained/internal/transport/server/websocket/middleware"

	"github.com/KenshiTech/unchained/internal/consts"
	"github.com/KenshiTech/unchained/internal/transport/server/websocket/store"
	"github.com/gorilla/websocket"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

func Kosk(conn *websocket.Conn, payload []byte) error {
	challenge := new(model.ChallengePacket).DeSia(&sia.Sia{Content: payload})

	hash, err := bls.Hash(challenge.Random[:])
	if err != nil {
		return err
	}

	_, err = middleware.IsMessageValid(conn, hash, challenge.Signature)
	if err != nil {
		return consts.ErrInvalidKosk
	}

	store.Challenges.Store(conn, *challenge)

	return nil
}
