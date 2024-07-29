package dto

import (
	"github.com/google/uuid"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

// RpcResponse is the response of a RPC request
type RpcResponse struct {
	// The ID of the request
	ID uuid.UUID `json:"id"`
	// The signature of the request
	Signature [48]byte
}

func (t *RpcResponse) Sia() sia.Sia {
	uuidBytes, err := t.ID.MarshalBinary()

	if err != nil {
		panic(err)
	}
	return sia.New().
		AddByteArray8(uuidBytes).
		AddByteArray8(t.Signature[:])
}

func (t *RpcResponse) FromSiaBytes(bytes []byte) *RpcResponse {
	s := sia.NewFromBytes(bytes)

	uuidBytes := s.ReadByteArray8()
	err := t.ID.UnmarshalBinary(uuidBytes)
	if err != nil {
		return nil
	}

	t.Signature = [48]byte{}
	copy(t.Signature[:], s.ReadByteArray8())

	return t
}
