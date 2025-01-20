package packet

import (
	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"golang.org/x/crypto/ed25519"
)

type Packet struct {
	Message   []byte
	Signer    ed25519.PublicKey
	Signature [64]byte
}

func New(message []byte) *Packet {
	p := &Packet{
		Signer:  crypto.Identity.Ed25519.PublicKey,
		Message: message,
	}

	return p.MustSign()
}

func (p *Packet) Sia() sia.Sia {
	return sia.New().
		AddByteArrayN(p.Message).
		AddByteArrayN(p.Signer).
		AddByteArrayN(p.Signature[:])
}

func (p *Packet) FromSia(sia sia.Sia) *Packet {
	length := len(sia.Bytes())
	messageLength := uint64(length - 32 - 64)

	copy(p.Message, sia.ReadByteArrayN(messageLength))
	copy(p.Signer, sia.ReadByteArrayN(32))
	copy(p.Signature[:], sia.ReadByteArrayN(64))

	return p
}

func (p *Packet) FromBytes(data []byte) *Packet {
	return p.FromSia(sia.NewFromBytes(data))
}

func (p *Packet) Sign() (*Packet, error) {
	signature, err := crypto.Identity.Ed25519.Sign(p.Message)
	if err != nil {
		return p, err
	}

	copy(p.Signature[:], signature)
	return p, nil
}

func (p *Packet) MustSign() *Packet {
	signed, err := p.Sign()
	if err != nil {
		utils.Logger.
			With("Error", err).
			Error("Cannot sign packet")
		panic(err)
	}

	return signed
}

func (p *Packet) IsValid() bool {
	return crypto.Identity.Ed25519.Verify(p.Signer, p.Message, p.Signature[:])
}
