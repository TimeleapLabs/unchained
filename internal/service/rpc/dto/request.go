package dto

import (
	"github.com/google/uuid"
	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
)

// RPCRequest is the request of a RPC request.
type RPCRequest struct {
	// The ID of the request
	ID uuid.UUID `json:"id"`
	// The signature of the request
	Signature [48]byte `json:"signature"`
	// Payment information
	TxHash string `json:"tx_hash"`
	// The method to be called
	Method string `json:"method"`
	// params to pass to the function
	Params []byte `json:"params"`
}

// NewRequest creates a new request with unique ID.
func NewRequest(method string, params []byte, signature [48]byte, txHash string) RPCRequest {
	taskID, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	return RPCRequest{
		ID:        taskID,
		Method:    method,
		Params:    params,
		Signature: signature,
		TxHash:    txHash,
	}
}

func (t *RPCRequest) Sia() sia.Sia {
	uuidBytes, err := t.ID.MarshalBinary()

	if err != nil {
		panic(err)
	}

	return sia.New().
		AddByteArray8(uuidBytes).
		AddByteArray8(t.Signature[:]).
		AddString8(t.TxHash).
		AddString8(t.Method).
		EmbedBytes(t.Params)
}

func (t *RPCRequest) FromSiaBytes(bytes []byte) *RPCRequest {
	s := sia.NewFromBytes(bytes)

	uuidBytes := s.ReadByteArray8()
	err := t.ID.UnmarshalBinary(uuidBytes)
	if err != nil {
		panic(err)
		// return nil
	}

	t.Signature = [48]byte{}
	copy(t.Signature[:], s.ReadByteArray8())

	t.TxHash = s.ReadString8()
	t.Method = s.ReadString8()
	t.Params = s.Bytes()[s.Offset():]

	return t
}
