package client

import (
	"context"
	"time"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
	"github.com/TimeleapLabs/unchained/internal/transport/client/handler"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

// NewRPC is a function that starts a new RPC client and connect to broker to consume events.
func NewRPC(handler handler.Handler) {
	incoming := conn.Read()

	go func() {
		utils.Logger.Info("Starting consumer from broker")

		for payload := range incoming {
			go func(payload []byte) {
				ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
				defer cancel()

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
					conn.Send(consts.OpCodeKoskResult, challenge)
				case consts.OpCodeAttestationBroadcast:
					handler.Attestation(ctx, payload[1:])
				case consts.OpCodeRPCRequest:
					handler.RPCRequest(ctx, payload[1:])
				default:
					utils.Logger.
						With("Code", payload[0]).
						Error("Unknown call code")
				}
			}(payload)
		}
	}()
}
