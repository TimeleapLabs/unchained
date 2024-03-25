package datasets

import (
	"github.com/KenshiTech/unchained/constants"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

//
//type TokenKey struct {
//	Name   string
//	Pair   string
//	Chain  string
//	Delta  int64
//	Invert bool
//	Cross  string
//}
//
//type AssetKey struct {
//	Token TokenKey
//	Block uint64
//}
//
//type PriceInfo struct {
//	Asset AssetKey
//	Price big.Int
//}
//
//type PriceReport struct {
//	PriceInfo PriceInfo
//	Signature [48]byte
//}

func (m *PriceReport) Protobuf() ([]byte, error) {
	protoModel, err := proto.Marshal(m)
	if err != nil {
		log.Err(err)
		return nil, constants.InternalError
	}

	return protoModel, nil
}

//
//type BroadcastPricePacket struct {
//	Info      PriceInfo
//	Signature [48]byte
//	Signer    bls.Signer
//}
