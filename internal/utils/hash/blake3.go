package hash

import (
	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
	"github.com/zeebo/blake3"
)

type Hashable interface {
	Sia() sia.Sia
}

func Hash(data Hashable) [32]byte {
	return blake3.Sum256(data.Sia().Bytes())
}
