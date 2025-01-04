package dto

import sia "github.com/pouya-eghbali/go-sia/v2/pkg"

// RegisterFunction is a DTO for registering a function.
type RegisterFunction struct {
	Plugin    string   `json:"plugin"`
	Functions []string `json:"function"`
	Runtime   string   `json:"runtime"`
}

func (t *RegisterFunction) Sia() sia.Sia {
	return sia.New().
		AddString8(t.Plugin).
		AddByteArray64(
			sia.NewSiaArray[string]().
				AddArray16(t.Functions, func(s *sia.ArraySia[string], v string) {
					s.AddString8(v)
				}).
				Bytes())
}

func (t *RegisterFunction) FromSiaBytes(bytes []byte) *RegisterFunction {
	s := sia.NewFromBytes(bytes)

	t.Plugin = s.ReadString8()
	t.Functions = sia.NewArrayFromBytes[string](s.ReadByteArray64()).
		ReadArray16(func(s *sia.ArraySia[string]) string {
			return s.ReadString8()
		})

	return t
}
