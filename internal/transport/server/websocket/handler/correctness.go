package handler

import (
	"github.com/KenshiTech/unchained/constants"
	"github.com/KenshiTech/unchained/crypto/bls"
	"github.com/KenshiTech/unchained/datasets"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/transport/server/websocket/middleware"
	"github.com/gorilla/websocket"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

func CorrectnessRecord(conn *websocket.Conn, payload []byte) ([]byte, error) {
	signer, err := middleware.CheckPublicKey(conn)
	if err != nil {
		return []byte{}, err
	}

	report := new(datasets.CorrectnessReport).DeSia(&sia.Sia{Content: payload})
	toHash := report.Correctness.Sia().Content
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
		return []byte{}, constants.ErrCantVerifyBls
	}

	ok, err := bls.Verify(signature, hash, pk)
	if err != nil {
		log.Logger.With("Error", err).Error("Can't verify bls")
		return []byte{}, constants.ErrCantVerifyBls
	}
	if !ok {
		return []byte{}, constants.ErrInvalidSignature
	}

	broadcastPacket := datasets.BroadcastCorrectnessPacket{
		Info:      report.Correctness,
		Signature: report.Signature,
		Signer:    *signer,
	}

	broadcastPayload := broadcastPacket.Sia().Content
	return broadcastPayload, nil
}
