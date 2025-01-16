package model

import (
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Signers []Signer

type Signer struct {
	ID    uint               `bson:"-"             gorm:"primarykey"`
	DocID primitive.ObjectID `bson:"_id,omitempty" gorm:"-"`

	Name       string `json:"name"`
	EvmAddress string `json:"evm_address"`
	PublicKey  []byte `json:"public_key"`
}

func (s *Signer) Sia() sia.Sia {
	return sia.New().
		AddString8(s.Name).
		AddString8(s.EvmAddress).
		AddByteArray8(s.PublicKey)
}

func (s *Signer) FromBytes(payload []byte) *Signer {
	siaMessage := sia.NewFromBytes(payload)
	return s.FromSia(siaMessage)
}

func (s *Signer) FromSia(sia sia.Sia) *Signer {
	s.Name = sia.ReadString8()
	s.EvmAddress = sia.ReadString8()
	s.PublicKey = sia.ReadByteArray8()

	return s
}
