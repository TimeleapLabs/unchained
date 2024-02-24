package client

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/KenshiTech/unchained/bls"
	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/constants"
	"github.com/KenshiTech/unchained/constants/opcodes"
	"github.com/KenshiTech/unchained/kosk"
	"github.com/KenshiTech/unchained/log"

	"github.com/gorilla/websocket"
	"github.com/vmihailenco/msgpack/v5"
)

var Client *websocket.Conn
var IsClientSocketClosed = false
var Done chan struct{}

func StartClient() {

	if !config.Config.IsSet("broker.uri") {
		return
	}

	brokerUrl := fmt.Sprintf(
		"%s/%s",
		config.Config.GetString("broker.uri"),
		constants.ProtocolVersion,
	)

	var err error
	Client, _, err = websocket.DefaultDialer.Dial(brokerUrl, nil)

	if err != nil {
		panic(err)
	}

	Done = make(chan struct{})

	hello := bls.Signer{
		Name:           config.Config.GetString("name"),
		PublicKey:      bls.ClientPublicKey.Bytes(),
		ShortPublicKey: bls.ClientShortPublicKey.Bytes(),
	}

	helloPayload, err := msgpack.Marshal(&hello)

	if err != nil {
		panic(err)
	}

	go func() {
		defer close(Done)

		for {
			_, payload, err := Client.ReadMessage()

			if err != nil || payload[0] == opcodes.Error {

				if err != nil {
					log.Logger.
						With("Error", err).
						Error("Read error")
				} else {
					log.Logger.
						With("Error", string(payload[1:])).
						Error("Broker error")
				}

				IsClientSocketClosed = true

				if websocket.IsUnexpectedCloseError(err) {
					for i := 1; i < 6; i++ {
						time.Sleep(time.Duration(i) * 3 * time.Second)
						Client, _, err = websocket.DefaultDialer.Dial(brokerUrl, nil)
						if err == nil {
							IsClientSocketClosed = false
							Client.WriteMessage(
								websocket.BinaryMessage,
								append([]byte{0}, helloPayload...),
							)
						}
					}
				}

				if IsClientSocketClosed {
					return
				} else {
					continue
				}
			}

			switch payload[0] {
			// TODO: Make a table of call codes
			case opcodes.Feedback:
				log.Logger.
					With("Feedback", string(payload[1:])).
					Info("Broker")

			case opcodes.KoskChallenge:
				// TODO: Refactor into a function
				// TODO: Check for errors!
				var challenge kosk.Challenge
				msgpack.Unmarshal(payload[1:], &challenge)

				signature, _ := bls.Sign(*bls.ClientSecretKey, challenge.Random[:])
				challenge.Signature = signature.Bytes()

				koskPayload, _ := msgpack.Marshal(challenge)

				Client.WriteMessage(
					websocket.BinaryMessage,
					append([]byte{opcodes.KoskResult}, koskPayload...),
				)

				if err != nil {
					log.Logger.
						With("Error", err).
						Error("Write error")
				}

			case opcodes.PriceReportBroadcast:
				ConsumePriceReport(payload[1:])

			case opcodes.EventLogBroadcast:
				ConsumeEventLog(payload[1:])

			default:
				log.Logger.
					With("Code", payload[0]).
					Info("Unknown call code")
			}

		}
	}()

	Client.WriteMessage(
		websocket.BinaryMessage,
		append([]byte{opcodes.Hello}, helloPayload...))
}

func closeConnection() {
	if config.Config.IsSet("broker.uri") {
		Client.Close()
	}
}

func ClientBlock() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	defer closeConnection()

	for {
		select {
		case <-Done:
			return
		case <-interrupt:

			if config.Config.IsSet("broker.uri") {
				err := Client.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))

				if err != nil {
					log.Logger.
						With("Error", err).
						Error("Connection closed")
					return
				}
			}

			select {
			case <-Done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
