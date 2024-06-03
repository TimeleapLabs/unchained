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

// ProofOfStakeStake is an auto generated low-level Go binding around an user-defined struct.
type ProofOfStakeStake struct {
	Amount *big.Int
	End    *big.Int
	Nfts   []*big.Int
	NftSum *big.Int
}

// SchnorrNftTransferTransfer is an auto generated low-level Go binding around an user-defined struct.
type SchnorrNftTransferTransfer struct {
	From  common.Address
	To    common.Address
	Id    *big.Int
	Nonce *big.Int
}

// SchnorrSignatureSignature is an auto generated low-level Go binding around an user-defined struct.
type SchnorrSignatureSignature struct {
	Rx *big.Int
	S  *big.Int
}

// SchnorrTransferOwnershipTransferOwnership is an auto generated low-level Go binding around an user-defined struct.
type SchnorrTransferOwnershipTransferOwnership struct {
	To    *big.Int
	Nonce *big.Int
}

// SchnorrTransferTransfer is an auto generated low-level Go binding around an user-defined struct.
type SchnorrTransferTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Nonce  *big.Int
}

// SetNftPricesNftPrices is an auto generated low-level Go binding around an user-defined struct.
type SetNftPricesNftPrices struct {
	Nfts   []*big.Int
	Prices []*big.Int
	Nonce  *big.Int
}

// SetSchnorrThresholdSchnorrThreshold is an auto generated low-level Go binding around an user-defined struct.
type SetSchnorrThresholdSchnorrThreshold struct {
	Threshold *big.Int
	Nonce     *big.Int
}

// ProofOfStakeMetaData contains all meta data concerning the ProofOfStake contract.
var ProofOfStakeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shcnorrOwner\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakingTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyProcessed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyStaked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DurationZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ElementAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ElementDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"IndexOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSchorrSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"InvalidSlice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NftNotInStake\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoStakeToExtend\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"Extended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"nfts\",\"type\":\"uint256[]\"}],\"name\":\"Increased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"nfts\",\"type\":\"uint256[]\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"nfts\",\"type\":\"uint256[]\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"eip712DomainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"extendStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getStake\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"nfts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"nftSum\",\"type\":\"uint256\"}],\"internalType\":\"structProofOfStake.Stake\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"nfts\",\"type\":\"uint256[]\"}],\"name\":\"increaseStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nftPrices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"processed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"eip712Hash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"internalType\":\"structSchnorrSignature.Signature\",\"name\":\"schnorrSignature\",\"type\":\"tuple\"}],\"name\":\"safeVerify\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"schnorrParticipationThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"nfts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structSetNftPrices.NftPrices\",\"name\":\"prices\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"internalType\":\"structSchnorrSignature.Signature\",\"name\":\"schnorrSignature\",\"type\":\"tuple\"}],\"name\":\"setNftPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structSetSchnorrThreshold.SchnorrThreshold\",\"name\":\"threshold\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"internalType\":\"structSchnorrSignature.Signature\",\"name\":\"schnorrSignature\",\"type\":\"tuple\"}],\"name\":\"setSchNorrParticipationThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"nfts\",\"type\":\"uint256[]\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftSum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structSchnorrTransfer.Transfer\",\"name\":\"txn\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"internalType\":\"structSchnorrSignature.Signature\",\"name\":\"schnorrSignature\",\"type\":\"tuple\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structSchnorrNftTransfer.Transfer\",\"name\":\"txn\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"internalType\":\"structSchnorrSignature.Signature\",\"name\":\"schnorrSignature\",\"type\":\"tuple\"}],\"name\":\"transferNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structSchnorrTransferOwnership.TransferOwnership\",\"name\":\"txn\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"internalType\":\"structSchnorrSignature.Signature\",\"name\":\"schnorrSignature\",\"type\":\"tuple\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"nfts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structSetNftPrices.NftPrices\",\"name\":\"prices\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"internalType\":\"structSchnorrSignature.Signature\",\"name\":\"schnorrSignature\",\"type\":\"tuple\"}],\"name\":\"verifySetNftPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structSetSchnorrThreshold.SchnorrThreshold\",\"name\":\"threshold\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"internalType\":\"structSchnorrSignature.Signature\",\"name\":\"schnorrSignature\",\"type\":\"tuple\"}],\"name\":\"verifySetSchnorrThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structSchnorrTransfer.Transfer\",\"name\":\"txn\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"internalType\":\"structSchnorrSignature.Signature\",\"name\":\"schnorrSignature\",\"type\":\"tuple\"}],\"name\":\"verifyTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structSchnorrNftTransfer.Transfer\",\"name\":\"txn\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"internalType\":\"structSchnorrSignature.Signature\",\"name\":\"schnorrSignature\",\"type\":\"tuple\"}],\"name\":\"verifyTransferNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ProofOfStakeABI is the input ABI used to generate the binding from.
// Deprecated: Use ProofOfStakeMetaData.ABI instead.
var ProofOfStakeABI = ProofOfStakeMetaData.ABI

// ProofOfStake is an auto generated Go binding around an Ethereum contract.
type ProofOfStake struct {
	ProofOfStakeCaller     // Read-only binding to the contract
	ProofOfStakeTransactor // Write-only binding to the contract
	ProofOfStakeFilterer   // Log filterer for contract events
}

// ProofOfStakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProofOfStakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofOfStakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProofOfStakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofOfStakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProofOfStakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofOfStakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProofOfStakeSession struct {
	Contract     *ProofOfStake     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProofOfStakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProofOfStakeCallerSession struct {
	Contract *ProofOfStakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ProofOfStakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProofOfStakeTransactorSession struct {
	Contract     *ProofOfStakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ProofOfStakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProofOfStakeRaw struct {
	Contract *ProofOfStake // Generic contract binding to access the raw methods on
}

// ProofOfStakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProofOfStakeCallerRaw struct {
	Contract *ProofOfStakeCaller // Generic read-only contract binding to access the raw methods on
}

// ProofOfStakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProofOfStakeTransactorRaw struct {
	Contract *ProofOfStakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProofOfStake creates a new instance of ProofOfStake, bound to a specific deployed contract.
func NewProofOfStake(address common.Address, backend bind.ContractBackend) (*ProofOfStake, error) {
	contract, err := bindProofOfStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProofOfStake{ProofOfStakeCaller: ProofOfStakeCaller{contract: contract}, ProofOfStakeTransactor: ProofOfStakeTransactor{contract: contract}, ProofOfStakeFilterer: ProofOfStakeFilterer{contract: contract}}, nil
}

// NewProofOfStakeCaller creates a new read-only instance of ProofOfStake, bound to a specific deployed contract.
func NewProofOfStakeCaller(address common.Address, caller bind.ContractCaller) (*ProofOfStakeCaller, error) {
	contract, err := bindProofOfStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProofOfStakeCaller{contract: contract}, nil
}

// NewProofOfStakeTransactor creates a new write-only instance of ProofOfStake, bound to a specific deployed contract.
func NewProofOfStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*ProofOfStakeTransactor, error) {
	contract, err := bindProofOfStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProofOfStakeTransactor{contract: contract}, nil
}

// NewProofOfStakeFilterer creates a new log filterer instance of ProofOfStake, bound to a specific deployed contract.
func NewProofOfStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*ProofOfStakeFilterer, error) {
	contract, err := bindProofOfStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProofOfStakeFilterer{contract: contract}, nil
}

// bindProofOfStake binds a generic wrapper to an already deployed contract.
func bindProofOfStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProofOfStakeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProofOfStake *ProofOfStakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProofOfStake.Contract.ProofOfStakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProofOfStake *ProofOfStakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProofOfStake.Contract.ProofOfStakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProofOfStake *ProofOfStakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProofOfStake.Contract.ProofOfStakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProofOfStake *ProofOfStakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProofOfStake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProofOfStake *ProofOfStakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProofOfStake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProofOfStake *ProofOfStakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProofOfStake.Contract.contract.Transact(opts, method, params...)
}

