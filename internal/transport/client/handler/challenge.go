package handler

import (
	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/crypto/kosk"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

func (h *consumer) Challenge(message []byte) *kosk.Challenge {
	challenge := new(kosk.Challenge).DeSia(&sia.Sia{Content: message})

	signature, _ := bls.Sign(*bls.ClientSecretKey, challenge.Random[:])
	challenge.Signature = signature.Bytes()

	return challenge
}

func (w worker) Challenge(message []byte) *kosk.Challenge {
	challenge := new(kosk.Challenge).DeSia(&sia.Sia{Content: message})

	signature, _ := bls.Sign(*bls.ClientSecretKey, challenge.Random[:])
	challenge.Signature = signature.Bytes()

	return challenge
}
