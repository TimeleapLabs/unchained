package middleware

import (
	"github.com/KenshiTech/unchained/internal/consts"
	"github.com/KenshiTech/unchained/internal/crypto"
	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/model"
	"github.com/KenshiTech/unchained/internal/transport/server/websocket/store"
	"github.com/KenshiTech/unchained/internal/utils"
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

	pk, err := bls.RecoverPublicKey(signer.PublicKey)
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
