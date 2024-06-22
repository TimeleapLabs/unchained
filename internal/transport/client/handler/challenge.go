package handler

import (
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/model"
)

func (h *consumer) Challenge(message []byte) []byte {
	challenge := new(model.ChallengePacket).FromBytes(message)

	signature, _ := crypto.Identity.Bls.Sign(challenge.Random[:])
	challenge.Signature = signature.Bytes()

	return challenge.Sia().Bytes()
}

func (w worker) Challenge(message []byte) []byte {
	challenge := new(model.ChallengePacket).FromBytes(message)

	signature, _ := crypto.Identity.Bls.Sign(challenge.Random[:])
	challenge.Signature = signature.Bytes()

	return challenge.Sia().Bytes()
}
