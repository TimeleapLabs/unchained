package model

import sia "github.com/pouya-eghbali/go-sia/v2/pkg"

const (
	LenOfChallenge = 128
	LenOfSignature = 48
)

type ChallengePacket struct {
	Passed    bool
	Random    [LenOfChallenge]byte
	Signature [LenOfSignature]byte
}

func (c *ChallengePacket) Sia() *sia.Sia {
	return new(sia.Sia).
		AddBool(c.Passed).
		AddByteArray8(c.Random[:]).
		AddByteArray8(c.Signature[:])
}

func (c *ChallengePacket) DeSia(sia *sia.Sia) *ChallengePacket {
	c.Passed = sia.ReadBool()
	copy(c.Random[:], sia.ReadByteArray8())
	copy(c.Signature[:], sia.ReadByteArray8())

	return c
}
