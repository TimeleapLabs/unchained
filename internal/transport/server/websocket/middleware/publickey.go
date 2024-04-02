package middleware

import (
	"github.com/KenshiTech/unchained/constants"
	"github.com/KenshiTech/unchained/datasets"
	"github.com/KenshiTech/unchained/transport/server/websocket/store"
	"github.com/gorilla/websocket"
)

func CheckPublicKey(conn *websocket.Conn) (*datasets.Signer, error) {
	challenge, ok := store.Challenges.Load(conn)
	if !ok || !challenge.Passed {
		return nil, constants.ErrMissingKosk
	}

	signer, ok := store.Signers.Load(conn)
	if !ok {
		return nil, constants.ErrMissingHello
	}

	return &signer, nil
}
