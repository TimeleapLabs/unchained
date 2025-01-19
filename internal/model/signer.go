package model

import (
	"crypto/ed25519"

	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Signers []Signer

type Signer struct {
	ID    uint               `bson:"-"             gorm:"primarykey"`
	DocID primitive.ObjectID `bson:"_id,omitempty" gorm:"-"`

	Name       string            `json:"name"`
	EvmAddress string            `json:"evm_address"`
	PublicKey  ed25519.PublicKey `json:"public_key"`
}

type Signature struct {
	PublicKey ed25519.PublicKey `json:"public_key"`
	Signature [64]byte          `json:"signature"`
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
	copy(s.PublicKey[:], sia.ReadByteArray8())
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

	pk := sia.ReadByteArray8()
	copy(s.PublicKey[:], pk)

	return s
}
