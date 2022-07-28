// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
)

// PriceOracleMetaData contains all meta data concerning the PriceOracle contract.
var PriceOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"amm\",\"type\":\"address\"}],\"name\":\"getIndexPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"amm\",\"type\":\"address\"}],\"name\":\"getMarkPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isIndexPrice\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"amm\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"beta\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"negative\",\"type\":\"bool\"}],\"name\":\"getMarkPriceAcc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"amm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"name\":\"getMarkPriceAfterSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isIndexPrice\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"amm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"name\":\"getMarkPriceInRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"resultBaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultQuoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isIndexPrice\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"amm\",\"type\":\"address\"}],\"name\":\"getMarketPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"source\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"amm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"name\":\"quoteFromAmmTwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"amm\",\"type\":\"address\"}],\"name\":\"setupTwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"updateAmmTwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PriceOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceOracleMetaData.ABI instead.
var PriceOracleABI = PriceOracleMetaData.ABI

// PriceOracle is an auto generated Go binding around an Ethereum contract.
type PriceOracle struct {
	PriceOracleCaller     // Read-only binding to the contract
	PriceOracleTransactor // Write-only binding to the contract
	PriceOracleFilterer   // Log filterer for contract events
}

// PriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceOracleSession struct {
	Contract     *PriceOracle      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceOracleCallerSession struct {
	Contract *PriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceOracleTransactorSession struct {
	Contract     *PriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceOracleRaw struct {
	Contract *PriceOracle // Generic contract binding to access the raw methods on
}

// PriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceOracleCallerRaw struct {
	Contract *PriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// PriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceOracleTransactorRaw struct {
	Contract *PriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceOracle creates a new instance of PriceOracle, bound to a specific deployed contract.
func NewPriceOracle(address common.Address, backend bind.ContractBackend) (*PriceOracle, error) {
	contract, err := bindPriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceOracle{PriceOracleCaller: PriceOracleCaller{contract: contract}, PriceOracleTransactor: PriceOracleTransactor{contract: contract}, PriceOracleFilterer: PriceOracleFilterer{contract: contract}}, nil
}

// NewPriceOracleCaller creates a new read-only instance of PriceOracle, bound to a specific deployed contract.
func NewPriceOracleCaller(address common.Address, caller bind.ContractCaller) (*PriceOracleCaller, error) {
	contract, err := bindPriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceOracleCaller{contract: contract}, nil
}

// NewPriceOracleTransactor creates a new write-only instance of PriceOracle, bound to a specific deployed contract.
func NewPriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceOracleTransactor, error) {
	contract, err := bindPriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceOracleTransactor{contract: contract}, nil
}

// NewPriceOracleFilterer creates a new log filterer instance of PriceOracle, bound to a specific deployed contract.
func NewPriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceOracleFilterer, error) {
	contract, err := bindPriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceOracleFilterer{contract: contract}, nil
}

// bindPriceOracle binds a generic wrapper to an already deployed contract.
func bindPriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceOracle *PriceOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceOracle.Contract.PriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceOracle *PriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOracle.Contract.PriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceOracle *PriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceOracle.Contract.PriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceOracle *PriceOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceOracle *PriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceOracle *PriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceOracle.Contract.contract.Transact(opts, method, params...)
}

