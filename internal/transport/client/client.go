package client

import (
	"context"
	"crypto/ed25519"
	"time"

	"github.com/TimeleapLabs/timeleap/internal/config"
	"github.com/TimeleapLabs/timeleap/internal/consts"
	"github.com/TimeleapLabs/timeleap/internal/transport/client/conn"
	"github.com/TimeleapLabs/timeleap/internal/transport/client/handler"
	"github.com/TimeleapLabs/timeleap/internal/transport/server/packet"
	"github.com/TimeleapLabs/timeleap/internal/utils"
	"github.com/btcsuite/btcutil/base58"
)

// NewRPC is a function that starts a new RPC client and connect to broker to consume events.
func NewRPC(handler handler.Handler) {
	incoming := conn.Read()

	brokerPubKeyBytes := base58.Decode(config.App.Network.Broker.PublicKey)
	brokerPubKey := ed25519.PublicKey(brokerPubKeyBytes)

	go func() {
		utils.Logger.Info("RPC client started")

		for payload := range incoming {
			go func(payload []byte) {
				ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
				defer cancel()

				p := new(packet.Packet).FromBytes(payload)

				if !p.IsValid() {
					utils.Logger.Error("Invalid packet")
					return
				}

				// verify sender is the broker specified in the config
				if !p.Signer.Equal(brokerPubKey) {
					utils.Logger.
						With("Signer", p.Signer).
						Error("Invalid signer")
					return
				}

				switch consts.OpCode(p.Message[0]) {
				case consts.OpCodeError:
					utils.Logger.
						With("Error", string(p.Message[1:])).
						Error("Broker")

				case consts.OpCodeFeedback:
					utils.Logger.
						With("Feedback", string(p.Message[1:])).
						Info("Broker")

				case consts.OpCodeAttestation:
					handler.Attestation(ctx, p.Message[1:])

				case consts.OpCodeRPCRequest:
					handler.RPCRequest(ctx, p.Message[1:])

				default:
					utils.Logger.
						With("Code", p.Message[0]).
						Error("Unknown call code")
				}
			}(payload)
		}
	}()
}
