package attestation

import (
	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
	"github.com/TimeleapLabs/timeleap/internal/model"
)

type Key struct {
	Hash  string
	Topic string
	Meta  string
}

type Signature struct {
	Signature [64]byte
	Signer    model.Signer
}

type SaveSignatureArgs struct {
	Info model.Attestation
	Hash [32]byte
}

func (s *Signature) Sia() sia.Sia {
	return sia.New().
		AddByteArray8(s.Signature[:]).
		EmbedBytes(s.Signer.Sia().Bytes())
}

func (s *Signature) FromBytes(payload []byte) *Signature {
	siaMessage := sia.NewFromBytes(payload)
	return s.FromSia(siaMessage)
}

func (s *Signature) FromSia(sia sia.Sia) *Signature {
	signature := sia.ReadByteArray8()
	copy(s.Signature[:], signature)

	s.Signer.FromSia(sia)
	return s
}
