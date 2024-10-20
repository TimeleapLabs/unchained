package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

type Proof struct {
	ID    uint               `bson:"-"             gorm:"primarykey"`
	DocID primitive.ObjectID `bson:"_id,omitempty" gorm:"-"`

	Hash      []byte    `bson:"hash"      json:"hash"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	Signature []byte    `bson:"signature" json:"signature"`

	Signers []Signer `bson:"signers" gorm:"many2many:proof_signers;" json:"signers"`
}

func (p *Proof) Sia() sia.Sia {
	signers := sia.NewSiaArray[Signer]().AddArray64(p.Signers, func(s *sia.ArraySia[Signer], item Signer) {
		s.EmbedBytes(item.Sia().Bytes())
	})

	return sia.New().
		AddByteArray8(p.Hash).
		AddInt64(p.Timestamp.Unix()).
		AddByteArray8(p.Signature).
		AddByteArray64(signers.Bytes())
}

func (p *Proof) FromBytes(payload []byte) *Proof {
	siaMessage := sia.NewFromBytes(payload)
	return p.FromSia(siaMessage)
}

func (p *Proof) FromSia(siaObj sia.Sia) *Proof {
	signers := sia.NewArrayFromBytes[Signer](siaObj.ReadByteArray64()).ReadArray64(func(s *sia.ArraySia[Signer]) Signer {
		signer := Signer{}
		signer.FromBytes(s.ReadByteArray64())
		return signer
	})

	p.Hash = siaObj.ReadByteArray8()
	p.Timestamp = time.Unix(siaObj.ReadInt64(), 0)
	p.Signature = siaObj.ReadByteArray8()
	p.Signers = signers
	return p
}

func NewProof(signers []Signer, signature []byte) *Proof {
	return &Proof{
		Hash:      Signers(signers).Bls(),
		Timestamp: time.Now(),
		Signature: signature,
		Signers:   signers,
	}
}
