package packet

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

func (c *ChallengePacket) Sia() sia.Sia {
	return sia.New().
		AddBool(c.Passed).
		AddByteArray8(c.Random[:]).
		AddByteArray8(c.Signature[:])
}

func (c *ChallengePacket) FromBytes(payload []byte) *ChallengePacket {
	siaMessage := sia.NewFromBytes(payload)
	c.Passed = siaMessage.ReadBool()
	copy(c.Random[:], siaMessage.ReadByteArray8())
	copy(c.Signature[:], siaMessage.ReadByteArray8())

	return c
}
