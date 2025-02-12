package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
)

type Proof struct {
	ID    uint               `bson:"-"`
	DocID primitive.ObjectID `bson:"_id,omitempty"`

	Hash      []byte    `bson:"hash"      json:"hash"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`

	Signatures []Signature `bson:"signatures" json:"signers"`
}

func (p *Proof) Sia() sia.Sia {
	s := sia.New().
		AddByteArray8(p.Hash).
		AddInt64(p.Timestamp.Unix())

	sarr := sia.NewArray[Signature](&s)
	sarr.AddArray64(p.Signatures, func(s *sia.ArraySia[Signature], item Signature) {
		s.EmbedBytes(item.Sia().Bytes())
	})

	return s
}

func (p *Proof) FromBytes(payload []byte) *Proof {
	siaMessage := sia.NewFromBytes(payload)
	return p.FromSia(siaMessage)
}

func (p *Proof) FromSia(s sia.Sia) *Proof {
	p.Hash = s.ReadByteArray8()
	p.Timestamp = time.Unix(s.ReadInt64(), 0)

	sarr := sia.NewArray[Signature](&s)
	signatures := sarr.ReadArray64(func(s *sia.ArraySia[Signature]) Signature {
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
