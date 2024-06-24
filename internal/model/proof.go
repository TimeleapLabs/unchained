package model

import (
	"time"

	"gorm.io/gorm"
)

type Proof struct {
	gorm.Model

	Hash      []byte    `bson:"hash"      json:"hash"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	Signature [48]byte  `bson:"signature" json:"signature"`
	Signers   []Signer  `bson:"signers"   json:"signers"`
	// list of signers (relations in Postgres, and an array in Mongo)
}
