package net

import (
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

func handleAtRoot(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, payload, err := c.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			break
		}

		var report uniswap.PriceReport
		msgpack.Unmarshal(payload, &report)

		toHash, err := msgpack.Marshal(&report.PriceInfo)

		if err != nil {
			continue
		}

		hash, err := bls.Hash(toHash)

		if err != nil {
			continue
		}

		signature, err := bls.RecoverSignature(report.Signature)

		if err != nil {
			continue
		}

		pk, err := bls.RecoverPublicKey(report.PublicKey)

		if err != nil {
			continue
		}

		ok, _ := bls.Verify(signature, hash, pk)

		var message []byte

		if ok {
			message = []byte("ok")
		} else {
			message = []byte("invalid signature")
		}

		err = c.WriteMessage(mt, message)
		if err != nil {
			fmt.Println("write:", err)
			break
		}
	}
}

func StartServer() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", handleAtRoot)
	log.Fatal(http.ListenAndServe(addr, nil))
}
