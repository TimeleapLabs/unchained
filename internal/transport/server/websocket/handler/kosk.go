package handler

import (
	"github.com/KenshiTech/unchained/internal/constants"
	"github.com/KenshiTech/unchained/internal/crypto/kosk"
	"github.com/KenshiTech/unchained/internal/log"
	"github.com/KenshiTech/unchained/internal/transport/server/websocket/store"
	"github.com/gorilla/websocket"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

func Kosk(conn *websocket.Conn, payload []byte) error {
	challenge := new(kosk.Challenge).DeSia(&sia.Sia{Content: payload})

	signer, ok := store.Signers.Load(conn)
	if !ok {
		return constants.ErrMissingHello
	}

	var err error
	challenge.Passed, err = kosk.VerifyChallenge(challenge.Random, signer.PublicKey, challenge.Signature)

	if err != nil {
		return constants.ErrInvalidKosk
	}

	if !challenge.Passed {
		log.Logger.Error("challenge is Passed")
		return constants.ErrInvalidKosk
	}

	store.Challenges.Store(conn, *challenge)
	return nil
}
