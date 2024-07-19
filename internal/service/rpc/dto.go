package rpc

import (
	"github.com/google/uuid"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

type RpcRequest struct {
	// The ID of the request
	ID uuid.UUID `json:"id"`
	// The signature of the request
	Signature [48]byte
	// The method to be called
	// Payment information
	TxHash string `json:"txHash"`
	Method string `json:"method"`
}

type TextToImageRpcRequestPramas struct {
	// The text to be converted to an image
	Prompt         string `json:"prompt"`
	NegativePrompt string `json:"negativePrompt"`
	// The model to be used
	Model string `json:"model"`
	// The weights of the model
	LoraWeights string `json:"loraWeights"`
	// The number of steps to run
	Steps uint8 `json:"steps"`
}

type TextToImageRpcRequest struct {
	RpcRequest
	TextToImageRpcRequestPramas
}

type RpcResponse struct {
	// The ID of the request
	ID uuid.UUID `json:"id"`
	// The signature of the request
	Signature [48]byte
}

type TextToImageRpcResponseParams struct {
	// The image in bytes
	Image []byte `json:"image"`
}

type TextToImageRpcResponse struct {
	RpcResponse
	TextToImageRpcResponseParams
}

type RegisterFunction struct {
	Function string `json:"function"`
}

func (t *TextToImageRpcRequest) Sia() sia.Sia {
	uuidBytes, err := t.ID.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return sia.New().
		AddByteArray8(uuidBytes).
		AddByteArray8(t.Signature[:]).
		AddString8(t.TxHash).
		AddString8(t.Method).
		AddString16(t.Prompt).
		AddString16(t.NegativePrompt).
		AddString8(t.Model).
		AddString8(t.LoraWeights).
		AddUInt8(t.Steps)
}

func (t *TextToImageRpcRequest) FromSiaBytes(bytes []byte) *TextToImageRpcRequest {
	s := sia.NewFromBytes(bytes)

	uuidBytes := s.ReadByteArray8()
	t.ID.UnmarshalBinary(uuidBytes)

	t.Signature = [48]byte{}
	copy(t.Signature[:], s.ReadByteArray8())

	t.TxHash = s.ReadString8()
	t.Method = s.ReadString8()
	t.Prompt = s.ReadString16()
	t.NegativePrompt = s.ReadString16()
	t.Model = s.ReadString8()
	t.LoraWeights = s.ReadString8()
	t.Steps = s.ReadUInt8()

	return t
}

func (t *TextToImageRpcResponse) Sia() sia.Sia {
	uuidBytes, err := t.ID.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return sia.New().
		AddByteArray8(uuidBytes).
		AddByteArray8(t.Signature[:]).
		AddByteArray32(t.Image)
}

func (t *TextToImageRpcResponse) FromSiaBytes(bytes []byte) *TextToImageRpcResponse {
	s := sia.NewFromBytes(bytes)

	uuidBytes := s.ReadByteArray8()
	t.ID.UnmarshalBinary(uuidBytes)

	t.Signature = [48]byte{}
	copy(t.Signature[:], s.ReadByteArray8())

	t.Image = s.ReadByteArray32()

	return t
}

func (t *RegisterFunction) Sia() sia.Sia {
	return sia.New().
		AddString8(t.Function)
}

func (t *RegisterFunction) FromSiaBytes(bytes []byte) *RegisterFunction {
	s := sia.NewFromBytes(bytes)

	t.Function = s.ReadString8()

	return t
}
