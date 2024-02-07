package uniswap

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/KenshiTech/unchained/address"
	"github.com/KenshiTech/unchained/bls"
	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/constants"
	"github.com/KenshiTech/unchained/db"
	"github.com/KenshiTech/unchained/ent"
	"github.com/KenshiTech/unchained/ent/signer"
	"github.com/KenshiTech/unchained/ethereum"
	"github.com/KenshiTech/unchained/kosk"
	"github.com/KenshiTech/unchained/utils"

	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/go-co-op/gocron/v2"
	"github.com/gorilla/websocket"
	"github.com/vmihailenco/msgpack/v5"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	lru "github.com/hashicorp/golang-lru/v2"
)

var etherUsdPairAddr = "0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640"

type Signer struct {
	Name           string
	PublicKey      [96]byte
	ShortPublicKey [48]byte
}
type Signature struct {
	Signature bls12381.G1Affine
	Signer    Signer
	Processed bool
}

type PriceInfo struct {
	Block uint64
	Price big.Int
}

type PriceReport struct {
	PriceInfo PriceInfo
	Signature [48]byte
}

var DebouncedSaveSignatures func(key uint64, arg uint64)

var priceCache *lru.Cache[uint64, big.Int]
var signatureCache *lru.Cache[uint64, []Signature]
var aggregateCache *lru.Cache[uint64, bls12381.G1Affine]

var twoOneNinetyTwo big.Int
var tenEighteen big.Int
var tenEighteenF big.Float
var lastBlock uint64
var lastPrice big.Int

func RecordSignature(signature bls12381.G1Affine, signer Signer, block uint64) {

	// TODO: Needs optimization
	if !priceCache.Contains(block) {
		blockNumber, _, err := GetPriceFromPair(etherUsdPairAddr, 6, true)

		if err != nil {
			return
		}

		lastBlock = *blockNumber
	}

	if lastBlock-block > 16 {
		return // Data too old
	}

	mu := new(sync.Mutex)
	mu.Lock()
	defer mu.Unlock()

	cached, ok := signatureCache.Get(block)
	packed := Signature{Signature: signature, Signer: signer, Processed: false}

	if !ok {
		signatureCache.Add(block, []Signature{packed})
		// TODO: This looks ugly
		DebouncedSaveSignatures(block, block)
		return
	}

	for _, item := range cached {
		if item.Signer.PublicKey == signer.PublicKey {
			return
		}
	}

	cached = append(cached, packed)
	signatureCache.Add(block, cached)

	DebouncedSaveSignatures(block, block)
}

func SaveSignatures(block uint64) {

	dbClient := db.GetClient()
	price, ok := priceCache.Get(block)

	if !ok {
		return
	}

	signatures, ok := signatureCache.Get(block)

	if !ok {
		return
	}

	ctx := context.Background()

	// TODO: Cache this
	datasetId, err := dbClient.DataSet.
		Create().
		SetName("uniswap::ethereum::ethereum").
		OnConflictColumns("name").
		UpdateName().
		ID(ctx)

	if err != nil {
		panic(err)
	}

	var newSigners []Signer
	var newSignatures []bls12381.G1Affine
	var keys [][]byte

	for i := range signatures {
		signature := signatures[i]
		keys = append(keys, signature.Signer.PublicKey[:])
		if !signature.Processed {
			newSignatures = append(newSignatures, signature.Signature)
			newSigners = append(newSigners, signature.Signer)
		}
	}

	err = dbClient.Signer.MapCreateBulk(newSigners, func(sc *ent.SignerCreate, i int) {
		signer := newSigners[i]
		sc.SetName(signer.Name).
			SetKey(signer.PublicKey[:]).
			SetShortkey(signer.ShortPublicKey[:]).
			SetPoints(0)
	}).
		OnConflictColumns("shortkey").
		UpdateName().
		UpdateKey().
		Update(func(su *ent.SignerUpsert) {
			su.AddPoints(1)
		}).
		Exec(ctx)

	if err != nil {
		panic(err)
	}

	signerIds, err := dbClient.Signer.
		Query().
		Where(signer.KeyIn(keys...)).
		IDs(ctx)

	if err != nil {
		return
	}

	var aggregate bls12381.G1Affine
	currentAggregate, ok := aggregateCache.Get(block)

	if ok {
		newSignatures = append(newSignatures, currentAggregate)
	}

	aggregate, err = bls.AggregateSignatures(newSignatures)

	if err != nil {
		return
	}

	signatureBytes := aggregate.Bytes()

	err = dbClient.AssetPrice.
		Create().
		SetBlock(block).
		SetPrice(&price).
		SetSignersCount(uint64(len(signatures))).
		SetSignature(signatureBytes[:]).
		AddDataSetIDs(datasetId).
		AddSignerIDs(signerIds...).
		OnConflictColumns("block").
		UpdateNewValues().
		Exec(ctx)

	if err != nil {
		panic(err)
	}

	for _, signature := range signatures {
		signature.Processed = true
	}

	aggregateCache.Add(block, aggregate)

}

