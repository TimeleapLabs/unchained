package handler

import (
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

func (h *consumer) Challenge(message []byte) []byte {
	challenge := new(model.ChallengePacket).FromBytes(message)

	signature, err := crypto.Identity.Bls.Sign(challenge.Random[:])
	if err != nil {
		utils.Logger.Error(err.Error())
		return nil
	}

	challenge.Signature = [48]byte(signature)

	return challenge.Sia().Bytes()
}

func (w *worker) Challenge(message []byte) []byte {
	challenge := new(model.ChallengePacket).FromBytes(message)

	signature, err := crypto.Identity.Bls.Sign(challenge.Random[:])
	if err != nil {
		utils.Logger.Error(err.Error())
		return nil
	}

	challenge.Signature = [48]byte(signature)

	return challenge.Sia().Bytes()
}
