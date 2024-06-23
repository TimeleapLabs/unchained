package model

import (
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

type Signer struct {
	Name           string
	EvmAddress     string
	PublicKey      [96]byte
	ShortPublicKey [48]byte
}

func (s *Signer) Sia() sia.Sia {
	return sia.New().
		AddString8(s.Name).
		AddString8(s.EvmAddress).
		AddByteArray8(s.PublicKey[:]).
		AddByteArray8(s.ShortPublicKey[:])
}

func (s *Signer) FromBytes(payload []byte) *Signer {
	siaMessage := sia.NewFromBytes(payload)
	return s.FromSia(siaMessage)
}

func (s *Signer) FromSia(sia sia.Sia) *Signer {
	s.Name = sia.ReadString8()
	s.EvmAddress = sia.ReadString8()
	copy(s.PublicKey[:], sia.ReadByteArray8())
	copy(s.ShortPublicKey[:], sia.ReadByteArray8())

	return s
}
