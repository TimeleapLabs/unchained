package middleware

import (
	"encoding/hex"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/store"
	"github.com/TimeleapLabs/unchained/internal/utils"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	"github.com/gorilla/websocket"
)

func IsMessageValid(conn *websocket.Conn, message bls12381.G1Affine, signature [48]byte) (model.Signer, error) {
	signer, ok := store.Signers.Load(conn)
	if !ok {
		return model.Signer{}, consts.ErrMissingHello
	}

	signatureBls, err := bls.RecoverSignature(signature)
	if err != nil {
		utils.Logger.With("Err", err).Error("Can't recover bls signature")
		return model.Signer{}, consts.ErrInternalError
	}

	publicKeyBytes, err := hex.DecodeString(signer.PublicKey)
	if err != nil {
		utils.Logger.Error("Can't decode public key: %v", err)
		return model.Signer{}, consts.ErrInternalError
	}

	pk, err := bls.RecoverPublicKey([96]byte(publicKeyBytes))
	if err != nil {
		utils.Logger.With("Err", err).Error("Can't recover pub key pub-key")
		return model.Signer{}, consts.ErrInternalError
	}

	ok, err = crypto.Identity.Bls.Verify(signatureBls, message, pk)
	if err != nil {
		utils.Logger.With("Err", err).Error("Can't verify bls")
		return model.Signer{}, consts.ErrCantVerifyBls
	}

	if !ok {
		return model.Signer{}, consts.ErrInvalidSignature
	}

	return signer, nil
}
