package handler

import (
	"github.com/KenshiTech/unchained/internal/consts"
	"github.com/KenshiTech/unchained/internal/crypto"
	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/model"
	"github.com/KenshiTech/unchained/internal/transport/server/websocket/middleware"
	"github.com/KenshiTech/unchained/internal/transport/server/websocket/store"
	"github.com/KenshiTech/unchained/internal/utils"
	"github.com/gorilla/websocket"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

func PriceReport(conn *websocket.Conn, payload []byte) ([]byte, error) {
	err := middleware.IsConnectionAuthenticated(conn)
	if err != nil {
		return []byte{}, err
	}

	signer, ok := store.Signers.Load(conn)
	if !ok {
		return nil, consts.ErrMissingHello
	}

	report := new(model.PriceReport).DeSia(&sia.Sia{Content: payload})
	toHash := report.PriceInfo.Sia().Content
	hash, err := bls.Hash(toHash)

	if err != nil {
		utils.Logger.Error("Can't hash bls: %v", err)
		return []byte{}, consts.ErrInternalError
	}

	signature, err := bls.RecoverSignature(report.Signature)
	if err != nil {
		utils.Logger.Error("Can't recover bls signature: %v", err)
		return []byte{}, consts.ErrInternalError
	}

	pk, err := bls.RecoverPublicKey(signer.PublicKey)
	if err != nil {
		utils.Logger.Error("Can't recover bls pub-key: %v", err)
		return []byte{}, consts.ErrInternalError
	}

	ok, err = crypto.Identity.Bls.Verify(signature, hash, pk)
	if err != nil {
		utils.Logger.Error("Can't recover bls pub-key: %v", err)
		return []byte{}, consts.ErrCantVerifyBls
	}
	if !ok {
		return []byte{}, consts.ErrInvalidSignature
	}

	priceInfo := model.BroadcastPricePacket{
		Info:      report.PriceInfo,
		Signature: report.Signature,
		Signer:    signer,
	}

	priceInfoByte := priceInfo.Sia().Content
	return priceInfoByte, nil
}
