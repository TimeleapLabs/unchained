package eip712

import (
	"math/big"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/crypto"

	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum/contracts"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

type Signer struct {
	domain apitypes.TypedDataDomain
}

func (s *Signer) SignTransferRequest(request *contracts.UnchainedStakingEIP712Transfer) (*contracts.UnchainedStakingSignature, error) {
	data := &apitypes.TypedData{
		Types:       Types,
		PrimaryType: "Transfer",
		Domain:      s.domain,
		Message: map[string]interface{}{
			"signer": config.App.Secret.EvmAddress,
			"from":   request.From,
			"to":     request.To,
			"amount": request.Amount,
			"nftIds": request.NftIds,
			"nonces": request.Nonces,
		},
	}

	dataBytes, err := TypedDataToByte(data)
	if err != nil {
		return nil, err
	}

	signedData, err := crypto.Identity.Eth.Sign(dataBytes)
	if err != nil {
		return nil, err
	}

	return NewUnchainedSignatureFromBytes(signedData), nil
}

func (s *Signer) SignSetParamsRequest(request *contracts.UnchainedStakingEIP712SetParams) (*contracts.UnchainedStakingSignature, error) {
	data := &apitypes.TypedData{
		Types:       Types,
		PrimaryType: "SetParams",
		Domain:      s.domain,
		Message: map[string]interface{}{
			"requester":  config.App.Secret.EvmAddress,
			"token":      request.Token,
			"nft":        request.Nft,
			"nftTracker": request.NftTracker,
			"threshold":  request.Threshold,
			"expiration": request.Expiration,
			"nonce":      request.Nonce,
		},
	}

	dataBytes, err := TypedDataToByte(data)
	if err != nil {
		return nil, err
	}

	signedData, err := crypto.Identity.Eth.Sign(dataBytes)
	if err != nil {
		return nil, err
	}

	return NewUnchainedSignatureFromBytes(signedData), nil
}

func (s *Signer) SignSetNftPriceRequest(request *contracts.UnchainedStakingEIP712SetNftPrice) (*contracts.UnchainedStakingSignature, error) {
	data := &apitypes.TypedData{
		Types:       Types,
		PrimaryType: "SetNftPrice",
		Domain:      s.domain,
		Message: map[string]interface{}{
			"requester": config.App.Secret.EvmAddress,
			"nftId":     request.NftId,
			"price":     request.Price,
			"nonce":     request.Nonce,
		},
	}

	dataBytes, err := TypedDataToByte(data)
	if err != nil {
		return nil, err
	}

	signedData, err := crypto.Identity.Eth.Sign(dataBytes)
	if err != nil {
		return nil, err
	}

	return NewUnchainedSignatureFromBytes(signedData), nil
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