// Eip712DomainHash is a free data retrieval call binding the contract method 0xf94dc4bc.
//
// Solidity: function eip712DomainHash() view returns(bytes32)
func (_ProofOfStake *ProofOfStakeCaller) Eip712DomainHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProofOfStake.contract.Call(opts, &out, "eip712DomainHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Eip712DomainHash is a free data retrieval call binding the contract method 0xf94dc4bc.
//
// Solidity: function eip712DomainHash() view returns(bytes32)
func (_ProofOfStake *ProofOfStakeSession) Eip712DomainHash() ([32]byte, error) {
	return _ProofOfStake.Contract.Eip712DomainHash(&_ProofOfStake.CallOpts)
}

// Eip712DomainHash is a free data retrieval call binding the contract method 0xf94dc4bc.
//
// Solidity: function eip712DomainHash() view returns(bytes32)
func (_ProofOfStake *ProofOfStakeCallerSession) Eip712DomainHash() ([32]byte, error) {
	return _ProofOfStake.Contract.Eip712DomainHash(&_ProofOfStake.CallOpts)
}

// GetStake is a free data retrieval call binding the contract method 0x7a766460.
//
// Solidity: function getStake(address user) view returns((uint256,uint256,uint256[],uint256))
func (_ProofOfStake *ProofOfStakeCaller) GetStake(opts *bind.CallOpts, user common.Address) (ProofOfStakeStake, error) {
	var out []interface{}
	err := _ProofOfStake.contract.Call(opts, &out, "getStake", user)

	if err != nil {
		return *new(ProofOfStakeStake), err
	}

	out0 := *abi.ConvertType(out[0], new(ProofOfStakeStake)).(*ProofOfStakeStake)

	return out0, err

}

// GetStake is a free data retrieval call binding the contract method 0x7a766460.
//
// Solidity: function getStake(address user) view returns((uint256,uint256,uint256[],uint256))
func (_ProofOfStake *ProofOfStakeSession) GetStake(user common.Address) (ProofOfStakeStake, error) {
	return _ProofOfStake.Contract.GetStake(&_ProofOfStake.CallOpts, user)
}

// GetStake is a free data retrieval call binding the contract method 0x7a766460.
//
// Solidity: function getStake(address user) view returns((uint256,uint256,uint256[],uint256))
func (_ProofOfStake *ProofOfStakeCallerSession) GetStake(user common.Address) (ProofOfStakeStake, error) {
	return _ProofOfStake.Contract.GetStake(&_ProofOfStake.CallOpts, user)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_ProofOfStake *ProofOfStakeCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ProofOfStake.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_ProofOfStake *ProofOfStakeSession) GetValidators() ([]common.Address, error) {
	return _ProofOfStake.Contract.GetValidators(&_ProofOfStake.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_ProofOfStake *ProofOfStakeCallerSession) GetValidators() ([]common.Address, error) {
	return _ProofOfStake.Contract.GetValidators(&_ProofOfStake.CallOpts)
}

// GetValidators0 is a free data retrieval call binding the contract method 0xbff02e20.
//
// Solidity: function getValidators(uint256 start, uint256 end) view returns(address[])
func (_ProofOfStake *ProofOfStakeCaller) GetValidators0(opts *bind.CallOpts, start *big.Int, end *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _ProofOfStake.contract.Call(opts, &out, "getValidators0", start, end)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators0 is a free data retrieval call binding the contract method 0xbff02e20.
//
// Solidity: function getValidators(uint256 start, uint256 end) view returns(address[])
func (_ProofOfStake *ProofOfStakeSession) GetValidators0(start *big.Int, end *big.Int) ([]common.Address, error) {
	return _ProofOfStake.Contract.GetValidators0(&_ProofOfStake.CallOpts, start, end)
}

// GetValidators0 is a free data retrieval call binding the contract method 0xbff02e20.
//
// Solidity: function getValidators(uint256 start, uint256 end) view returns(address[])
func (_ProofOfStake *ProofOfStakeCallerSession) GetValidators0(start *big.Int, end *big.Int) ([]common.Address, error) {
	return _ProofOfStake.Contract.GetValidators0(&_ProofOfStake.CallOpts, start, end)
}

// NftPrices is a free data retrieval call binding the contract method 0xd9f10a2b.
//
// Solidity: function nftPrices(uint256 ) view returns(uint256)
func (_ProofOfStake *ProofOfStakeCaller) NftPrices(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ProofOfStake.contract.Call(opts, &out, "nftPrices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftPrices is a free data retrieval call binding the contract method 0xd9f10a2b.
//
// Solidity: function nftPrices(uint256 ) view returns(uint256)
func (_ProofOfStake *ProofOfStakeSession) NftPrices(arg0 *big.Int) (*big.Int, error) {
	return _ProofOfStake.Contract.NftPrices(&_ProofOfStake.CallOpts, arg0)
}

// NftPrices is a free data retrieval call binding the contract method 0xd9f10a2b.
//
// Solidity: function nftPrices(uint256 ) view returns(uint256)
func (_ProofOfStake *ProofOfStakeCallerSession) NftPrices(arg0 *big.Int) (*big.Int, error) {
	return _ProofOfStake.Contract.NftPrices(&_ProofOfStake.CallOpts, arg0)
}

// NftToken is a free data retrieval call binding the contract method 0xd06fcba8.
//
// Solidity: function nftToken() view returns(address)
func (_ProofOfStake *ProofOfStakeCaller) NftToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProofOfStake.contract.Call(opts, &out, "nftToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftToken is a free data retrieval call binding the contract method 0xd06fcba8.
//
// Solidity: function nftToken() view returns(address)
func (_ProofOfStake *ProofOfStakeSession) NftToken() (common.Address, error) {
	return _ProofOfStake.Contract.NftToken(&_ProofOfStake.CallOpts)
}

// NftToken is a free data retrieval call binding the contract method 0xd06fcba8.
//
// Solidity: function nftToken() view returns(address)
func (_ProofOfStake *ProofOfStakeCallerSession) NftToken() (common.Address, error) {
	return _ProofOfStake.Contract.NftToken(&_ProofOfStake.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(uint256)
func (_ProofOfStake *ProofOfStakeCaller) Owner(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProofOfStake.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(uint256)
func (_ProofOfStake *ProofOfStakeSession) Owner() (*big.Int, error) {
	return _ProofOfStake.Contract.Owner(&_ProofOfStake.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(uint256)
func (_ProofOfStake *ProofOfStakeCallerSession) Owner() (*big.Int, error) {
	return _ProofOfStake.Contract.Owner(&_ProofOfStake.CallOpts)
}

// Processed is a free data retrieval call binding the contract method 0xc1f0808a.
//
// Solidity: function processed(bytes32 ) view returns(bool)
func (_ProofOfStake *ProofOfStakeCaller) Processed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ProofOfStake.contract.Call(opts, &out, "processed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Processed is a free data retrieval call binding the contract method 0xc1f0808a.
//
// Solidity: function processed(bytes32 ) view returns(bool)
func (_ProofOfStake *ProofOfStakeSession) Processed(arg0 [32]byte) (bool, error) {
	return _ProofOfStake.Contract.Processed(&_ProofOfStake.CallOpts, arg0)
}

// Processed is a free data retrieval call binding the contract method 0xc1f0808a.
//
// Solidity: function processed(bytes32 ) view returns(bool)
func (_ProofOfStake *ProofOfStakeCallerSession) Processed(arg0 [32]byte) (bool, error) {
	return _ProofOfStake.Contract.Processed(&_ProofOfStake.CallOpts, arg0)
}

// SafeVerify is a free data retrieval call binding the contract method 0xac4185e8.
//
// Solidity: function safeVerify(bytes32 eip712Hash, (uint256,uint256) schnorrSignature) view returns()
func (_ProofOfStake *ProofOfStakeCaller) SafeVerify(opts *bind.CallOpts, eip712Hash [32]byte, schnorrSignature SchnorrSignatureSignature) error {
	var out []interface{}
	err := _ProofOfStake.contract.Call(opts, &out, "safeVerify", eip712Hash, schnorrSignature)

	if err != nil {
		return err
	}

	return err

}

// SafeVerify is a free data retrieval call binding the contract method 0xac4185e8.
//
// Solidity: function safeVerify(bytes32 eip712Hash, (uint256,uint256) schnorrSignature) view returns()
func (_ProofOfStake *ProofOfStakeSession) SafeVerify(eip712Hash [32]byte, schnorrSignature SchnorrSignatureSignature) error {
	return _ProofOfStake.Contract.SafeVerify(&_ProofOfStake.CallOpts, eip712Hash, schnorrSignature)
}

// SafeVerify is a free data retrieval call binding the contract method 0xac4185e8.
//
// Solidity: function safeVerify(bytes32 eip712Hash, (uint256,uint256) schnorrSignature) view returns()
func (_ProofOfStake *ProofOfStakeCallerSession) SafeVerify(eip712Hash [32]byte, schnorrSignature SchnorrSignatureSignature) error {
	return _ProofOfStake.Contract.SafeVerify(&_ProofOfStake.CallOpts, eip712Hash, schnorrSignature)
}

// SchnorrParticipationThreshold is a free data retrieval call binding the contract method 0x82e51aca.
//
// Solidity: function schnorrParticipationThreshold() view returns(uint256)
func (_ProofOfStake *ProofOfStakeCaller) SchnorrParticipationThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProofOfStake.contract.Call(opts, &out, "schnorrParticipationThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SchnorrParticipationThreshold is a free data retrieval call binding the contract method 0x82e51aca.
//
// Solidity: function schnorrParticipationThreshold() view returns(uint256)
func (_ProofOfStake *ProofOfStakeSession) SchnorrParticipationThreshold() (*big.Int, error) {
	return _ProofOfStake.Contract.SchnorrParticipationThreshold(&_ProofOfStake.CallOpts)
}

// SchnorrParticipationThreshold is a free data retrieval call binding the contract method 0x82e51aca.
//
// Solidity: function schnorrParticipationThreshold() view returns(uint256)
func (_ProofOfStake *ProofOfStakeCallerSession) SchnorrParticipationThreshold() (*big.Int, error) {
	return _ProofOfStake.Contract.SchnorrParticipationThreshold(&_ProofOfStake.CallOpts)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256 amount, uint256 end, uint256 nftSum)
func (_ProofOfStake *ProofOfStakeCaller) Stakes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount *big.Int
	End    *big.Int
	NftSum *big.Int
}, error) {
	var out []interface{}
	err := _ProofOfStake.contract.Call(opts, &out, "stakes", arg0)

	outstruct := new(struct {
		Amount *big.Int
		End    *big.Int
		NftSum *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NftSum = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256 amount, uint256 end, uint256 nftSum)
func (_ProofOfStake *ProofOfStakeSession) Stakes(arg0 common.Address) (struct {
	Amount *big.Int
	End    *big.Int
	NftSum *big.Int
}, error) {
	return _ProofOfStake.Contract.Stakes(&_ProofOfStake.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256 amount, uint256 end, uint256 nftSum)
func (_ProofOfStake *ProofOfStakeCallerSession) Stakes(arg0 common.Address) (struct {
	Amount *big.Int
	End    *big.Int
	NftSum *big.Int
}, error) {
	return _ProofOfStake.Contract.Stakes(&_ProofOfStake.CallOpts, arg0)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_ProofOfStake *ProofOfStakeCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProofOfStake.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_ProofOfStake *ProofOfStakeSession) StakingToken() (common.Address, error) {
	return _ProofOfStake.Contract.StakingToken(&_ProofOfStake.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_ProofOfStake *ProofOfStakeCallerSession) StakingToken() (common.Address, error) {
	return _ProofOfStake.Contract.StakingToken(&_ProofOfStake.CallOpts)
}

// ExtendStake is a paid mutator transaction binding the contract method 0x7e49627d.
//
// Solidity: function extendStake(uint256 duration) returns()
func (_ProofOfStake *ProofOfStakeTransactor) ExtendStake(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _ProofOfStake.contract.Transact(opts, "extendStake", duration)
}

// ExtendStake is a paid mutator transaction binding the contract method 0x7e49627d.
//
// Solidity: function extendStake(uint256 duration) returns()
func (_ProofOfStake *ProofOfStakeSession) ExtendStake(duration *big.Int) (*types.Transaction, error) {
	return _ProofOfStake.Contract.ExtendStake(&_ProofOfStake.TransactOpts, duration)
}

// ExtendStake is a paid mutator transaction binding the contract method 0x7e49627d.
//
// Solidity: function extendStake(uint256 duration) returns()
func (_ProofOfStake *ProofOfStakeTransactorSession) ExtendStake(duration *big.Int) (*types.Transaction, error) {
	return _ProofOfStake.Contract.ExtendStake(&_ProofOfStake.TransactOpts, duration)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0x0062ad9d.
//
// Solidity: function increaseStake(uint256 amount, uint256[] nfts) returns()
func (_ProofOfStake *ProofOfStakeTransactor) IncreaseStake(opts *bind.TransactOpts, amount *big.Int, nfts []*big.Int) (*types.Transaction, error) {
	return _ProofOfStake.contract.Transact(opts, "increaseStake", amount, nfts)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0x0062ad9d.
//
// Solidity: function increaseStake(uint256 amount, uint256[] nfts) returns()
func (_ProofOfStake *ProofOfStakeSession) IncreaseStake(amount *big.Int, nfts []*big.Int) (*types.Transaction, error) {
	return _ProofOfStake.Contract.IncreaseStake(&_ProofOfStake.TransactOpts, amount, nfts)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0x0062ad9d.
//
// Solidity: function increaseStake(uint256 amount, uint256[] nfts) returns()
func (_ProofOfStake *ProofOfStakeTransactorSession) IncreaseStake(amount *big.Int, nfts []*big.Int) (*types.Transaction, error) {
	return _ProofOfStake.Contract.IncreaseStake(&_ProofOfStake.TransactOpts, amount, nfts)
}

// SetNftPrices is a paid mutator transaction binding the contract method 0x556935eb.
//
// Solidity: function setNftPrices((uint256[],uint256[],uint256) prices, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeTransactor) SetNftPrices(opts *bind.TransactOpts, prices SetNftPricesNftPrices, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.contract.Transact(opts, "setNftPrices", prices, schnorrSignature)
}

// SetNftPrices is a paid mutator transaction binding the contract method 0x556935eb.
//
// Solidity: function setNftPrices((uint256[],uint256[],uint256) prices, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeSession) SetNftPrices(prices SetNftPricesNftPrices, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.Contract.SetNftPrices(&_ProofOfStake.TransactOpts, prices, schnorrSignature)
}

// SetNftPrices is a paid mutator transaction binding the contract method 0x556935eb.
//
// Solidity: function setNftPrices((uint256[],uint256[],uint256) prices, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeTransactorSession) SetNftPrices(prices SetNftPricesNftPrices, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.Contract.SetNftPrices(&_ProofOfStake.TransactOpts, prices, schnorrSignature)
}

// SetSchNorrParticipationThreshold is a paid mutator transaction binding the contract method 0x41bd25f5.
//
// Solidity: function setSchNorrParticipationThreshold((uint256,uint256) threshold, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeTransactor) SetSchNorrParticipationThreshold(opts *bind.TransactOpts, threshold SetSchnorrThresholdSchnorrThreshold, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.contract.Transact(opts, "setSchNorrParticipationThreshold", threshold, schnorrSignature)
}

// SetSchNorrParticipationThreshold is a paid mutator transaction binding the contract method 0x41bd25f5.
//
// Solidity: function setSchNorrParticipationThreshold((uint256,uint256) threshold, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeSession) SetSchNorrParticipationThreshold(threshold SetSchnorrThresholdSchnorrThreshold, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.Contract.SetSchNorrParticipationThreshold(&_ProofOfStake.TransactOpts, threshold, schnorrSignature)
}

// SetSchNorrParticipationThreshold is a paid mutator transaction binding the contract method 0x41bd25f5.
//
// Solidity: function setSchNorrParticipationThreshold((uint256,uint256) threshold, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeTransactorSession) SetSchNorrParticipationThreshold(threshold SetSchnorrThresholdSchnorrThreshold, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.Contract.SetSchNorrParticipationThreshold(&_ProofOfStake.TransactOpts, threshold, schnorrSignature)
}

// Stake is a paid mutator transaction binding the contract method 0x9debdddc.
//
// Solidity: function stake(uint256 amount, uint256 duration, uint256[] nfts) returns()
func (_ProofOfStake *ProofOfStakeTransactor) Stake(opts *bind.TransactOpts, amount *big.Int, duration *big.Int, nfts []*big.Int) (*types.Transaction, error) {
	return _ProofOfStake.contract.Transact(opts, "stake", amount, duration, nfts)
}

// Stake is a paid mutator transaction binding the contract method 0x9debdddc.
//
// Solidity: function stake(uint256 amount, uint256 duration, uint256[] nfts) returns()
func (_ProofOfStake *ProofOfStakeSession) Stake(amount *big.Int, duration *big.Int, nfts []*big.Int) (*types.Transaction, error) {
	return _ProofOfStake.Contract.Stake(&_ProofOfStake.TransactOpts, amount, duration, nfts)
}

// Stake is a paid mutator transaction binding the contract method 0x9debdddc.
//
// Solidity: function stake(uint256 amount, uint256 duration, uint256[] nfts) returns()
func (_ProofOfStake *ProofOfStakeTransactorSession) Stake(amount *big.Int, duration *big.Int, nfts []*big.Int) (*types.Transaction, error) {
	return _ProofOfStake.Contract.Stake(&_ProofOfStake.TransactOpts, amount, duration, nfts)
}

// Transfer is a paid mutator transaction binding the contract method 0x423ae347.
//
// Solidity: function transfer((address,address,uint256,uint256) txn, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeTransactor) Transfer(opts *bind.TransactOpts, txn SchnorrTransferTransfer, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.contract.Transact(opts, "transfer", txn, schnorrSignature)
}

// Transfer is a paid mutator transaction binding the contract method 0x423ae347.
//
// Solidity: function transfer((address,address,uint256,uint256) txn, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeSession) Transfer(txn SchnorrTransferTransfer, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.Contract.Transfer(&_ProofOfStake.TransactOpts, txn, schnorrSignature)
}

// Transfer is a paid mutator transaction binding the contract method 0x423ae347.
//
// Solidity: function transfer((address,address,uint256,uint256) txn, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeTransactorSession) Transfer(txn SchnorrTransferTransfer, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.Contract.Transfer(&_ProofOfStake.TransactOpts, txn, schnorrSignature)
}

// TransferNft is a paid mutator transaction binding the contract method 0x9571efb9.
//
// Solidity: function transferNft((address,address,uint256,uint256) txn, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeTransactor) TransferNft(opts *bind.TransactOpts, txn SchnorrNftTransferTransfer, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.contract.Transact(opts, "transferNft", txn, schnorrSignature)
}

// TransferNft is a paid mutator transaction binding the contract method 0x9571efb9.
//
// Solidity: function transferNft((address,address,uint256,uint256) txn, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeSession) TransferNft(txn SchnorrNftTransferTransfer, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.Contract.TransferNft(&_ProofOfStake.TransactOpts, txn, schnorrSignature)
}

// TransferNft is a paid mutator transaction binding the contract method 0x9571efb9.
//
// Solidity: function transferNft((address,address,uint256,uint256) txn, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeTransactorSession) TransferNft(txn SchnorrNftTransferTransfer, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.Contract.TransferNft(&_ProofOfStake.TransactOpts, txn, schnorrSignature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x7b8d534c.
//
// Solidity: function transferOwnership((uint256,uint256) txn, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeTransactor) TransferOwnership(opts *bind.TransactOpts, txn SchnorrTransferOwnershipTransferOwnership, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.contract.Transact(opts, "transferOwnership", txn, schnorrSignature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x7b8d534c.
//
// Solidity: function transferOwnership((uint256,uint256) txn, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeSession) TransferOwnership(txn SchnorrTransferOwnershipTransferOwnership, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.Contract.TransferOwnership(&_ProofOfStake.TransactOpts, txn, schnorrSignature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x7b8d534c.
//
// Solidity: function transferOwnership((uint256,uint256) txn, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeTransactorSession) TransferOwnership(txn SchnorrTransferOwnershipTransferOwnership, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.Contract.TransferOwnership(&_ProofOfStake.TransactOpts, txn, schnorrSignature)
}

// VerifySetNftPrice is a paid mutator transaction binding the contract method 0xecefbdbe.
//
// Solidity: function verifySetNftPrice((uint256[],uint256[],uint256) prices, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeTransactor) VerifySetNftPrice(opts *bind.TransactOpts, prices SetNftPricesNftPrices, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.contract.Transact(opts, "verifySetNftPrice", prices, schnorrSignature)
}

// VerifySetNftPrice is a paid mutator transaction binding the contract method 0xecefbdbe.
//
// Solidity: function verifySetNftPrice((uint256[],uint256[],uint256) prices, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeSession) VerifySetNftPrice(prices SetNftPricesNftPrices, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.Contract.VerifySetNftPrice(&_ProofOfStake.TransactOpts, prices, schnorrSignature)
}

// VerifySetNftPrice is a paid mutator transaction binding the contract method 0xecefbdbe.
//
// Solidity: function verifySetNftPrice((uint256[],uint256[],uint256) prices, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeTransactorSession) VerifySetNftPrice(prices SetNftPricesNftPrices, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.Contract.VerifySetNftPrice(&_ProofOfStake.TransactOpts, prices, schnorrSignature)
}

// VerifySetSchnorrThreshold is a paid mutator transaction binding the contract method 0x9a1bc8b8.
//
// Solidity: function verifySetSchnorrThreshold((uint256,uint256) threshold, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeTransactor) VerifySetSchnorrThreshold(opts *bind.TransactOpts, threshold SetSchnorrThresholdSchnorrThreshold, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.contract.Transact(opts, "verifySetSchnorrThreshold", threshold, schnorrSignature)
}

// VerifySetSchnorrThreshold is a paid mutator transaction binding the contract method 0x9a1bc8b8.
//
// Solidity: function verifySetSchnorrThreshold((uint256,uint256) threshold, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeSession) VerifySetSchnorrThreshold(threshold SetSchnorrThresholdSchnorrThreshold, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.Contract.VerifySetSchnorrThreshold(&_ProofOfStake.TransactOpts, threshold, schnorrSignature)
}

// VerifySetSchnorrThreshold is a paid mutator transaction binding the contract method 0x9a1bc8b8.
//
// Solidity: function verifySetSchnorrThreshold((uint256,uint256) threshold, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeTransactorSession) VerifySetSchnorrThreshold(threshold SetSchnorrThresholdSchnorrThreshold, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.Contract.VerifySetSchnorrThreshold(&_ProofOfStake.TransactOpts, threshold, schnorrSignature)
}

