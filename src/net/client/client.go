package client

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/KenshiTech/unchained/src/config"
	"github.com/KenshiTech/unchained/src/constants"
	"github.com/KenshiTech/unchained/src/constants/opcodes"
	"github.com/KenshiTech/unchained/src/consumers"
	"github.com/KenshiTech/unchained/src/crypto/bls"
	"github.com/KenshiTech/unchained/src/kosk"
	"github.com/KenshiTech/unchained/src/log"
	"github.com/KenshiTech/unchained/src/net/shared"

	"github.com/gorilla/websocket"

	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

var Done chan struct{}

func StartClient() {
	if config.App.Broker.URI == "" {
		return
	}

	brokerURL := fmt.Sprintf(
		"%s/%s",
		config.App.Broker.URI,
		constants.ProtocolVersion,
	)

	var err error
	shared.Client, _, err = websocket.DefaultDialer.Dial(brokerURL, nil)
	if err != nil {
		panic(err)
	}

	Done = make(chan struct{})

	hello := bls.ClientSigner
	helloPayload := hello.Sia().Content

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
				challenge := new(kosk.Challenge).DeSia(&sia.Sia{Content: payload[1:]})
				signature, _ := bls.Sign(*bls.ClientSecretKey, challenge.Random[:])
				challenge.Signature = signature.Bytes()

				koskPayload := challenge.Sia().Content
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
	if shared.Client != nil && config.App.Broker.URI != "" {
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
