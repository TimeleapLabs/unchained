package handler

import (
	"github.com/KenshiTech/unchained/constants"
	"github.com/KenshiTech/unchained/crypto/bls"
	"github.com/KenshiTech/unchained/datasets"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/transport/server/websocket/middleware"
	"github.com/KenshiTech/unchained/transport/server/websocket/store"
	"github.com/gorilla/websocket"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

func PriceReport(conn *websocket.Conn, payload []byte) ([]byte, error) {
	err := middleware.CheckPublicKey(conn)
	if err != nil {
		return []byte{}, err
	}

	signer, ok := store.Signers.Load(conn)
	if !ok {
		return nil, constants.ErrMissingHello
	}

	report := new(datasets.PriceReport).DeSia(&sia.Sia{Content: payload})
	toHash := report.PriceInfo.Sia().Content
	hash, err := bls.Hash(toHash)

	if err != nil {
		log.Logger.Error("Can't hash bls: %v", err)
		return []byte{}, constants.ErrInternalError
	}

	signature, err := bls.RecoverSignature(report.Signature)
	if err != nil {
		log.Logger.Error("Can't recover bls signature: %v", err)
		return []byte{}, constants.ErrInternalError
	}

	pk, err := bls.RecoverPublicKey(signer.PublicKey)
	if err != nil {
		log.Logger.Error("Can't recover bls pub-key: %v", err)
		return []byte{}, constants.ErrInternalError
	}

	ok, err = bls.Verify(signature, hash, pk)
	if err != nil {
		log.Logger.Error("Can't recover bls pub-key: %v", err)
		return []byte{}, constants.ErrCantVerifyBls
	}
	if !ok {
		return []byte{}, constants.ErrInvalidSignature
	}

	priceInfo := datasets.BroadcastPricePacket{
		Info:      report.PriceInfo,
		Signature: report.Signature,
		Signer:    signer,
	}

	priceInfoByte := priceInfo.Sia().Content
	return priceInfoByte, nil
}
