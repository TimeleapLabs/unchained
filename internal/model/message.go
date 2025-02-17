package model

import (
	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
	"github.com/google/uuid"
)

type Message struct {
	ID        uuid.UUID
	Timestamp uint64
	Topic     string
	Meta      []byte
}

func (c *Message) Sia() sia.Sia {
	return sia.New().
		AddByteArray8(c.ID[:]).
		AddUInt64(c.Timestamp).
		AddString16(c.Topic).
		AddByteArray32(c.Meta)
}

func (c *Message) FromBytes(payload []byte) *Message {
	siaMessage := sia.NewFromBytes(payload)
	return c.FromSia(siaMessage)
}

func (c *Message) FromSia(sia sia.Sia) *Message {
	c.ID = uuid.UUID(sia.ReadByteArray8())
	c.Timestamp = sia.ReadUInt64()
	c.Topic = sia.ReadString16()
	c.Meta = sia.ReadByteArray32()
	return c
}
