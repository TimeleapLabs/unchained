package packet

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/store"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/gorilla/websocket"
)

func SignPacket(message []byte) ([]byte, error) {
	signature, err := crypto.Identity.Ed25519.Sign(message)
	if err != nil {
		return nil, err
	}
	return append(message, signature...), nil
}

func MustSignPacket(message []byte) []byte {
	signed, err := SignPacket(message)
	if err != nil {
		utils.Logger.
			With("Error", err).
			Error("Cannot sign packet")
		panic(err)
	}

	return signed
}

func IsPacketValid(conn *websocket.Conn, message []byte) (model.Signer, [64]byte, error) {
	signer, ok := store.Signers.Load(conn)
	if !ok {
		return model.Signer{}, [64]byte{}, consts.ErrMissingHello
	}

	signature := [64]byte{}
	copy(signature[:], message[len(message)-64:])

	if ok = crypto.Identity.Ed25519.Verify(signature[:], message, signer.PublicKey); !ok {
		return model.Signer{}, [64]byte{}, consts.ErrInvalidSignature
	}

	return signer, signature, nil
}
