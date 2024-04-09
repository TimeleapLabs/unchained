package eip712

import (
	"fmt"
	"math/big"

	"github.com/KenshiTech/unchained/ethereum"
	"github.com/KenshiTech/unchained/ethereum/contracts"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

var domain apitypes.TypedDataDomain

func bytesToUnchainedSignature(b []byte) *contracts.UnchainedStakingSignature {
	return &contracts.UnchainedStakingSignature{
		V: b[64],
		R: [32]byte(b[:32]),
		S: [32]byte(b[32:64]),
	}
}

func signEip712Message(evmSigner *ethereum.EvmSigner, data *apitypes.TypedData) (*contracts.UnchainedStakingSignature, error) {
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

	return bytesToUnchainedSignature(signature), nil
}

func SignTransferRequest(evmSigner *ethereum.EvmSigner, request *contracts.UnchainedStakingEIP712Transfer) (*contracts.UnchainedStakingSignature, error) {
	data := &apitypes.TypedData{
		Types:       Types,
		PrimaryType: "Transfer",
		Domain:      domain,
		Message: map[string]interface{}{
			"signer": evmSigner.Address,
			"from":   request.From,
			"to":     request.To,
			"amount": request.Amount,
			"nftIds": request.NftIds,
			"nonces": request.Nonces,
		},
	}

	return signEip712Message(evmSigner, data)
}

func SignSetParamsRequest(evmSigner *ethereum.EvmSigner, request *contracts.UnchainedStakingEIP712SetParams) (*contracts.UnchainedStakingSignature, error) {
	data := &apitypes.TypedData{
		Types:       Types,
		PrimaryType: "SetParams",
		Domain:      domain,
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

	return signEip712Message(evmSigner, data)
}

func SignSetNftPriceRequest(evmSigner *ethereum.EvmSigner, request *contracts.UnchainedStakingEIP712SetNftPrice) (*contracts.UnchainedStakingSignature, error) {
	data := &apitypes.TypedData{
		Types:       Types,
		PrimaryType: "SetNftPrice",
		Domain:      domain,
		Message: map[string]interface{}{
			"requester": evmSigner.Address,
			"nftId":     request.NftId,
			"price":     request.Price,
			"nonce":     request.Nonce,
		},
	}

	return signEip712Message(evmSigner, data)
}

func InitDomain(chainID *big.Int, verifyingContract string) {
	domain = apitypes.TypedDataDomain{
		Name:              "Unchained",
		Version:           "1.0.0",
		ChainId:           math.NewHexOrDecimal256(chainID.Int64()),
		VerifyingContract: verifyingContract,
	}
}
