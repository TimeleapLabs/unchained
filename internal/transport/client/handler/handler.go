package handler

import (
	"github.com/TimeleapLabs/unchained/internal/crypto/kosk"
)

type Handler interface {
	Challenge(message []byte) *kosk.Challenge
	CorrectnessReport(message []byte)
	EventLog(message []byte)
	PriceReport(message []byte)
}
