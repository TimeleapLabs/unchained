package middleware

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/store"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/gorilla/websocket"
)

func IsMessageValid(conn *websocket.Conn, message []byte, signature [48]byte) (model.Signer, error) {
	signer, ok := store.Signers.Load(conn)
	if !ok {
		return model.Signer{}, consts.ErrMissingHello
	}

	ok, err := crypto.Identity.Bls.Verify(signature[:], message, signer.PublicKey[:])
	if err != nil {
		utils.Logger.With("Err", err).Error("Can't verify bls")
		return model.Signer{}, consts.ErrCantVerifyBls
	}

	if !ok {
		return model.Signer{}, consts.ErrInvalidSignature
	}

	return signer, nil
}
