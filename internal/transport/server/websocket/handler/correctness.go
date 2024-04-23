package handler

import (
	"github.com/TimeleapLabs/unchained/internal/constants"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/datasets"
	"github.com/TimeleapLabs/unchained/internal/log"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/middleware"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/store"
	"github.com/gorilla/websocket"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

func CorrectnessRecord(conn *websocket.Conn, payload []byte) ([]byte, error) {
	err := middleware.IsConnectionAuthenticated(conn)
	if err != nil {
		return []byte{}, err
	}

	signer, ok := store.Signers.Load(conn)
	if !ok {
		return nil, constants.ErrMissingHello
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

	ok, err = crypto.Identity.Bls.Verify(signature, hash, pk)
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
		Signer:    signer,
	}

	broadcastPayload := broadcastPacket.Sia().Content
	return broadcastPayload, nil
}
