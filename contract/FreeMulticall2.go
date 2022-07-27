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

// Multicall2Call is an auto generated low-level Go binding around an user-defined struct.
type Multicall2Call struct {
	Target   common.Address
	CallData []byte
}

// Multicall2Result is an auto generated low-level Go binding around an user-defined struct.
type Multicall2Result struct {
	Success    bool
	ReturnData []byte
}

// FreeMulticall2MetaData contains all meta data concerning the FreeMulticall2 contract.
var FreeMulticall2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"blockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockCoinbase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gaslimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryAggregate\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryBlockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall2.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FreeMulticall2ABI is the input ABI used to generate the binding from.
// Deprecated: Use FreeMulticall2MetaData.ABI instead.
var FreeMulticall2ABI = FreeMulticall2MetaData.ABI

// FreeMulticall2 is an auto generated Go binding around an Ethereum contract.
type FreeMulticall2 struct {
	FreeMulticall2Caller     // Read-only binding to the contract
	FreeMulticall2Transactor // Write-only binding to the contract
	FreeMulticall2Filterer   // Log filterer for contract events
}

// FreeMulticall2Caller is an auto generated read-only Go binding around an Ethereum contract.
type FreeMulticall2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreeMulticall2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FreeMulticall2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreeMulticall2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FreeMulticall2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreeMulticall2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FreeMulticall2Session struct {
	Contract     *FreeMulticall2   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FreeMulticall2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FreeMulticall2CallerSession struct {
	Contract *FreeMulticall2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// FreeMulticall2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FreeMulticall2TransactorSession struct {
	Contract     *FreeMulticall2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// FreeMulticall2Raw is an auto generated low-level Go binding around an Ethereum contract.
type FreeMulticall2Raw struct {
	Contract *FreeMulticall2 // Generic contract binding to access the raw methods on
}

// FreeMulticall2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FreeMulticall2CallerRaw struct {
	Contract *FreeMulticall2Caller // Generic read-only contract binding to access the raw methods on
}

// FreeMulticall2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FreeMulticall2TransactorRaw struct {
	Contract *FreeMulticall2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFreeMulticall2 creates a new instance of FreeMulticall2, bound to a specific deployed contract.
func NewFreeMulticall2(address common.Address, backend bind.ContractBackend) (*FreeMulticall2, error) {
	contract, err := bindFreeMulticall2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FreeMulticall2{FreeMulticall2Caller: FreeMulticall2Caller{contract: contract}, FreeMulticall2Transactor: FreeMulticall2Transactor{contract: contract}, FreeMulticall2Filterer: FreeMulticall2Filterer{contract: contract}}, nil
}

// NewFreeMulticall2Caller creates a new read-only instance of FreeMulticall2, bound to a specific deployed contract.
func NewFreeMulticall2Caller(address common.Address, caller bind.ContractCaller) (*FreeMulticall2Caller, error) {
	contract, err := bindFreeMulticall2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FreeMulticall2Caller{contract: contract}, nil
}

// NewFreeMulticall2Transactor creates a new write-only instance of FreeMulticall2, bound to a specific deployed contract.
func NewFreeMulticall2Transactor(address common.Address, transactor bind.ContractTransactor) (*FreeMulticall2Transactor, error) {
	contract, err := bindFreeMulticall2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FreeMulticall2Transactor{contract: contract}, nil
}

// NewFreeMulticall2Filterer creates a new log filterer instance of FreeMulticall2, bound to a specific deployed contract.
func NewFreeMulticall2Filterer(address common.Address, filterer bind.ContractFilterer) (*FreeMulticall2Filterer, error) {
	contract, err := bindFreeMulticall2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FreeMulticall2Filterer{contract: contract}, nil
}

// bindFreeMulticall2 binds a generic wrapper to an already deployed contract.
func bindFreeMulticall2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FreeMulticall2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FreeMulticall2 *FreeMulticall2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FreeMulticall2.Contract.FreeMulticall2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FreeMulticall2 *FreeMulticall2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreeMulticall2.Contract.FreeMulticall2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FreeMulticall2 *FreeMulticall2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FreeMulticall2.Contract.FreeMulticall2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FreeMulticall2 *FreeMulticall2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FreeMulticall2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FreeMulticall2 *FreeMulticall2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreeMulticall2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FreeMulticall2 *FreeMulticall2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FreeMulticall2.Contract.contract.Transact(opts, method, params...)
}