// GetIndexPrice is a free data retrieval call binding the contract method 0xb263e010.
//
// Solidity: function getIndexPrice(address amm) view returns(uint256)
func (_PriceOracle *PriceOracleCaller) GetIndexPrice(opts *bind.CallOpts, amm common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceOracle.contract.Call(opts, &out, "getIndexPrice", amm)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIndexPrice is a free data retrieval call binding the contract method 0xb263e010.
//
// Solidity: function getIndexPrice(address amm) view returns(uint256)
func (_PriceOracle *PriceOracleSession) GetIndexPrice(amm common.Address) (*big.Int, error) {
	return _PriceOracle.Contract.GetIndexPrice(&_PriceOracle.CallOpts, amm)
}

// GetIndexPrice is a free data retrieval call binding the contract method 0xb263e010.
//
// Solidity: function getIndexPrice(address amm) view returns(uint256)
func (_PriceOracle *PriceOracleCallerSession) GetIndexPrice(amm common.Address) (*big.Int, error) {
	return _PriceOracle.Contract.GetIndexPrice(&_PriceOracle.CallOpts, amm)
}

// GetMarkPrice is a free data retrieval call binding the contract method 0x5d6f9c14.
//
// Solidity: function getMarkPrice(address amm) view returns(uint256 price, bool isIndexPrice)
func (_PriceOracle *PriceOracleCaller) GetMarkPrice(opts *bind.CallOpts, amm common.Address) (struct {
	Price        *big.Int
	IsIndexPrice bool
}, error) {
	var out []interface{}
	err := _PriceOracle.contract.Call(opts, &out, "getMarkPrice", amm)

	outstruct := new(struct {
		Price        *big.Int
		IsIndexPrice bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IsIndexPrice = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetMarkPrice is a free data retrieval call binding the contract method 0x5d6f9c14.
//
// Solidity: function getMarkPrice(address amm) view returns(uint256 price, bool isIndexPrice)
func (_PriceOracle *PriceOracleSession) GetMarkPrice(amm common.Address) (struct {
	Price        *big.Int
	IsIndexPrice bool
}, error) {
	return _PriceOracle.Contract.GetMarkPrice(&_PriceOracle.CallOpts, amm)
}

// GetMarkPrice is a free data retrieval call binding the contract method 0x5d6f9c14.
//
// Solidity: function getMarkPrice(address amm) view returns(uint256 price, bool isIndexPrice)
func (_PriceOracle *PriceOracleCallerSession) GetMarkPrice(amm common.Address) (struct {
	Price        *big.Int
	IsIndexPrice bool
}, error) {
	return _PriceOracle.Contract.GetMarkPrice(&_PriceOracle.CallOpts, amm)
}

// GetMarkPriceAcc is a free data retrieval call binding the contract method 0x4c0000dc.
//
// Solidity: function getMarkPriceAcc(address amm, uint8 beta, uint256 quoteAmount, bool negative) view returns(uint256 baseAmount)
func (_PriceOracle *PriceOracleCaller) GetMarkPriceAcc(opts *bind.CallOpts, amm common.Address, beta uint8, quoteAmount *big.Int, negative bool) (*big.Int, error) {
	var out []interface{}
	err := _PriceOracle.contract.Call(opts, &out, "getMarkPriceAcc", amm, beta, quoteAmount, negative)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarkPriceAcc is a free data retrieval call binding the contract method 0x4c0000dc.
//
// Solidity: function getMarkPriceAcc(address amm, uint8 beta, uint256 quoteAmount, bool negative) view returns(uint256 baseAmount)
func (_PriceOracle *PriceOracleSession) GetMarkPriceAcc(amm common.Address, beta uint8, quoteAmount *big.Int, negative bool) (*big.Int, error) {
	return _PriceOracle.Contract.GetMarkPriceAcc(&_PriceOracle.CallOpts, amm, beta, quoteAmount, negative)
}

// GetMarkPriceAcc is a free data retrieval call binding the contract method 0x4c0000dc.
//
// Solidity: function getMarkPriceAcc(address amm, uint8 beta, uint256 quoteAmount, bool negative) view returns(uint256 baseAmount)
func (_PriceOracle *PriceOracleCallerSession) GetMarkPriceAcc(amm common.Address, beta uint8, quoteAmount *big.Int, negative bool) (*big.Int, error) {
	return _PriceOracle.Contract.GetMarkPriceAcc(&_PriceOracle.CallOpts, amm, beta, quoteAmount, negative)
}

// GetMarkPriceAfterSwap is a free data retrieval call binding the contract method 0xba4433c0.
//
// Solidity: function getMarkPriceAfterSwap(address amm, uint256 quoteAmount, uint256 baseAmount) view returns(uint256 price, bool isIndexPrice)
func (_PriceOracle *PriceOracleCaller) GetMarkPriceAfterSwap(opts *bind.CallOpts, amm common.Address, quoteAmount *big.Int, baseAmount *big.Int) (struct {
	Price        *big.Int
	IsIndexPrice bool
}, error) {
	var out []interface{}
	err := _PriceOracle.contract.Call(opts, &out, "getMarkPriceAfterSwap", amm, quoteAmount, baseAmount)

	outstruct := new(struct {
		Price        *big.Int
		IsIndexPrice bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IsIndexPrice = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetMarkPriceAfterSwap is a free data retrieval call binding the contract method 0xba4433c0.
//
// Solidity: function getMarkPriceAfterSwap(address amm, uint256 quoteAmount, uint256 baseAmount) view returns(uint256 price, bool isIndexPrice)
func (_PriceOracle *PriceOracleSession) GetMarkPriceAfterSwap(amm common.Address, quoteAmount *big.Int, baseAmount *big.Int) (struct {
	Price        *big.Int
	IsIndexPrice bool
}, error) {
	return _PriceOracle.Contract.GetMarkPriceAfterSwap(&_PriceOracle.CallOpts, amm, quoteAmount, baseAmount)
}

// GetMarkPriceAfterSwap is a free data retrieval call binding the contract method 0xba4433c0.
//
// Solidity: function getMarkPriceAfterSwap(address amm, uint256 quoteAmount, uint256 baseAmount) view returns(uint256 price, bool isIndexPrice)
func (_PriceOracle *PriceOracleCallerSession) GetMarkPriceAfterSwap(amm common.Address, quoteAmount *big.Int, baseAmount *big.Int) (struct {
	Price        *big.Int
	IsIndexPrice bool
}, error) {
	return _PriceOracle.Contract.GetMarkPriceAfterSwap(&_PriceOracle.CallOpts, amm, quoteAmount, baseAmount)
}

// GetMarkPriceInRatio is a free data retrieval call binding the contract method 0x7e650a05.
//
// Solidity: function getMarkPriceInRatio(address amm, uint256 quoteAmount, uint256 baseAmount) view returns(uint256 resultBaseAmount, uint256 resultQuoteAmount, bool isIndexPrice)
func (_PriceOracle *PriceOracleCaller) GetMarkPriceInRatio(opts *bind.CallOpts, amm common.Address, quoteAmount *big.Int, baseAmount *big.Int) (struct {
	ResultBaseAmount  *big.Int
	ResultQuoteAmount *big.Int
	IsIndexPrice      bool
}, error) {
	var out []interface{}
	err := _PriceOracle.contract.Call(opts, &out, "getMarkPriceInRatio", amm, quoteAmount, baseAmount)

	outstruct := new(struct {
		ResultBaseAmount  *big.Int
		ResultQuoteAmount *big.Int
		IsIndexPrice      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ResultBaseAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ResultQuoteAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsIndexPrice = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// GetMarkPriceInRatio is a free data retrieval call binding the contract method 0x7e650a05.
//
// Solidity: function getMarkPriceInRatio(address amm, uint256 quoteAmount, uint256 baseAmount) view returns(uint256 resultBaseAmount, uint256 resultQuoteAmount, bool isIndexPrice)
func (_PriceOracle *PriceOracleSession) GetMarkPriceInRatio(amm common.Address, quoteAmount *big.Int, baseAmount *big.Int) (struct {
	ResultBaseAmount  *big.Int
	ResultQuoteAmount *big.Int
	IsIndexPrice      bool
}, error) {
	return _PriceOracle.Contract.GetMarkPriceInRatio(&_PriceOracle.CallOpts, amm, quoteAmount, baseAmount)
}

// GetMarkPriceInRatio is a free data retrieval call binding the contract method 0x7e650a05.
//
// Solidity: function getMarkPriceInRatio(address amm, uint256 quoteAmount, uint256 baseAmount) view returns(uint256 resultBaseAmount, uint256 resultQuoteAmount, bool isIndexPrice)
func (_PriceOracle *PriceOracleCallerSession) GetMarkPriceInRatio(amm common.Address, quoteAmount *big.Int, baseAmount *big.Int) (struct {
	ResultBaseAmount  *big.Int
	ResultQuoteAmount *big.Int
	IsIndexPrice      bool
}, error) {
	return _PriceOracle.Contract.GetMarkPriceInRatio(&_PriceOracle.CallOpts, amm, quoteAmount, baseAmount)
}

// GetMarketPrice is a free data retrieval call binding the contract method 0x9c6f3026.
//
// Solidity: function getMarketPrice(address amm) view returns(uint256)
func (_PriceOracle *PriceOracleCaller) GetMarketPrice(opts *bind.CallOpts, amm common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceOracle.contract.Call(opts, &out, "getMarketPrice", amm)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketPrice is a free data retrieval call binding the contract method 0x9c6f3026.
//
// Solidity: function getMarketPrice(address amm) view returns(uint256)
func (_PriceOracle *PriceOracleSession) GetMarketPrice(amm common.Address) (*big.Int, error) {
	return _PriceOracle.Contract.GetMarketPrice(&_PriceOracle.CallOpts, amm)
}

// GetMarketPrice is a free data retrieval call binding the contract method 0x9c6f3026.
//
// Solidity: function getMarketPrice(address amm) view returns(uint256)
func (_PriceOracle *PriceOracleCallerSession) GetMarketPrice(amm common.Address) (*big.Int, error) {
	return _PriceOracle.Contract.GetMarketPrice(&_PriceOracle.CallOpts, amm)
}

// Quote is a free data retrieval call binding the contract method 0xb6466384.
//
// Solidity: function quote(address baseToken, address quoteToken, uint256 baseAmount) view returns(uint256 quoteAmount, uint8 source)
func (_PriceOracle *PriceOracleCaller) Quote(opts *bind.CallOpts, baseToken common.Address, quoteToken common.Address, baseAmount *big.Int) (struct {
	QuoteAmount *big.Int
	Source      uint8
}, error) {
	var out []interface{}
	err := _PriceOracle.contract.Call(opts, &out, "quote", baseToken, quoteToken, baseAmount)

	outstruct := new(struct {
		QuoteAmount *big.Int
		Source      uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.QuoteAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Source = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// Quote is a free data retrieval call binding the contract method 0xb6466384.
//
// Solidity: function quote(address baseToken, address quoteToken, uint256 baseAmount) view returns(uint256 quoteAmount, uint8 source)
func (_PriceOracle *PriceOracleSession) Quote(baseToken common.Address, quoteToken common.Address, baseAmount *big.Int) (struct {
	QuoteAmount *big.Int
	Source      uint8
}, error) {
	return _PriceOracle.Contract.Quote(&_PriceOracle.CallOpts, baseToken, quoteToken, baseAmount)
}

// Quote is a free data retrieval call binding the contract method 0xb6466384.
//
// Solidity: function quote(address baseToken, address quoteToken, uint256 baseAmount) view returns(uint256 quoteAmount, uint8 source)
func (_PriceOracle *PriceOracleCallerSession) Quote(baseToken common.Address, quoteToken common.Address, baseAmount *big.Int) (struct {
	QuoteAmount *big.Int
	Source      uint8
}, error) {
	return _PriceOracle.Contract.Quote(&_PriceOracle.CallOpts, baseToken, quoteToken, baseAmount)
}

// QuoteFromAmmTwap is a free data retrieval call binding the contract method 0xd480731c.
//
// Solidity: function quoteFromAmmTwap(address amm, uint256 baseAmount) view returns(uint256 quoteAmount)
func (_PriceOracle *PriceOracleCaller) QuoteFromAmmTwap(opts *bind.CallOpts, amm common.Address, baseAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PriceOracle.contract.Call(opts, &out, "quoteFromAmmTwap", amm, baseAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteFromAmmTwap is a free data retrieval call binding the contract method 0xd480731c.
//
// Solidity: function quoteFromAmmTwap(address amm, uint256 baseAmount) view returns(uint256 quoteAmount)
func (_PriceOracle *PriceOracleSession) QuoteFromAmmTwap(amm common.Address, baseAmount *big.Int) (*big.Int, error) {
	return _PriceOracle.Contract.QuoteFromAmmTwap(&_PriceOracle.CallOpts, amm, baseAmount)
}

// QuoteFromAmmTwap is a free data retrieval call binding the contract method 0xd480731c.
//
// Solidity: function quoteFromAmmTwap(address amm, uint256 baseAmount) view returns(uint256 quoteAmount)
func (_PriceOracle *PriceOracleCallerSession) QuoteFromAmmTwap(amm common.Address, baseAmount *big.Int) (*big.Int, error) {
	return _PriceOracle.Contract.QuoteFromAmmTwap(&_PriceOracle.CallOpts, amm, baseAmount)
}

// SetupTwap is a paid mutator transaction binding the contract method 0xf65770ba.
//
// Solidity: function setupTwap(address amm) returns()
func (_PriceOracle *PriceOracleTransactor) SetupTwap(opts *bind.TransactOpts, amm common.Address) (*types.Transaction, error) {
	return _PriceOracle.contract.Transact(opts, "setupTwap", amm)
}

// SetupTwap is a paid mutator transaction binding the contract method 0xf65770ba.
//
// Solidity: function setupTwap(address amm) returns()
func (_PriceOracle *PriceOracleSession) SetupTwap(amm common.Address) (*types.Transaction, error) {
	return _PriceOracle.Contract.SetupTwap(&_PriceOracle.TransactOpts, amm)
}

// SetupTwap is a paid mutator transaction binding the contract method 0xf65770ba.
//
// Solidity: function setupTwap(address amm) returns()
func (_PriceOracle *PriceOracleTransactorSession) SetupTwap(amm common.Address) (*types.Transaction, error) {
	return _PriceOracle.Contract.SetupTwap(&_PriceOracle.TransactOpts, amm)
}

// UpdateAmmTwap is a paid mutator transaction binding the contract method 0x422bb010.
//
// Solidity: function updateAmmTwap(address pair) returns()
func (_PriceOracle *PriceOracleTransactor) UpdateAmmTwap(opts *bind.TransactOpts, pair common.Address) (*types.Transaction, error) {
	return _PriceOracle.contract.Transact(opts, "updateAmmTwap", pair)
}

// UpdateAmmTwap is a paid mutator transaction binding the contract method 0x422bb010.
//
// Solidity: function updateAmmTwap(address pair) returns()
func (_PriceOracle *PriceOracleSession) UpdateAmmTwap(pair common.Address) (*types.Transaction, error) {
	return _PriceOracle.Contract.UpdateAmmTwap(&_PriceOracle.TransactOpts, pair)
}

// UpdateAmmTwap is a paid mutator transaction binding the contract method 0x422bb010.
//
// Solidity: function updateAmmTwap(address pair) returns()
func (_PriceOracle *PriceOracleTransactorSession) UpdateAmmTwap(pair common.Address) (*types.Transaction, error) {
	return _PriceOracle.Contract.UpdateAmmTwap(&_PriceOracle.TransactOpts, pair)
}
