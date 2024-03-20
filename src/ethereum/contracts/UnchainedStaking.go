// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// UnchainedStakingEIP712SetParams is an auto generated low-level Go binding around an user-defined struct.
type UnchainedStakingEIP712SetParams struct {
	Requester  common.Address
	Token      common.Address
	Nft        common.Address
	Threshold  *big.Int
	Expiration *big.Int
	Nonce      *big.Int
}

// UnchainedStakingEIP712SetParamsKey is an auto generated low-level Go binding around an user-defined struct.
type UnchainedStakingEIP712SetParamsKey struct {
	Token      common.Address
	Nft        common.Address
	Threshold  *big.Int
	Expiration *big.Int
	Nonce      *big.Int
}

// UnchainedStakingEIP712SetSigner is an auto generated low-level Go binding around an user-defined struct.
type UnchainedStakingEIP712SetSigner struct {
	Staker common.Address
	Signer common.Address
}

// UnchainedStakingEIP712Transfer is an auto generated low-level Go binding around an user-defined struct.
type UnchainedStakingEIP712Transfer struct {
	Signer common.Address
	From   common.Address
	To     common.Address
	Amount *big.Int
	Nonces []*big.Int
}

// UnchainedStakingEIP712TransferKey is an auto generated low-level Go binding around an user-defined struct.
type UnchainedStakingEIP712TransferKey struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Nonces []*big.Int
}

// UnchainedStakingParamsInfo is an auto generated low-level Go binding around an user-defined struct.
type UnchainedStakingParamsInfo struct {
	Token      common.Address
	Nft        common.Address
	Threshold  *big.Int
	Expiration *big.Int
	Voted      *big.Int
	Nonce      *big.Int
	Accepted   bool
}

// UnchainedStakingSignature is an auto generated low-level Go binding around an user-defined struct.
type UnchainedStakingSignature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// UnchainedStakingStake is an auto generated low-level Go binding around an user-defined struct.
type UnchainedStakingStake struct {
	Amount *big.Int
	Unlock *big.Int
	NftIds []*big.Int
}

// UnchainedStakingTransferInfo is an auto generated low-level Go binding around an user-defined struct.
type UnchainedStakingTransferInfo struct {
	From     common.Address
	To       common.Address
	Amount   *big.Int
	Voted    *big.Int
	Accepted bool
	Nonces   []*big.Int
}

// UnchainedStakingMetaData contains all meta data concerning the UnchainedStaking contract.
var UnchainedStakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"consensusLock\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressInUse\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"AlreadyAccused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyStaked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"AlreadyVoted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlsNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DurationZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Forbidden\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"NonceUsed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"NotConsumer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotUnlocked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"StakeExpiresBeforeVote\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"TopicExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"VotingPowerZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongEIP712Signature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongNFT\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"accused\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"accuser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"incident\",\"type\":\"bytes32\"}],\"name\":\"Accused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"from\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"}],\"name\":\"BlsAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlock\",\"type\":\"uint256\"}],\"name\":\"Extended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"ParamsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SignerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"nftIds\",\"type\":\"uint256[]\"}],\"name\":\"StakeIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"nftIds\",\"type\":\"uint256[]\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferIn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"nonces\",\"type\":\"uint256[]\"}],\"name\":\"TransferOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"nftIds\",\"type\":\"uint256[]\"}],\"name\":\"UnStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"VotedForParams\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"evm\",\"type\":\"address\"}],\"name\":\"blsAddressOf\",\"outputs\":[{\"internalType\":\"bytes20\",\"name\":\"\",\"type\":\"bytes20\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"bls\",\"type\":\"bytes20\"}],\"name\":\"evmAddressOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"extend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getParams\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"accepted\",\"type\":\"bool\"}],\"internalType\":\"structUnchainedStaking.ParamsInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structUnchainedStaking.EIP712SetParamsKey\",\"name\":\"key\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"}],\"name\":\"getRequestedSetParams\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"nonces\",\"type\":\"uint256[]\"}],\"internalType\":\"structUnchainedStaking.EIP712TransferKey\",\"name\":\"key\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"transferer\",\"type\":\"address\"}],\"name\":\"getRequestedTransferOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structUnchainedStaking.EIP712SetParamsKey\",\"name\":\"key\",\"type\":\"tuple\"}],\"name\":\"getSetParamsData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"accepted\",\"type\":\"bool\"}],\"internalType\":\"structUnchainedStaking.ParamsInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalVotingPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"nonces\",\"type\":\"uint256[]\"}],\"internalType\":\"structUnchainedStaking.EIP712TransferKey\",\"name\":\"key\",\"type\":\"tuple\"}],\"name\":\"getTransferOutData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voted\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"accepted\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"nonces\",\"type\":\"uint256[]\"}],\"internalType\":\"structUnchainedStaking.TransferInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransferredIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"nftIds\",\"type\":\"uint256[]\"}],\"name\":\"increaseStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"blsAddress\",\"type\":\"bytes20\"}],\"name\":\"setBlsAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structUnchainedStaking.EIP712SetParams[]\",\"name\":\"eip712SetParams\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structUnchainedStaking.Signature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"setParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"internalType\":\"structUnchainedStaking.EIP712SetSigner\",\"name\":\"eip712SetSigner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structUnchainedStaking.Signature\",\"name\":\"stakerSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structUnchainedStaking.Signature\",\"name\":\"signerSignature\",\"type\":\"tuple\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"signerToStaker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"nftIds\",\"type\":\"uint256[]\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"evm\",\"type\":\"address\"}],\"name\":\"stakeOf\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"nftIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structUnchainedStaking.Stake\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"bls\",\"type\":\"bytes20\"}],\"name\":\"stakeOf\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"nftIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structUnchainedStaking.Stake\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"stakerToSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"nonces\",\"type\":\"uint256[]\"}],\"internalType\":\"structUnchainedStaking.EIP712Transfer[]\",\"name\":\"eip712Transferes\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structUnchainedStaking.Signature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"transferOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structUnchainedStaking.EIP712SetParams\",\"name\":\"eip712SetParam\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structUnchainedStaking.Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"internalType\":\"structUnchainedStaking.EIP712SetSigner\",\"name\":\"eip712SetSigner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structUnchainedStaking.Signature\",\"name\":\"stakerSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structUnchainedStaking.Signature\",\"name\":\"signerSignature\",\"type\":\"tuple\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"nonces\",\"type\":\"uint256[]\"}],\"internalType\":\"structUnchainedStaking.EIP712Transfer\",\"name\":\"eip712Transfer\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structUnchainedStaking.Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UnchainedStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use UnchainedStakingMetaData.ABI instead.
var UnchainedStakingABI = UnchainedStakingMetaData.ABI

// UnchainedStaking is an auto generated Go binding around an Ethereum contract.
type UnchainedStaking struct {
	UnchainedStakingCaller     // Read-only binding to the contract
	UnchainedStakingTransactor // Write-only binding to the contract
	UnchainedStakingFilterer   // Log filterer for contract events
}

// UnchainedStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnchainedStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnchainedStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnchainedStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnchainedStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnchainedStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnchainedStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnchainedStakingSession struct {
	Contract     *UnchainedStaking // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnchainedStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnchainedStakingCallerSession struct {
	Contract *UnchainedStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// UnchainedStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnchainedStakingTransactorSession struct {
	Contract     *UnchainedStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// UnchainedStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnchainedStakingRaw struct {
	Contract *UnchainedStaking // Generic contract binding to access the raw methods on
}

// UnchainedStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnchainedStakingCallerRaw struct {
	Contract *UnchainedStakingCaller // Generic read-only contract binding to access the raw methods on
}

// UnchainedStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnchainedStakingTransactorRaw struct {
	Contract *UnchainedStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnchainedStaking creates a new instance of UnchainedStaking, bound to a specific deployed contract.
func NewUnchainedStaking(address common.Address, backend bind.ContractBackend) (*UnchainedStaking, error) {
	contract, err := bindUnchainedStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UnchainedStaking{UnchainedStakingCaller: UnchainedStakingCaller{contract: contract}, UnchainedStakingTransactor: UnchainedStakingTransactor{contract: contract}, UnchainedStakingFilterer: UnchainedStakingFilterer{contract: contract}}, nil
}

// NewUnchainedStakingCaller creates a new read-only instance of UnchainedStaking, bound to a specific deployed contract.
func NewUnchainedStakingCaller(address common.Address, caller bind.ContractCaller) (*UnchainedStakingCaller, error) {
	contract, err := bindUnchainedStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnchainedStakingCaller{contract: contract}, nil
}

// NewUnchainedStakingTransactor creates a new write-only instance of UnchainedStaking, bound to a specific deployed contract.
func NewUnchainedStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*UnchainedStakingTransactor, error) {
	contract, err := bindUnchainedStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnchainedStakingTransactor{contract: contract}, nil
}

// NewUnchainedStakingFilterer creates a new log filterer instance of UnchainedStaking, bound to a specific deployed contract.
func NewUnchainedStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*UnchainedStakingFilterer, error) {
	contract, err := bindUnchainedStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnchainedStakingFilterer{contract: contract}, nil
}

// bindUnchainedStaking binds a generic wrapper to an already deployed contract.
func bindUnchainedStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UnchainedStakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnchainedStaking *UnchainedStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnchainedStaking.Contract.UnchainedStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnchainedStaking *UnchainedStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnchainedStaking.Contract.UnchainedStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnchainedStaking *UnchainedStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnchainedStaking.Contract.UnchainedStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnchainedStaking *UnchainedStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnchainedStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnchainedStaking *UnchainedStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnchainedStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnchainedStaking *UnchainedStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnchainedStaking.Contract.contract.Transact(opts, method, params...)
}

