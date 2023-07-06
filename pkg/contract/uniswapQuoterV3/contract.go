// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapQuoterV3

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

// UniswapQuoterV3MetaData contains all meta data concerning the UniswapQuoterV3 contract.
var UniswapQuoterV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"quoteExactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"name\":\"quoteExactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"quoteExactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"name\":\"quoteExactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UniswapQuoterV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapQuoterV3MetaData.ABI instead.
var UniswapQuoterV3ABI = UniswapQuoterV3MetaData.ABI

// UniswapQuoterV3 is an auto generated Go binding around an Ethereum contract.
type UniswapQuoterV3 struct {
	UniswapQuoterV3Caller     // Read-only binding to the contract
	UniswapQuoterV3Transactor // Write-only binding to the contract
	UniswapQuoterV3Filterer   // Log filterer for contract events
}

// UniswapQuoterV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapQuoterV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapQuoterV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapQuoterV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapQuoterV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapQuoterV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapQuoterV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapQuoterV3Session struct {
	Contract     *UniswapQuoterV3  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapQuoterV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapQuoterV3CallerSession struct {
	Contract *UniswapQuoterV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// UniswapQuoterV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapQuoterV3TransactorSession struct {
	Contract     *UniswapQuoterV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// UniswapQuoterV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapQuoterV3Raw struct {
	Contract *UniswapQuoterV3 // Generic contract binding to access the raw methods on
}

// UniswapQuoterV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapQuoterV3CallerRaw struct {
	Contract *UniswapQuoterV3Caller // Generic read-only contract binding to access the raw methods on
}

// UniswapQuoterV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapQuoterV3TransactorRaw struct {
	Contract *UniswapQuoterV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapQuoterV3 creates a new instance of UniswapQuoterV3, bound to a specific deployed contract.
func NewUniswapQuoterV3(address common.Address, backend bind.ContractBackend) (*UniswapQuoterV3, error) {
	contract, err := bindUniswapQuoterV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapQuoterV3{UniswapQuoterV3Caller: UniswapQuoterV3Caller{contract: contract}, UniswapQuoterV3Transactor: UniswapQuoterV3Transactor{contract: contract}, UniswapQuoterV3Filterer: UniswapQuoterV3Filterer{contract: contract}}, nil
}

// NewUniswapQuoterV3Caller creates a new read-only instance of UniswapQuoterV3, bound to a specific deployed contract.
func NewUniswapQuoterV3Caller(address common.Address, caller bind.ContractCaller) (*UniswapQuoterV3Caller, error) {
	contract, err := bindUniswapQuoterV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapQuoterV3Caller{contract: contract}, nil
}

// NewUniswapQuoterV3Transactor creates a new write-only instance of UniswapQuoterV3, bound to a specific deployed contract.
func NewUniswapQuoterV3Transactor(address common.Address, transactor bind.ContractTransactor) (*UniswapQuoterV3Transactor, error) {
	contract, err := bindUniswapQuoterV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapQuoterV3Transactor{contract: contract}, nil
}

// NewUniswapQuoterV3Filterer creates a new log filterer instance of UniswapQuoterV3, bound to a specific deployed contract.
func NewUniswapQuoterV3Filterer(address common.Address, filterer bind.ContractFilterer) (*UniswapQuoterV3Filterer, error) {
	contract, err := bindUniswapQuoterV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapQuoterV3Filterer{contract: contract}, nil
}

