package model

import (
	"time"

	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/utils"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssetPriceDataFrame struct {
	ID    uint               `bson:"-"             gorm:"primarykey"`
	DocID primitive.ObjectID `bson:"_id,omitempty" gorm:"-"`

	Hash      string     `bson:"hash"      json:"hash"`
	Timestamp time.Time  `bson:"timestamp" json:"timestamp"`
	Data      AssetPrice `bson:"data"      gorm:"embedded"  json:"data"`
}

type AssetPrice struct {
	Pair         string `gorm:"uniqueIndex:idx_pair_name_chain_block" json:"pair"`
	Name         string `gorm:"uniqueIndex:idx_pair_name_chain_block" json:"name"`
	Chain        string `gorm:"uniqueIndex:idx_pair_name_chain_block" json:"chain"`
	Block        uint64 `gorm:"uniqueIndex:idx_pair_name_chain_block" json:"block"`
	Price        int64  `json:"price"`
	SignersCount uint64 `json:"signers_count"`
	Signature    []byte `json:"signature"`
	Consensus    bool   `json:"consensus"`
	Voted        int64  `json:"voted"`
}

func (c *AssetPrice) Sia() sia.Sia {
	return sia.New().
		AddString8(c.Pair).
		AddString8(c.Name).
		AddString8(c.Chain).
		AddUInt64(c.Block).
		AddInt64(c.Price).
		AddUInt64(c.SignersCount).
		AddByteArray8(c.Signature).
		AddBool(c.Consensus).
		AddInt64(c.Voted)
}

func (c *AssetPrice) FromBytes(payload []byte) *AssetPrice {
	siaMessage := sia.NewFromBytes(payload)
	return c.FromSia(siaMessage)
}

func (c *AssetPrice) FromSia(siaObj sia.Sia) *AssetPrice {
	c.Pair = siaObj.ReadString8()
	c.Name = siaObj.ReadString8()
	c.Chain = siaObj.ReadString8()
	c.Block = siaObj.ReadUInt64()
	c.Price = siaObj.ReadInt64()
	c.SignersCount = siaObj.ReadUInt64()
	copy(c.Signature, siaObj.ReadByteArray8())
	c.Consensus = siaObj.ReadBool()
	c.Voted = siaObj.ReadInt64()

	return c
}

func (c *AssetPrice) Bls() *bls12381.G1Affine {
	hash, err := bls.Hash(c.Sia().Bytes())
	if err != nil {
		utils.Logger.Error("Can't hash bls: %v", err)
		return &bls12381.G1Affine{}
	}

	return &hash
}
