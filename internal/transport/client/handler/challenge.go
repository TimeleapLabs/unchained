package handler

import (
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/crypto/kosk"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

func (h *consumer) Challenge(message []byte) *kosk.Challenge {
	challenge := new(kosk.Challenge).DeSia(&sia.Sia{Content: message})

	signature, _ := bls.Sign(*crypto.Identity.Bls.SecretKey, challenge.Random[:])
	challenge.Signature = signature.Bytes()

	return challenge
}

func (w worker) Challenge(message []byte) *kosk.Challenge {
	challenge := new(kosk.Challenge).DeSia(&sia.Sia{Content: message})

	signature, _ := bls.Sign(*crypto.Identity.Bls.SecretKey, challenge.Random[:])
	challenge.Signature = signature.Bytes()

	return challenge
}
