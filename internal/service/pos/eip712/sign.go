package eip712

import (
	"fmt"
	"math/big"

	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum/contracts"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

type Signer struct {
	domain apitypes.TypedDataDomain
}

func (s *Signer) bytesToUnchainedSignature(signature []byte) *contracts.UnchainedStakingSignature {
	return &contracts.UnchainedStakingSignature{
		V: signature[64],
		R: [32]byte(signature[:32]),
		S: [32]byte(signature[32:64]),
	}
}

func (s *Signer) signEip712Message(evmSigner *ethereum.EvmSigner, data *apitypes.TypedData) (*contracts.UnchainedStakingSignature, error) {
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

	signature, err := crypto.Sign(messageHash, evmSigner.PrivateKey)
	if err != nil {
		return nil, err
	}

	if signature[64] < 27 {
		signature[64] += 27
	}

	return s.bytesToUnchainedSignature(signature), nil
}

func (s *Signer) SignTransferRequest(evmSigner *ethereum.EvmSigner, request *contracts.UnchainedStakingEIP712Transfer) (*contracts.UnchainedStakingSignature, error) {
	data := &apitypes.TypedData{
		Types:       Types,
		PrimaryType: "Transfer",
		Domain:      s.domain,
		Message: map[string]interface{}{
			"signer": evmSigner.Address,
			"from":   request.From,
			"to":     request.To,
			"amount": request.Amount,
			"nftIds": request.NftIds,
			"nonces": request.Nonces,
		},
	}

	return s.signEip712Message(evmSigner, data)
}

func (s *Signer) SignSetParamsRequest(evmSigner *ethereum.EvmSigner, request *contracts.UnchainedStakingEIP712SetParams) (*contracts.UnchainedStakingSignature, error) {
	data := &apitypes.TypedData{
		Types:       Types,
		PrimaryType: "SetParams",
		Domain:      s.domain,
		Message: map[string]interface{}{
			"requester":  evmSigner.Address,
			"token":      request.Token,
			"nft":        request.Nft,
			"nftTracker": request.NftTracker,
			"threshold":  request.Threshold,
			"expiration": request.Expiration,
			"nonce":      request.Nonce,
		},
	}

	return s.signEip712Message(evmSigner, data)
}

func (s *Signer) SignSetNftPriceRequest(evmSigner *ethereum.EvmSigner, request *contracts.UnchainedStakingEIP712SetNftPrice) (*contracts.UnchainedStakingSignature, error) {
	data := &apitypes.TypedData{
		Types:       Types,
		PrimaryType: "SetNftPrice",
		Domain:      s.domain,
		Message: map[string]interface{}{
			"requester": evmSigner.Address,
			"nftId":     request.NftId,
			"price":     request.Price,
			"nonce":     request.Nonce,
		},
	}

	return s.signEip712Message(evmSigner, data)
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
