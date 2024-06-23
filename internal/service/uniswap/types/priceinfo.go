package types

import (
	"math/big"

	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/utils"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

type PriceInfo struct {
	Asset AssetKey
	Price big.Int
}

func (p *PriceInfo) Sia() sia.Sia {
	return sia.New().
		EmbedBytes(p.Asset.Sia().Bytes()).
		AddBigInt(&p.Price)
}

func (p *PriceInfo) FromBytes(payload []byte) *PriceInfo {
	siaMessage := sia.NewFromBytes(payload)
	return p.FromSia(siaMessage)
}

func (p *PriceInfo) FromSia(sia sia.Sia) *PriceInfo {
	p.Asset.FromSia(sia)
	p.Price = *sia.ReadBigInt()

	return p
}

func (p *PriceInfo) Bls() (bls12381.G1Affine, error) {
	hash, err := bls.Hash(p.Sia().Bytes())
	if err != nil {
		utils.Logger.With("err", err).Error("Can't hash bls")
		return bls12381.G1Affine{}, err
	}

	return hash, err
}
