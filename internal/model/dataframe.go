package model

import "time"

type DataFrame struct {
	Hash      []byte
	Timestamp time.Time
	Data      interface{}
}
