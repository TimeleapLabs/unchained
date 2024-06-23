package correctness

import (
	"github.com/TimeleapLabs/unchained/internal/model"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

type Key struct {
	Hash    string
	Topic   string
	Correct bool
}

type Signature struct {
	Signature bls12381.G1Affine
	Signer    model.Signer
}

func (s *Signature) Sia() sia.Sia {
	return sia.New().
		AddByteArray8(s.Signature.Marshal()).
		EmbedBytes(s.Signer.Sia().Bytes())
}

func (s *Signature) FromBytes(payload []byte) *Signature {
	siaMessage := sia.NewFromBytes(payload)
	return s.FromSia(siaMessage)
}

func (s *Signature) FromSia(sia sia.Sia) *Signature {
	err := s.Signature.Unmarshal(sia.ReadByteArray8())

	if err != nil {
		s.Signature = bls12381.G1Affine{}
	}

	s.Signer.FromSia(sia)

	return s
}
