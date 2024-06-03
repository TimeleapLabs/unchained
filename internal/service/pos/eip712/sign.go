package eip712

import (
	"fmt"
	"math/big"

	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

type Signer struct {
	domain apitypes.TypedDataDomain
}

// TODO: Rewrite to use Schnorr signature scheme
func (s *Signer) SignEip712Message(evmSigner *ethereum.Signer, data *apitypes.TypedData) ([]byte, error) {
	domainSeparator, err := data.HashStruct("EIP712Domain", data.Domain.Map())
	if err != nil {
		return nil, err
	}

	typedDataHash, err := data.HashStruct(data.PrimaryType, data.Message)
	if err != nil {
		return nil, err
	}

	message := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	messageHash := crypto.Keccak256(message)

	// This should be replaced with Schnorr signature scheme
	signature, err := crypto.Sign(messageHash, evmSigner.PrivateKey)
	if err != nil {
		return nil, err
	}

	if signature[64] < 27 {
		signature[64] += 27
	}

	return signature, nil
}

func New(chainID *big.Int, verifyingContract string) *Signer {
	return &Signer{
		domain: apitypes.TypedDataDomain{
			Name:              "Unchained",
			Version:           "1.0.0",
			ChainId:           math.NewHexOrDecimal256(chainID.Int64()),
			VerifyingContract: verifyingContract,
		},
	}
}
