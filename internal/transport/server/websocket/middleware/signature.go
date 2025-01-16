package middleware

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/store"
	"github.com/gorilla/websocket"
)

// IsMessageValid checks if the message's signature belong to signer or not.
func IsMessageValid(conn *websocket.Conn, message []byte, signature [48]byte) (model.Signer, error) {
	signer, ok := store.Signers.Load(conn)
	if !ok {
		return model.Signer{}, consts.ErrMissingHello
	}

	if ok = crypto.Identity.Ed25519.Verify(signature[:], message, signer.PublicKey); !ok {
		return model.Signer{}, consts.ErrInvalidSignature
	}

	return signer, nil
}
