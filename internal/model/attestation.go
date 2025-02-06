package model

import (
	"encoding/json"

	"github.com/TimeleapLabs/timeleap/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"

	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
)

type AttestationDataFrame struct {
	ID    uint               `bson:"-"             gorm:"primarykey"`
	DocID primitive.ObjectID `bson:"_id,omitempty" gorm:"-"`

	Hash string      `bson:"hash" gorm:"uniqueIndex:idx_topic_hash" json:"hash"`
	Data Attestation `bson:"data" gorm:"embedded"                   json:"data"`
}

type Attestation struct {
	Timestamp uint64
	Topic     []byte `gorm:"uniqueIndex:idx_topic_hash"`
	Meta      map[string]interface{}
}

func (c *Attestation) Sia() sia.Sia {
	json, err := json.Marshal(c.Meta)
	if err != nil {
		utils.Logger.With("Err", err).Error("Cannot marshal meta")
		panic(err)
	}

	return sia.New().
		AddUInt64(c.Timestamp).
		AddByteArray8(c.Topic).
		AddByteArray32(json)
}

func (c *Attestation) FromBytes(payload []byte) *Attestation {
	siaMessage := sia.NewFromBytes(payload)
	return c.FromSia(siaMessage)
}

func (c *Attestation) FromSia(sia sia.Sia) *Attestation {
	c.Timestamp = sia.ReadUInt64()
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
