// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pancakeSwapFactory

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

// PancakeSwapFactoryMetaData contains all meta data concerning the PancakeSwapFactory contract.
var PancakeSwapFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"INIT_CODE_PAIR_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PancakeSwapFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use PancakeSwapFactoryMetaData.ABI instead.
var PancakeSwapFactoryABI = PancakeSwapFactoryMetaData.ABI

// PancakeSwapFactory is an auto generated Go binding around an Ethereum contract.
type PancakeSwapFactory struct {
	PancakeSwapFactoryCaller     // Read-only binding to the contract
	PancakeSwapFactoryTransactor // Write-only binding to the contract
	PancakeSwapFactoryFilterer   // Log filterer for contract events
}

// PancakeSwapFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PancakeSwapFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeSwapFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PancakeSwapFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeSwapFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PancakeSwapFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeSwapFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PancakeSwapFactorySession struct {
	Contract     *PancakeSwapFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PancakeSwapFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PancakeSwapFactoryCallerSession struct {
	Contract *PancakeSwapFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// PancakeSwapFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PancakeSwapFactoryTransactorSession struct {
	Contract     *PancakeSwapFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// PancakeSwapFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PancakeSwapFactoryRaw struct {
	Contract *PancakeSwapFactory // Generic contract binding to access the raw methods on
}

// PancakeSwapFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PancakeSwapFactoryCallerRaw struct {
	Contract *PancakeSwapFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// PancakeSwapFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PancakeSwapFactoryTransactorRaw struct {
	Contract *PancakeSwapFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPancakeSwapFactory creates a new instance of PancakeSwapFactory, bound to a specific deployed contract.
func NewPancakeSwapFactory(address common.Address, backend bind.ContractBackend) (*PancakeSwapFactory, error) {
	contract, err := bindPancakeSwapFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PancakeSwapFactory{PancakeSwapFactoryCaller: PancakeSwapFactoryCaller{contract: contract}, PancakeSwapFactoryTransactor: PancakeSwapFactoryTransactor{contract: contract}, PancakeSwapFactoryFilterer: PancakeSwapFactoryFilterer{contract: contract}}, nil
}

// NewPancakeSwapFactoryCaller creates a new read-only instance of PancakeSwapFactory, bound to a specific deployed contract.
func NewPancakeSwapFactoryCaller(address common.Address, caller bind.ContractCaller) (*PancakeSwapFactoryCaller, error) {
	contract, err := bindPancakeSwapFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeSwapFactoryCaller{contract: contract}, nil
}

// NewPancakeSwapFactoryTransactor creates a new write-only instance of PancakeSwapFactory, bound to a specific deployed contract.
func NewPancakeSwapFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*PancakeSwapFactoryTransactor, error) {
	contract, err := bindPancakeSwapFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeSwapFactoryTransactor{contract: contract}, nil
}

// NewPancakeSwapFactoryFilterer creates a new log filterer instance of PancakeSwapFactory, bound to a specific deployed contract.
func NewPancakeSwapFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*PancakeSwapFactoryFilterer, error) {
	contract, err := bindPancakeSwapFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PancakeSwapFactoryFilterer{contract: contract}, nil
}

// bindPancakeSwapFactory binds a generic wrapper to an already deployed contract.
func bindPancakeSwapFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PancakeSwapFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeSwapFactory *PancakeSwapFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeSwapFactory.Contract.PancakeSwapFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeSwapFactory *PancakeSwapFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeSwapFactory.Contract.PancakeSwapFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeSwapFactory *PancakeSwapFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeSwapFactory.Contract.PancakeSwapFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeSwapFactory *PancakeSwapFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeSwapFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeSwapFactory *PancakeSwapFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeSwapFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeSwapFactory *PancakeSwapFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeSwapFactory.Contract.contract.Transact(opts, method, params...)
}

