package model

import (
	"encoding/hex"
	"time"

	"github.com/TimeleapLabs/unchained/internal/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"

	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

type Proof struct {
	ID    uint               `bson:"-"             gorm:"primarykey"`
	DocID primitive.ObjectID `bson:"_id,omitempty" gorm:"-"`

	Hash      string    `bson:"hash"      json:"hash"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	Signature string    `bson:"signature" json:"signature"`

	Signers []Signer `bson:"signers" gorm:"many2many:proof_signers;" json:"signers"`
}

func (p *Proof) Sia() sia.Sia {
	signers := sia.NewSiaArray[Signer]().AddArray64(p.Signers, func(s *sia.ArraySia[Signer], item Signer) {
		s.EmbedBytes(item.Sia().Bytes())
	})

	hashBytes, err := hex.DecodeString(p.Hash)
	if err != nil {
		utils.Logger.Error("Can't decode hash: %v", err)
		return sia.New()
	}

	signatureBytes, err := hex.DecodeString(p.Signature)
	if err != nil {
		utils.Logger.Error("Can't decode signature: %v", err)
		return sia.New()
	}

	return sia.New().
		AddByteArray8(hashBytes).
		AddInt64(p.Timestamp.Unix()).
		AddByteArray8(signatureBytes).
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

	p.Hash = hex.EncodeToString(siaObj.ReadByteArray8())
	p.Timestamp = time.Unix(siaObj.ReadInt64(), 0)
	p.Signature = hex.EncodeToString(siaObj.ReadByteArray8())
	p.Signers = signers
	return p
}

func NewProof(signers []Signer, signature []byte) *Proof {
	return &Proof{
		Hash:      hex.EncodeToString(Signers(signers).Bls()),
		Timestamp: time.Now(),
		Signature: hex.EncodeToString(signature),
		Signers:   signers,
	}
}
