package handler

import (
	"github.com/KenshiTech/unchained/internal/constants"
	"github.com/KenshiTech/unchained/internal/crypto/kosk"
	"github.com/KenshiTech/unchained/internal/datasets"
	"github.com/KenshiTech/unchained/internal/log"
	"github.com/KenshiTech/unchained/internal/transport/server/websocket/store"
	"github.com/gorilla/websocket"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

func Hello(conn *websocket.Conn, payload []byte) ([]byte, error) {
	signer := new(datasets.Signer).DeSia(&sia.Sia{Content: payload})

	if signer.Name == "" {
		log.Logger.Error("Signer name is empty Or public key is invalid")
		return []byte{}, constants.ErrInvalidConfig
	}

	store.Signers.Range(func(conn *websocket.Conn, signerInMap datasets.Signer) bool {
		publicKeyInUse := signerInMap.PublicKey == signer.PublicKey
		if publicKeyInUse {
			Close(conn)
		}
		return !publicKeyInUse
	})

	store.Signers.Store(conn, *signer)

	// Start KOSK verification
	challenge := kosk.Challenge{Random: kosk.NewChallenge()}
	store.Challenges.Store(conn, challenge)
	koskPayload := challenge.Sia().Content

	return koskPayload, nil
}
