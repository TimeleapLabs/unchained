package middleware

import (
	"crypto/ed25519"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/store"
	"github.com/gorilla/websocket"
)

// IsMessageValid checks if the message's signature belong to signer or not.
func IsMessageValid(conn *websocket.Conn, message []byte, signature [64]byte) (model.Signer, error) {
	signer, ok := store.Signers.Load(conn)
	if !ok {
		return model.Signer{}, consts.ErrMissingHello
	}

	pk := ed25519.PublicKey{}
	copy(pk, signer.PublicKey[:])

	if ok = crypto.Identity.Ed25519.Verify(signature[:], message, pk); !ok {
		return model.Signer{}, consts.ErrInvalidSignature
	}

	return signer, nil
}