// VerifyTransfer is a paid mutator transaction binding the contract method 0x69e1f231.
//
// Solidity: function verifyTransfer((address,address,uint256,uint256) txn, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeTransactor) VerifyTransfer(opts *bind.TransactOpts, txn SchnorrTransferTransfer, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.contract.Transact(opts, "verifyTransfer", txn, schnorrSignature)
}

// VerifyTransfer is a paid mutator transaction binding the contract method 0x69e1f231.
//
// Solidity: function verifyTransfer((address,address,uint256,uint256) txn, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeSession) VerifyTransfer(txn SchnorrTransferTransfer, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.Contract.VerifyTransfer(&_ProofOfStake.TransactOpts, txn, schnorrSignature)
}

// VerifyTransfer is a paid mutator transaction binding the contract method 0x69e1f231.
//
// Solidity: function verifyTransfer((address,address,uint256,uint256) txn, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeTransactorSession) VerifyTransfer(txn SchnorrTransferTransfer, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.Contract.VerifyTransfer(&_ProofOfStake.TransactOpts, txn, schnorrSignature)
}

// VerifyTransferNft is a paid mutator transaction binding the contract method 0x65951984.
//
// Solidity: function verifyTransferNft((address,address,uint256,uint256) txn, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeTransactor) VerifyTransferNft(opts *bind.TransactOpts, txn SchnorrNftTransferTransfer, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.contract.Transact(opts, "verifyTransferNft", txn, schnorrSignature)
}

