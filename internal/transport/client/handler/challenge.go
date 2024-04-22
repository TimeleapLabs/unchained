package handler

import (
	"github.com/KenshiTech/unchained/internal/crypto"
	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/model"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

func (h *consumer) Challenge(message []byte) *model.ChallengePacket {
	challenge := new(model.ChallengePacket).DeSia(&sia.Sia{Content: message})

	signature, _ := bls.Sign(*crypto.Identity.Bls.SecretKey, challenge.Random[:])
	challenge.Signature = signature.Bytes()

	return challenge
}

func (w worker) Challenge(message []byte) *model.ChallengePacket {
	challenge := new(model.ChallengePacket).DeSia(&sia.Sia{Content: message})

	signature, _ := bls.Sign(*crypto.Identity.Bls.SecretKey, challenge.Random[:])
	challenge.Signature = signature.Bytes()

	return challenge
}