// TODO: Each pair should have its own LRU-Cache
func GetPriceFromCache(block uint64) (big.Int, bool) {
	return priceCache.Get(block)
}

func GetPriceFromPair(pairAddr string, decimalDif int64, inverse bool) (*uint64, *big.Int, error) {
	blockNumber, err := ethereum.GetBlockNumber()

	if err != nil {
		ethereum.RefreshRPC()
		return nil, nil, err
	}

	if blockNumber == lastBlock {
		return &blockNumber, &lastPrice, nil
	}

	pair, err := ethereum.GetNewUniV3Contract(pairAddr, false)

	if err != nil {
		ethereum.RefreshRPC()
		return nil, nil, err
	}

	data, err := pair.Slot0(
		&bind.CallOpts{
			BlockNumber: big.NewInt(int64(blockNumber)),
		})

	if err != nil {
		ethereum.RefreshRPC()
		return nil, nil, err
	}

	lastPrice = *priceFromSqrtX96(data.SqrtPriceX96, 6, true)
	priceCache.Add(blockNumber, lastPrice)

	return &blockNumber, &lastPrice, nil
}

func priceFromSqrtX96(sqrtPriceX96 *big.Int, decimalDif int64, inverse bool) *big.Int {
	var priceX96 big.Int
	var raw big.Int
	var price big.Int
	var factor big.Int

	// const raw = (fetchedSqrtPriceX96**2 / 2**192) * 10**6;
	priceX96.Exp(sqrtPriceX96, big.NewInt(2), nil)
	raw.Div(&priceX96, &twoOneNinetyTwo)

	if inverse {
		factor.Exp(big.NewInt(10), big.NewInt(36-decimalDif), nil)
		price.Div(&factor, &raw)
	} else {
		// TODO: needs work
		factor.Exp(big.NewInt(10), big.NewInt(decimalDif), nil)
		price.Div(&raw, &factor)
	}
	return &price
}

