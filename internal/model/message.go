package model

import (
	"encoding/json"

	"github.com/TimeleapLabs/timeleap/internal/utils"

	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
)

type Message struct {
	Timestamp uint64
	Topic     []byte
	Meta      map[string]interface{}
}

func (c *Message) Sia() sia.Sia {
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

func (c *Message) FromBytes(payload []byte) *Message {
	siaMessage := sia.NewFromBytes(payload)
	return c.FromSia(siaMessage)
}

func (c *Message) FromSia(sia sia.Sia) *Message {
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
