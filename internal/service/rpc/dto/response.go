package dto

import (
	"github.com/google/uuid"
	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
)

// RPCResponse is the response of a RPC request.
type RPCResponse struct {
	// The ID of the request
	ID uuid.UUID `json:"id"`
	// The error of the function
	Error uint64 `json:"error"`
	// The response of the function
	Response []byte `json:"response"`
}

func (t *RPCResponse) Sia() sia.Sia {
	uuidBytes, err := t.ID.MarshalBinary()
	if err != nil {
		panic(err)
	}

	return sia.New().
		AddByteArray8(uuidBytes).
		AddUInt64(t.Error).
		EmbedBytes(t.Response)
}

func (t *RPCResponse) FromSiaBytes(bytes []byte) *RPCResponse {
	s := sia.NewFromBytes(bytes)

	uuidBytes := s.ReadByteArray8()
	err := t.ID.UnmarshalBinary(uuidBytes)
	if err != nil {
		return nil
	}

	t.Error = s.ReadUInt64()
	t.Response = s.Bytes()[s.Offset():]

	return t
}
