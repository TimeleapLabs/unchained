package net

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/KenshiTech/unchained/src/config"
	"github.com/KenshiTech/unchained/src/constants"
	"github.com/KenshiTech/unchained/src/constants/opcodes"
	"github.com/KenshiTech/unchained/src/crypto/bls"
	"github.com/KenshiTech/unchained/src/datasets"
	"github.com/KenshiTech/unchained/src/kosk"
	"github.com/KenshiTech/unchained/src/log"
	"github.com/KenshiTech/unchained/src/net/repository"

	"github.com/gorilla/websocket"
	"github.com/puzpuzpuz/xsync/v3"

	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

var challenges *xsync.MapOf[*websocket.Conn, kosk.Challenge]
var signers *xsync.MapOf[*websocket.Conn, datasets.Signer]
var upgrader = websocket.Upgrader{} // use default options

func processKosk(conn *websocket.Conn, payload []byte) error {
	challenge := new(kosk.Challenge).DeSia(&sia.Sia{Content: payload})

	signer, ok := signers.Load(conn)
	if !ok {
		return constants.ErrMissingHello
	}

	var err error
	challenge.Passed, err = kosk.VerifyChallenge(challenge.Random, signer.PublicKey, challenge.Signature)

	if err != nil {
		return constants.ErrInvalidKosk
	}

	if !challenge.Passed {
		log.Logger.Error("challenge is Passed")
		return constants.ErrInvalidKosk
	}

	challenges.Store(conn, *challenge)
	return nil
}

func processHello(conn *websocket.Conn, payload []byte) ([]byte, error) {
	signer := new(datasets.Signer).DeSia(&sia.Sia{Content: payload})

	if signer.Name == "" {
		log.Logger.Error("Signer name is empty Or public key is invalid")
		return []byte{}, constants.ErrInvalidConfig
	}

	signers.Range(func(conn *websocket.Conn, signerInMap datasets.Signer) bool {
		publicKeyInUse := signerInMap.PublicKey == signer.PublicKey
		if publicKeyInUse {
			Close(conn)
		}
		return !publicKeyInUse
	})

	signers.Store(conn, *signer)

	// Start KOSK verification
	challenge := kosk.Challenge{Random: kosk.NewChallenge()}
	challenges.Store(conn, challenge)
	koskPayload := challenge.Sia().Content

	return koskPayload, nil
}

func checkPublicKey(conn *websocket.Conn) (*datasets.Signer, error) {
	challenge, ok := challenges.Load(conn)
	if !ok || !challenge.Passed {
		return nil, constants.ErrMissingKosk
	}

	signer, ok := signers.Load(conn)
	if !ok {
		return nil, constants.ErrMissingHello
	}

	return &signer, nil
}

// TODO: Can we use any part of this?
func processPriceReport(conn *websocket.Conn, payload []byte) ([]byte, error) {
	signer, err := checkPublicKey(conn)
	if err != nil {
		return []byte{}, err
	}

	report := new(datasets.PriceReport).DeSia(&sia.Sia{Content: payload})
	toHash := report.PriceInfo.Sia().Content
	hash, err := bls.Hash(toHash)

	if err != nil {
		log.Logger.Error("Can't hash bls: %v", err)
		return []byte{}, constants.ErrInternalError
	}

	signature, err := bls.RecoverSignature(report.Signature)
	if err != nil {
		log.Logger.Error("Can't recover bls signature: %v", err)
		return []byte{}, constants.ErrInternalError
	}

	pk, err := bls.RecoverPublicKey(signer.PublicKey)
	if err != nil {
		log.Logger.Error("Can't recover bls pub-key: %v", err)
		return []byte{}, constants.ErrInternalError
	}

	ok, err := bls.Verify(signature, hash, pk)
	if err != nil {
		log.Logger.Error("Can't recover bls pub-key: %v", err)
		return []byte{}, constants.ErrCantVerifyBls
	}
	if !ok {
		return []byte{}, constants.ErrInvalidSignature
	}

	priceInfo := datasets.BroadcastPricePacket{
		Info:      report.PriceInfo,
		Signature: report.Signature,
		Signer:    *signer,
	}

	priceInfoByte := priceInfo.Sia().Content
	return priceInfoByte, nil
}

func processEventLog(conn *websocket.Conn, payload []byte) ([]byte, error) {
	signer, err := checkPublicKey(conn)
	if err != nil {
		return []byte{}, err
	}

	report := new(datasets.EventLogReport).DeSia(&sia.Sia{Content: payload})
	toHash := report.EventLog.Sia().Content
	hash, err := bls.Hash(toHash)

	if err != nil {
		log.Logger.Error("Can't hash bls: %v", err)
		return []byte{}, constants.ErrInternalError
	}

	signature, err := bls.RecoverSignature(report.Signature)
	if err != nil {
		log.Logger.Error("Can't recover bls signature: %v", err)
		return []byte{}, constants.ErrInternalError
	}

	pk, err := bls.RecoverPublicKey(signer.PublicKey)
	if err != nil {
		log.Logger.Error("Can't recover bls pub-key: %v", err)
		return []byte{}, constants.ErrCantVerifyBls
	}

	ok, err := bls.Verify(signature, hash, pk)
	if err != nil {
		log.Logger.Error("Can't recover bls pub-key: %v", err)
		return []byte{}, constants.ErrCantVerifyBls
	}
	if !ok {
		return []byte{}, constants.ErrInvalidSignature
	}

	broadcastPacket := datasets.BroadcastEventPacket{
		Info:      report.EventLog,
		Signature: report.Signature,
		Signer:    *signer,
	}

	broadcastPayload := broadcastPacket.Sia().Content
	return broadcastPayload, nil
}

