package uniswap

import (
	"context"
	"fmt"
	"math/big"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/KenshiTech/unchained/bls"
	"github.com/KenshiTech/unchained/contracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-co-op/gocron/v2"
	"github.com/gorilla/websocket"
	"github.com/njones/base58"
	"github.com/vmihailenco/msgpack/v5"
)

var addr = "65.108.48.32:9123"

type PriceInfo struct {
	Block uint64
	Price big.Int
}

type PriceReport struct {
	PriceInfo PriceInfo
	Signature []byte
	PublicKey []byte
}

func Work() {

	packedUrl := url.URL{Scheme: "ws", Host: addr, Path: "/"}
	wsClient, _, err := websocket.DefaultDialer.Dial(packedUrl.String(), nil)

	if err != nil {
		panic(err)
	}

	defer wsClient.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := wsClient.ReadMessage()
			if err != nil {
				fmt.Println("Read error:", err)
				return
			}
			fmt.Printf("Unchained feedback: %s\n", message)
		}
	}()

	encoding := base58.BitcoinEncoding
	client, err := ethclient.Dial("https://eth.llamarpc.com")

	if err != nil {
		panic(err)
	}

	scheduler, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	var twoOneNinetyTwo big.Int
	var tenEighteen big.Int
	var tenEighteenF big.Float
	var lastBlock uint64

	twoOneNinetyTwo.Exp(big.NewInt(2), big.NewInt(192), nil)
	tenEighteen.Exp(big.NewInt(10), big.NewInt(18), nil)
	tenEighteenF.SetInt(&tenEighteen)

	address := common.HexToAddress("0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640")
	uniV3, _ := contracts.NewUniV3(address, client)

	sk, pk, err := bls.GenerateKeyPair()

	if err != nil {
		panic(err)
	}

	pkBytes := pk.Bytes()
	var encoded [96]byte
	_ = encoding.Encode(encoded[:], pkBytes[:])

	fmt.Printf("Public Key: %s\n", encoded)

	_, err = scheduler.NewJob(
		gocron.DurationJob(5*time.Second),
		gocron.NewTask(
			func(decimalDif int64, inverse bool) {
				blockNumber, _ := client.BlockNumber(context.Background())

				if lastBlock == blockNumber {
					return
				}

				lastBlock = blockNumber
				data, _ := uniV3.Slot0(nil)

				var priceX96 big.Int
				var raw big.Int
				var price big.Int
				var factor big.Int
				var priceF big.Float

				// const raw = (fetchedSqrtPriceX96**2 / 2**192) * 10**6;
				priceX96.Exp(data.SqrtPriceX96, big.NewInt(2), nil)
				raw.Div(&priceX96, &twoOneNinetyTwo)

				if inverse {
					factor.Exp(big.NewInt(10), big.NewInt(36-decimalDif), nil)
					price.Div(&factor, &raw)
				} else {
					// TODO: needs work
					factor.Exp(big.NewInt(10), big.NewInt(decimalDif), nil)
					price.Div(&raw, &factor)
				}

				priceF.Quo(new(big.Float).SetInt(&price), &tenEighteenF)
				fmt.Printf("%d -> $%.18f\n", blockNumber, &priceF)

				priceInfo := PriceInfo{Price: price, Block: blockNumber}

				toHash, err := msgpack.Marshal(&priceInfo)
				if err != nil {
					panic(err)
				}

				signature, _ := bls.Sign(*sk, toHash)
				compressedSignature := signature.Bytes()

				priceReport := PriceReport{
					PriceInfo: priceInfo,
					Signature: compressedSignature[:],
					PublicKey: pkBytes[:],
				}

				payload, err := msgpack.Marshal(&priceReport)
				if err != nil {
					panic(err)
				}

				//fmt.Printf("%x\n", payload)
				wsClient.WriteMessage(websocket.BinaryMessage, payload)

				// fmt.Printf("%x\n", signature.Bytes())
				// ok, _ := bls.Verify(signature, hash, *pk)
				// fmt.Printf("Is OK? %t\n", ok)
			},
			int64(6),
			true,
		),
	)

	if err != nil {
		panic(err)
	}

	scheduler.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for {
		select {
		case <-done:
			return
		case <-interrupt:

			err := wsClient.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))

			if err != nil {
				fmt.Println("write close:", err)
				return
			}

			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
