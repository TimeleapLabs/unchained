package model

import (
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
)

type Signer struct {
	Name           string
	EvmAddress     string
	PublicKey      [96]byte
	ShortPublicKey [48]byte
}

type Signature struct {
	Signature bls12381.G1Affine
	Signer    Signer
}

func (s *Signature) Sia() sia.Sia {
	return sia.New().
		AddByteArray8(s.Signature.Marshal()).
		EmbedBytes(s.Signer.Sia().Bytes())
}

func (s *Signature) FromBytes(payload []byte) *Signature {
	siaMessage := sia.NewFromBytes(payload)
	return s.FromSia(siaMessage)
}

func (s *Signature) FromSia(sia sia.Sia) *Signature {
	err := s.Signature.Unmarshal(sia.ReadByteArray8())

	if err != nil {
		s.Signature = bls12381.G1Affine{}
	}

	s.Signer.FromSia(sia)

	return s
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
