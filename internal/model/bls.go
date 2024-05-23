package model

import (
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

type Signers []Signer

type Signer struct {
	Name           string
	EvmAddress     string
	PublicKey      [96]byte
	ShortPublicKey [48]byte
}

type Signature struct {
	Signature []byte
	Signer    Signer
}

func (s *Signature) Sia() sia.Sia {
	return sia.New().
		AddByteArray8(s.Signature).
		EmbedBytes(s.Signer.Sia().Bytes())
}

func (s *Signature) FromBytes(payload []byte) *Signature {
	siaMessage := sia.NewFromBytes(payload)
	return s.FromSia(siaMessage)
}

func (s *Signature) FromSia(sia sia.Sia) *Signature {
	s.Signature = sia.ReadByteArray8()
	s.Signer.FromSia(sia)

	return s
}

func (s Signers) Sia() sia.Sia {
	return new(sia.ArraySia[Signer]).
		AddArray8(s, func(s *sia.ArraySia[Signer], item Signer) {
			s.EmbedBytes(item.Sia().Bytes())
		})
}

// func (s Signers) FromBytes(payload []byte) Signers {
//	signers := Signers{}
//
//	siaArray := sia.ArraySia[Signer]{
//		sia.NewFromBytes(payload),
//	}
//
//	ReadArray8(func(s *sia.ArraySia[Signer]) Signer {
//		signers = append(signers, Signer)
//	})
//}

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
