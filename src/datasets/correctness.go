package datasets

import (
	"github.com/KenshiTech/unchained/crypto/bls"
)

type Correctness struct {
	Timestamp uint64
	Hash      [64]byte
	Topic     [64]byte
	Correct   bool
}

type CorrectnessReport struct {
	Correctness
	Signature [48]byte
}

type BroadcastCorrectnessPacket struct {
	Info      Correctness
	Signature [48]byte
	Signer    bls.Signer
}
