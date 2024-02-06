package net

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/KenshiTech/unchained/bls"
	"github.com/KenshiTech/unchained/constants"
	"github.com/KenshiTech/unchained/kosk"
	"github.com/KenshiTech/unchained/plugins/uniswap"

	"github.com/gorilla/websocket"
	"github.com/puzpuzpuz/xsync/v3"
	"github.com/vmihailenco/msgpack/v5"
)

var challenges *xsync.MapOf[*websocket.Conn, kosk.Challenge]
var signers *xsync.MapOf[*websocket.Conn, uniswap.Signer]
var upgrader = websocket.Upgrader{} // use default options
var addr = "0.0.0.0:9123"           // TODO: port should be passed as a cli option

func processKosk(conn *websocket.Conn, messageType int, payload []byte) error {
	var challenge kosk.Challenge
	err := msgpack.Unmarshal(payload, &challenge)

	if err != nil {
		err = conn.WriteMessage(messageType, append([]byte{2}, []byte("packet.invalid")...))
		if err != nil {
			fmt.Println("write:", err)
			return err
		}
		return nil
	}

	signer, ok := signers.Load(conn)

	if !ok {
		conn.WriteMessage(messageType, append([]byte{2}, []byte("hello.missing")...))
		return errors.New("hello.missing")
	}

	challenge.Passed, err = kosk.VerifyChallenge(
		challenge.Random,
		signer.PublicKey,
		challenge.Signature,
	)

	if err != nil || !challenge.Passed {
		conn.WriteMessage(messageType, append([]byte{2}, []byte("kosk.invalid")...))
		return errors.New("kosk.invalid")
	}

	conn.WriteMessage(messageType, append([]byte{2}, []byte("kosk.ok")...))
	challenges.Store(conn, challenge)

	return nil
}

func processHello(conn *websocket.Conn, messageType int, payload []byte) error {

	var signer uniswap.Signer
	err := msgpack.Unmarshal(payload, &signer)

	if err != nil {
		// TODO: what's the best way of doing this?
		err = conn.WriteMessage(messageType, append([]byte{2}, []byte("packet.invalid")...))
		if err != nil {
			fmt.Println("write:", err)
			return err
		}
		return nil
	}

	if signer.Name == "" || len(signer.PublicKey) != 96 {
		conn.WriteMessage(messageType, append([]byte{2}, []byte("conf.invalid")...))
		return errors.New("conf.invalid")
	}

	signers.Store(conn, signer)
	err = conn.WriteMessage(messageType, append([]byte{2}, []byte("conf.ok")...))

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
		conn.WriteMessage(messageType, append([]byte{5}, []byte("kosk.error")...))
		return err
	}

	err = conn.WriteMessage(messageType, append([]byte{4}, koskPayload...))

	if err != nil {
		conn.WriteMessage(messageType, append([]byte{5}, []byte("kosk.error")...))
		return err
	}

	return nil
}

func processPriceReport(conn *websocket.Conn, messageType int, payload []byte) error {

	challenge, ok := challenges.Load(conn)

	if !ok || !challenge.Passed {
		conn.WriteMessage(messageType, append([]byte{2}, []byte("kosk.missing")...))
		return errors.New("kosk.missing")
	}

	signer, ok := signers.Load(conn)

	if !ok {
		conn.WriteMessage(messageType, append([]byte{2}, []byte("hello.missing")...))
		return errors.New("hello.missing")
	}

	var report uniswap.PriceReport
	err := msgpack.Unmarshal(payload, &report)

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

	ok, _ = bls.Verify(signature, hash, pk)

	message := []byte("signature.invalid")
	if ok {
		message = []byte("signature.accepted")
		uniswap.RecordSignature(
			signature,
			signer,
			report.PriceInfo.Block,
		)
	}

	err = conn.WriteMessage(messageType, append([]byte{2}, message...))

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

	for {
		messageType, payload, err := conn.ReadMessage()

		if err != nil {
			fmt.Println("read:", err)
			break
		}

		switch payload[0] {
		// TODO: Make a table of call codes
		case 0:
			err := processHello(conn, messageType, payload[1:])

			if err != nil {
				fmt.Println("write:", err)
			}

		case 1:
			err := processPriceReport(conn, messageType, payload[1:])

			if err != nil {
				fmt.Println("write:", err)
			}

		case 3:
			err := processKosk(conn, messageType, payload[1:])

			if err != nil {
				fmt.Println("write:", err)
			}

		default:
			err = conn.WriteMessage(
				messageType,
				append([]byte{2}, []byte("Instruction not supported")...),
			)
			if err != nil {
				fmt.Println("write:", err)
			}
		}
	}
}

func StartServer() {
	flag.Parse()
	log.SetFlags(0)
	versionedRoot := fmt.Sprintf("/%s", constants.Version)
	http.HandleFunc(versionedRoot, handleAtRoot)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func init() {
	signers = xsync.NewMapOf[*websocket.Conn, uniswap.Signer]()
	challenges = xsync.NewMapOf[*websocket.Conn, kosk.Challenge]()
}
