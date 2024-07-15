package handler

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/transport/server/packet"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/store"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/gorilla/websocket"
)

func Hello(conn *websocket.Conn, payload []byte) ([]byte, error) {
	signer := new(model.Signer).FromBytes(payload)

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
	challenge := packet.ChallengePacket{Random: utils.NewChallenge()}
	store.Challenges.Store(conn, challenge)

	return challenge.Sia().Bytes(), nil
}