func Start() {

	brokerUrl := fmt.Sprintf(
		"%s/%s",
		config.Config.GetString("broker"),
		constants.ProtocolVersion,
	)

	wsClient, _, err := websocket.DefaultDialer.Dial(brokerUrl, nil)

	if err != nil {
		panic(err)
	}

	defer wsClient.Close()

	done := make(chan struct{})

	scheduler, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	var sk *big.Int
	var pk *bls12381.G2Affine
	var pkBytes [96]byte

	if config.Secrets.InConfig("secretKey") {

		decoded := base58.Decode(config.Secrets.GetString("secretKey"))

		sk = new(big.Int)
		sk.SetBytes(decoded)

		pk = bls.GetPublicKey(sk)
		pkBytes = pk.Bytes()

	} else {
		sk, pk, err = bls.GenerateKeyPair()
		pkBytes = pk.Bytes()

		config.Secrets.Set("secretKey", base58.Encode(sk.Bytes()))
		config.Secrets.Set("publicKey", base58.Encode(pkBytes[:]))

		err := config.Secrets.WriteConfig()

		if err != nil {
			panic(err)
		}
	}

	spk := bls.GetShortPublicKey(sk)
	spkBytes := spk.Bytes()

	if err != nil {
		panic(err)
	}

	pkStr := address.Calculate(pkBytes[:])
	fmt.Printf("Unchained Address: %s\n", pkStr)

	hello := Signer{
		Name:           config.Config.GetString("name"),
		PublicKey:      pkBytes,
		ShortPublicKey: spkBytes,
	}

	helloPayload, err := msgpack.Marshal(&hello)

	if err != nil {
		panic(err)
	}

	isSocketClosed := false

	go func() {
		defer close(done)

		for {
			_, payload, err := wsClient.ReadMessage()

			if err != nil || payload[0] == 5 {

				if err != nil {
					fmt.Println("Read error:", err)
				} else {
					fmt.Printf("Broker error: %s\n", payload[1:])
				}

				isSocketClosed = true

				if websocket.IsUnexpectedCloseError(err) {
					for i := 1; i < 6; i++ {
						time.Sleep(time.Duration(i) * 3 * time.Second)
						wsClient, _, err = websocket.DefaultDialer.Dial(brokerUrl, nil)
						if err == nil {
							isSocketClosed = false
							wsClient.WriteMessage(
								websocket.BinaryMessage,
								append([]byte{0}, helloPayload...),
							)
						}
					}
				}

				if isSocketClosed {
					return
				}
			}

			if err != nil {
				switch payload[0] {
				// TODO: Make a table of call codes
				case 2:
					fmt.Printf("Unchained feedback: %s\n", payload[1:])

				case 4:
					// TODO: Refactor into a function
					// TODO: Check for errors!
					var challenge kosk.Challenge
					msgpack.Unmarshal(payload[1:], &challenge)

					signature, _ := bls.Sign(*sk, challenge.Random[:])
					challenge.Signature = signature.Bytes()

					koskPayload, _ := msgpack.Marshal(challenge)

					wsClient.WriteMessage(
						websocket.BinaryMessage,
						append([]byte{3}, koskPayload...),
					)

					if err != nil {
						fmt.Println("write:", err)
					}

				default:
					fmt.Printf("Received unknown call code: %d\n", payload[0])
				}
			}

		}
	}()

	wsClient.WriteMessage(websocket.BinaryMessage, append([]byte{0}, helloPayload...))

	_, err = scheduler.NewJob(
		gocron.DurationJob(5*time.Second),
		gocron.NewTask(
			func() {

				if isSocketClosed {
					return
				}

				blockNumber, price, err := GetPriceFromPair(etherUsdPairAddr, 6, true)

				if err != nil {
					return
				}

				if lastBlock == *blockNumber {
					return
				}

				lastBlock = *blockNumber

				var priceF big.Float
				priceF.Quo(new(big.Float).SetInt(price), &tenEighteenF)

				fmt.Printf("%d -> $%.18f\n", *blockNumber, &priceF)

				priceInfo := PriceInfo{Price: *price, Block: *blockNumber}
				toHash, err := msgpack.Marshal(&priceInfo)

				if err != nil {
					panic(err)
				}

				signature, _ := bls.Sign(*sk, toHash)
				compressedSignature := signature.Bytes()

				priceReport := PriceReport{
					PriceInfo: priceInfo,
					Signature: compressedSignature,
				}

				payload, err := msgpack.Marshal(&priceReport)

				if err != nil {
					panic(err)
				}

				//fmt.Printf("%x\n", payload)
				if !isSocketClosed {
					wsClient.WriteMessage(websocket.BinaryMessage, append([]byte{1}, payload...))
				}

				// fmt.Printf("%x\n", signature.Bytes())
				// ok, _ := bls.Verify(signature, hash, *pk)
				// fmt.Printf("Is OK? %t\n", ok)
			},
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

func init() {

	DebouncedSaveSignatures = utils.Debounce[uint64, uint64](5*time.Second, SaveSignatures)

	twoOneNinetyTwo.Exp(big.NewInt(2), big.NewInt(192), nil)
	tenEighteen.Exp(big.NewInt(10), big.NewInt(18), nil)
	tenEighteenF.SetInt(&tenEighteen)

	var err error
	priceCache, err = lru.New[uint64, big.Int](24)

	if err != nil {
		panic(err)
	}

	signatureCache, err = lru.New[uint64, []Signature](24)

	if err != nil {
		panic(err)
	}

	aggregateCache, err = lru.New[uint64, bls12381.G1Affine](24)

	if err != nil {
		panic(err)
	}
}
