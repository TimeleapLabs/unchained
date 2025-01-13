package dto

import (
	"github.com/google/uuid"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
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
	// The signature of the request
	Signature [48]byte `json:"signature"`
}

// NewRequest creates a new request with unique ID.
func NewRequest(method string, params []byte, signature [48]byte) RPCRequest {
	taskID, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	return RPCRequest{
		ID:        taskID,
		Method:    method,
		Params:    params,
		Signature: signature,
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
		AddByteArray64(t.Params).
		AddByteArray8(t.Signature[:])
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
	t.Params = s.ReadByteArray64()

	t.Signature = [48]byte{}
	copy(t.Signature[:], s.ReadByteArray8())

	return t
}
