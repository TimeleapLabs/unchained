package handler

import (
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/transport/server/packet"
)

// Challenge is a method that is used to sign the challenge packet and return the signed packet to verify the client.
func (h *consumer) Challenge(message []byte) []byte {
	challenge := new(packet.ChallengePacket).FromBytes(message)

	signature, _ := crypto.Identity.Bls.Sign(challenge.Random[:])
	challenge.Signature = signature.Bytes()

	return challenge.Sia().Bytes()
}

// Challenge is a method that is used to sign the challenge packet and return the signed packet to verify the client.
func (w worker) Challenge(message []byte) []byte {
	challenge := new(packet.ChallengePacket).FromBytes(message)

	signature, _ := crypto.Identity.Bls.Sign(challenge.Random[:])
	challenge.Signature = signature.Bytes()

	return challenge.Sia().Bytes()
}
