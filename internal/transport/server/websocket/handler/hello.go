package handler

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/store"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/gorilla/websocket"
)

func (h Handler) Hello(conn *websocket.Conn, payload []byte) ([]byte, error) {
	signer := new(model.Signer).FromBytes(payload)

	if signer.Name == "" {
		utils.Logger.Error("Signer name is empty Or public key is invalid")
		return []byte{}, consts.ErrInvalidConfig
	}

	preConn, preConnExist := h.clientRepository.GetByPublicKey(signer.PublicKey)
	if preConnExist {
		Close(preConn)
	}

	h.clientRepository.Set(conn, *signer)

	// Start KOSK verification
	challenge := model.ChallengePacket{Random: utils.NewChallenge()}
	store.Challenges.Store(conn, challenge)

	return challenge.Sia().Bytes(), nil
}
