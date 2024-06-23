package model

import "time"

type DataFrame struct {
	Hash      []byte      `bson:"hash"      json:"hash"`
	Timestamp time.Time   `bson:"timestamp" json:"timestamp"`
	Data      interface{} `bson:"data"      gorm:"embedded"  json:"data"`
}
