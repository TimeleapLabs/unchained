package eip712

import (
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

var Types = apitypes.Types{
	"EIP712Transfer": {
		{Name: "signer", Type: "address"},
		{Name: "from", Type: "address"},
		{Name: "to", Type: "address"},
		{Name: "amount", Type: "uint256"},
		{Name: "nftIds", Type: "uint256[]"},
		{Name: "nonces", Type: "uint256[]"},
	},

	"EIP712SetParams": {
		{Name: "requester", Type: "address"},
		{Name: "token", Type: "address"},
		{Name: "nft", Type: "address"},
		{Name: "nftTracker", Type: "address"},
		{Name: "threshold", Type: "uint256"},
		{Name: "expiration", Type: "uint256"},
		{Name: "nonce", Type: "uint256"},
	},

	"EIP712SetNftPrice": {
		{Name: "requester", Type: "address"},
		{Name: "nftId", Type: "uint256"},
		{Name: "price", Type: "uint256"},
		{Name: "nonce", Type: "uint256"},
	},
}
