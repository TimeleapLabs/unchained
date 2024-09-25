package model

import (
	"encoding/hex"

	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/utils"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Signers []Signer

type Signer struct {
	ID    uint               `bson:"-"             gorm:"primarykey"`
	DocID primitive.ObjectID `bson:"_id,omitempty" gorm:"-"`

	Name           string `json:"name"`
	EvmAddress     string `json:"evm_address"`
	PublicKey      string `json:"public_key"`
	ShortPublicKey string `json:"short_public_key"`
}

func (s *Signer) Sia() sia.Sia {
	publicKeyBytes, err := hex.DecodeString(s.PublicKey)
	if err != nil {
		utils.Logger.Error("Can't decode public key: %v", err)
		return sia.New()
	}

	shortPublicKey, err := hex.DecodeString(s.ShortPublicKey)
	if err != nil {
		utils.Logger.Error("Can't decode short public key: %v", err)
		return sia.New()
	}

	return sia.New().
		AddString8(s.Name).
		AddString8(s.EvmAddress).
		AddByteArray8(publicKeyBytes).
		AddByteArray8(shortPublicKey)
}

func (s *Signer) FromBytes(payload []byte) *Signer {
	siaMessage := sia.NewFromBytes(payload)
	return s.FromSia(siaMessage)
}

func (s *Signer) FromSia(sia sia.Sia) *Signer {
	s.Name = sia.ReadString8()
	s.EvmAddress = sia.ReadString8()
	s.PublicKey = hex.EncodeToString(sia.ReadByteArray8())
	s.ShortPublicKey = hex.EncodeToString(sia.ReadByteArray8())

	return s
}

func (s *Signer) Bls() bls12381.G1Affine {
	hash, err := bls.Hash(s.Sia().Bytes())
	if err != nil {
		utils.Logger.Error("Can't hash bls: %v", err)
		return bls12381.G1Affine{}
	}

	return hash
}

func (s Signers) Bls() []byte {
	bytes := []byte{}

	for _, signer := range s {
		hash, err := bls.Hash(signer.Sia().Bytes())
		if err != nil {
			utils.Logger.Error("Can't hash bls: %v", err)
			return bytes
		}

		hashBytes := hash.Bytes()
		bytes = append(bytes, hashBytes[:]...)
	}

	return bytes
}