// BlsAddressOf is a free data retrieval call binding the contract method 0x9e95a1a5.
//
// Solidity: function blsAddressOf(address evm) view returns(bytes20)
func (_UnchainedStaking *UnchainedStakingCaller) BlsAddressOf(opts *bind.CallOpts, evm common.Address) ([20]byte, error) {
	var out []interface{}
	err := _UnchainedStaking.contract.Call(opts, &out, "blsAddressOf", evm)

	if err != nil {
		return *new([20]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([20]byte)).(*[20]byte)

	return out0, err

}

// BlsAddressOf is a free data retrieval call binding the contract method 0x9e95a1a5.
//
// Solidity: function blsAddressOf(address evm) view returns(bytes20)
func (_UnchainedStaking *UnchainedStakingSession) BlsAddressOf(evm common.Address) ([20]byte, error) {
	return _UnchainedStaking.Contract.BlsAddressOf(&_UnchainedStaking.CallOpts, evm)
}

// BlsAddressOf is a free data retrieval call binding the contract method 0x9e95a1a5.
//
// Solidity: function blsAddressOf(address evm) view returns(bytes20)
func (_UnchainedStaking *UnchainedStakingCallerSession) BlsAddressOf(evm common.Address) ([20]byte, error) {
	return _UnchainedStaking.Contract.BlsAddressOf(&_UnchainedStaking.CallOpts, evm)
}

// EvmAddressOf is a free data retrieval call binding the contract method 0x40043ead.
//
// Solidity: function evmAddressOf(bytes20 bls) view returns(address)
func (_UnchainedStaking *UnchainedStakingCaller) EvmAddressOf(opts *bind.CallOpts, bls [20]byte) (common.Address, error) {
	var out []interface{}
	err := _UnchainedStaking.contract.Call(opts, &out, "evmAddressOf", bls)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EvmAddressOf is a free data retrieval call binding the contract method 0x40043ead.
//
// Solidity: function evmAddressOf(bytes20 bls) view returns(address)
func (_UnchainedStaking *UnchainedStakingSession) EvmAddressOf(bls [20]byte) (common.Address, error) {
	return _UnchainedStaking.Contract.EvmAddressOf(&_UnchainedStaking.CallOpts, bls)
}

// EvmAddressOf is a free data retrieval call binding the contract method 0x40043ead.
//
// Solidity: function evmAddressOf(bytes20 bls) view returns(address)
func (_UnchainedStaking *UnchainedStakingCallerSession) EvmAddressOf(bls [20]byte) (common.Address, error) {
	return _UnchainedStaking.Contract.EvmAddressOf(&_UnchainedStaking.CallOpts, bls)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_UnchainedStaking *UnchainedStakingCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnchainedStaking.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_UnchainedStaking *UnchainedStakingSession) GetChainId() (*big.Int, error) {
	return _UnchainedStaking.Contract.GetChainId(&_UnchainedStaking.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_UnchainedStaking *UnchainedStakingCallerSession) GetChainId() (*big.Int, error) {
	return _UnchainedStaking.Contract.GetChainId(&_UnchainedStaking.CallOpts)
}

// GetConsensusThreshold is a free data retrieval call binding the contract method 0xd42791e1.
//
// Solidity: function getConsensusThreshold() view returns(uint256)
func (_UnchainedStaking *UnchainedStakingCaller) GetConsensusThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnchainedStaking.contract.Call(opts, &out, "getConsensusThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConsensusThreshold is a free data retrieval call binding the contract method 0xd42791e1.
//
// Solidity: function getConsensusThreshold() view returns(uint256)
func (_UnchainedStaking *UnchainedStakingSession) GetConsensusThreshold() (*big.Int, error) {
	return _UnchainedStaking.Contract.GetConsensusThreshold(&_UnchainedStaking.CallOpts)
}

// GetConsensusThreshold is a free data retrieval call binding the contract method 0xd42791e1.
//
// Solidity: function getConsensusThreshold() view returns(uint256)
func (_UnchainedStaking *UnchainedStakingCallerSession) GetConsensusThreshold() (*big.Int, error) {
	return _UnchainedStaking.Contract.GetConsensusThreshold(&_UnchainedStaking.CallOpts)
}

// GetParams is a free data retrieval call binding the contract method 0x5e615a6b.
//
// Solidity: function getParams() view returns((address,address,uint256,uint256,uint256,uint256,bool))
func (_UnchainedStaking *UnchainedStakingCaller) GetParams(opts *bind.CallOpts) (UnchainedStakingParamsInfo, error) {
	var out []interface{}
	err := _UnchainedStaking.contract.Call(opts, &out, "getParams")

	if err != nil {
		return *new(UnchainedStakingParamsInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(UnchainedStakingParamsInfo)).(*UnchainedStakingParamsInfo)

	return out0, err

}

// GetParams is a free data retrieval call binding the contract method 0x5e615a6b.
//
// Solidity: function getParams() view returns((address,address,uint256,uint256,uint256,uint256,bool))
func (_UnchainedStaking *UnchainedStakingSession) GetParams() (UnchainedStakingParamsInfo, error) {
	return _UnchainedStaking.Contract.GetParams(&_UnchainedStaking.CallOpts)
}

// GetParams is a free data retrieval call binding the contract method 0x5e615a6b.
//
// Solidity: function getParams() view returns((address,address,uint256,uint256,uint256,uint256,bool))
func (_UnchainedStaking *UnchainedStakingCallerSession) GetParams() (UnchainedStakingParamsInfo, error) {
	return _UnchainedStaking.Contract.GetParams(&_UnchainedStaking.CallOpts)
}

// GetRequestedSetParams is a free data retrieval call binding the contract method 0x628beb07.
//
// Solidity: function getRequestedSetParams((address,address,uint256,uint256,uint256) key, address requester) view returns(bool)
func (_UnchainedStaking *UnchainedStakingCaller) GetRequestedSetParams(opts *bind.CallOpts, key UnchainedStakingEIP712SetParamsKey, requester common.Address) (bool, error) {
	var out []interface{}
	err := _UnchainedStaking.contract.Call(opts, &out, "getRequestedSetParams", key, requester)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetRequestedSetParams is a free data retrieval call binding the contract method 0x628beb07.
//
// Solidity: function getRequestedSetParams((address,address,uint256,uint256,uint256) key, address requester) view returns(bool)
func (_UnchainedStaking *UnchainedStakingSession) GetRequestedSetParams(key UnchainedStakingEIP712SetParamsKey, requester common.Address) (bool, error) {
	return _UnchainedStaking.Contract.GetRequestedSetParams(&_UnchainedStaking.CallOpts, key, requester)
}

// GetRequestedSetParams is a free data retrieval call binding the contract method 0x628beb07.
//
// Solidity: function getRequestedSetParams((address,address,uint256,uint256,uint256) key, address requester) view returns(bool)
func (_UnchainedStaking *UnchainedStakingCallerSession) GetRequestedSetParams(key UnchainedStakingEIP712SetParamsKey, requester common.Address) (bool, error) {
	return _UnchainedStaking.Contract.GetRequestedSetParams(&_UnchainedStaking.CallOpts, key, requester)
}

// GetRequestedTransferOut is a free data retrieval call binding the contract method 0x7a04820e.
//
// Solidity: function getRequestedTransferOut((address,address,uint256,uint256[]) key, address transferer) view returns(bool)
func (_UnchainedStaking *UnchainedStakingCaller) GetRequestedTransferOut(opts *bind.CallOpts, key UnchainedStakingEIP712TransferKey, transferer common.Address) (bool, error) {
	var out []interface{}
	err := _UnchainedStaking.contract.Call(opts, &out, "getRequestedTransferOut", key, transferer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetRequestedTransferOut is a free data retrieval call binding the contract method 0x7a04820e.
//
// Solidity: function getRequestedTransferOut((address,address,uint256,uint256[]) key, address transferer) view returns(bool)
func (_UnchainedStaking *UnchainedStakingSession) GetRequestedTransferOut(key UnchainedStakingEIP712TransferKey, transferer common.Address) (bool, error) {
	return _UnchainedStaking.Contract.GetRequestedTransferOut(&_UnchainedStaking.CallOpts, key, transferer)
}

// GetRequestedTransferOut is a free data retrieval call binding the contract method 0x7a04820e.
//
// Solidity: function getRequestedTransferOut((address,address,uint256,uint256[]) key, address transferer) view returns(bool)
func (_UnchainedStaking *UnchainedStakingCallerSession) GetRequestedTransferOut(key UnchainedStakingEIP712TransferKey, transferer common.Address) (bool, error) {
	return _UnchainedStaking.Contract.GetRequestedTransferOut(&_UnchainedStaking.CallOpts, key, transferer)
}

// GetSetParamsData is a free data retrieval call binding the contract method 0x5ad235f6.
//
// Solidity: function getSetParamsData((address,address,uint256,uint256,uint256) key) view returns((address,address,uint256,uint256,uint256,uint256,bool))
func (_UnchainedStaking *UnchainedStakingCaller) GetSetParamsData(opts *bind.CallOpts, key UnchainedStakingEIP712SetParamsKey) (UnchainedStakingParamsInfo, error) {
	var out []interface{}
	err := _UnchainedStaking.contract.Call(opts, &out, "getSetParamsData", key)

	if err != nil {
		return *new(UnchainedStakingParamsInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(UnchainedStakingParamsInfo)).(*UnchainedStakingParamsInfo)

	return out0, err

}

// GetSetParamsData is a free data retrieval call binding the contract method 0x5ad235f6.
//
// Solidity: function getSetParamsData((address,address,uint256,uint256,uint256) key) view returns((address,address,uint256,uint256,uint256,uint256,bool))
func (_UnchainedStaking *UnchainedStakingSession) GetSetParamsData(key UnchainedStakingEIP712SetParamsKey) (UnchainedStakingParamsInfo, error) {
	return _UnchainedStaking.Contract.GetSetParamsData(&_UnchainedStaking.CallOpts, key)
}

// GetSetParamsData is a free data retrieval call binding the contract method 0x5ad235f6.
//
// Solidity: function getSetParamsData((address,address,uint256,uint256,uint256) key) view returns((address,address,uint256,uint256,uint256,uint256,bool))
func (_UnchainedStaking *UnchainedStakingCallerSession) GetSetParamsData(key UnchainedStakingEIP712SetParamsKey) (UnchainedStakingParamsInfo, error) {
	return _UnchainedStaking.Contract.GetSetParamsData(&_UnchainedStaking.CallOpts, key)
}

// GetTotalVotingPower is a free data retrieval call binding the contract method 0x11acc1a7.
//
// Solidity: function getTotalVotingPower() view returns(uint256)
func (_UnchainedStaking *UnchainedStakingCaller) GetTotalVotingPower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnchainedStaking.contract.Call(opts, &out, "getTotalVotingPower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalVotingPower is a free data retrieval call binding the contract method 0x11acc1a7.
//
// Solidity: function getTotalVotingPower() view returns(uint256)
func (_UnchainedStaking *UnchainedStakingSession) GetTotalVotingPower() (*big.Int, error) {
	return _UnchainedStaking.Contract.GetTotalVotingPower(&_UnchainedStaking.CallOpts)
}

// GetTotalVotingPower is a free data retrieval call binding the contract method 0x11acc1a7.
//
// Solidity: function getTotalVotingPower() view returns(uint256)
func (_UnchainedStaking *UnchainedStakingCallerSession) GetTotalVotingPower() (*big.Int, error) {
	return _UnchainedStaking.Contract.GetTotalVotingPower(&_UnchainedStaking.CallOpts)
}

// GetTransferOutData is a free data retrieval call binding the contract method 0x2efc5265.
//
// Solidity: function getTransferOutData((address,address,uint256,uint256[]) key) view returns((address,address,uint256,uint256,bool,uint256[]))
func (_UnchainedStaking *UnchainedStakingCaller) GetTransferOutData(opts *bind.CallOpts, key UnchainedStakingEIP712TransferKey) (UnchainedStakingTransferInfo, error) {
	var out []interface{}
	err := _UnchainedStaking.contract.Call(opts, &out, "getTransferOutData", key)

	if err != nil {
		return *new(UnchainedStakingTransferInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(UnchainedStakingTransferInfo)).(*UnchainedStakingTransferInfo)

	return out0, err

}

// GetTransferOutData is a free data retrieval call binding the contract method 0x2efc5265.
//
// Solidity: function getTransferOutData((address,address,uint256,uint256[]) key) view returns((address,address,uint256,uint256,bool,uint256[]))
func (_UnchainedStaking *UnchainedStakingSession) GetTransferOutData(key UnchainedStakingEIP712TransferKey) (UnchainedStakingTransferInfo, error) {
	return _UnchainedStaking.Contract.GetTransferOutData(&_UnchainedStaking.CallOpts, key)
}

// GetTransferOutData is a free data retrieval call binding the contract method 0x2efc5265.
//
// Solidity: function getTransferOutData((address,address,uint256,uint256[]) key) view returns((address,address,uint256,uint256,bool,uint256[]))
func (_UnchainedStaking *UnchainedStakingCallerSession) GetTransferOutData(key UnchainedStakingEIP712TransferKey) (UnchainedStakingTransferInfo, error) {
	return _UnchainedStaking.Contract.GetTransferOutData(&_UnchainedStaking.CallOpts, key)
}

// GetTransferredIn is a free data retrieval call binding the contract method 0x5be979b6.
//
// Solidity: function getTransferredIn() view returns(uint256)
func (_UnchainedStaking *UnchainedStakingCaller) GetTransferredIn(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnchainedStaking.contract.Call(opts, &out, "getTransferredIn")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransferredIn is a free data retrieval call binding the contract method 0x5be979b6.
//
// Solidity: function getTransferredIn() view returns(uint256)
func (_UnchainedStaking *UnchainedStakingSession) GetTransferredIn() (*big.Int, error) {
	return _UnchainedStaking.Contract.GetTransferredIn(&_UnchainedStaking.CallOpts)
}

// GetTransferredIn is a free data retrieval call binding the contract method 0x5be979b6.
//
// Solidity: function getTransferredIn() view returns(uint256)
func (_UnchainedStaking *UnchainedStakingCallerSession) GetTransferredIn() (*big.Int, error) {
	return _UnchainedStaking.Contract.GetTransferredIn(&_UnchainedStaking.CallOpts)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) view returns(bytes4)
func (_UnchainedStaking *UnchainedStakingCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _UnchainedStaking.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) view returns(bytes4)
func (_UnchainedStaking *UnchainedStakingSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _UnchainedStaking.Contract.OnERC721Received(&_UnchainedStaking.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) view returns(bytes4)
func (_UnchainedStaking *UnchainedStakingCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _UnchainedStaking.Contract.OnERC721Received(&_UnchainedStaking.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UnchainedStaking *UnchainedStakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UnchainedStaking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UnchainedStaking *UnchainedStakingSession) Owner() (common.Address, error) {
	return _UnchainedStaking.Contract.Owner(&_UnchainedStaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UnchainedStaking *UnchainedStakingCallerSession) Owner() (common.Address, error) {
	return _UnchainedStaking.Contract.Owner(&_UnchainedStaking.CallOpts)
}

// SignerToStaker is a free data retrieval call binding the contract method 0xad5a98c5.
//
// Solidity: function signerToStaker(address signer) view returns(address)
func (_UnchainedStaking *UnchainedStakingCaller) SignerToStaker(opts *bind.CallOpts, signer common.Address) (common.Address, error) {
	var out []interface{}
	err := _UnchainedStaking.contract.Call(opts, &out, "signerToStaker", signer)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerToStaker is a free data retrieval call binding the contract method 0xad5a98c5.
//
// Solidity: function signerToStaker(address signer) view returns(address)
func (_UnchainedStaking *UnchainedStakingSession) SignerToStaker(signer common.Address) (common.Address, error) {
	return _UnchainedStaking.Contract.SignerToStaker(&_UnchainedStaking.CallOpts, signer)
}

// SignerToStaker is a free data retrieval call binding the contract method 0xad5a98c5.
//
// Solidity: function signerToStaker(address signer) view returns(address)
func (_UnchainedStaking *UnchainedStakingCallerSession) SignerToStaker(signer common.Address) (common.Address, error) {
	return _UnchainedStaking.Contract.SignerToStaker(&_UnchainedStaking.CallOpts, signer)
}

// StakeOf is a free data retrieval call binding the contract method 0x42623360.
//
// Solidity: function stakeOf(address evm) view returns((uint256,uint256,uint256[]))
func (_UnchainedStaking *UnchainedStakingCaller) StakeOf(opts *bind.CallOpts, evm common.Address) (UnchainedStakingStake, error) {
	var out []interface{}
	err := _UnchainedStaking.contract.Call(opts, &out, "stakeOf", evm)

	if err != nil {
		return *new(UnchainedStakingStake), err
	}

	out0 := *abi.ConvertType(out[0], new(UnchainedStakingStake)).(*UnchainedStakingStake)

	return out0, err

}

// StakeOf is a free data retrieval call binding the contract method 0x42623360.
//
// Solidity: function stakeOf(address evm) view returns((uint256,uint256,uint256[]))
func (_UnchainedStaking *UnchainedStakingSession) StakeOf(evm common.Address) (UnchainedStakingStake, error) {
	return _UnchainedStaking.Contract.StakeOf(&_UnchainedStaking.CallOpts, evm)
}

// StakeOf is a free data retrieval call binding the contract method 0x42623360.
//
// Solidity: function stakeOf(address evm) view returns((uint256,uint256,uint256[]))
func (_UnchainedStaking *UnchainedStakingCallerSession) StakeOf(evm common.Address) (UnchainedStakingStake, error) {
	return _UnchainedStaking.Contract.StakeOf(&_UnchainedStaking.CallOpts, evm)
}

// StakeOf0 is a free data retrieval call binding the contract method 0xf5056014.
//
// Solidity: function stakeOf(bytes20 bls) view returns((uint256,uint256,uint256[]))
func (_UnchainedStaking *UnchainedStakingCaller) StakeOf0(opts *bind.CallOpts, bls [20]byte) (UnchainedStakingStake, error) {
	var out []interface{}
	err := _UnchainedStaking.contract.Call(opts, &out, "stakeOf0", bls)

	if err != nil {
		return *new(UnchainedStakingStake), err
	}

	out0 := *abi.ConvertType(out[0], new(UnchainedStakingStake)).(*UnchainedStakingStake)

	return out0, err

}

// StakeOf0 is a free data retrieval call binding the contract method 0xf5056014.
//
// Solidity: function stakeOf(bytes20 bls) view returns((uint256,uint256,uint256[]))
func (_UnchainedStaking *UnchainedStakingSession) StakeOf0(bls [20]byte) (UnchainedStakingStake, error) {
	return _UnchainedStaking.Contract.StakeOf0(&_UnchainedStaking.CallOpts, bls)
}

// StakeOf0 is a free data retrieval call binding the contract method 0xf5056014.
//
// Solidity: function stakeOf(bytes20 bls) view returns((uint256,uint256,uint256[]))
func (_UnchainedStaking *UnchainedStakingCallerSession) StakeOf0(bls [20]byte) (UnchainedStakingStake, error) {
	return _UnchainedStaking.Contract.StakeOf0(&_UnchainedStaking.CallOpts, bls)
}

// StakerToSigner is a free data retrieval call binding the contract method 0xc7bcae78.
//
// Solidity: function stakerToSigner(address staker) view returns(address)
func (_UnchainedStaking *UnchainedStakingCaller) StakerToSigner(opts *bind.CallOpts, staker common.Address) (common.Address, error) {
	var out []interface{}
	err := _UnchainedStaking.contract.Call(opts, &out, "stakerToSigner", staker)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerToSigner is a free data retrieval call binding the contract method 0xc7bcae78.
//
// Solidity: function stakerToSigner(address staker) view returns(address)
func (_UnchainedStaking *UnchainedStakingSession) StakerToSigner(staker common.Address) (common.Address, error) {
	return _UnchainedStaking.Contract.StakerToSigner(&_UnchainedStaking.CallOpts, staker)
}

// StakerToSigner is a free data retrieval call binding the contract method 0xc7bcae78.
//
// Solidity: function stakerToSigner(address staker) view returns(address)
func (_UnchainedStaking *UnchainedStakingCallerSession) StakerToSigner(staker common.Address) (common.Address, error) {
	return _UnchainedStaking.Contract.StakerToSigner(&_UnchainedStaking.CallOpts, staker)
}

// Verify is a free data retrieval call binding the contract method 0x34c9667c.
//
// Solidity: function verify((address,address,address,uint256,uint256,uint256) eip712SetParam, (uint8,bytes32,bytes32) signature) view returns(bool)
func (_UnchainedStaking *UnchainedStakingCaller) Verify(opts *bind.CallOpts, eip712SetParam UnchainedStakingEIP712SetParams, signature UnchainedStakingSignature) (bool, error) {
	var out []interface{}
	err := _UnchainedStaking.contract.Call(opts, &out, "verify", eip712SetParam, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x34c9667c.
//
// Solidity: function verify((address,address,address,uint256,uint256,uint256) eip712SetParam, (uint8,bytes32,bytes32) signature) view returns(bool)
func (_UnchainedStaking *UnchainedStakingSession) Verify(eip712SetParam UnchainedStakingEIP712SetParams, signature UnchainedStakingSignature) (bool, error) {
	return _UnchainedStaking.Contract.Verify(&_UnchainedStaking.CallOpts, eip712SetParam, signature)
}

// Verify is a free data retrieval call binding the contract method 0x34c9667c.
//
// Solidity: function verify((address,address,address,uint256,uint256,uint256) eip712SetParam, (uint8,bytes32,bytes32) signature) view returns(bool)
func (_UnchainedStaking *UnchainedStakingCallerSession) Verify(eip712SetParam UnchainedStakingEIP712SetParams, signature UnchainedStakingSignature) (bool, error) {
	return _UnchainedStaking.Contract.Verify(&_UnchainedStaking.CallOpts, eip712SetParam, signature)
}

// Verify0 is a free data retrieval call binding the contract method 0x7856cc70.
//
// Solidity: function verify((address,address) eip712SetSigner, (uint8,bytes32,bytes32) stakerSignature, (uint8,bytes32,bytes32) signerSignature) view returns(bool)
func (_UnchainedStaking *UnchainedStakingCaller) Verify0(opts *bind.CallOpts, eip712SetSigner UnchainedStakingEIP712SetSigner, stakerSignature UnchainedStakingSignature, signerSignature UnchainedStakingSignature) (bool, error) {
	var out []interface{}
	err := _UnchainedStaking.contract.Call(opts, &out, "verify0", eip712SetSigner, stakerSignature, signerSignature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify0 is a free data retrieval call binding the contract method 0x7856cc70.
//
// Solidity: function verify((address,address) eip712SetSigner, (uint8,bytes32,bytes32) stakerSignature, (uint8,bytes32,bytes32) signerSignature) view returns(bool)
func (_UnchainedStaking *UnchainedStakingSession) Verify0(eip712SetSigner UnchainedStakingEIP712SetSigner, stakerSignature UnchainedStakingSignature, signerSignature UnchainedStakingSignature) (bool, error) {
	return _UnchainedStaking.Contract.Verify0(&_UnchainedStaking.CallOpts, eip712SetSigner, stakerSignature, signerSignature)
}

// Verify0 is a free data retrieval call binding the contract method 0x7856cc70.
//
// Solidity: function verify((address,address) eip712SetSigner, (uint8,bytes32,bytes32) stakerSignature, (uint8,bytes32,bytes32) signerSignature) view returns(bool)
func (_UnchainedStaking *UnchainedStakingCallerSession) Verify0(eip712SetSigner UnchainedStakingEIP712SetSigner, stakerSignature UnchainedStakingSignature, signerSignature UnchainedStakingSignature) (bool, error) {
	return _UnchainedStaking.Contract.Verify0(&_UnchainedStaking.CallOpts, eip712SetSigner, stakerSignature, signerSignature)
}

// Verify1 is a free data retrieval call binding the contract method 0xf48d0be7.
//
// Solidity: function verify((address,address,address,uint256,uint256[]) eip712Transfer, (uint8,bytes32,bytes32) signature) view returns(bool)
func (_UnchainedStaking *UnchainedStakingCaller) Verify1(opts *bind.CallOpts, eip712Transfer UnchainedStakingEIP712Transfer, signature UnchainedStakingSignature) (bool, error) {
	var out []interface{}
	err := _UnchainedStaking.contract.Call(opts, &out, "verify1", eip712Transfer, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify1 is a free data retrieval call binding the contract method 0xf48d0be7.
//
// Solidity: function verify((address,address,address,uint256,uint256[]) eip712Transfer, (uint8,bytes32,bytes32) signature) view returns(bool)
func (_UnchainedStaking *UnchainedStakingSession) Verify1(eip712Transfer UnchainedStakingEIP712Transfer, signature UnchainedStakingSignature) (bool, error) {
	return _UnchainedStaking.Contract.Verify1(&_UnchainedStaking.CallOpts, eip712Transfer, signature)
}

// Verify1 is a free data retrieval call binding the contract method 0xf48d0be7.
//
// Solidity: function verify((address,address,address,uint256,uint256[]) eip712Transfer, (uint8,bytes32,bytes32) signature) view returns(bool)
func (_UnchainedStaking *UnchainedStakingCallerSession) Verify1(eip712Transfer UnchainedStakingEIP712Transfer, signature UnchainedStakingSignature) (bool, error) {
	return _UnchainedStaking.Contract.Verify1(&_UnchainedStaking.CallOpts, eip712Transfer, signature)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 duration) returns()
func (_UnchainedStaking *UnchainedStakingTransactor) Extend(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _UnchainedStaking.contract.Transact(opts, "extend", duration)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 duration) returns()
func (_UnchainedStaking *UnchainedStakingSession) Extend(duration *big.Int) (*types.Transaction, error) {
	return _UnchainedStaking.Contract.Extend(&_UnchainedStaking.TransactOpts, duration)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 duration) returns()
func (_UnchainedStaking *UnchainedStakingTransactorSession) Extend(duration *big.Int) (*types.Transaction, error) {
	return _UnchainedStaking.Contract.Extend(&_UnchainedStaking.TransactOpts, duration)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0x0062ad9d.
//
// Solidity: function increaseStake(uint256 amount, uint256[] nftIds) returns()
func (_UnchainedStaking *UnchainedStakingTransactor) IncreaseStake(opts *bind.TransactOpts, amount *big.Int, nftIds []*big.Int) (*types.Transaction, error) {
	return _UnchainedStaking.contract.Transact(opts, "increaseStake", amount, nftIds)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0x0062ad9d.
//
// Solidity: function increaseStake(uint256 amount, uint256[] nftIds) returns()
func (_UnchainedStaking *UnchainedStakingSession) IncreaseStake(amount *big.Int, nftIds []*big.Int) (*types.Transaction, error) {
	return _UnchainedStaking.Contract.IncreaseStake(&_UnchainedStaking.TransactOpts, amount, nftIds)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0x0062ad9d.
//
// Solidity: function increaseStake(uint256 amount, uint256[] nftIds) returns()
func (_UnchainedStaking *UnchainedStakingTransactorSession) IncreaseStake(amount *big.Int, nftIds []*big.Int) (*types.Transaction, error) {
	return _UnchainedStaking.Contract.IncreaseStake(&_UnchainedStaking.TransactOpts, amount, nftIds)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x1171bda9.
//
// Solidity: function recoverERC20(address token, address recipient, uint256 amount) returns()
func (_UnchainedStaking *UnchainedStakingTransactor) RecoverERC20(opts *bind.TransactOpts, token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UnchainedStaking.contract.Transact(opts, "recoverERC20", token, recipient, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x1171bda9.
//
// Solidity: function recoverERC20(address token, address recipient, uint256 amount) returns()
func (_UnchainedStaking *UnchainedStakingSession) RecoverERC20(token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UnchainedStaking.Contract.RecoverERC20(&_UnchainedStaking.TransactOpts, token, recipient, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x1171bda9.
//
// Solidity: function recoverERC20(address token, address recipient, uint256 amount) returns()
func (_UnchainedStaking *UnchainedStakingTransactorSession) RecoverERC20(token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UnchainedStaking.Contract.RecoverERC20(&_UnchainedStaking.TransactOpts, token, recipient, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UnchainedStaking *UnchainedStakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnchainedStaking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UnchainedStaking *UnchainedStakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _UnchainedStaking.Contract.RenounceOwnership(&_UnchainedStaking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UnchainedStaking *UnchainedStakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UnchainedStaking.Contract.RenounceOwnership(&_UnchainedStaking.TransactOpts)
}

// SetBlsAddress is a paid mutator transaction binding the contract method 0x3e5d4675.
//
// Solidity: function setBlsAddress(bytes20 blsAddress) returns()
func (_UnchainedStaking *UnchainedStakingTransactor) SetBlsAddress(opts *bind.TransactOpts, blsAddress [20]byte) (*types.Transaction, error) {
	return _UnchainedStaking.contract.Transact(opts, "setBlsAddress", blsAddress)
}

// SetBlsAddress is a paid mutator transaction binding the contract method 0x3e5d4675.
//
// Solidity: function setBlsAddress(bytes20 blsAddress) returns()
func (_UnchainedStaking *UnchainedStakingSession) SetBlsAddress(blsAddress [20]byte) (*types.Transaction, error) {
	return _UnchainedStaking.Contract.SetBlsAddress(&_UnchainedStaking.TransactOpts, blsAddress)
}

// SetBlsAddress is a paid mutator transaction binding the contract method 0x3e5d4675.
//
// Solidity: function setBlsAddress(bytes20 blsAddress) returns()
func (_UnchainedStaking *UnchainedStakingTransactorSession) SetBlsAddress(blsAddress [20]byte) (*types.Transaction, error) {
	return _UnchainedStaking.Contract.SetBlsAddress(&_UnchainedStaking.TransactOpts, blsAddress)
}

// SetParams is a paid mutator transaction binding the contract method 0xca023b5b.
//
// Solidity: function setParams((address,address,address,uint256,uint256,uint256)[] eip712SetParams, (uint8,bytes32,bytes32)[] signatures) returns()
func (_UnchainedStaking *UnchainedStakingTransactor) SetParams(opts *bind.TransactOpts, eip712SetParams []UnchainedStakingEIP712SetParams, signatures []UnchainedStakingSignature) (*types.Transaction, error) {
	return _UnchainedStaking.contract.Transact(opts, "setParams", eip712SetParams, signatures)
}

// SetParams is a paid mutator transaction binding the contract method 0xca023b5b.
//
// Solidity: function setParams((address,address,address,uint256,uint256,uint256)[] eip712SetParams, (uint8,bytes32,bytes32)[] signatures) returns()
func (_UnchainedStaking *UnchainedStakingSession) SetParams(eip712SetParams []UnchainedStakingEIP712SetParams, signatures []UnchainedStakingSignature) (*types.Transaction, error) {
	return _UnchainedStaking.Contract.SetParams(&_UnchainedStaking.TransactOpts, eip712SetParams, signatures)
}

// SetParams is a paid mutator transaction binding the contract method 0xca023b5b.
//
// Solidity: function setParams((address,address,address,uint256,uint256,uint256)[] eip712SetParams, (uint8,bytes32,bytes32)[] signatures) returns()
func (_UnchainedStaking *UnchainedStakingTransactorSession) SetParams(eip712SetParams []UnchainedStakingEIP712SetParams, signatures []UnchainedStakingSignature) (*types.Transaction, error) {
	return _UnchainedStaking.Contract.SetParams(&_UnchainedStaking.TransactOpts, eip712SetParams, signatures)
}

// SetSigner is a paid mutator transaction binding the contract method 0xd9d35a4d.
//
// Solidity: function setSigner((address,address) eip712SetSigner, (uint8,bytes32,bytes32) stakerSignature, (uint8,bytes32,bytes32) signerSignature) returns()
func (_UnchainedStaking *UnchainedStakingTransactor) SetSigner(opts *bind.TransactOpts, eip712SetSigner UnchainedStakingEIP712SetSigner, stakerSignature UnchainedStakingSignature, signerSignature UnchainedStakingSignature) (*types.Transaction, error) {
	return _UnchainedStaking.contract.Transact(opts, "setSigner", eip712SetSigner, stakerSignature, signerSignature)
}

// SetSigner is a paid mutator transaction binding the contract method 0xd9d35a4d.
//
// Solidity: function setSigner((address,address) eip712SetSigner, (uint8,bytes32,bytes32) stakerSignature, (uint8,bytes32,bytes32) signerSignature) returns()
func (_UnchainedStaking *UnchainedStakingSession) SetSigner(eip712SetSigner UnchainedStakingEIP712SetSigner, stakerSignature UnchainedStakingSignature, signerSignature UnchainedStakingSignature) (*types.Transaction, error) {
	return _UnchainedStaking.Contract.SetSigner(&_UnchainedStaking.TransactOpts, eip712SetSigner, stakerSignature, signerSignature)
}

// SetSigner is a paid mutator transaction binding the contract method 0xd9d35a4d.
//
// Solidity: function setSigner((address,address) eip712SetSigner, (uint8,bytes32,bytes32) stakerSignature, (uint8,bytes32,bytes32) signerSignature) returns()
func (_UnchainedStaking *UnchainedStakingTransactorSession) SetSigner(eip712SetSigner UnchainedStakingEIP712SetSigner, stakerSignature UnchainedStakingSignature, signerSignature UnchainedStakingSignature) (*types.Transaction, error) {
	return _UnchainedStaking.Contract.SetSigner(&_UnchainedStaking.TransactOpts, eip712SetSigner, stakerSignature, signerSignature)
}

// Stake is a paid mutator transaction binding the contract method 0x9debdddc.
//
// Solidity: function stake(uint256 duration, uint256 amount, uint256[] nftIds) returns()
func (_UnchainedStaking *UnchainedStakingTransactor) Stake(opts *bind.TransactOpts, duration *big.Int, amount *big.Int, nftIds []*big.Int) (*types.Transaction, error) {
	return _UnchainedStaking.contract.Transact(opts, "stake", duration, amount, nftIds)
}

// Stake is a paid mutator transaction binding the contract method 0x9debdddc.
//
// Solidity: function stake(uint256 duration, uint256 amount, uint256[] nftIds) returns()
func (_UnchainedStaking *UnchainedStakingSession) Stake(duration *big.Int, amount *big.Int, nftIds []*big.Int) (*types.Transaction, error) {
	return _UnchainedStaking.Contract.Stake(&_UnchainedStaking.TransactOpts, duration, amount, nftIds)
}

// Stake is a paid mutator transaction binding the contract method 0x9debdddc.
//
// Solidity: function stake(uint256 duration, uint256 amount, uint256[] nftIds) returns()
func (_UnchainedStaking *UnchainedStakingTransactorSession) Stake(duration *big.Int, amount *big.Int, nftIds []*big.Int) (*types.Transaction, error) {
	return _UnchainedStaking.Contract.Stake(&_UnchainedStaking.TransactOpts, duration, amount, nftIds)
}

// TransferIn is a paid mutator transaction binding the contract method 0xd65014c3.
//
// Solidity: function transferIn(uint256 amount) returns()
func (_UnchainedStaking *UnchainedStakingTransactor) TransferIn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _UnchainedStaking.contract.Transact(opts, "transferIn", amount)
}

// TransferIn is a paid mutator transaction binding the contract method 0xd65014c3.
//
// Solidity: function transferIn(uint256 amount) returns()
func (_UnchainedStaking *UnchainedStakingSession) TransferIn(amount *big.Int) (*types.Transaction, error) {
	return _UnchainedStaking.Contract.TransferIn(&_UnchainedStaking.TransactOpts, amount)
}

// TransferIn is a paid mutator transaction binding the contract method 0xd65014c3.
//
// Solidity: function transferIn(uint256 amount) returns()
func (_UnchainedStaking *UnchainedStakingTransactorSession) TransferIn(amount *big.Int) (*types.Transaction, error) {
	return _UnchainedStaking.Contract.TransferIn(&_UnchainedStaking.TransactOpts, amount)
}

// TransferOut is a paid mutator transaction binding the contract method 0x6a340cc5.
//
// Solidity: function transferOut((address,address,address,uint256,uint256[])[] eip712Transferes, (uint8,bytes32,bytes32)[] signatures) returns()
func (_UnchainedStaking *UnchainedStakingTransactor) TransferOut(opts *bind.TransactOpts, eip712Transferes []UnchainedStakingEIP712Transfer, signatures []UnchainedStakingSignature) (*types.Transaction, error) {
	return _UnchainedStaking.contract.Transact(opts, "transferOut", eip712Transferes, signatures)
}

// TransferOut is a paid mutator transaction binding the contract method 0x6a340cc5.
//
// Solidity: function transferOut((address,address,address,uint256,uint256[])[] eip712Transferes, (uint8,bytes32,bytes32)[] signatures) returns()
func (_UnchainedStaking *UnchainedStakingSession) TransferOut(eip712Transferes []UnchainedStakingEIP712Transfer, signatures []UnchainedStakingSignature) (*types.Transaction, error) {
	return _UnchainedStaking.Contract.TransferOut(&_UnchainedStaking.TransactOpts, eip712Transferes, signatures)
}

// TransferOut is a paid mutator transaction binding the contract method 0x6a340cc5.
//
// Solidity: function transferOut((address,address,address,uint256,uint256[])[] eip712Transferes, (uint8,bytes32,bytes32)[] signatures) returns()
func (_UnchainedStaking *UnchainedStakingTransactorSession) TransferOut(eip712Transferes []UnchainedStakingEIP712Transfer, signatures []UnchainedStakingSignature) (*types.Transaction, error) {
	return _UnchainedStaking.Contract.TransferOut(&_UnchainedStaking.TransactOpts, eip712Transferes, signatures)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UnchainedStaking *UnchainedStakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UnchainedStaking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UnchainedStaking *UnchainedStakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UnchainedStaking.Contract.TransferOwnership(&_UnchainedStaking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UnchainedStaking *UnchainedStakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UnchainedStaking.Contract.TransferOwnership(&_UnchainedStaking.TransactOpts, newOwner)
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_UnchainedStaking *UnchainedStakingTransactor) Unstake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnchainedStaking.contract.Transact(opts, "unstake")
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_UnchainedStaking *UnchainedStakingSession) Unstake() (*types.Transaction, error) {
	return _UnchainedStaking.Contract.Unstake(&_UnchainedStaking.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_UnchainedStaking *UnchainedStakingTransactorSession) Unstake() (*types.Transaction, error) {
	return _UnchainedStaking.Contract.Unstake(&_UnchainedStaking.TransactOpts)
}

// UnchainedStakingAccusedIterator is returned from FilterAccused and is used to iterate over the raw logs and unpacked data for Accused events raised by the UnchainedStaking contract.
type UnchainedStakingAccusedIterator struct {
	Event *UnchainedStakingAccused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnchainedStakingAccusedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnchainedStakingAccused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnchainedStakingAccused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnchainedStakingAccusedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnchainedStakingAccusedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnchainedStakingAccused represents a Accused event raised by the UnchainedStaking contract.
type UnchainedStakingAccused struct {
	Accused  common.Address
	Accuser  common.Address
	Amount   *big.Int
	Voted    *big.Int
	Incident [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAccused is a free log retrieval operation binding the contract event 0xdc67704e99fa42c847e50d5d0ca44644110b5ac921c8cb738d66369a3ae5423a.
//
// Solidity: event Accused(address accused, address accuser, uint256 amount, uint256 voted, bytes32 incident)
func (_UnchainedStaking *UnchainedStakingFilterer) FilterAccused(opts *bind.FilterOpts) (*UnchainedStakingAccusedIterator, error) {

	logs, sub, err := _UnchainedStaking.contract.FilterLogs(opts, "Accused")
	if err != nil {
		return nil, err
	}
	return &UnchainedStakingAccusedIterator{contract: _UnchainedStaking.contract, event: "Accused", logs: logs, sub: sub}, nil
}

// WatchAccused is a free log subscription operation binding the contract event 0xdc67704e99fa42c847e50d5d0ca44644110b5ac921c8cb738d66369a3ae5423a.
//
// Solidity: event Accused(address accused, address accuser, uint256 amount, uint256 voted, bytes32 incident)
func (_UnchainedStaking *UnchainedStakingFilterer) WatchAccused(opts *bind.WatchOpts, sink chan<- *UnchainedStakingAccused) (event.Subscription, error) {

	logs, sub, err := _UnchainedStaking.contract.WatchLogs(opts, "Accused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnchainedStakingAccused)
				if err := _UnchainedStaking.contract.UnpackLog(event, "Accused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccused is a log parse operation binding the contract event 0xdc67704e99fa42c847e50d5d0ca44644110b5ac921c8cb738d66369a3ae5423a.
//
// Solidity: event Accused(address accused, address accuser, uint256 amount, uint256 voted, bytes32 incident)
func (_UnchainedStaking *UnchainedStakingFilterer) ParseAccused(log types.Log) (*UnchainedStakingAccused, error) {
	event := new(UnchainedStakingAccused)
	if err := _UnchainedStaking.contract.UnpackLog(event, "Accused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnchainedStakingBlsAddressChangedIterator is returned from FilterBlsAddressChanged and is used to iterate over the raw logs and unpacked data for BlsAddressChanged events raised by the UnchainedStaking contract.
type UnchainedStakingBlsAddressChangedIterator struct {
	Event *UnchainedStakingBlsAddressChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnchainedStakingBlsAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnchainedStakingBlsAddressChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnchainedStakingBlsAddressChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnchainedStakingBlsAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnchainedStakingBlsAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnchainedStakingBlsAddressChanged represents a BlsAddressChanged event raised by the UnchainedStaking contract.
type UnchainedStakingBlsAddressChanged struct {
	User common.Address
	From [32]byte
	To   [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBlsAddressChanged is a free log retrieval operation binding the contract event 0xa5c20a3e40dbfce0ccdadcb27e2f561e84ddf0618a41338cc1acb1524780ff39.
//
// Solidity: event BlsAddressChanged(address user, bytes32 from, bytes32 to)
func (_UnchainedStaking *UnchainedStakingFilterer) FilterBlsAddressChanged(opts *bind.FilterOpts) (*UnchainedStakingBlsAddressChangedIterator, error) {

	logs, sub, err := _UnchainedStaking.contract.FilterLogs(opts, "BlsAddressChanged")
	if err != nil {
		return nil, err
	}
	return &UnchainedStakingBlsAddressChangedIterator{contract: _UnchainedStaking.contract, event: "BlsAddressChanged", logs: logs, sub: sub}, nil
}

// WatchBlsAddressChanged is a free log subscription operation binding the contract event 0xa5c20a3e40dbfce0ccdadcb27e2f561e84ddf0618a41338cc1acb1524780ff39.
//
// Solidity: event BlsAddressChanged(address user, bytes32 from, bytes32 to)
func (_UnchainedStaking *UnchainedStakingFilterer) WatchBlsAddressChanged(opts *bind.WatchOpts, sink chan<- *UnchainedStakingBlsAddressChanged) (event.Subscription, error) {

	logs, sub, err := _UnchainedStaking.contract.WatchLogs(opts, "BlsAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnchainedStakingBlsAddressChanged)
				if err := _UnchainedStaking.contract.UnpackLog(event, "BlsAddressChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlsAddressChanged is a log parse operation binding the contract event 0xa5c20a3e40dbfce0ccdadcb27e2f561e84ddf0618a41338cc1acb1524780ff39.
//
// Solidity: event BlsAddressChanged(address user, bytes32 from, bytes32 to)
func (_UnchainedStaking *UnchainedStakingFilterer) ParseBlsAddressChanged(log types.Log) (*UnchainedStakingBlsAddressChanged, error) {
	event := new(UnchainedStakingBlsAddressChanged)
	if err := _UnchainedStaking.contract.UnpackLog(event, "BlsAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnchainedStakingExtendedIterator is returned from FilterExtended and is used to iterate over the raw logs and unpacked data for Extended events raised by the UnchainedStaking contract.
type UnchainedStakingExtendedIterator struct {
	Event *UnchainedStakingExtended // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnchainedStakingExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnchainedStakingExtended)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnchainedStakingExtended)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnchainedStakingExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnchainedStakingExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnchainedStakingExtended represents a Extended event raised by the UnchainedStaking contract.
type UnchainedStakingExtended struct {
	User   common.Address
	Unlock *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExtended is a free log retrieval operation binding the contract event 0xa29fc12cda82ff659de006abb10fa5ee256d922af1661e395e5f2fb6b004387e.
//
// Solidity: event Extended(address user, uint256 unlock)
func (_UnchainedStaking *UnchainedStakingFilterer) FilterExtended(opts *bind.FilterOpts) (*UnchainedStakingExtendedIterator, error) {

	logs, sub, err := _UnchainedStaking.contract.FilterLogs(opts, "Extended")
	if err != nil {
		return nil, err
	}
	return &UnchainedStakingExtendedIterator{contract: _UnchainedStaking.contract, event: "Extended", logs: logs, sub: sub}, nil
}

// WatchExtended is a free log subscription operation binding the contract event 0xa29fc12cda82ff659de006abb10fa5ee256d922af1661e395e5f2fb6b004387e.
//
// Solidity: event Extended(address user, uint256 unlock)
func (_UnchainedStaking *UnchainedStakingFilterer) WatchExtended(opts *bind.WatchOpts, sink chan<- *UnchainedStakingExtended) (event.Subscription, error) {

	logs, sub, err := _UnchainedStaking.contract.WatchLogs(opts, "Extended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnchainedStakingExtended)
				if err := _UnchainedStaking.contract.UnpackLog(event, "Extended", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExtended is a log parse operation binding the contract event 0xa29fc12cda82ff659de006abb10fa5ee256d922af1661e395e5f2fb6b004387e.
//
// Solidity: event Extended(address user, uint256 unlock)
func (_UnchainedStaking *UnchainedStakingFilterer) ParseExtended(log types.Log) (*UnchainedStakingExtended, error) {
	event := new(UnchainedStakingExtended)
	if err := _UnchainedStaking.contract.UnpackLog(event, "Extended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnchainedStakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UnchainedStaking contract.
type UnchainedStakingOwnershipTransferredIterator struct {
	Event *UnchainedStakingOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnchainedStakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnchainedStakingOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnchainedStakingOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnchainedStakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnchainedStakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnchainedStakingOwnershipTransferred represents a OwnershipTransferred event raised by the UnchainedStaking contract.
type UnchainedStakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UnchainedStaking *UnchainedStakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UnchainedStakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UnchainedStaking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UnchainedStakingOwnershipTransferredIterator{contract: _UnchainedStaking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UnchainedStaking *UnchainedStakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UnchainedStakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UnchainedStaking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnchainedStakingOwnershipTransferred)
				if err := _UnchainedStaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UnchainedStaking *UnchainedStakingFilterer) ParseOwnershipTransferred(log types.Log) (*UnchainedStakingOwnershipTransferred, error) {
	event := new(UnchainedStakingOwnershipTransferred)
	if err := _UnchainedStaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnchainedStakingParamsChangedIterator is returned from FilterParamsChanged and is used to iterate over the raw logs and unpacked data for ParamsChanged events raised by the UnchainedStaking contract.
type UnchainedStakingParamsChangedIterator struct {
	Event *UnchainedStakingParamsChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnchainedStakingParamsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnchainedStakingParamsChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnchainedStakingParamsChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnchainedStakingParamsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnchainedStakingParamsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnchainedStakingParamsChanged represents a ParamsChanged event raised by the UnchainedStaking contract.
type UnchainedStakingParamsChanged struct {
	Token      common.Address
	Nft        common.Address
	Threshold  *big.Int
	Expiration *big.Int
	Voted      *big.Int
	Nonce      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterParamsChanged is a free log retrieval operation binding the contract event 0xcbe90a96ab941d4bcb02f4a5eda44c0ac9b4cbb16bb7d84db6ce65e8a574f029.
//
// Solidity: event ParamsChanged(address token, address nft, uint256 threshold, uint256 expiration, uint256 voted, uint256 nonce)
func (_UnchainedStaking *UnchainedStakingFilterer) FilterParamsChanged(opts *bind.FilterOpts) (*UnchainedStakingParamsChangedIterator, error) {

	logs, sub, err := _UnchainedStaking.contract.FilterLogs(opts, "ParamsChanged")
	if err != nil {
		return nil, err
	}
	return &UnchainedStakingParamsChangedIterator{contract: _UnchainedStaking.contract, event: "ParamsChanged", logs: logs, sub: sub}, nil
}

// WatchParamsChanged is a free log subscription operation binding the contract event 0xcbe90a96ab941d4bcb02f4a5eda44c0ac9b4cbb16bb7d84db6ce65e8a574f029.
//
// Solidity: event ParamsChanged(address token, address nft, uint256 threshold, uint256 expiration, uint256 voted, uint256 nonce)
func (_UnchainedStaking *UnchainedStakingFilterer) WatchParamsChanged(opts *bind.WatchOpts, sink chan<- *UnchainedStakingParamsChanged) (event.Subscription, error) {

	logs, sub, err := _UnchainedStaking.contract.WatchLogs(opts, "ParamsChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnchainedStakingParamsChanged)
				if err := _UnchainedStaking.contract.UnpackLog(event, "ParamsChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseParamsChanged is a log parse operation binding the contract event 0xcbe90a96ab941d4bcb02f4a5eda44c0ac9b4cbb16bb7d84db6ce65e8a574f029.
//
// Solidity: event ParamsChanged(address token, address nft, uint256 threshold, uint256 expiration, uint256 voted, uint256 nonce)
func (_UnchainedStaking *UnchainedStakingFilterer) ParseParamsChanged(log types.Log) (*UnchainedStakingParamsChanged, error) {
	event := new(UnchainedStakingParamsChanged)
	if err := _UnchainedStaking.contract.UnpackLog(event, "ParamsChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnchainedStakingSignerChangedIterator is returned from FilterSignerChanged and is used to iterate over the raw logs and unpacked data for SignerChanged events raised by the UnchainedStaking contract.
type UnchainedStakingSignerChangedIterator struct {
	Event *UnchainedStakingSignerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnchainedStakingSignerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnchainedStakingSignerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnchainedStakingSignerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnchainedStakingSignerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnchainedStakingSignerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnchainedStakingSignerChanged represents a SignerChanged event raised by the UnchainedStaking contract.
type UnchainedStakingSignerChanged struct {
	Staker common.Address
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignerChanged is a free log retrieval operation binding the contract event 0xeeb293e1f8f3a9db91ade748726387ed1352ca78f5430c5f06fe3d1e1ad50579.
//
// Solidity: event SignerChanged(address staker, address signer)
func (_UnchainedStaking *UnchainedStakingFilterer) FilterSignerChanged(opts *bind.FilterOpts) (*UnchainedStakingSignerChangedIterator, error) {

	logs, sub, err := _UnchainedStaking.contract.FilterLogs(opts, "SignerChanged")
	if err != nil {
		return nil, err
	}
	return &UnchainedStakingSignerChangedIterator{contract: _UnchainedStaking.contract, event: "SignerChanged", logs: logs, sub: sub}, nil
}

// WatchSignerChanged is a free log subscription operation binding the contract event 0xeeb293e1f8f3a9db91ade748726387ed1352ca78f5430c5f06fe3d1e1ad50579.
//
// Solidity: event SignerChanged(address staker, address signer)
func (_UnchainedStaking *UnchainedStakingFilterer) WatchSignerChanged(opts *bind.WatchOpts, sink chan<- *UnchainedStakingSignerChanged) (event.Subscription, error) {

	logs, sub, err := _UnchainedStaking.contract.WatchLogs(opts, "SignerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnchainedStakingSignerChanged)
				if err := _UnchainedStaking.contract.UnpackLog(event, "SignerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignerChanged is a log parse operation binding the contract event 0xeeb293e1f8f3a9db91ade748726387ed1352ca78f5430c5f06fe3d1e1ad50579.
//
// Solidity: event SignerChanged(address staker, address signer)
func (_UnchainedStaking *UnchainedStakingFilterer) ParseSignerChanged(log types.Log) (*UnchainedStakingSignerChanged, error) {
	event := new(UnchainedStakingSignerChanged)
	if err := _UnchainedStaking.contract.UnpackLog(event, "SignerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnchainedStakingStakeIncreasedIterator is returned from FilterStakeIncreased and is used to iterate over the raw logs and unpacked data for StakeIncreased events raised by the UnchainedStaking contract.
type UnchainedStakingStakeIncreasedIterator struct {
	Event *UnchainedStakingStakeIncreased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnchainedStakingStakeIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnchainedStakingStakeIncreased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnchainedStakingStakeIncreased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnchainedStakingStakeIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnchainedStakingStakeIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnchainedStakingStakeIncreased represents a StakeIncreased event raised by the UnchainedStaking contract.
type UnchainedStakingStakeIncreased struct {
	User   common.Address
	Amount *big.Int
	NftIds []*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeIncreased is a free log retrieval operation binding the contract event 0x26e4d06799c60ba22ec168ee4a9314ed451cf9c78dbf6a5f7bfeaf3c84688f58.
//
// Solidity: event StakeIncreased(address user, uint256 amount, uint256[] nftIds)
func (_UnchainedStaking *UnchainedStakingFilterer) FilterStakeIncreased(opts *bind.FilterOpts) (*UnchainedStakingStakeIncreasedIterator, error) {

	logs, sub, err := _UnchainedStaking.contract.FilterLogs(opts, "StakeIncreased")
	if err != nil {
		return nil, err
	}
	return &UnchainedStakingStakeIncreasedIterator{contract: _UnchainedStaking.contract, event: "StakeIncreased", logs: logs, sub: sub}, nil
}

// WatchStakeIncreased is a free log subscription operation binding the contract event 0x26e4d06799c60ba22ec168ee4a9314ed451cf9c78dbf6a5f7bfeaf3c84688f58.
//
// Solidity: event StakeIncreased(address user, uint256 amount, uint256[] nftIds)
func (_UnchainedStaking *UnchainedStakingFilterer) WatchStakeIncreased(opts *bind.WatchOpts, sink chan<- *UnchainedStakingStakeIncreased) (event.Subscription, error) {

	logs, sub, err := _UnchainedStaking.contract.WatchLogs(opts, "StakeIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnchainedStakingStakeIncreased)
				if err := _UnchainedStaking.contract.UnpackLog(event, "StakeIncreased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeIncreased is a log parse operation binding the contract event 0x26e4d06799c60ba22ec168ee4a9314ed451cf9c78dbf6a5f7bfeaf3c84688f58.
//
// Solidity: event StakeIncreased(address user, uint256 amount, uint256[] nftIds)
func (_UnchainedStaking *UnchainedStakingFilterer) ParseStakeIncreased(log types.Log) (*UnchainedStakingStakeIncreased, error) {
	event := new(UnchainedStakingStakeIncreased)
	if err := _UnchainedStaking.contract.UnpackLog(event, "StakeIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnchainedStakingStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the UnchainedStaking contract.
type UnchainedStakingStakedIterator struct {
	Event *UnchainedStakingStaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnchainedStakingStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnchainedStakingStaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnchainedStakingStaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnchainedStakingStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnchainedStakingStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnchainedStakingStaked represents a Staked event raised by the UnchainedStaking contract.
type UnchainedStakingStaked struct {
	User   common.Address
	Unlock *big.Int
	Amount *big.Int
	NftIds []*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x5df5de9ccb680fe3d60088f6d4c3b6d535074c704699377046c743a5b276e171.
//
// Solidity: event Staked(address user, uint256 unlock, uint256 amount, uint256[] nftIds)
func (_UnchainedStaking *UnchainedStakingFilterer) FilterStaked(opts *bind.FilterOpts) (*UnchainedStakingStakedIterator, error) {

	logs, sub, err := _UnchainedStaking.contract.FilterLogs(opts, "Staked")
	if err != nil {
		return nil, err
	}
	return &UnchainedStakingStakedIterator{contract: _UnchainedStaking.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x5df5de9ccb680fe3d60088f6d4c3b6d535074c704699377046c743a5b276e171.
//
// Solidity: event Staked(address user, uint256 unlock, uint256 amount, uint256[] nftIds)
func (_UnchainedStaking *UnchainedStakingFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *UnchainedStakingStaked) (event.Subscription, error) {

	logs, sub, err := _UnchainedStaking.contract.WatchLogs(opts, "Staked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnchainedStakingStaked)
				if err := _UnchainedStaking.contract.UnpackLog(event, "Staked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaked is a log parse operation binding the contract event 0x5df5de9ccb680fe3d60088f6d4c3b6d535074c704699377046c743a5b276e171.
//
// Solidity: event Staked(address user, uint256 unlock, uint256 amount, uint256[] nftIds)
func (_UnchainedStaking *UnchainedStakingFilterer) ParseStaked(log types.Log) (*UnchainedStakingStaked, error) {
	event := new(UnchainedStakingStaked)
	if err := _UnchainedStaking.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnchainedStakingTransferInIterator is returned from FilterTransferIn and is used to iterate over the raw logs and unpacked data for TransferIn events raised by the UnchainedStaking contract.
type UnchainedStakingTransferInIterator struct {
	Event *UnchainedStakingTransferIn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnchainedStakingTransferInIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnchainedStakingTransferIn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnchainedStakingTransferIn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnchainedStakingTransferInIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnchainedStakingTransferInIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnchainedStakingTransferIn represents a TransferIn event raised by the UnchainedStaking contract.
type UnchainedStakingTransferIn struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferIn is a free log retrieval operation binding the contract event 0x20bdd1a9edc2be1c156c5768ce9e83b8a08cd73923158e955da8718d6c7bbbfc.
//
// Solidity: event TransferIn(address from, uint256 amount)
func (_UnchainedStaking *UnchainedStakingFilterer) FilterTransferIn(opts *bind.FilterOpts) (*UnchainedStakingTransferInIterator, error) {

	logs, sub, err := _UnchainedStaking.contract.FilterLogs(opts, "TransferIn")
	if err != nil {
		return nil, err
	}
	return &UnchainedStakingTransferInIterator{contract: _UnchainedStaking.contract, event: "TransferIn", logs: logs, sub: sub}, nil
}

// WatchTransferIn is a free log subscription operation binding the contract event 0x20bdd1a9edc2be1c156c5768ce9e83b8a08cd73923158e955da8718d6c7bbbfc.
//
// Solidity: event TransferIn(address from, uint256 amount)
func (_UnchainedStaking *UnchainedStakingFilterer) WatchTransferIn(opts *bind.WatchOpts, sink chan<- *UnchainedStakingTransferIn) (event.Subscription, error) {

	logs, sub, err := _UnchainedStaking.contract.WatchLogs(opts, "TransferIn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnchainedStakingTransferIn)
				if err := _UnchainedStaking.contract.UnpackLog(event, "TransferIn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferIn is a log parse operation binding the contract event 0x20bdd1a9edc2be1c156c5768ce9e83b8a08cd73923158e955da8718d6c7bbbfc.
//
// Solidity: event TransferIn(address from, uint256 amount)
func (_UnchainedStaking *UnchainedStakingFilterer) ParseTransferIn(log types.Log) (*UnchainedStakingTransferIn, error) {
	event := new(UnchainedStakingTransferIn)
	if err := _UnchainedStaking.contract.UnpackLog(event, "TransferIn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnchainedStakingTransferOutIterator is returned from FilterTransferOut and is used to iterate over the raw logs and unpacked data for TransferOut events raised by the UnchainedStaking contract.
type UnchainedStakingTransferOutIterator struct {
	Event *UnchainedStakingTransferOut // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnchainedStakingTransferOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnchainedStakingTransferOut)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnchainedStakingTransferOut)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnchainedStakingTransferOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnchainedStakingTransferOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnchainedStakingTransferOut represents a TransferOut event raised by the UnchainedStaking contract.
type UnchainedStakingTransferOut struct {
	To     common.Address
	Amount *big.Int
	Nonces []*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferOut is a free log retrieval operation binding the contract event 0xa68a27e36f77abe383c0eef9bf2f0a3bd5b45fccf9c149a130d262c47963e74d.
//
// Solidity: event TransferOut(address to, uint256 amount, uint256[] nonces)
func (_UnchainedStaking *UnchainedStakingFilterer) FilterTransferOut(opts *bind.FilterOpts) (*UnchainedStakingTransferOutIterator, error) {

	logs, sub, err := _UnchainedStaking.contract.FilterLogs(opts, "TransferOut")
	if err != nil {
		return nil, err
	}
	return &UnchainedStakingTransferOutIterator{contract: _UnchainedStaking.contract, event: "TransferOut", logs: logs, sub: sub}, nil
}

// WatchTransferOut is a free log subscription operation binding the contract event 0xa68a27e36f77abe383c0eef9bf2f0a3bd5b45fccf9c149a130d262c47963e74d.
//
// Solidity: event TransferOut(address to, uint256 amount, uint256[] nonces)
func (_UnchainedStaking *UnchainedStakingFilterer) WatchTransferOut(opts *bind.WatchOpts, sink chan<- *UnchainedStakingTransferOut) (event.Subscription, error) {

	logs, sub, err := _UnchainedStaking.contract.WatchLogs(opts, "TransferOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnchainedStakingTransferOut)
				if err := _UnchainedStaking.contract.UnpackLog(event, "TransferOut", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferOut is a log parse operation binding the contract event 0xa68a27e36f77abe383c0eef9bf2f0a3bd5b45fccf9c149a130d262c47963e74d.
//
// Solidity: event TransferOut(address to, uint256 amount, uint256[] nonces)
func (_UnchainedStaking *UnchainedStakingFilterer) ParseTransferOut(log types.Log) (*UnchainedStakingTransferOut, error) {
	event := new(UnchainedStakingTransferOut)
	if err := _UnchainedStaking.contract.UnpackLog(event, "TransferOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnchainedStakingUnStakedIterator is returned from FilterUnStaked and is used to iterate over the raw logs and unpacked data for UnStaked events raised by the UnchainedStaking contract.
type UnchainedStakingUnStakedIterator struct {
	Event *UnchainedStakingUnStaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnchainedStakingUnStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnchainedStakingUnStaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnchainedStakingUnStaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnchainedStakingUnStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnchainedStakingUnStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnchainedStakingUnStaked represents a UnStaked event raised by the UnchainedStaking contract.
type UnchainedStakingUnStaked struct {
	User   common.Address
	Amount *big.Int
	NftIds []*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnStaked is a free log retrieval operation binding the contract event 0xef4ce4a0205d268e0effbc76aaabb2ad2509ec58a2e9013645347d3c3cd9be42.
//
// Solidity: event UnStaked(address user, uint256 amount, uint256[] nftIds)
func (_UnchainedStaking *UnchainedStakingFilterer) FilterUnStaked(opts *bind.FilterOpts) (*UnchainedStakingUnStakedIterator, error) {

	logs, sub, err := _UnchainedStaking.contract.FilterLogs(opts, "UnStaked")
	if err != nil {
		return nil, err
	}
	return &UnchainedStakingUnStakedIterator{contract: _UnchainedStaking.contract, event: "UnStaked", logs: logs, sub: sub}, nil
}

// WatchUnStaked is a free log subscription operation binding the contract event 0xef4ce4a0205d268e0effbc76aaabb2ad2509ec58a2e9013645347d3c3cd9be42.
//
// Solidity: event UnStaked(address user, uint256 amount, uint256[] nftIds)
func (_UnchainedStaking *UnchainedStakingFilterer) WatchUnStaked(opts *bind.WatchOpts, sink chan<- *UnchainedStakingUnStaked) (event.Subscription, error) {

	logs, sub, err := _UnchainedStaking.contract.WatchLogs(opts, "UnStaked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnchainedStakingUnStaked)
				if err := _UnchainedStaking.contract.UnpackLog(event, "UnStaked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnStaked is a log parse operation binding the contract event 0xef4ce4a0205d268e0effbc76aaabb2ad2509ec58a2e9013645347d3c3cd9be42.
//
// Solidity: event UnStaked(address user, uint256 amount, uint256[] nftIds)
func (_UnchainedStaking *UnchainedStakingFilterer) ParseUnStaked(log types.Log) (*UnchainedStakingUnStaked, error) {
	event := new(UnchainedStakingUnStaked)
	if err := _UnchainedStaking.contract.UnpackLog(event, "UnStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnchainedStakingVotedForParamsIterator is returned from FilterVotedForParams and is used to iterate over the raw logs and unpacked data for VotedForParams events raised by the UnchainedStaking contract.
type UnchainedStakingVotedForParamsIterator struct {
	Event *UnchainedStakingVotedForParams // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnchainedStakingVotedForParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnchainedStakingVotedForParams)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnchainedStakingVotedForParams)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnchainedStakingVotedForParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnchainedStakingVotedForParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnchainedStakingVotedForParams represents a VotedForParams event raised by the UnchainedStaking contract.
type UnchainedStakingVotedForParams struct {
	User  common.Address
	Nonce *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVotedForParams is a free log retrieval operation binding the contract event 0x6143a6e12175d81f2e3478633048270835198f154641368d807deeb1b8b641b1.
//
// Solidity: event VotedForParams(address user, uint256 nonce)
func (_UnchainedStaking *UnchainedStakingFilterer) FilterVotedForParams(opts *bind.FilterOpts) (*UnchainedStakingVotedForParamsIterator, error) {

	logs, sub, err := _UnchainedStaking.contract.FilterLogs(opts, "VotedForParams")
	if err != nil {
		return nil, err
	}
	return &UnchainedStakingVotedForParamsIterator{contract: _UnchainedStaking.contract, event: "VotedForParams", logs: logs, sub: sub}, nil
}

// WatchVotedForParams is a free log subscription operation binding the contract event 0x6143a6e12175d81f2e3478633048270835198f154641368d807deeb1b8b641b1.
//
// Solidity: event VotedForParams(address user, uint256 nonce)
func (_UnchainedStaking *UnchainedStakingFilterer) WatchVotedForParams(opts *bind.WatchOpts, sink chan<- *UnchainedStakingVotedForParams) (event.Subscription, error) {

	logs, sub, err := _UnchainedStaking.contract.WatchLogs(opts, "VotedForParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnchainedStakingVotedForParams)
				if err := _UnchainedStaking.contract.UnpackLog(event, "VotedForParams", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVotedForParams is a log parse operation binding the contract event 0x6143a6e12175d81f2e3478633048270835198f154641368d807deeb1b8b641b1.
//
// Solidity: event VotedForParams(address user, uint256 nonce)
func (_UnchainedStaking *UnchainedStakingFilterer) ParseVotedForParams(log types.Log) (*UnchainedStakingVotedForParams, error) {
	event := new(UnchainedStakingVotedForParams)
	if err := _UnchainedStaking.contract.UnpackLog(event, "VotedForParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
