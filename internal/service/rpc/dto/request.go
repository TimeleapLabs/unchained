package dto

import (
	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
	"github.com/google/uuid"
)

// RPCRequest is the request of a RPC request.
type RPCRequest struct {
	// The ID of the request
	ID uuid.UUID `json:"id"`
	// The plugin to be called
	Plugin string `json:"plugin"`
	// The method to be called
	Method string `json:"method"`
	// The timeout of the request
	Timeout int `json:"timeout"`
	// params to pass to the function
	Params []byte `json:"params"`
}

// NewRequest creates a new request with unique ID.
func NewRequest(method string, params []byte) RPCRequest {
	taskID, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	return RPCRequest{
		ID:     taskID,
		Method: method,
		Params: params,
	}
}

func (t *RPCRequest) Sia() sia.Sia {
	uuidBytes, err := t.ID.MarshalBinary()

	if err != nil {
		panic(err)
	}

	return sia.New().
		AddByteArray8(uuidBytes).
		AddString8(t.Plugin).
		AddString8(t.Method).
		AddInt64(int64(t.Timeout)).
		AddByteArrayN(t.Params)
}

func (t *RPCRequest) FromSiaBytes(bytes []byte) *RPCRequest {
	s := sia.NewFromBytes(bytes)

	uuidBytes := s.ReadByteArray8()
	err := t.ID.UnmarshalBinary(uuidBytes)
	if err != nil {
		panic(err)
		// return nil
	}

	t.Plugin = s.ReadString8()
	t.Method = s.ReadString8()
	t.Timeout = int(s.ReadInt64())
	t.Params = s.Bytes()[s.Offset():]

	return t
}
