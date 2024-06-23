package types

import sia "github.com/pouya-eghbali/go-sia/v2/pkg"

type TokenKey struct {
	Name   string
	Pair   string
	Chain  string
	Delta  int64
	Invert bool
	Cross  string
}

func (t *TokenKey) Sia() sia.Sia {
	return sia.New().
		AddString8(t.Name).
		AddString8(t.Pair).
		AddString8(t.Chain).
		AddInt64(t.Delta).
		AddBool(t.Invert).
		AddString8(t.Cross)
}

func (t *TokenKey) FromSia(sia sia.Sia) *TokenKey {
	t.Name = sia.ReadString8()
	t.Pair = sia.ReadString8()
	t.Chain = sia.ReadString8()
	t.Delta = sia.ReadInt64()
	t.Invert = sia.ReadBool()
	t.Cross = sia.ReadString8()

	return t
}
