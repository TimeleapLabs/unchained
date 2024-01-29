package uniswap

import (
	"fmt"
	"math/big"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/KenshiTech/unchained/bls"
	"github.com/KenshiTech/unchained/ethereum"

	btcutil "github.com/btcsuite/btcutil/base58"
	"github.com/go-co-op/gocron/v2"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"github.com/vmihailenco/msgpack/v5"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
)

var addr = "shinobi.brokers.kenshi.io"

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

	packedUrl := url.URL{Scheme: "wss", Host: addr, Path: "/"}
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

	etherUsdPairAddr := "0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640"
	etherUsdPair, err := ethereum.GetNewUniV3Contract(etherUsdPairAddr, false)

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

	var sk *big.Int
	var pk *bls12381.G1Affine

	if viper.InConfig("secretKey") {

		decoded := btcutil.Decode(viper.GetString("secretKey"))

		sk = new(big.Int)
		sk.SetBytes(decoded)

		pk = bls.GetPublicKey(sk)
	} else {
		sk, pk, err = bls.GenerateKeyPair()
	}

	if err != nil {
		panic(err)
	}

	pkBytes := pk.Bytes()
	pkStr := btcutil.Encode(pkBytes[:])

	fmt.Printf("Public Key: %s\n", pkStr)

	_, err = scheduler.NewJob(
		gocron.DurationJob(5*time.Second),
		gocron.NewTask(
			func(decimalDif int64, inverse bool) {
				blockNumber, err := ethereum.GetBlockNumber()

				if err != nil {
					etherUsdPair, err = ethereum.GetNewUniV3Contract(etherUsdPairAddr, true)

					if err != nil {
						panic(err)
					}

					return
				}

				if lastBlock == blockNumber {
					return
				}

				lastBlock = blockNumber
				data, err := etherUsdPair.Slot0(nil)

				if err != nil {
					etherUsdPair, err = ethereum.GetNewUniV3Contract(etherUsdPairAddr, true)

					if err != nil {
						panic(err)
					}

					return
				}

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
