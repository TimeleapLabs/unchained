package model

import (
	"encoding/json"
	"math/big"

	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/utils"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"

	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

type EventLogArg struct {
	Name  string `json:"Name"`
	Type  string `json:"Type"`
	Value any    `json:"Value"`
}

type EventLog struct {
	LogIndex uint64
	Block    uint64
	Address  string
	Event    string
	Chain    string
	TxHash   [32]byte
	Args     []EventLogArg

	Consensus    bool
	SignersCount uint64
	SignerIDs    []int
	Signers      []Signer
	Signature    []byte
	Voted        *big.Int
}

func (e *EventLog) Sia() sia.Sia {
	argsEncoded, err := json.Marshal(e.Args)

	if err != nil {
		panic(err)
	}

	return sia.New().
		AddUInt64(e.LogIndex).
		AddUInt64(e.Block).
		AddString8(e.Address).
		AddString8(e.Event).
		AddString8(e.Chain).
		AddByteArray8(e.TxHash[:]).
		AddByteArray16(argsEncoded)
}

func (e *EventLog) FromBytes(payload []byte) *EventLog {
	siaMessage := sia.NewFromBytes(payload)
	return e.FromSia(siaMessage)
}

func (e *EventLog) FromSia(sia sia.Sia) *EventLog {
	e.LogIndex = sia.ReadUInt64()
	e.Block = sia.ReadUInt64()
	e.Address = sia.ReadString8()
	e.Event = sia.ReadString8()
	e.Chain = sia.ReadString8()
	copy(e.TxHash[:], sia.ReadByteArray8())

	argsEncoded := sia.ReadByteArray16()
	err := json.Unmarshal(argsEncoded, &e.Args)

	if err != nil {
		panic(err)
	}

	return e
}

func (e *EventLog) Bls() (bls12381.G1Affine, error) {
	hash, err := bls.Hash(e.Sia().Bytes())
	if err != nil {
		utils.Logger.Error("Can't hash bls: %v", err)
		return bls12381.G1Affine{}, err
	}

	return hash, err
}