// VerifyTransferNft is a paid mutator transaction binding the contract method 0x65951984.
//
// Solidity: function verifyTransferNft((address,address,uint256,uint256) txn, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeSession) VerifyTransferNft(txn SchnorrNftTransferTransfer, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.Contract.VerifyTransferNft(&_ProofOfStake.TransactOpts, txn, schnorrSignature)
}

// VerifyTransferNft is a paid mutator transaction binding the contract method 0x65951984.
//
// Solidity: function verifyTransferNft((address,address,uint256,uint256) txn, (uint256,uint256) schnorrSignature) returns()
func (_ProofOfStake *ProofOfStakeTransactorSession) VerifyTransferNft(txn SchnorrNftTransferTransfer, schnorrSignature SchnorrSignatureSignature) (*types.Transaction, error) {
	return _ProofOfStake.Contract.VerifyTransferNft(&_ProofOfStake.TransactOpts, txn, schnorrSignature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ProofOfStake *ProofOfStakeTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProofOfStake.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ProofOfStake *ProofOfStakeSession) Withdraw() (*types.Transaction, error) {
	return _ProofOfStake.Contract.Withdraw(&_ProofOfStake.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ProofOfStake *ProofOfStakeTransactorSession) Withdraw() (*types.Transaction, error) {
	return _ProofOfStake.Contract.Withdraw(&_ProofOfStake.TransactOpts)
}

// ProofOfStakeExtendedIterator is returned from FilterExtended and is used to iterate over the raw logs and unpacked data for Extended events raised by the ProofOfStake contract.
type ProofOfStakeExtendedIterator struct {
	Event *ProofOfStakeExtended // Event containing the contract specifics and raw log

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
func (it *ProofOfStakeExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofOfStakeExtended)
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
		it.Event = new(ProofOfStakeExtended)
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
func (it *ProofOfStakeExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofOfStakeExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofOfStakeExtended represents a Extended event raised by the ProofOfStake contract.
type ProofOfStakeExtended struct {
	User common.Address
	End  *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterExtended is a free log retrieval operation binding the contract event 0xa29fc12cda82ff659de006abb10fa5ee256d922af1661e395e5f2fb6b004387e.
//
// Solidity: event Extended(address indexed user, uint256 end)
func (_ProofOfStake *ProofOfStakeFilterer) FilterExtended(opts *bind.FilterOpts, user []common.Address) (*ProofOfStakeExtendedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ProofOfStake.contract.FilterLogs(opts, "Extended", userRule)
	if err != nil {
		return nil, err
	}
	return &ProofOfStakeExtendedIterator{contract: _ProofOfStake.contract, event: "Extended", logs: logs, sub: sub}, nil
}

// WatchExtended is a free log subscription operation binding the contract event 0xa29fc12cda82ff659de006abb10fa5ee256d922af1661e395e5f2fb6b004387e.
//
// Solidity: event Extended(address indexed user, uint256 end)
func (_ProofOfStake *ProofOfStakeFilterer) WatchExtended(opts *bind.WatchOpts, sink chan<- *ProofOfStakeExtended, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ProofOfStake.contract.WatchLogs(opts, "Extended", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofOfStakeExtended)
				if err := _ProofOfStake.contract.UnpackLog(event, "Extended", log); err != nil {
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
// Solidity: event Extended(address indexed user, uint256 end)
func (_ProofOfStake *ProofOfStakeFilterer) ParseExtended(log types.Log) (*ProofOfStakeExtended, error) {
	event := new(ProofOfStakeExtended)
	if err := _ProofOfStake.contract.UnpackLog(event, "Extended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofOfStakeIncreasedIterator is returned from FilterIncreased and is used to iterate over the raw logs and unpacked data for Increased events raised by the ProofOfStake contract.
type ProofOfStakeIncreasedIterator struct {
	Event *ProofOfStakeIncreased // Event containing the contract specifics and raw log

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
func (it *ProofOfStakeIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofOfStakeIncreased)
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
		it.Event = new(ProofOfStakeIncreased)
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
func (it *ProofOfStakeIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofOfStakeIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofOfStakeIncreased represents a Increased event raised by the ProofOfStake contract.
type ProofOfStakeIncreased struct {
	User   common.Address
	Amount *big.Int
	Nfts   []*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIncreased is a free log retrieval operation binding the contract event 0x2db19fadbf40ed77512048f708e78f52f7727cc5af099c414320de1c2137aa88.
//
// Solidity: event Increased(address indexed user, uint256 amount, uint256[] nfts)
func (_ProofOfStake *ProofOfStakeFilterer) FilterIncreased(opts *bind.FilterOpts, user []common.Address) (*ProofOfStakeIncreasedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ProofOfStake.contract.FilterLogs(opts, "Increased", userRule)
	if err != nil {
		return nil, err
	}
	return &ProofOfStakeIncreasedIterator{contract: _ProofOfStake.contract, event: "Increased", logs: logs, sub: sub}, nil
}

// WatchIncreased is a free log subscription operation binding the contract event 0x2db19fadbf40ed77512048f708e78f52f7727cc5af099c414320de1c2137aa88.
//
// Solidity: event Increased(address indexed user, uint256 amount, uint256[] nfts)
func (_ProofOfStake *ProofOfStakeFilterer) WatchIncreased(opts *bind.WatchOpts, sink chan<- *ProofOfStakeIncreased, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ProofOfStake.contract.WatchLogs(opts, "Increased", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofOfStakeIncreased)
				if err := _ProofOfStake.contract.UnpackLog(event, "Increased", log); err != nil {
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

// ParseIncreased is a log parse operation binding the contract event 0x2db19fadbf40ed77512048f708e78f52f7727cc5af099c414320de1c2137aa88.
//
// Solidity: event Increased(address indexed user, uint256 amount, uint256[] nfts)
func (_ProofOfStake *ProofOfStakeFilterer) ParseIncreased(log types.Log) (*ProofOfStakeIncreased, error) {
	event := new(ProofOfStakeIncreased)
	if err := _ProofOfStake.contract.UnpackLog(event, "Increased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofOfStakeStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the ProofOfStake contract.
type ProofOfStakeStakedIterator struct {
	Event *ProofOfStakeStaked // Event containing the contract specifics and raw log

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
func (it *ProofOfStakeStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofOfStakeStaked)
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
		it.Event = new(ProofOfStakeStaked)
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
func (it *ProofOfStakeStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofOfStakeStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofOfStakeStaked represents a Staked event raised by the ProofOfStake contract.
type ProofOfStakeStaked struct {
	User   common.Address
	Amount *big.Int
	End    *big.Int
	Nfts   []*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x5df5de9ccb680fe3d60088f6d4c3b6d535074c704699377046c743a5b276e171.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 end, uint256[] nfts)
func (_ProofOfStake *ProofOfStakeFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*ProofOfStakeStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ProofOfStake.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &ProofOfStakeStakedIterator{contract: _ProofOfStake.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x5df5de9ccb680fe3d60088f6d4c3b6d535074c704699377046c743a5b276e171.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 end, uint256[] nfts)
func (_ProofOfStake *ProofOfStakeFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *ProofOfStakeStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ProofOfStake.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofOfStakeStaked)
				if err := _ProofOfStake.contract.UnpackLog(event, "Staked", log); err != nil {
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
// Solidity: event Staked(address indexed user, uint256 amount, uint256 end, uint256[] nfts)
func (_ProofOfStake *ProofOfStakeFilterer) ParseStaked(log types.Log) (*ProofOfStakeStaked, error) {
	event := new(ProofOfStakeStaked)
	if err := _ProofOfStake.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofOfStakeWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the ProofOfStake contract.
type ProofOfStakeWithdrawnIterator struct {
	Event *ProofOfStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *ProofOfStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofOfStakeWithdrawn)
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
		it.Event = new(ProofOfStakeWithdrawn)
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
func (it *ProofOfStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofOfStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofOfStakeWithdrawn represents a Withdrawn event raised by the ProofOfStake contract.
type ProofOfStakeWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Nfts   []*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd40a9786b597b88b3426158112e3930e09cb031c138a150e367cfe17ff20e302.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount, uint256[] nfts)
func (_ProofOfStake *ProofOfStakeFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*ProofOfStakeWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ProofOfStake.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &ProofOfStakeWithdrawnIterator{contract: _ProofOfStake.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd40a9786b597b88b3426158112e3930e09cb031c138a150e367cfe17ff20e302.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount, uint256[] nfts)
func (_ProofOfStake *ProofOfStakeFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ProofOfStakeWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ProofOfStake.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofOfStakeWithdrawn)
				if err := _ProofOfStake.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd40a9786b597b88b3426158112e3930e09cb031c138a150e367cfe17ff20e302.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount, uint256[] nfts)
func (_ProofOfStake *ProofOfStakeFilterer) ParseWithdrawn(log types.Log) (*ProofOfStakeWithdrawn, error) {
	event := new(ProofOfStakeWithdrawn)
	if err := _ProofOfStake.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
