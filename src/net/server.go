package net

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/constants"
	"github.com/KenshiTech/unchained/constants/opcodes"
	"github.com/KenshiTech/unchained/crypto/bls"
	"github.com/KenshiTech/unchained/datasets"
	"github.com/KenshiTech/unchained/kosk"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/net/consumer"
	"github.com/KenshiTech/unchained/net/repository"

	"github.com/gorilla/websocket"
	"github.com/puzpuzpuz/xsync/v3"
	"github.com/vmihailenco/msgpack/v5"
)

var challenges *xsync.MapOf[*websocket.Conn, kosk.Challenge]
var signers *xsync.MapOf[*websocket.Conn, bls.Signer]
var upgrader = websocket.Upgrader{} // use default options

func processKosk(conn *websocket.Conn, messageType int, payload []byte) error {
	var challenge kosk.Challenge
	err := msgpack.Unmarshal(payload, &challenge)

	if err != nil {
		err = conn.WriteMessage(
			messageType,
			append(
				[]byte{opcodes.Feedback},
				[]byte("packet.invalid")...),
		)

		if err != nil {
			fmt.Println("write:", err)
			return err
		}

		return nil
	}

	signer, ok := signers.Load(conn)

	if !ok {
		conn.WriteMessage(
			messageType,
			append(
				[]byte{opcodes.Feedback},
				[]byte("hello.missing")...),
		)
		return errors.New("hello.missing")
	}

	challenge.Passed, err = kosk.VerifyChallenge(
		challenge.Random,
		signer.PublicKey,
		challenge.Signature,
	)

	if err != nil || !challenge.Passed {
		conn.WriteMessage(
			messageType,
			append(
				[]byte{opcodes.Error},
				[]byte("kosk.invalid")...),
		)
		return errors.New("kosk.invalid")
	}

	conn.WriteMessage(
		messageType,
		append(
			[]byte{opcodes.Feedback},
			[]byte("kosk.ok")...),
	)

	challenges.Store(conn, challenge)

	return nil
}

func processHello(conn *websocket.Conn, messageType int, payload []byte) error {

	var signer bls.Signer
	err := msgpack.Unmarshal(payload, &signer)

	if err != nil {
		// TODO: what's the best way of doing this?
		err = conn.WriteMessage(
			messageType,
			append(
				[]byte{opcodes.Feedback},
				[]byte("packet.invalid")...),
		)

		if err != nil {
			fmt.Println("write:", err)
			return err
		}

		return nil
	}

	if signer.Name == "" || len(signer.PublicKey) != 96 {
		conn.WriteMessage(
			messageType,
			append(
				[]byte{opcodes.Error},
				[]byte("conf.invalid")...),
		)
		return errors.New("conf.invalid")
	}

	signers.Range(func(conn *websocket.Conn, signerInMap bls.Signer) bool {
		publicKeyInUse := signerInMap.PublicKey == signer.PublicKey
		if publicKeyInUse {
			err := conn.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))

			conn.Close()

			if err != nil {
				log.Logger.
					With("Error", err).
					Error("Connection closed")
			}
		}
		return !publicKeyInUse
	})

	signers.Store(conn, signer)

	err = conn.WriteMessage(
		messageType,
		append(
			[]byte{opcodes.Feedback},
			[]byte("conf.ok")...),
	)

	if err != nil {
		fmt.Println("write:", err)
		return err
	}

	// Start KOSK verification

	challenge := kosk.Challenge{Random: kosk.NewChallenge()}
	challenges.Store(conn, challenge)
	koskPayload, err := msgpack.Marshal(challenge)

	// TODO: Client should hang on error
	if err != nil {
		conn.WriteMessage(
			messageType,
			append(
				[]byte{opcodes.Error},
				[]byte("kosk.error")...),
		)
		return err
	}

	err = conn.WriteMessage(
		messageType,
		append(
			[]byte{opcodes.KoskChallenge},
			koskPayload...),
	)

	if err != nil {
		conn.WriteMessage(
			messageType,
			append(
				[]byte{opcodes.Error},
				[]byte("kosk.error")...),
		)
		return err
	}

	return nil
}

func checkPublicKey(conn *websocket.Conn, messageType int) (*bls.Signer, error) {
	challenge, ok := challenges.Load(conn)

	if !ok || !challenge.Passed {
		conn.WriteMessage(
			messageType,
			append(
				[]byte{opcodes.Feedback},
				[]byte("kosk.missing")...),
		)
		return nil, errors.New("kosk.missing")
	}

	signer, ok := signers.Load(conn)

	if !ok {
		conn.WriteMessage(
			messageType,
			append(
				[]byte{opcodes.Feedback},
				[]byte("hello.missing")...),
		)
		return nil, errors.New("hello.missing")
	}

	return &signer, nil
}

