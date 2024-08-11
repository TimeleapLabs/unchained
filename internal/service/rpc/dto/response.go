package dto

import (
	"github.com/google/uuid"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

// RPCResponse is the response of a RPC request.
type RPCResponse struct {
	// The ID of the request
	ID uuid.UUID `json:"id"`
	// The signature of the request
	Signature [48]byte
	// The response of the function
	Response []byte `json:"response"`
	// The error of the function
	Error string `json:"error"`
}

func (t *RPCResponse) Sia() sia.Sia {
	uuidBytes, err := t.ID.MarshalBinary()

	if err != nil {
		panic(err)
	}
	return sia.New().
		AddByteArray8(uuidBytes).
		AddByteArray8(t.Signature[:]).
		AddByteArray8(t.Response).
		AddString8(t.Error)
}

func (t *RPCResponse) FromSiaBytes(bytes []byte) *RPCResponse {
	s := sia.NewFromBytes(bytes)

	uuidBytes := s.ReadByteArray8()
	err := t.ID.UnmarshalBinary(uuidBytes)
	if err != nil {
		return nil
	}

	t.Signature = [48]byte{}
	copy(t.Signature[:], s.ReadByteArray8())

	t.Response = s.ReadByteArray8()

	t.Error = s.ReadString8()

	return t
}
