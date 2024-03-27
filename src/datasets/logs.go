package datasets

import (
	"encoding/json"
	"github.com/KenshiTech/unchained/constants"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

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

type EventLogArgs []*EventLogArg

func (m EventLogArgs) Array() []EventLogArg {
	result := []EventLogArg{}
	for _, i := range m {
		result = append(result, *i)
	}

	return result
}

func ConvertInterfaceToAny(v interface{}) (*anypb.Any, error) {
	anyValue := &anypb.Any{}

	bytes, err := json.Marshal(v)
	if err != nil {
		log.Err(err)
		return nil, constants.ErrInternalError
	}
	bytesValue := wrapperspb.Bytes(bytes)

	err = anypb.MarshalFrom(anyValue, bytesValue, proto.MarshalOptions{})
	if err != nil {
		log.Err(err)
		return nil, constants.ErrInternalError
	}

	return anyValue, nil
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