// TODO: Can we use any part of this?
func processPriceReport(conn *websocket.Conn, messageType int, payload []byte) error {

	signer, err := checkPublicKey(conn, messageType)

	if err != nil {
		return err
	}

	var report datasets.PriceReport
	err = msgpack.Unmarshal(payload, &report)

	if err != nil {
		return nil
	}

	toHash, err := msgpack.Marshal(&report.PriceInfo)

	if err != nil {
		return nil
	}

	hash, err := bls.Hash(toHash)

	if err != nil {
		return nil
	}

	signature, err := bls.RecoverSignature(report.Signature)

	if err != nil {
		return nil
	}

	pk, err := bls.RecoverPublicKey(signer.PublicKey)

	if err != nil {
		return nil
	}

	ok, _ := bls.Verify(signature, hash, pk)

	message := []byte("signature.invalid")
	if ok {
		message = []byte("signature.accepted")

		broadcastPacket := datasets.BroadcastPricePacket{
			Info:      report.PriceInfo,
			Signature: report.Signature,
			Signer:    *signer,
		}

		broadcastPayload, err := msgpack.Marshal(&broadcastPacket)

		// TODO: Handle this error properly
		// TODO: Maybe notify the peer so they can resend
		if err != nil {
			log.Logger.
				With("Error", err).
				Error("Cannot marshal the broadcast packet")
		} else {
			// TODO: Safe to use 'go' here?
			consumer.Broadcast(
				append(
					[]byte{opcodes.PriceReportBroadcast},
					broadcastPayload...,
				),
			)
		}
	}

	err = conn.WriteMessage(
		messageType,
		append(
			[]byte{opcodes.Feedback},
			message...),
	)

	if err != nil {
		fmt.Println("write:", err)
		return err
	}

	return nil
}

func processEventLog(conn *websocket.Conn, messageType int, payload []byte) error {

	signer, err := checkPublicKey(conn, messageType)

	if err != nil {
		return err
	}

	var report datasets.EventLogReport
	err = msgpack.Unmarshal(payload, &report)

	if err != nil {
		return nil
	}

	toHash, err := msgpack.Marshal(&report.EventLog)

	if err != nil {
		return nil
	}

	hash, err := bls.Hash(toHash)

	if err != nil {
		return nil
	}

	signature, err := bls.RecoverSignature(report.Signature)

	if err != nil {
		return nil
	}

	pk, err := bls.RecoverPublicKey(signer.PublicKey)

	if err != nil {
		return nil
	}

	ok, _ := bls.Verify(signature, hash, pk)

	message := []byte("signature.invalid")
	if ok {
		message = []byte("signature.accepted")
		broadcastPacket := datasets.BroadcastEventPacket{
			Info:      report.EventLog,
			Signature: report.Signature,
			Signer:    *signer,
		}

		broadcastPayload, err := msgpack.Marshal(&broadcastPacket)

		// TODO: Handle this error properly
		// TODO: Maybe notify the peer so they can resend
		if err != nil {
			log.Logger.
				With("Error", err).
				Error("Cannot marshal the broadcast packet")
		} else {
			// TODO: Safe to use 'go' here?
			consumer.Broadcast(
				append(
					[]byte{opcodes.EventLogBroadcast},
					broadcastPayload...,
				))
		}
	}

	err = conn.WriteMessage(
		messageType,
		append(
			[]byte{opcodes.Feedback},
			message...),
	)

	if err != nil {
		fmt.Println("write:", err)
		return err
	}

	return nil
}

func handleAtRoot(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Print("upgrade:", err)
		return
	}

	defer conn.Close()
	defer signers.Delete(conn)
	defer challenges.Delete(conn)
	defer repository.Consumers.Delete(conn)
	defer repository.BroadcastMutex.Delete(conn)

	for {
		messageType, payload, err := conn.ReadMessage()

		if err != nil {
			fmt.Println("read:", err)
			break
		}

		switch payload[0] {
		// TODO: Make a table of call codes
		case opcodes.Hello:
			err := processHello(conn, messageType, payload[1:])

			if err != nil {
				fmt.Println("write:", err)
			}

		case opcodes.PriceReport:
			// TODO: Maybe this is unnecessary
			if payload[1] == 0 {
				err := processPriceReport(conn, messageType, payload[2:])

				if err != nil {
					fmt.Println("write:", err)
				}

			} else {
				conn.WriteMessage(
					messageType,
					append(
						[]byte{opcodes.Feedback},
						[]byte("Dataset not supported")...),
				)
			}

		case opcodes.EventLog:
			// TODO: Maybe this is unnecessary
			if payload[1] == 0 {
				err := processEventLog(conn, messageType, payload[2:])

				if err != nil {
					fmt.Println("write:", err)
				}

			} else {
				conn.WriteMessage(
					messageType,
					append(
						[]byte{opcodes.Feedback},
						[]byte("Dataset not supported")...),
				)
			}

		case opcodes.KoskResult:
			err := processKosk(conn, messageType, payload[1:])

			if err != nil {
				fmt.Println("write:", err)
			}

		case opcodes.RegisterConsumer:
			// TODO: Consumers must specify what they're subscribing to
			repository.Consumers.Store(conn, true)
			repository.BroadcastMutex.Store(conn, new(sync.Mutex))

		default:
			err = conn.WriteMessage(
				messageType,
				append(
					[]byte{opcodes.Error},
					[]byte(fmt.Sprintf("Instruction not supported: %x", payload[0]))...),
			)
			if err != nil {
				fmt.Println("write:", err)
			}
		}
	}
}

func StartServer() {
	flag.Parse()
	versionedRoot := fmt.Sprintf("/%s", constants.ProtocolVersion)
	http.HandleFunc(versionedRoot, handleAtRoot)
	addr := config.Config.GetString("broker.bind")
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Logger.
			With("err", err).
			Error("Cound't start the websocket server")
		os.Exit(1)
	}
}

func init() {
	signers = xsync.NewMapOf[*websocket.Conn, bls.Signer]()
	challenges = xsync.NewMapOf[*websocket.Conn, kosk.Challenge]()
}