// bindUniswapQuoterV3 binds a generic wrapper to an already deployed contract.
func bindUniswapQuoterV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniswapQuoterV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapQuoterV3 *UniswapQuoterV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapQuoterV3.Contract.UniswapQuoterV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapQuoterV3 *UniswapQuoterV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapQuoterV3.Contract.UniswapQuoterV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapQuoterV3 *UniswapQuoterV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapQuoterV3.Contract.UniswapQuoterV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapQuoterV3 *UniswapQuoterV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapQuoterV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapQuoterV3 *UniswapQuoterV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapQuoterV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapQuoterV3 *UniswapQuoterV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapQuoterV3.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_UniswapQuoterV3 *UniswapQuoterV3Caller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapQuoterV3.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_UniswapQuoterV3 *UniswapQuoterV3Session) WETH9() (common.Address, error) {
	return _UniswapQuoterV3.Contract.WETH9(&_UniswapQuoterV3.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_UniswapQuoterV3 *UniswapQuoterV3CallerSession) WETH9() (common.Address, error) {
	return _UniswapQuoterV3.Contract.WETH9(&_UniswapQuoterV3.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_UniswapQuoterV3 *UniswapQuoterV3Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapQuoterV3.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_UniswapQuoterV3 *UniswapQuoterV3Session) Factory() (common.Address, error) {
	return _UniswapQuoterV3.Contract.Factory(&_UniswapQuoterV3.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_UniswapQuoterV3 *UniswapQuoterV3CallerSession) Factory() (common.Address, error) {
	return _UniswapQuoterV3.Contract.Factory(&_UniswapQuoterV3.CallOpts)
}

// QuoteExactInput is a free data retrieval call binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) view returns(uint256 amountOut)
func (_UniswapQuoterV3 *UniswapQuoterV3Caller) QuoteExactInput(opts *bind.CallOpts, path []byte, amountIn *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UniswapQuoterV3.contract.Call(opts, &out, "quoteExactInput", path, amountIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteExactInput is a free data retrieval call binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) view returns(uint256 amountOut)
func (_UniswapQuoterV3 *UniswapQuoterV3Session) QuoteExactInput(path []byte, amountIn *big.Int) (*big.Int, error) {
	return _UniswapQuoterV3.Contract.QuoteExactInput(&_UniswapQuoterV3.CallOpts, path, amountIn)
}

// QuoteExactInput is a free data retrieval call binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) view returns(uint256 amountOut)
func (_UniswapQuoterV3 *UniswapQuoterV3CallerSession) QuoteExactInput(path []byte, amountIn *big.Int) (*big.Int, error) {
	return _UniswapQuoterV3.Contract.QuoteExactInput(&_UniswapQuoterV3.CallOpts, path, amountIn)
}

// QuoteExactInputSingle is a free data retrieval call binding the contract method 0xf7729d43.
//
// Solidity: function quoteExactInputSingle(address tokenIn, address tokenOut, uint24 fee, uint256 amountIn, uint160 sqrtPriceLimitX96) view returns(uint256 amountOut)
func (_UniswapQuoterV3 *UniswapQuoterV3Caller) QuoteExactInputSingle(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address, fee *big.Int, amountIn *big.Int, sqrtPriceLimitX96 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UniswapQuoterV3.contract.Call(opts, &out, "quoteExactInputSingle", tokenIn, tokenOut, fee, amountIn, sqrtPriceLimitX96)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteExactInputSingle is a free data retrieval call binding the contract method 0xf7729d43.
//
// Solidity: function quoteExactInputSingle(address tokenIn, address tokenOut, uint24 fee, uint256 amountIn, uint160 sqrtPriceLimitX96) view returns(uint256 amountOut)
func (_UniswapQuoterV3 *UniswapQuoterV3Session) QuoteExactInputSingle(tokenIn common.Address, tokenOut common.Address, fee *big.Int, amountIn *big.Int, sqrtPriceLimitX96 *big.Int) (*big.Int, error) {
	return _UniswapQuoterV3.Contract.QuoteExactInputSingle(&_UniswapQuoterV3.CallOpts, tokenIn, tokenOut, fee, amountIn, sqrtPriceLimitX96)
}

// QuoteExactInputSingle is a free data retrieval call binding the contract method 0xf7729d43.
//
// Solidity: function quoteExactInputSingle(address tokenIn, address tokenOut, uint24 fee, uint256 amountIn, uint160 sqrtPriceLimitX96) view returns(uint256 amountOut)
func (_UniswapQuoterV3 *UniswapQuoterV3CallerSession) QuoteExactInputSingle(tokenIn common.Address, tokenOut common.Address, fee *big.Int, amountIn *big.Int, sqrtPriceLimitX96 *big.Int) (*big.Int, error) {
	return _UniswapQuoterV3.Contract.QuoteExactInputSingle(&_UniswapQuoterV3.CallOpts, tokenIn, tokenOut, fee, amountIn, sqrtPriceLimitX96)
}

// QuoteExactOutput is a free data retrieval call binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) view returns(uint256 amountIn)
func (_UniswapQuoterV3 *UniswapQuoterV3Caller) QuoteExactOutput(opts *bind.CallOpts, path []byte, amountOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UniswapQuoterV3.contract.Call(opts, &out, "quoteExactOutput", path, amountOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteExactOutput is a free data retrieval call binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) view returns(uint256 amountIn)
func (_UniswapQuoterV3 *UniswapQuoterV3Session) QuoteExactOutput(path []byte, amountOut *big.Int) (*big.Int, error) {
	return _UniswapQuoterV3.Contract.QuoteExactOutput(&_UniswapQuoterV3.CallOpts, path, amountOut)
}

// QuoteExactOutput is a free data retrieval call binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) view returns(uint256 amountIn)
func (_UniswapQuoterV3 *UniswapQuoterV3CallerSession) QuoteExactOutput(path []byte, amountOut *big.Int) (*big.Int, error) {
	return _UniswapQuoterV3.Contract.QuoteExactOutput(&_UniswapQuoterV3.CallOpts, path, amountOut)
}

// QuoteExactOutputSingle is a free data retrieval call binding the contract method 0x30d07f21.
//
// Solidity: function quoteExactOutputSingle(address tokenIn, address tokenOut, uint24 fee, uint256 amountOut, uint160 sqrtPriceLimitX96) view returns(uint256 amountIn)
func (_UniswapQuoterV3 *UniswapQuoterV3Caller) QuoteExactOutputSingle(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address, fee *big.Int, amountOut *big.Int, sqrtPriceLimitX96 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UniswapQuoterV3.contract.Call(opts, &out, "quoteExactOutputSingle", tokenIn, tokenOut, fee, amountOut, sqrtPriceLimitX96)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteExactOutputSingle is a free data retrieval call binding the contract method 0x30d07f21.
//
// Solidity: function quoteExactOutputSingle(address tokenIn, address tokenOut, uint24 fee, uint256 amountOut, uint160 sqrtPriceLimitX96) view returns(uint256 amountIn)
func (_UniswapQuoterV3 *UniswapQuoterV3Session) QuoteExactOutputSingle(tokenIn common.Address, tokenOut common.Address, fee *big.Int, amountOut *big.Int, sqrtPriceLimitX96 *big.Int) (*big.Int, error) {
	return _UniswapQuoterV3.Contract.QuoteExactOutputSingle(&_UniswapQuoterV3.CallOpts, tokenIn, tokenOut, fee, amountOut, sqrtPriceLimitX96)
}

// QuoteExactOutputSingle is a free data retrieval call binding the contract method 0x30d07f21.
//
// Solidity: function quoteExactOutputSingle(address tokenIn, address tokenOut, uint24 fee, uint256 amountOut, uint160 sqrtPriceLimitX96) view returns(uint256 amountIn)
func (_UniswapQuoterV3 *UniswapQuoterV3CallerSession) QuoteExactOutputSingle(tokenIn common.Address, tokenOut common.Address, fee *big.Int, amountOut *big.Int, sqrtPriceLimitX96 *big.Int) (*big.Int, error) {
	return _UniswapQuoterV3.Contract.QuoteExactOutputSingle(&_UniswapQuoterV3.CallOpts, tokenIn, tokenOut, fee, amountOut, sqrtPriceLimitX96)
}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_UniswapQuoterV3 *UniswapQuoterV3Caller) UniswapV3SwapCallback(opts *bind.CallOpts, amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	var out []interface{}
	err := _UniswapQuoterV3.contract.Call(opts, &out, "uniswapV3SwapCallback", amount0Delta, amount1Delta, path)

	if err != nil {
		return err
	}

	return err

}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_UniswapQuoterV3 *UniswapQuoterV3Session) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	return _UniswapQuoterV3.Contract.UniswapV3SwapCallback(&_UniswapQuoterV3.CallOpts, amount0Delta, amount1Delta, path)
}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_UniswapQuoterV3 *UniswapQuoterV3CallerSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	return _UniswapQuoterV3.Contract.UniswapV3SwapCallback(&_UniswapQuoterV3.CallOpts, amount0Delta, amount1Delta, path)
}
