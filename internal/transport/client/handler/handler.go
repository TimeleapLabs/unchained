package handler

import (
	"github.com/KenshiTech/unchained/internal/model"
)

type Handler interface {
	Challenge(message []byte) *model.ChallengePacket
	CorrectnessReport(message []byte)
	EventLog(message []byte)
	PriceReport(message []byte)
}
