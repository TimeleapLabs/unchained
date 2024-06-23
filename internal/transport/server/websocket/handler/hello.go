package handler

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/transport/server/packet"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/store"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/gorilla/websocket"
)

// Hello handler store the new client in the Signers map and send it a challenge packet.
func Hello(conn *websocket.Conn, payload []byte) {
	utils.Logger.With("IP", conn.RemoteAddr().String()).Info("New Client Registered")
	signer := new(model.Signer).FromBytes(payload)

	if signer.Name == "" {
		utils.Logger.Error("Signer name is empty Or public key is invalid")
		SendError(conn, consts.OpCodeError, consts.ErrInvalidConfig)
		return
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

	SendMessage(conn, consts.OpCodeFeedback, "conf.ok")
	Send(conn, consts.OpCodeKoskChallenge, challenge.Sia().Bytes())
}