func processCorrectnessRecord(conn *websocket.Conn, payload []byte) ([]byte, error) {
	signer, err := checkPublicKey(conn)
	if err != nil {
		return []byte{}, err
	}

	report := new(datasets.CorrectnessReport).DeSia(&sia.Sia{Content: payload})
	toHash := report.Correctness.Sia().Content
	hash, err := bls.Hash(toHash)

	if err != nil {
		log.Logger.Error("Can't hash bls: %v", err)
		return []byte{}, constants.ErrInternalError
	}

	signature, err := bls.RecoverSignature(report.Signature)
	if err != nil {
		log.Logger.Error("Can't recover bls signature: %v", err)
		return []byte{}, constants.ErrInternalError
	}

	pk, err := bls.RecoverPublicKey(signer.PublicKey)
	if err != nil {
		log.Logger.Error("Can't recover bls pub-key: %v", err)
		return []byte{}, constants.ErrCantVerifyBls
	}

	ok, err := bls.Verify(signature, hash, pk)
	if err != nil {
		log.Logger.With("Error", err).Error("Can't verify bls")
		return []byte{}, constants.ErrCantVerifyBls
	}
	if !ok {
		return []byte{}, constants.ErrInvalidSignature
	}

	broadcastPacket := datasets.BroadcastCorrectnessPacket{
		Info:      report.Correctness,
		Signature: report.Signature,
		Signer:    *signer,
	}

	broadcastPayload := broadcastPacket.Sia().Content
	return broadcastPayload, nil
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Logger.Error("Can't upgrade connection: %v", err)
		return
	}

	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Logger.Error("Can't close connection: %v", err)
		}
	}(conn)

	defer signers.Delete(conn)
	defer challenges.Delete(conn)
	defer repository.Consumers.Delete(conn)
	defer repository.BroadcastMutex.Delete(conn)

	for {
		messageType, payload, err := conn.ReadMessage()
		if err != nil {
			log.Logger.Error("Can't read message: %v", err)
			break
		}

		switch opcodes.OpCode(payload[0]) {
		// TODO: Make a table of call codes
		case opcodes.Hello:
			result, err := processHello(conn, payload[1:])
			if err != nil {
				SendError(conn, messageType, opcodes.Error, err)
			}

			SendMessage(conn, messageType, opcodes.Feedback, "conf.ok")
			Send(conn, messageType, opcodes.KoskChallenge, result)

		case opcodes.PriceReport:
			result, err := processPriceReport(conn, payload[1:])
			if err != nil {
				SendError(conn, messageType, opcodes.Error, err)
			}

			BroadcastPayload(opcodes.PriceReportBroadcast, result)
			SendMessage(conn, messageType, opcodes.Feedback, "signature.accepted")

		case opcodes.EventLog:
			result, err := processEventLog(conn, payload[1:])
			if err != nil {
				SendError(conn, messageType, opcodes.Error, err)
			}

			BroadcastPayload(opcodes.EventLogBroadcast, result)
			SendMessage(conn, messageType, opcodes.Feedback, "signature.accepted")

		case opcodes.CorrectnessReport:
			result, err := processCorrectnessRecord(conn, payload[1:])
			if err != nil {
				SendError(conn, messageType, opcodes.Error, err)
			}

			BroadcastPayload(opcodes.CorrectnessReportBroadcast, result)
			SendMessage(conn, messageType, opcodes.Feedback, "signature.accepted")

		case opcodes.KoskResult:
			err := processKosk(conn, payload[1:])
			if err != nil {
				SendError(conn, messageType, opcodes.Error, err)
			}
			SendMessage(conn, messageType, opcodes.Feedback, "kosk.ok")

		case opcodes.RegisterConsumer:
			// TODO: Consumers must specify what they're subscribing to
			repository.Consumers.Store(conn, true)
			repository.BroadcastMutex.Store(conn, new(sync.Mutex))

		default:
			SendError(conn, messageType, opcodes.Error, constants.ErrNotSupportedInstruction)
		}
	}
}

func StartServer() {
	flag.Parse()
	versionedRoot := fmt.Sprintf("/%s", constants.ProtocolVersion)
	http.HandleFunc(versionedRoot, rootHandler)
	addr := config.App.Broker.Bind

	readHeaderTimeoutInSecond := 3
	server := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: time.Duration(readHeaderTimeoutInSecond) * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func init() {
	signers = xsync.NewMapOf[*websocket.Conn, datasets.Signer]()
	challenges = xsync.NewMapOf[*websocket.Conn, kosk.Challenge]()
}
