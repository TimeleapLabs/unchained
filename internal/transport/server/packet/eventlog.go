package packet

import (
	"github.com/TimeleapLabs/unchained/internal/model"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

type EventLogReportPacket struct {
	model.EventLog
	Signature [48]byte
}

func (e *EventLogReportPacket) Sia() sia.Sia {
	return sia.New().
		EmbedBytes(e.EventLog.Sia().Bytes()).
		AddByteArray8(e.Signature[:])
}

func (e *EventLogReportPacket) FromBytes(payload []byte) *EventLogReportPacket {
	siaMessage := sia.NewFromBytes(payload)
	e.EventLog.FromSia(siaMessage)
	copy(e.Signature[:], siaMessage.ReadByteArray8())

	return e
}

type BroadcastEventPacket struct {
	Info      model.EventLog
	Signature [48]byte
	Signer    model.Signer
}

func (b *BroadcastEventPacket) Sia() sia.Sia {
	return sia.New().
		EmbedBytes(b.Info.Sia().Bytes()).
		AddByteArray8(b.Signature[:]).
		EmbedBytes(b.Signer.Sia().Bytes())
}

func (b *BroadcastEventPacket) FromBytes(payload []byte) *BroadcastEventPacket {
	siaMessage := sia.NewFromBytes(payload)

	b.Info.FromSia(siaMessage)
	copy(b.Signature[:], siaMessage.ReadByteArray8())
	b.Signer.FromSia(siaMessage)

	return b
}
