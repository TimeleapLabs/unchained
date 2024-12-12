package dto

import sia "github.com/TimeleapLabs/go-sia/v2/pkg"

// RegisterFunction is a DTO for registering a function.
type RegisterFunction struct {
	Function string `json:"function"`
	Runtime  string `json:"runtime"`
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
