package eip712

import (
	"fmt"

	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum/contracts"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func NewUnchainedSignatureFromBytes(signature []byte) *contracts.UnchainedStakingSignature {
	return &contracts.UnchainedStakingSignature{
		V: signature[64],
		R: [32]byte(signature[:32]),
		S: [32]byte(signature[32:64]),
	}
}

func TypedDataToByte(data *apitypes.TypedData) ([]byte, error) {
	domainSeparator, err := data.HashStruct("EIP712Domain", data.Domain.Map())
	if err != nil {
		return nil, err
	}

	typedDataHash, err := data.HashStruct(data.PrimaryType, data.Message)
	if err != nil {
		return nil, err
	}
	message := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	messageHash := ethCrypto.Keccak256(message)

	return messageHash, nil
}