// Aggregate is a free data retrieval call binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) view returns(uint256 blockNumber, bytes[] returnData)
func (_FreeMulticall2 *FreeMulticall2Caller) Aggregate(opts *bind.CallOpts, calls []Multicall2Call) (struct {
	BlockNumber *big.Int
	ReturnData  [][]byte
}, error) {
	var out []interface{}
	err := _FreeMulticall2.contract.Call(opts, &out, "aggregate", calls)

	outstruct := new(struct {
		BlockNumber *big.Int
		ReturnData  [][]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReturnData = *abi.ConvertType(out[1], new([][]byte)).(*[][]byte)

	return *outstruct, err

}

// Aggregate is a free data retrieval call binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) view returns(uint256 blockNumber, bytes[] returnData)
func (_FreeMulticall2 *FreeMulticall2Session) Aggregate(calls []Multicall2Call) (struct {
	BlockNumber *big.Int
	ReturnData  [][]byte
}, error) {
	return _FreeMulticall2.Contract.Aggregate(&_FreeMulticall2.CallOpts, calls)
}

// Aggregate is a free data retrieval call binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) view returns(uint256 blockNumber, bytes[] returnData)
func (_FreeMulticall2 *FreeMulticall2CallerSession) Aggregate(calls []Multicall2Call) (struct {
	BlockNumber *big.Int
	ReturnData  [][]byte
}, error) {
	return _FreeMulticall2.Contract.Aggregate(&_FreeMulticall2.CallOpts, calls)
}

// BlockAndAggregate is a free data retrieval call binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) view returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_FreeMulticall2 *FreeMulticall2Caller) BlockAndAggregate(opts *bind.CallOpts, calls []Multicall2Call) (struct {
	BlockNumber *big.Int
	BlockHash   [32]byte
	ReturnData  []Multicall2Result
}, error) {
	var out []interface{}
	err := _FreeMulticall2.contract.Call(opts, &out, "blockAndAggregate", calls)

	outstruct := new(struct {
		BlockNumber *big.Int
		BlockHash   [32]byte
		ReturnData  []Multicall2Result
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BlockHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.ReturnData = *abi.ConvertType(out[2], new([]Multicall2Result)).(*[]Multicall2Result)

	return *outstruct, err

}

// BlockAndAggregate is a free data retrieval call binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) view returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_FreeMulticall2 *FreeMulticall2Session) BlockAndAggregate(calls []Multicall2Call) (struct {
	BlockNumber *big.Int
	BlockHash   [32]byte
	ReturnData  []Multicall2Result
}, error) {
	return _FreeMulticall2.Contract.BlockAndAggregate(&_FreeMulticall2.CallOpts, calls)
}

