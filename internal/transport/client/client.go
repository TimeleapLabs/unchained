package client

import (
	"github.com/KenshiTech/unchained/internal/constants/opcodes"
	"github.com/KenshiTech/unchained/internal/log"
	"github.com/KenshiTech/unchained/internal/transport/client/conn"
	"github.com/KenshiTech/unchained/internal/transport/client/handler"
)

func Consume(handler handler.Handler) {
	incoming := conn.Read()

	go func() {
		log.Logger.Info("Starting consumer from broker")

		for payload := range incoming {
			switch opcodes.OpCode(payload[0]) {
			case opcodes.Error:
				log.Logger.
					With("Error", string(payload[1:])).
					Error("Broker")

			case opcodes.Feedback:
				log.Logger.
					With("Feedback", string(payload[1:])).
					Warn("Broker")

			case opcodes.KoskChallenge:
				challenge := handler.Challenge(payload[1:])
				conn.Send(opcodes.KoskResult, challenge.Sia().Content)

			case opcodes.PriceReportBroadcast:
				go handler.PriceReport(payload[1:])

			case opcodes.EventLogBroadcast:
				go handler.EventLog(payload[1:])

			case opcodes.CorrectnessReportBroadcast:
				go handler.CorrectnessReport(payload[1:])

			default:
				log.Logger.
					With("Code", payload[0]).
					Info("Unknown call code")
			}
		}
	}()
}
