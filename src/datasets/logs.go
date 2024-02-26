package datasets

import (
	"encoding/json"

	"github.com/vmihailenco/msgpack/v5"
)

type EventLogArg struct {
	Name  string
	Type  string
	Value any
}

var _ msgpack.CustomEncoder = (*EventLogArg)(nil)

// TODO: this can be improved
func (eventLog *EventLogArg) EncodeMsgpack(enc *msgpack.Encoder) error {
	encoded, err := json.Marshal(eventLog)
	if err != nil {
		return err
	}
	return enc.EncodeBytes(encoded)
}

var _ msgpack.CustomDecoder = (*EventLogArg)(nil)

func (eventLog *EventLogArg) DecodeMsgpack(dec *msgpack.Decoder) error {
	bytes, err := dec.DecodeBytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, eventLog)
}

type EventLog struct {
	LogIndex uint64
	Block    uint64
	Address  string
	Event    string
	Chain    string
	TxHash   [32]byte
	Args     []EventLogArg
}

type EventLogReport struct {
	EventLog
	Signature [48]byte
}

type BroadcastEventPacket struct {
	Info      EventLog
	Signature [48]byte
	Signers   [][]byte
}
