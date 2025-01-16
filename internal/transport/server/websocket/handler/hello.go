package handler

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/store"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/gorilla/websocket"
)

// Hello handler store the new client in the Signers map.
func Hello(conn *websocket.Conn, payload []byte) error {
	utils.Logger.With("IP", conn.RemoteAddr().String()).Info("New Client Registered")
	signer := new(model.Signer).FromBytes(payload)

	if signer.Name == "" {
		utils.Logger.Error("Signer name is empty or public key is invalid")
		return consts.ErrInvalidConfig
	}

	store.Signers.Store(conn, *signer)
	return nil
}
