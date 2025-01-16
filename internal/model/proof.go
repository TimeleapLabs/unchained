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

	Signatures []Signature `bson:"signatures" gorm:"many2many:proof_signatures;" json:"signers"`
}

func (p *Proof) Sia() sia.Sia {
	signatures := sia.NewSiaArray[Signature]().AddArray64(p.Signatures, func(s *sia.ArraySia[Signature], item Signature) {
		s.EmbedBytes(item.Sia().Bytes())
	})

	return sia.New().
		AddByteArray8(p.Hash).
		AddInt64(p.Timestamp.Unix()).
		AddByteArray64(signatures.Bytes())
}

func (p *Proof) FromBytes(payload []byte) *Proof {
	siaMessage := sia.NewFromBytes(payload)
	return p.FromSia(siaMessage)
}

func (p *Proof) FromSia(siaObj sia.Sia) *Proof {
	p.Hash = siaObj.ReadByteArray8()
	p.Timestamp = time.Unix(siaObj.ReadInt64(), 0)

	signatureBytes := siaObj.ReadByteArray64()
	signatures := sia.NewArrayFromBytes[Signature](signatureBytes).ReadArray64(func(s *sia.ArraySia[Signature]) Signature {
		signature := new(Signature)
		signature.FromBytes(s.Bytes())
		return *signature
	})

	p.Signatures = signatures

	return p
}

func NewProof(signatures []Signature, hash []byte) *Proof {
	return &Proof{
		Hash:       hash,
		Timestamp:  time.Now(),
		Signatures: signatures,
	}
}
