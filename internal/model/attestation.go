package model

import (
	"encoding/json"
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
	Meta         map[string]interface{}
}

func (c *Attestation) Sia() sia.Sia {
	json, err := json.Marshal(c.Meta)
	if err != nil {
		utils.Logger.With("Err", err).Error("Cannot marshal meta")
		panic(err)
	}

	return sia.New().
		AddUInt64(c.Timestamp).
		AddByteArray8(c.Hash).
		AddByteArray8(c.Topic).
		AddByteArray32(json)
}

func (c *Attestation) FromBytes(payload []byte) *Attestation {
	siaMessage := sia.NewFromBytes(payload)
	return c.FromSia(siaMessage)
}

func (c *Attestation) FromSia(sia sia.Sia) *Attestation {
	c.Timestamp = sia.ReadUInt64()
	c.Hash = sia.ReadByteArray8()
	c.Topic = sia.ReadByteArray8()
	c.Meta = make(map[string]interface{})

	jsonBytes := sia.ReadByteArray32()
	err := json.Unmarshal(jsonBytes, &c.Meta)
	if err != nil {
		utils.Logger.With("Err", err).Error("Cannot unmarshal meta")
		panic(err)
	}

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