// BlockAndAggregate is a free data retrieval call binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) view returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_FreeMulticall2 *FreeMulticall2CallerSession) BlockAndAggregate(calls []Multicall2Call) (struct {
	BlockNumber *big.Int
	BlockHash   [32]byte
	ReturnData  []Multicall2Result
}, error) {
	return _FreeMulticall2.Contract.BlockAndAggregate(&_FreeMulticall2.CallOpts, calls)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_FreeMulticall2 *FreeMulticall2Caller) GetBlockHash(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _FreeMulticall2.contract.Call(opts, &out, "getBlockHash", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_FreeMulticall2 *FreeMulticall2Session) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _FreeMulticall2.Contract.GetBlockHash(&_FreeMulticall2.CallOpts, blockNumber)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_FreeMulticall2 *FreeMulticall2CallerSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _FreeMulticall2.Contract.GetBlockHash(&_FreeMulticall2.CallOpts, blockNumber)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_FreeMulticall2 *FreeMulticall2Caller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FreeMulticall2.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_FreeMulticall2 *FreeMulticall2Session) GetBlockNumber() (*big.Int, error) {
	return _FreeMulticall2.Contract.GetBlockNumber(&_FreeMulticall2.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_FreeMulticall2 *FreeMulticall2CallerSession) GetBlockNumber() (*big.Int, error) {
	return _FreeMulticall2.Contract.GetBlockNumber(&_FreeMulticall2.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_FreeMulticall2 *FreeMulticall2Caller) GetCurrentBlockCoinbase(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FreeMulticall2.contract.Call(opts, &out, "getCurrentBlockCoinbase")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_FreeMulticall2 *FreeMulticall2Session) GetCurrentBlockCoinbase() (common.Address, error) {
	return _FreeMulticall2.Contract.GetCurrentBlockCoinbase(&_FreeMulticall2.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_FreeMulticall2 *FreeMulticall2CallerSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _FreeMulticall2.Contract.GetCurrentBlockCoinbase(&_FreeMulticall2.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_FreeMulticall2 *FreeMulticall2Caller) GetCurrentBlockDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FreeMulticall2.contract.Call(opts, &out, "getCurrentBlockDifficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_FreeMulticall2 *FreeMulticall2Session) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _FreeMulticall2.Contract.GetCurrentBlockDifficulty(&_FreeMulticall2.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_FreeMulticall2 *FreeMulticall2CallerSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _FreeMulticall2.Contract.GetCurrentBlockDifficulty(&_FreeMulticall2.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_FreeMulticall2 *FreeMulticall2Caller) GetCurrentBlockGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FreeMulticall2.contract.Call(opts, &out, "getCurrentBlockGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_FreeMulticall2 *FreeMulticall2Session) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _FreeMulticall2.Contract.GetCurrentBlockGasLimit(&_FreeMulticall2.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_FreeMulticall2 *FreeMulticall2CallerSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _FreeMulticall2.Contract.GetCurrentBlockGasLimit(&_FreeMulticall2.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_FreeMulticall2 *FreeMulticall2Caller) GetCurrentBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FreeMulticall2.contract.Call(opts, &out, "getCurrentBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_FreeMulticall2 *FreeMulticall2Session) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _FreeMulticall2.Contract.GetCurrentBlockTimestamp(&_FreeMulticall2.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_FreeMulticall2 *FreeMulticall2CallerSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _FreeMulticall2.Contract.GetCurrentBlockTimestamp(&_FreeMulticall2.CallOpts)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_FreeMulticall2 *FreeMulticall2Caller) GetEthBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FreeMulticall2.contract.Call(opts, &out, "getEthBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_FreeMulticall2 *FreeMulticall2Session) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _FreeMulticall2.Contract.GetEthBalance(&_FreeMulticall2.CallOpts, addr)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_FreeMulticall2 *FreeMulticall2CallerSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _FreeMulticall2.Contract.GetEthBalance(&_FreeMulticall2.CallOpts, addr)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_FreeMulticall2 *FreeMulticall2Caller) GetLastBlockHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FreeMulticall2.contract.Call(opts, &out, "getLastBlockHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_FreeMulticall2 *FreeMulticall2Session) GetLastBlockHash() ([32]byte, error) {
	return _FreeMulticall2.Contract.GetLastBlockHash(&_FreeMulticall2.CallOpts)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_FreeMulticall2 *FreeMulticall2CallerSession) GetLastBlockHash() ([32]byte, error) {
	return _FreeMulticall2.Contract.GetLastBlockHash(&_FreeMulticall2.CallOpts)
}

// TryAggregate is a free data retrieval call binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) view returns((bool,bytes)[] returnData)
func (_FreeMulticall2 *FreeMulticall2Caller) TryAggregate(opts *bind.CallOpts, requireSuccess bool, calls []Multicall2Call) ([]Multicall2Result, error) {
	var out []interface{}
	err := _FreeMulticall2.contract.Call(opts, &out, "tryAggregate", requireSuccess, calls)

	if err != nil {
		return *new([]Multicall2Result), err
	}

	out0 := *abi.ConvertType(out[0], new([]Multicall2Result)).(*[]Multicall2Result)

	return out0, err

}

// TryAggregate is a free data retrieval call binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) view returns((bool,bytes)[] returnData)
func (_FreeMulticall2 *FreeMulticall2Session) TryAggregate(requireSuccess bool, calls []Multicall2Call) ([]Multicall2Result, error) {
	return _FreeMulticall2.Contract.TryAggregate(&_FreeMulticall2.CallOpts, requireSuccess, calls)
}

// TryAggregate is a free data retrieval call binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) view returns((bool,bytes)[] returnData)
func (_FreeMulticall2 *FreeMulticall2CallerSession) TryAggregate(requireSuccess bool, calls []Multicall2Call) ([]Multicall2Result, error) {
	return _FreeMulticall2.Contract.TryAggregate(&_FreeMulticall2.CallOpts, requireSuccess, calls)
}

// TryBlockAndAggregate is a free data retrieval call binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) view returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_FreeMulticall2 *FreeMulticall2Caller) TryBlockAndAggregate(opts *bind.CallOpts, requireSuccess bool, calls []Multicall2Call) (struct {
	BlockNumber *big.Int
	BlockHash   [32]byte
	ReturnData  []Multicall2Result
}, error) {
	var out []interface{}
	err := _FreeMulticall2.contract.Call(opts, &out, "tryBlockAndAggregate", requireSuccess, calls)

	outstruct := new(struct {
		BlockNumber *big.Int
		BlockHash   [32]byte
		ReturnData  []Multicall2Result
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BlockHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.ReturnData = *abi.ConvertType(out[2], new([]Multicall2Result)).(*[]Multicall2Result)

	return *outstruct, err

}

// TryBlockAndAggregate is a free data retrieval call binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) view returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_FreeMulticall2 *FreeMulticall2Session) TryBlockAndAggregate(requireSuccess bool, calls []Multicall2Call) (struct {
	BlockNumber *big.Int
	BlockHash   [32]byte
	ReturnData  []Multicall2Result
}, error) {
	return _FreeMulticall2.Contract.TryBlockAndAggregate(&_FreeMulticall2.CallOpts, requireSuccess, calls)
}

// TryBlockAndAggregate is a free data retrieval call binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) view returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_FreeMulticall2 *FreeMulticall2CallerSession) TryBlockAndAggregate(requireSuccess bool, calls []Multicall2Call) (struct {
	BlockNumber *big.Int
	BlockHash   [32]byte
	ReturnData  []Multicall2Result
}, error) {
	return _FreeMulticall2.Contract.TryBlockAndAggregate(&_FreeMulticall2.CallOpts, requireSuccess, calls)
}
