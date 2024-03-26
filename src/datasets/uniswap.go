package datasets

import (
	"github.com/KenshiTech/unchained/constants"
	"github.com/KenshiTech/unchained/crypto/bls"
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

func (m *TokenKey) Protobuf() ([]byte, error) {
	protoModel, err := proto.Marshal(m)
	if err != nil {
		log.Err(err)
		return nil, constants.ErrInternalError
	}

	return protoModel, nil
}

func (m *PriceInfo) Protobuf() ([]byte, error) {
	protoModel, err := proto.Marshal(m)
	if err != nil {
		log.Err(err)
		return nil, constants.ErrInternalError
	}

	return protoModel, nil
}

func NewSigner(input *bls.Signer) *Signer {
	return &Signer{
		Name:           input.Name,
		EvmWallet:      input.EvmWallet,
		PublicKey:      input.PublicKey[:],
		ShortPublicKey: input.ShortPublicKey[:],
	}
}

func (m *PriceReport) Protobuf() ([]byte, error) {
	protoModel, err := proto.Marshal(m)
	if err != nil {
		log.Err(err)
		return nil, constants.ErrInternalError
	}

	return protoModel, nil
}

func (m *BroadcastPricePacket) Protobuf() ([]byte, error) {
	protoModel, err := proto.Marshal(m)
	if err != nil {
		log.Err(err)
		return nil, constants.ErrInternalError
	}

	return protoModel, nil
}
