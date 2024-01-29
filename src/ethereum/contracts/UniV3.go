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

// UniV3MetaData contains all meta data concerning the UniV3 contract.
var UniV3MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"slot0\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"uint160\",\"name\":\"sqrtPriceX96\"},{\"type\":\"int24\",\"name\":\"tick\"},{\"type\":\"uint16\",\"name\":\"observationIndex\"},{\"type\":\"uint16\",\"name\":\"observationCardinality\"},{\"type\":\"uint16\",\"name\":\"observationCardinalityNext\"},{\"type\":\"uint8\",\"name\":\"feeProtocol\"},{\"type\":\"bool\",\"name\":\"unlocked\"}]}]",
}

// UniV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use UniV3MetaData.ABI instead.
var UniV3ABI = UniV3MetaData.ABI

// UniV3 is an auto generated Go binding around an Ethereum contract.
type UniV3 struct {
	UniV3Caller     // Read-only binding to the contract
	UniV3Transactor // Write-only binding to the contract
	UniV3Filterer   // Log filterer for contract events
}

// UniV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type UniV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type UniV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniV3Session struct {
	Contract     *UniV3            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniV3CallerSession struct {
	Contract *UniV3Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UniV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniV3TransactorSession struct {
	Contract     *UniV3Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type UniV3Raw struct {
	Contract *UniV3 // Generic contract binding to access the raw methods on
}

// UniV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniV3CallerRaw struct {
	Contract *UniV3Caller // Generic read-only contract binding to access the raw methods on
}

// UniV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniV3TransactorRaw struct {
	Contract *UniV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewUniV3 creates a new instance of UniV3, bound to a specific deployed contract.
func NewUniV3(address common.Address, backend bind.ContractBackend) (*UniV3, error) {
	contract, err := bindUniV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniV3{UniV3Caller: UniV3Caller{contract: contract}, UniV3Transactor: UniV3Transactor{contract: contract}, UniV3Filterer: UniV3Filterer{contract: contract}}, nil
}

// NewUniV3Caller creates a new read-only instance of UniV3, bound to a specific deployed contract.
func NewUniV3Caller(address common.Address, caller bind.ContractCaller) (*UniV3Caller, error) {
	contract, err := bindUniV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniV3Caller{contract: contract}, nil
}

// NewUniV3Transactor creates a new write-only instance of UniV3, bound to a specific deployed contract.
func NewUniV3Transactor(address common.Address, transactor bind.ContractTransactor) (*UniV3Transactor, error) {
	contract, err := bindUniV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniV3Transactor{contract: contract}, nil
}

// NewUniV3Filterer creates a new log filterer instance of UniV3, bound to a specific deployed contract.
func NewUniV3Filterer(address common.Address, filterer bind.ContractFilterer) (*UniV3Filterer, error) {
	contract, err := bindUniV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniV3Filterer{contract: contract}, nil
}

// bindUniV3 binds a generic wrapper to an already deployed contract.
func bindUniV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniV3 *UniV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniV3.Contract.UniV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniV3 *UniV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniV3.Contract.UniV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniV3 *UniV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniV3.Contract.UniV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniV3 *UniV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniV3 *UniV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniV3 *UniV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniV3.Contract.contract.Transact(opts, method, params...)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (_UniV3 *UniV3Caller) Slot0(opts *bind.CallOpts) (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	var out []interface{}
	err := _UniV3.contract.Call(opts, &out, "slot0")

	outstruct := new(struct {
		SqrtPriceX96               *big.Int
		Tick                       *big.Int
		ObservationIndex           uint16
		ObservationCardinality     uint16
		ObservationCardinalityNext uint16
		FeeProtocol                uint8
		Unlocked                   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SqrtPriceX96 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Tick = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ObservationIndex = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.ObservationCardinality = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.ObservationCardinalityNext = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.FeeProtocol = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.Unlocked = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (_UniV3 *UniV3Session) Slot0() (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	return _UniV3.Contract.Slot0(&_UniV3.CallOpts)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (_UniV3 *UniV3CallerSession) Slot0() (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	return _UniV3.Contract.Slot0(&_UniV3.CallOpts)
}
