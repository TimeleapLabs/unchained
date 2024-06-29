package model

import (
	"encoding/json"
	"time"

	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"

	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

type EventLogArg struct {
	Name  string `json:"Name"`
	Type  string `json:"Type"`
	Value any    `json:"Value"`
}

type EventLogDataFrame struct {
	ID    uint               `bson:"-"             gorm:"primarykey"`
	DocID primitive.ObjectID `bson:"_id,omitempty" gorm:"-"`

	Hash      string    `bson:"hash"      json:"hash"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	Data      EventLog  `bson:"data"      gorm:"embedded"  json:"data"`
}

type EventLog struct {
	LogIndex uint64
	Block    uint64 `gorm:"uniqueIndex:idx_block_tx_index"`
	Address  string
	Event    string
	Chain    string
	TxHash   [32]byte      `gorm:"uniqueIndex:idx_block_tx_index"`
	Args     []EventLogArg `gorm:"type:jsonb"`

	Consensus    bool
	SignersCount uint64
	Signature    []byte
	Voted        int64
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

func (e *EventLog) Bls() *bls12381.G1Affine {
	hash, err := bls.Hash(e.Sia().Bytes())
	if err != nil {
		utils.Logger.Error("Can't hash bls: %v", err)
		return &bls12381.G1Affine{}
	}

	return &hash
}
