package client

import (
	"github.com/KenshiTech/unchained/internal/consts"
	"github.com/KenshiTech/unchained/internal/transport/client/conn"
	"github.com/KenshiTech/unchained/internal/transport/client/handler"
	"github.com/KenshiTech/unchained/internal/utils"
)

func NewRPC(handler handler.Handler) {
	incoming := conn.Read()

	go func() {
		utils.Logger.Info("Starting consumer from broker")

		for payload := range incoming {
			switch consts.OpCode(payload[0]) {
			case consts.OpCodeError:
				utils.Logger.
					With("Error", string(payload[1:])).
					Error("Broker")

			case consts.OpCodeFeedback:
				utils.Logger.
					With("Feedback", string(payload[1:])).
					Info("Broker")

			case consts.OpCodeKoskChallenge:
				challenge := handler.Challenge(payload[1:])
				conn.Send(consts.OpCodeKoskResult, challenge.Sia().Content)

			case consts.OpCodePriceReportBroadcast:
				go handler.PriceReport(payload[1:])

			case consts.OpCodeEventLogBroadcast:
				go handler.EventLog(payload[1:])

			case consts.OpCodeCorrectnessReportBroadcast:
				go handler.CorrectnessReport(payload[1:])

			default:
				utils.Logger.
					With("Code", payload[0]).
					Error("Unknown call code")
			}
		}
	}()
}
