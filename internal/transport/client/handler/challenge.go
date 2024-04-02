package handler

import (
	"github.com/KenshiTech/unchained/crypto/bls"
	"github.com/KenshiTech/unchained/crypto/kosk"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

func (h *Handler) Challenge(message []byte) *kosk.Challenge {
	challenge := new(kosk.Challenge).DeSia(&sia.Sia{Content: message})

	signature, _ := bls.Sign(*bls.ClientSecretKey, challenge.Random[:])
	challenge.Signature = signature.Bytes()

	return challenge
}
