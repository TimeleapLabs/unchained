package model

import (
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
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
	err := s.Signature.Unmarshal(siaMessage.ReadByteArray8())

	if err != nil {
		s.Signature = bls12381.G1Affine{}
	}

	s.Signer.FromBytes(payload)

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
	s.Name = siaMessage.ReadString8()
	s.EvmAddress = siaMessage.ReadString8()
	copy(s.PublicKey[:], siaMessage.ReadByteArray8())
	copy(s.ShortPublicKey[:], siaMessage.ReadByteArray8())

	return s
}
