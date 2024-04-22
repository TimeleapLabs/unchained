package handler

import (
	"github.com/KenshiTech/unchained/internal/consts"
	"github.com/KenshiTech/unchained/internal/model"
	"github.com/KenshiTech/unchained/internal/transport/server/websocket/store"
	"github.com/KenshiTech/unchained/internal/utils"
	"github.com/gorilla/websocket"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

func Hello(conn *websocket.Conn, payload []byte) ([]byte, error) {
	signer := new(model.Signer).DeSia(&sia.Sia{Content: payload})

	if signer.Name == "" {
		utils.Logger.Error("Signer name is empty Or public key is invalid")
		return []byte{}, consts.ErrInvalidConfig
	}

	store.Signers.Range(func(conn *websocket.Conn, signerInMap model.Signer) bool {
		publicKeyInUse := signerInMap.PublicKey == signer.PublicKey
		if publicKeyInUse {
			Close(conn)
		}
		return !publicKeyInUse
	})

	store.Signers.Store(conn, *signer)

	// Start KOSK verification
	challenge := model.ChallengePacket{Random: utils.NewChallenge()}
	store.Challenges.Store(conn, challenge)
	koskPayload := challenge.Sia().Content

	return koskPayload, nil
}
