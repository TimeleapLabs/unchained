package model

import (
	"crypto/ed25519"

	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
)

type Signers []Signer

type Signer struct {
	ID         uint
	Name       string
	EvmAddress string
	PublicKey  ed25519.PublicKey
}

type Signature struct {
	PublicKey ed25519.PublicKey
	Signature [64]byte
}

func (s *Signature) Sia() sia.Sia {
	return sia.New().
		AddByteArray8(s.PublicKey[:]).
		AddByteArray8(s.Signature[:])
}

func (s *Signature) FromBytes(payload []byte) *Signature {
	siaMessage := sia.NewFromBytes(payload)
	return s.FromSia(siaMessage)
}

func (s *Signature) FromSia(sia sia.Sia) *Signature {
	s.PublicKey = ed25519.PublicKey(sia.ReadByteArray8())
	copy(s.Signature[:], sia.ReadByteArray8())

	return s
}

func (s *Signer) Sia() sia.Sia {
	return sia.New().
		AddString8(s.Name).
		AddString8(s.EvmAddress).
		AddByteArray8(s.PublicKey[:])
}

func (s *Signer) FromBytes(payload []byte) *Signer {
	siaMessage := sia.NewFromBytes(payload)
	return s.FromSia(siaMessage)
}

func (s *Signer) FromSia(sia sia.Sia) *Signer {
	s.Name = sia.ReadString8()
	s.EvmAddress = sia.ReadString8()
	s.PublicKey = ed25519.PublicKey(sia.ReadByteArray8())

	return s
}
