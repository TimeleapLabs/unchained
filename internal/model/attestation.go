package model

import (
	"time"

	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"

	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

type AttestationDataFrame struct {
	ID    uint               `bson:"-"             gorm:"primarykey"`
	DocID primitive.ObjectID `bson:"_id,omitempty" gorm:"-"`

	Hash      string      `bson:"hash"      json:"hash"`
	Timestamp time.Time   `bson:"timestamp" json:"timestamp"`
	Data      Attestation `bson:"data"      gorm:"embedded"  json:"data"`
}

type Attestation struct {
	SignersCount uint64
	Signature    []byte
	Consensus    bool
	Voted        int64
	Timestamp    uint64
	Hash         []byte `gorm:"uniqueIndex:idx_topic_hash"`
	Topic        []byte `gorm:"uniqueIndex:idx_topic_hash"`
	Correct      bool
}

func (c *Attestation) Sia() sia.Sia {
	return sia.New().
		AddUInt64(c.Timestamp).
		AddByteArray8(c.Hash).
		AddByteArray8(c.Topic).
		AddBool(c.Correct)
}

func (c *Attestation) FromBytes(payload []byte) *Attestation {
	siaMessage := sia.NewFromBytes(payload)
	return c.FromSia(siaMessage)
}

func (c *Attestation) FromSia(sia sia.Sia) *Attestation {
	c.Timestamp = sia.ReadUInt64()
	c.Hash = sia.ReadByteArray8()
	c.Topic = sia.ReadByteArray8()
	c.Correct = sia.ReadBool()

	return c
}

func (c *Attestation) Bls() *bls12381.G1Affine {
	hash, err := bls.Hash(c.Sia().Bytes())
	if err != nil {
		utils.Logger.With("Err", err).Error("Cannot hash bls")
		return &bls12381.G1Affine{}
	}

	return &hash
}
