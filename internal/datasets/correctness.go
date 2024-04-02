package datasets

import (
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

type Correctness struct {
	Timestamp uint64
	Hash      [64]byte
	Topic     [64]byte
	Correct   bool
}

type CorrectnessReport struct {
	Correctness
	Signature [48]byte
}

type BroadcastCorrectnessPacket struct {
	Info      Correctness
	Signature [48]byte
	Signer    Signer
}

func (c *Correctness) Sia() *sia.Sia {
	return new(sia.Sia).
		AddUInt64(c.Timestamp).
		AddByteArray8(c.Hash[:]).
		AddByteArray8(c.Topic[:]).
		AddBool(c.Correct)
}

func (c *Correctness) DeSia(sia *sia.Sia) *Correctness {
	c.Timestamp = sia.ReadUInt64()
	copy(c.Hash[:], sia.ReadByteArray8())
	copy(c.Topic[:], sia.ReadByteArray8())
	c.Correct = sia.ReadBool()

	return c
}

func (c *CorrectnessReport) Sia() *sia.Sia {
	return new(sia.Sia).
		EmbedSia(c.Correctness.Sia()).
		AddByteArray8(c.Signature[:])
}

func (c *CorrectnessReport) DeSia(sia *sia.Sia) *CorrectnessReport {
	c.Correctness.DeSia(sia)
	copy(c.Signature[:], sia.ReadByteArray8())

	return c
}

func (b *BroadcastCorrectnessPacket) Sia() *sia.Sia {
	return new(sia.Sia).
		EmbedSia(b.Info.Sia()).
		AddByteArray8(b.Signature[:]).
		EmbedSia(b.Signer.Sia())
}

func (b *BroadcastCorrectnessPacket) DeSia(sia *sia.Sia) *BroadcastCorrectnessPacket {
	b.Info.DeSia(sia)
	copy(b.Signature[:], sia.ReadByteArray8())
	b.Signer.DeSia(sia)

	return b
}