// INITCODEPAIRHASH is a free data retrieval call binding the contract method 0x5855a25a.
//
// Solidity: function INIT_CODE_PAIR_HASH() view returns(bytes32)
func (_PancakeSwapFactory *PancakeSwapFactoryCaller) INITCODEPAIRHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PancakeSwapFactory.contract.Call(opts, &out, "INIT_CODE_PAIR_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// INITCODEPAIRHASH is a free data retrieval call binding the contract method 0x5855a25a.
//
// Solidity: function INIT_CODE_PAIR_HASH() view returns(bytes32)
func (_PancakeSwapFactory *PancakeSwapFactorySession) INITCODEPAIRHASH() ([32]byte, error) {
	return _PancakeSwapFactory.Contract.INITCODEPAIRHASH(&_PancakeSwapFactory.CallOpts)
}

// INITCODEPAIRHASH is a free data retrieval call binding the contract method 0x5855a25a.
//
// Solidity: function INIT_CODE_PAIR_HASH() view returns(bytes32)
func (_PancakeSwapFactory *PancakeSwapFactoryCallerSession) INITCODEPAIRHASH() ([32]byte, error) {
	return _PancakeSwapFactory.Contract.INITCODEPAIRHASH(&_PancakeSwapFactory.CallOpts)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_PancakeSwapFactory *PancakeSwapFactoryCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PancakeSwapFactory.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_PancakeSwapFactory *PancakeSwapFactorySession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _PancakeSwapFactory.Contract.AllPairs(&_PancakeSwapFactory.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_PancakeSwapFactory *PancakeSwapFactoryCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _PancakeSwapFactory.Contract.AllPairs(&_PancakeSwapFactory.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_PancakeSwapFactory *PancakeSwapFactoryCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeSwapFactory.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_PancakeSwapFactory *PancakeSwapFactorySession) AllPairsLength() (*big.Int, error) {
	return _PancakeSwapFactory.Contract.AllPairsLength(&_PancakeSwapFactory.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_PancakeSwapFactory *PancakeSwapFactoryCallerSession) AllPairsLength() (*big.Int, error) {
	return _PancakeSwapFactory.Contract.AllPairsLength(&_PancakeSwapFactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_PancakeSwapFactory *PancakeSwapFactoryCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeSwapFactory.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_PancakeSwapFactory *PancakeSwapFactorySession) FeeTo() (common.Address, error) {
	return _PancakeSwapFactory.Contract.FeeTo(&_PancakeSwapFactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_PancakeSwapFactory *PancakeSwapFactoryCallerSession) FeeTo() (common.Address, error) {
	return _PancakeSwapFactory.Contract.FeeTo(&_PancakeSwapFactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_PancakeSwapFactory *PancakeSwapFactoryCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeSwapFactory.contract.Call(opts, &out, "feeToSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_PancakeSwapFactory *PancakeSwapFactorySession) FeeToSetter() (common.Address, error) {
	return _PancakeSwapFactory.Contract.FeeToSetter(&_PancakeSwapFactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_PancakeSwapFactory *PancakeSwapFactoryCallerSession) FeeToSetter() (common.Address, error) {
	return _PancakeSwapFactory.Contract.FeeToSetter(&_PancakeSwapFactory.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_PancakeSwapFactory *PancakeSwapFactoryCaller) GetPair(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _PancakeSwapFactory.contract.Call(opts, &out, "getPair", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_PancakeSwapFactory *PancakeSwapFactorySession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _PancakeSwapFactory.Contract.GetPair(&_PancakeSwapFactory.CallOpts, arg0, arg1)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_PancakeSwapFactory *PancakeSwapFactoryCallerSession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _PancakeSwapFactory.Contract.GetPair(&_PancakeSwapFactory.CallOpts, arg0, arg1)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_PancakeSwapFactory *PancakeSwapFactoryTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _PancakeSwapFactory.contract.Transact(opts, "createPair", tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_PancakeSwapFactory *PancakeSwapFactorySession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _PancakeSwapFactory.Contract.CreatePair(&_PancakeSwapFactory.TransactOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_PancakeSwapFactory *PancakeSwapFactoryTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _PancakeSwapFactory.Contract.CreatePair(&_PancakeSwapFactory.TransactOpts, tokenA, tokenB)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_PancakeSwapFactory *PancakeSwapFactoryTransactor) SetFeeTo(opts *bind.TransactOpts, _feeTo common.Address) (*types.Transaction, error) {
	return _PancakeSwapFactory.contract.Transact(opts, "setFeeTo", _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_PancakeSwapFactory *PancakeSwapFactorySession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _PancakeSwapFactory.Contract.SetFeeTo(&_PancakeSwapFactory.TransactOpts, _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_PancakeSwapFactory *PancakeSwapFactoryTransactorSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _PancakeSwapFactory.Contract.SetFeeTo(&_PancakeSwapFactory.TransactOpts, _feeTo)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_PancakeSwapFactory *PancakeSwapFactoryTransactor) SetFeeToSetter(opts *bind.TransactOpts, _feeToSetter common.Address) (*types.Transaction, error) {
	return _PancakeSwapFactory.contract.Transact(opts, "setFeeToSetter", _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_PancakeSwapFactory *PancakeSwapFactorySession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _PancakeSwapFactory.Contract.SetFeeToSetter(&_PancakeSwapFactory.TransactOpts, _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_PancakeSwapFactory *PancakeSwapFactoryTransactorSession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _PancakeSwapFactory.Contract.SetFeeToSetter(&_PancakeSwapFactory.TransactOpts, _feeToSetter)
}

// PancakeSwapFactoryPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the PancakeSwapFactory contract.
type PancakeSwapFactoryPairCreatedIterator struct {
	Event *PancakeSwapFactoryPairCreated // Event containing the contract specifics and raw log

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
func (it *PancakeSwapFactoryPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeSwapFactoryPairCreated)
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
		it.Event = new(PancakeSwapFactoryPairCreated)
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
func (it *PancakeSwapFactoryPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeSwapFactoryPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeSwapFactoryPairCreated represents a PairCreated event raised by the PancakeSwapFactory contract.
type PancakeSwapFactoryPairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_PancakeSwapFactory *PancakeSwapFactoryFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*PancakeSwapFactoryPairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _PancakeSwapFactory.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &PancakeSwapFactoryPairCreatedIterator{contract: _PancakeSwapFactory.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_PancakeSwapFactory *PancakeSwapFactoryFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *PancakeSwapFactoryPairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _PancakeSwapFactory.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeSwapFactoryPairCreated)
				if err := _PancakeSwapFactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_PancakeSwapFactory *PancakeSwapFactoryFilterer) ParsePairCreated(log types.Log) (*PancakeSwapFactoryPairCreated, error) {
	event := new(PancakeSwapFactoryPairCreated)
	if err := _PancakeSwapFactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
