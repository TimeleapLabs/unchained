package rpc

import (
	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
)

type TextToImageRPCRequestParams struct {
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

type TextToImageRPCResponseParams struct {
	// The image in bytes
	Image []byte `json:"image"`
}

func (t *TextToImageRPCRequestParams) Sia() sia.Sia {
	return sia.New().
		AddString16(t.Prompt).
		AddString16(t.NegativePrompt).
		AddString8(t.Model).
		AddString8(t.LoraWeights).
		AddUInt8(t.Steps)
}

func (t *TextToImageRPCRequestParams) FromSiaBytes(bytes []byte) *TextToImageRPCRequestParams {
	s := sia.NewFromBytes(bytes)

	t.Prompt = s.ReadString16()
	t.NegativePrompt = s.ReadString16()
	t.Model = s.ReadString8()
	t.LoraWeights = s.ReadString8()
	t.Steps = s.ReadUInt8()

	return t
}

func (t *TextToImageRPCResponseParams) Sia() sia.Sia {
	return sia.New().
		AddByteArray32(t.Image)
}

func (t *TextToImageRPCResponseParams) FromSiaBytes(bytes []byte) *TextToImageRPCResponseParams {
	s := sia.NewFromBytes(bytes)
	t.Image = s.ReadByteArray32()

	return t
}
