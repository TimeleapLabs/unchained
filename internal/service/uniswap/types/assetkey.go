package types

import sia "github.com/pouya-eghbali/go-sia/v2/pkg"

type AssetKey struct {
	Token TokenKey
	Block uint64
}

func (a *AssetKey) Sia() sia.Sia {
	return sia.New().
		EmbedBytes(a.Token.Sia().Bytes()).
		AddUInt64(a.Block)
}

func (a *AssetKey) FromSia(sia sia.Sia) *AssetKey {
	a.Token.FromSia(sia)
	a.Block = sia.ReadUInt64()

	return a
}
