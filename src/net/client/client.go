package client

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/KenshiTech/unchained/datasets"

	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/constants"
	"github.com/KenshiTech/unchained/constants/opcodes"
	"github.com/KenshiTech/unchained/consumers"
	"github.com/KenshiTech/unchained/crypto/bls"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/net/shared"

	"github.com/gorilla/websocket"
)

var Done chan struct{}

func StartClient() {
	if !config.Config.IsSet("broker.uri") {
		return
	}

	brokerURL := fmt.Sprintf(
		"%s/%s",
		config.Config.GetString("broker.uri"),
		constants.ProtocolVersion,
	)

	var err error
	shared.Client, _, err = websocket.DefaultDialer.Dial(brokerURL, nil)
	if err != nil {
		panic(err)
	}

	Done = make(chan struct{})

	hello := bls.ClientSigner
	helloPayload, err := datasets.NewSigner(&hello).Protobuf()
	if err != nil {
		panic(err)
	}

	go func() {
		defer close(Done)

		for {
			_, payload, err := shared.Client.ReadMessage()
			if err != nil || payload[0] == byte(opcodes.Error) {
				if err != nil {
					log.Logger.
						With("Error", err).
						Error("Read error")
				} else {
					log.Logger.
						With("Error", string(payload[1:])).
						Error("Broker error")
				}

				ReConnect(err, brokerURL, helloPayload)
				if shared.IsClientSocketClosed {
					return
				}

				continue
			}

			switch opcodes.OpCode(payload[0]) {
			// TODO: Make a table of call codes
			case opcodes.Feedback:
				log.Logger.
					With("Feedback", string(payload[1:])).
					Info("Broker")

			case opcodes.KoskChallenge:
				// TODO: Refactor into a function
				// TODO: Check for errors!
				challenge, err := datasets.NewChallenge(payload[1:])
				if err != nil {
					log.Logger.
						With("Error", err).
						Error("Can't unmarshal challenge")
					continue
				}

				signature, _ := bls.Sign(*bls.ClientSecretKey, challenge.Random)
				signatureByte := signature.Bytes()
				challenge.Signature = signatureByte[:]

				koskPayload, err := challenge.Protobuf()
				if err != nil {
					log.Logger.
						With("Error", err).
						Error("Can't marshal challenge")
					continue
				}

				shared.Send(opcodes.KoskResult, koskPayload)
			case opcodes.PriceReportBroadcast:
				go consumers.ConsumePriceReport(payload[1:])

			case opcodes.EventLogBroadcast:
				go consumers.ConsumeEventLog(payload[1:])

			case opcodes.CorrectnessReportBroadcast:
				go consumers.ConsumeCorrectnessReport(payload[1:])

			default:
				log.Logger.
					With("Code", payload[0]).
					Info("Unknown call code")
			}
		}
	}()

	shared.Send(opcodes.Hello, helloPayload)
}

func ReConnect(err error, brokerURL string, helloMessageByte []byte) {
	shared.IsClientSocketClosed = true

	if websocket.IsUnexpectedCloseError(err) {
		for i := 1; i < 6; i++ {
			time.Sleep(time.Duration(i) * 3 * time.Second)
			shared.Client, _, err = websocket.DefaultDialer.Dial(brokerURL, nil)
			if err == nil {
				shared.IsClientSocketClosed = false
				shared.Send(opcodes.Hello, helloMessageByte)
			}
		}
	}
}

func closeConnection() {
	if shared.Client != nil && config.Config.IsSet("broker.uri") {
		err := shared.Client.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			log.Logger.
				With("Error", err).
				Error("Can't sent close packet")
		}

		err = shared.Client.Close()
		if err != nil {
			log.Logger.
				With("Error", err).
				Error("Connection closed")
		}
	}
}

func Listen() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	defer closeConnection()

	for {
		select {
		case <-Done:
			return
		case <-interrupt:

			select {
			case <-Done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
