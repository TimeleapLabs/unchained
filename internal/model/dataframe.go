package model

import (
	"time"

	"gorm.io/gorm"
)

type DataFrame struct {
	gorm.Model
	Hash      []byte      `bson:"hash"      json:"hash"`
	Timestamp time.Time   `bson:"timestamp" json:"timestamp"`
	Data      interface{} `bson:"data"      gorm:"embedded"  json:"data"`
}
