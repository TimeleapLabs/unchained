package handler

import (
	"github.com/KenshiTech/unchained/internal/crypto"
	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/model"
)

func (h *consumer) Challenge(message []byte) []byte {
	challenge := new(model.ChallengePacket).FromBytes(message)

	signature, _ := bls.Sign(*crypto.Identity.Bls.SecretKey, challenge.Random[:])
	challenge.Signature = signature.Bytes()

	return challenge.Sia().Bytes()
}

func (w worker) Challenge(message []byte) []byte {
	challenge := new(model.ChallengePacket).FromBytes(message)

	signature, _ := bls.Sign(*crypto.Identity.Bls.SecretKey, challenge.Random[:])
	challenge.Signature = signature.Bytes()

	return challenge.Sia().Bytes()
}
