package dto

import (
	"github.com/google/uuid"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

// RpcRequest is the request of a RPC request
type RpcRequest struct {
	// The ID of the request
	ID uuid.UUID `json:"id"`
	// The signature of the request
	Signature [48]byte
	// The method to be called
	// Payment information
	TxHash string `json:"txHash"`
	Method string `json:"method"`
	Params []byte `json:"params"`
}

func (t *RpcRequest) Sia() sia.Sia {
	uuidBytes, err := t.ID.MarshalBinary()

	if err != nil {
		panic(err)
	}
	return sia.New().
		AddByteArray8(uuidBytes).
		AddByteArray8(t.Signature[:]).
		AddString8(t.TxHash).
		AddString8(t.Method).
		AddByteArray8(t.Params)
}

func (t *RpcRequest) FromSiaBytes(bytes []byte) *RpcRequest {
	s := sia.NewFromBytes(bytes)

	uuidBytes := s.ReadByteArray8()
	err := t.ID.UnmarshalBinary(uuidBytes)
	if err != nil {
		return nil
	}

	t.Signature = [48]byte{}
	copy(t.Signature[:], s.ReadByteArray8())

	t.TxHash = s.ReadString8()
	t.Method = s.ReadString8()
	t.Params = s.ReadByteArray8()

	return t
}
