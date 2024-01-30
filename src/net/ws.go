package net

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/KenshiTech/unchained/bls"
	"github.com/KenshiTech/unchained/plugins/uniswap"

	"github.com/gorilla/websocket"
	"github.com/vmihailenco/msgpack/v5"
)

var upgrader = websocket.Upgrader{} // use default options
var addr = "0.0.0.0:9123"

var signers = make(map[*websocket.Conn]uniswap.Signer)

func processHello(conn *websocket.Conn, messageType int, payload []byte) error {

	var signer uniswap.Signer
	err := msgpack.Unmarshal(payload, &signer)

	if err != nil {
		err = conn.WriteMessage(messageType, []byte("packet.invalid"))
		if err != nil {
			fmt.Println("write:", err)
			return err
		}
		return nil
	}

	signers[conn] = signer
	err = conn.WriteMessage(messageType, []byte("conf.ok"))

	if err != nil {
		fmt.Println("write:", err)
		return err
	}

	return nil
}

func processPriceReport(conn *websocket.Conn, messageType int, payload []byte) error {

	signer, ok := signers[conn]

	if !ok {
		conn.WriteMessage(messageType, []byte("hello.missing"))
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

	if err != nil {
		fmt.Println("Faileded")
	}

	uniswap.RecordSignature(
		signature,
		signer,
		report.PriceInfo.Block,
	)

	message := []byte("signature.invalid")

	if ok {
		message = []byte("signature.accepted")
	}

	err = conn.WriteMessage(messageType, message)

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
	defer delete(signers, conn)

	for {
		messageType, payload, err := conn.ReadMessage()

		if err != nil {
			fmt.Println("read:", err)
			break
		}

		switch payload[0] {
		case 0:
			err := processHello(conn, messageType, payload[1:])

			if err != nil {
				fmt.Println("write:", err)
				break
			}

		case 1:
			err := processPriceReport(conn, messageType, payload[1:])

			if err != nil {
				fmt.Println("write:", err)
				break
			}
		default:
			err = conn.WriteMessage(messageType, []byte("Instruction not supported"))
			if err != nil {
				fmt.Println("write:", err)
				break
			}
		}
	}
}

func StartServer() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", handleAtRoot)
	log.Fatal(http.ListenAndServe(addr, nil))
}
