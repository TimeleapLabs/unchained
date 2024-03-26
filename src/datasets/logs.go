package datasets

import (
	"github.com/KenshiTech/unchained/constants"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

//type EventLogArg struct {
//	Name  string
//	Type  string
//	Value any
//}
//
//var _ msgpack.CustomEncoder = (*EventLogArg)(nil)
//
//// TODO: this can be improved
//func (eventLog *EventLogArg) EncodeMsgpack(enc *msgpack.Encoder) error {
//	encoded, err := json.Marshal(eventLog)
//	if err != nil {
//		return err
//	}
//	return enc.EncodeBytes(encoded)
//}
//
//var _ msgpack.CustomDecoder = (*EventLogArg)(nil)
//
//func (eventLog *EventLogArg) DecodeMsgpack(dec *msgpack.Decoder) error {
//	bytes, err := dec.DecodeBytes()
//	if err != nil {
//		return err
//	}
//	return json.Unmarshal(bytes, eventLog)
//}

func (m *EventLog) Protobuf() ([]byte, error) {
	protoModel, err := proto.Marshal(m)
	if err != nil {
		log.Err(err)
		return nil, constants.ErrInternalError
	}

	return protoModel, nil
}

func (m *EventLogArg) Protobuf() ([]byte, error) {
	protoModel, err := proto.Marshal(m)
	if err != nil {
		log.Err(err)
		return nil, constants.ErrInternalError
	}

	return protoModel, nil
}

func (m *EventLogReport) Protobuf() ([]byte, error) {
	protoModel, err := proto.Marshal(m)
	if err != nil {
		log.Err(err)
		return nil, constants.ErrInternalError
	}

	return protoModel, nil
}

func (m *BroadcastEventPacket) Protobuf() ([]byte, error) {
	protoModel, err := proto.Marshal(m)
	if err != nil {
		log.Err(err)
		return nil, constants.ErrInternalError
	}

	return protoModel, nil
}

//type EventLog struct {
//	LogIndex uint64
//	Block    uint64
//	Address  string
//	Event    string
//	Chain    string
//	TxHash   [32]byte
//	Args     []EventLogArg
//}
//
//type EventLogReport struct {
//	EventLog
//	Signature [48]byte
//}
//
//type BroadcastEventPacket struct {
//	Info      EventLog
//	Signature [48]byte
//	Signer    bls.Signer
//}
