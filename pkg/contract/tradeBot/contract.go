// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tradeBot

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

// TradeBotTradeLog is an auto generated low-level Go binding around an user-defined struct.
type TradeBotTradeLog struct {
	Id          *big.Int
	PairAddress common.Address
	Token0      common.Address
	Token1      common.Address
	AmountIn    *big.Int
	AmountOut   *big.Int
	Timestamp   *big.Int
	IsBuy       bool
	Price       *big.Int
}

// TradeBotMetaData contains all meta data concerning the TradeBot contract.
var TradeBotMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wbnb\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"busd\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usdt\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usdc\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BUSD\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WBNB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getTradeLogById\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pairAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isBuy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structTradeBot.TradeLog\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTradeLogCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"getTradeLogs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pairAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isBuy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structTradeBot.TradeLog[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintChiToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sendBnbBack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendTokenBack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"sendTokenBackAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"setWhite\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// TradeBotABI is the input ABI used to generate the binding from.
// Deprecated: Use TradeBotMetaData.ABI instead.
var TradeBotABI = TradeBotMetaData.ABI

// TradeBot is an auto generated Go binding around an Ethereum contract.
type TradeBot struct {
	TradeBotCaller     // Read-only binding to the contract
	TradeBotTransactor // Write-only binding to the contract
	TradeBotFilterer   // Log filterer for contract events
}

// TradeBotCaller is an auto generated read-only Go binding around an Ethereum contract.
type TradeBotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeBotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TradeBotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeBotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TradeBotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeBotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TradeBotSession struct {
	Contract     *TradeBot         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradeBotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TradeBotCallerSession struct {
	Contract *TradeBotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TradeBotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TradeBotTransactorSession struct {
	Contract     *TradeBotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TradeBotRaw is an auto generated low-level Go binding around an Ethereum contract.
type TradeBotRaw struct {
	Contract *TradeBot // Generic contract binding to access the raw methods on
}

// TradeBotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TradeBotCallerRaw struct {
	Contract *TradeBotCaller // Generic read-only contract binding to access the raw methods on
}

// TradeBotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TradeBotTransactorRaw struct {
	Contract *TradeBotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTradeBot creates a new instance of TradeBot, bound to a specific deployed contract.
func NewTradeBot(address common.Address, backend bind.ContractBackend) (*TradeBot, error) {
	contract, err := bindTradeBot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TradeBot{TradeBotCaller: TradeBotCaller{contract: contract}, TradeBotTransactor: TradeBotTransactor{contract: contract}, TradeBotFilterer: TradeBotFilterer{contract: contract}}, nil
}

// NewTradeBotCaller creates a new read-only instance of TradeBot, bound to a specific deployed contract.
func NewTradeBotCaller(address common.Address, caller bind.ContractCaller) (*TradeBotCaller, error) {
	contract, err := bindTradeBot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TradeBotCaller{contract: contract}, nil
}

// NewTradeBotTransactor creates a new write-only instance of TradeBot, bound to a specific deployed contract.
func NewTradeBotTransactor(address common.Address, transactor bind.ContractTransactor) (*TradeBotTransactor, error) {
	contract, err := bindTradeBot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TradeBotTransactor{contract: contract}, nil
}

// NewTradeBotFilterer creates a new log filterer instance of TradeBot, bound to a specific deployed contract.
func NewTradeBotFilterer(address common.Address, filterer bind.ContractFilterer) (*TradeBotFilterer, error) {
	contract, err := bindTradeBot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TradeBotFilterer{contract: contract}, nil
}

// bindTradeBot binds a generic wrapper to an already deployed contract.
func bindTradeBot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TradeBotMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TradeBot *TradeBotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TradeBot.Contract.TradeBotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TradeBot *TradeBotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradeBot.Contract.TradeBotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TradeBot *TradeBotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TradeBot.Contract.TradeBotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TradeBot *TradeBotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TradeBot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TradeBot *TradeBotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradeBot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TradeBot *TradeBotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TradeBot.Contract.contract.Transact(opts, method, params...)
}

// BUSD is a free data retrieval call binding the contract method 0x484f4ea9.
//
// Solidity: function BUSD() view returns(address)
func (_TradeBot *TradeBotCaller) BUSD(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TradeBot.contract.Call(opts, &out, "BUSD")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BUSD is a free data retrieval call binding the contract method 0x484f4ea9.
//
// Solidity: function BUSD() view returns(address)
func (_TradeBot *TradeBotSession) BUSD() (common.Address, error) {
	return _TradeBot.Contract.BUSD(&_TradeBot.CallOpts)
}

// BUSD is a free data retrieval call binding the contract method 0x484f4ea9.
//
// Solidity: function BUSD() view returns(address)
func (_TradeBot *TradeBotCallerSession) BUSD() (common.Address, error) {
	return _TradeBot.Contract.BUSD(&_TradeBot.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_TradeBot *TradeBotCaller) USDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TradeBot.contract.Call(opts, &out, "USDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_TradeBot *TradeBotSession) USDC() (common.Address, error) {
	return _TradeBot.Contract.USDC(&_TradeBot.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_TradeBot *TradeBotCallerSession) USDC() (common.Address, error) {
	return _TradeBot.Contract.USDC(&_TradeBot.CallOpts)
}

// USDT is a free data retrieval call binding the contract method 0xc54e44eb.
//
// Solidity: function USDT() view returns(address)
func (_TradeBot *TradeBotCaller) USDT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TradeBot.contract.Call(opts, &out, "USDT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDT is a free data retrieval call binding the contract method 0xc54e44eb.
//
// Solidity: function USDT() view returns(address)
func (_TradeBot *TradeBotSession) USDT() (common.Address, error) {
	return _TradeBot.Contract.USDT(&_TradeBot.CallOpts)
}

// USDT is a free data retrieval call binding the contract method 0xc54e44eb.
//
// Solidity: function USDT() view returns(address)
func (_TradeBot *TradeBotCallerSession) USDT() (common.Address, error) {
	return _TradeBot.Contract.USDT(&_TradeBot.CallOpts)
}

// WBNB is a free data retrieval call binding the contract method 0x8dd95002.
//
// Solidity: function WBNB() view returns(address)
func (_TradeBot *TradeBotCaller) WBNB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TradeBot.contract.Call(opts, &out, "WBNB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WBNB is a free data retrieval call binding the contract method 0x8dd95002.
//
// Solidity: function WBNB() view returns(address)
func (_TradeBot *TradeBotSession) WBNB() (common.Address, error) {
	return _TradeBot.Contract.WBNB(&_TradeBot.CallOpts)
}

// WBNB is a free data retrieval call binding the contract method 0x8dd95002.
//
// Solidity: function WBNB() view returns(address)
func (_TradeBot *TradeBotCallerSession) WBNB() (common.Address, error) {
	return _TradeBot.Contract.WBNB(&_TradeBot.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0x61e0b77f.
//
// Solidity: function getPair(address router, address token0, address token1) view returns(address pair)
func (_TradeBot *TradeBotCaller) GetPair(opts *bind.CallOpts, router common.Address, token0 common.Address, token1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _TradeBot.contract.Call(opts, &out, "getPair", router, token0, token1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0x61e0b77f.
//
// Solidity: function getPair(address router, address token0, address token1) view returns(address pair)
func (_TradeBot *TradeBotSession) GetPair(router common.Address, token0 common.Address, token1 common.Address) (common.Address, error) {
	return _TradeBot.Contract.GetPair(&_TradeBot.CallOpts, router, token0, token1)
}

// GetPair is a free data retrieval call binding the contract method 0x61e0b77f.
//
// Solidity: function getPair(address router, address token0, address token1) view returns(address pair)
func (_TradeBot *TradeBotCallerSession) GetPair(router common.Address, token0 common.Address, token1 common.Address) (common.Address, error) {
	return _TradeBot.Contract.GetPair(&_TradeBot.CallOpts, router, token0, token1)
}

// GetTradeLogById is a free data retrieval call binding the contract method 0x616c121a.
//
// Solidity: function getTradeLogById(uint256 id) view returns((uint256,address,address,address,uint256,uint256,uint256,bool,uint256))
func (_TradeBot *TradeBotCaller) GetTradeLogById(opts *bind.CallOpts, id *big.Int) (TradeBotTradeLog, error) {
	var out []interface{}
	err := _TradeBot.contract.Call(opts, &out, "getTradeLogById", id)

	if err != nil {
		return *new(TradeBotTradeLog), err
	}

	out0 := *abi.ConvertType(out[0], new(TradeBotTradeLog)).(*TradeBotTradeLog)

	return out0, err

}

// GetTradeLogById is a free data retrieval call binding the contract method 0x616c121a.
//
// Solidity: function getTradeLogById(uint256 id) view returns((uint256,address,address,address,uint256,uint256,uint256,bool,uint256))
func (_TradeBot *TradeBotSession) GetTradeLogById(id *big.Int) (TradeBotTradeLog, error) {
	return _TradeBot.Contract.GetTradeLogById(&_TradeBot.CallOpts, id)
}

// GetTradeLogById is a free data retrieval call binding the contract method 0x616c121a.
//
// Solidity: function getTradeLogById(uint256 id) view returns((uint256,address,address,address,uint256,uint256,uint256,bool,uint256))
func (_TradeBot *TradeBotCallerSession) GetTradeLogById(id *big.Int) (TradeBotTradeLog, error) {
	return _TradeBot.Contract.GetTradeLogById(&_TradeBot.CallOpts, id)
}

// GetTradeLogCount is a free data retrieval call binding the contract method 0xa5900788.
//
// Solidity: function getTradeLogCount() view returns(uint256)
func (_TradeBot *TradeBotCaller) GetTradeLogCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TradeBot.contract.Call(opts, &out, "getTradeLogCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTradeLogCount is a free data retrieval call binding the contract method 0xa5900788.
//
// Solidity: function getTradeLogCount() view returns(uint256)
func (_TradeBot *TradeBotSession) GetTradeLogCount() (*big.Int, error) {
	return _TradeBot.Contract.GetTradeLogCount(&_TradeBot.CallOpts)
}

// GetTradeLogCount is a free data retrieval call binding the contract method 0xa5900788.
//
// Solidity: function getTradeLogCount() view returns(uint256)
func (_TradeBot *TradeBotCallerSession) GetTradeLogCount() (*big.Int, error) {
	return _TradeBot.Contract.GetTradeLogCount(&_TradeBot.CallOpts)
}

// GetTradeLogs is a free data retrieval call binding the contract method 0x57037df6.
//
// Solidity: function getTradeLogs(uint256 pageIndex, uint256 pageSize) view returns((uint256,address,address,address,uint256,uint256,uint256,bool,uint256)[])
func (_TradeBot *TradeBotCaller) GetTradeLogs(opts *bind.CallOpts, pageIndex *big.Int, pageSize *big.Int) ([]TradeBotTradeLog, error) {
	var out []interface{}
	err := _TradeBot.contract.Call(opts, &out, "getTradeLogs", pageIndex, pageSize)

	if err != nil {
		return *new([]TradeBotTradeLog), err
	}

	out0 := *abi.ConvertType(out[0], new([]TradeBotTradeLog)).(*[]TradeBotTradeLog)

	return out0, err

}

// GetTradeLogs is a free data retrieval call binding the contract method 0x57037df6.
//
// Solidity: function getTradeLogs(uint256 pageIndex, uint256 pageSize) view returns((uint256,address,address,address,uint256,uint256,uint256,bool,uint256)[])
func (_TradeBot *TradeBotSession) GetTradeLogs(pageIndex *big.Int, pageSize *big.Int) ([]TradeBotTradeLog, error) {
	return _TradeBot.Contract.GetTradeLogs(&_TradeBot.CallOpts, pageIndex, pageSize)
}

// GetTradeLogs is a free data retrieval call binding the contract method 0x57037df6.
//
// Solidity: function getTradeLogs(uint256 pageIndex, uint256 pageSize) view returns((uint256,address,address,address,uint256,uint256,uint256,bool,uint256)[])
func (_TradeBot *TradeBotCallerSession) GetTradeLogs(pageIndex *big.Int, pageSize *big.Int) ([]TradeBotTradeLog, error) {
	return _TradeBot.Contract.GetTradeLogs(&_TradeBot.CallOpts, pageIndex, pageSize)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TradeBot *TradeBotCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TradeBot.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TradeBot *TradeBotSession) Owner() (common.Address, error) {
	return _TradeBot.Contract.Owner(&_TradeBot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TradeBot *TradeBotCallerSession) Owner() (common.Address, error) {
	return _TradeBot.Contract.Owner(&_TradeBot.CallOpts)
}

// MintChiToken is a paid mutator transaction binding the contract method 0x84c469f1.
//
// Solidity: function mintChiToken(uint256 amount) returns()
func (_TradeBot *TradeBotTransactor) MintChiToken(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TradeBot.contract.Transact(opts, "mintChiToken", amount)
}

// MintChiToken is a paid mutator transaction binding the contract method 0x84c469f1.
//
// Solidity: function mintChiToken(uint256 amount) returns()
func (_TradeBot *TradeBotSession) MintChiToken(amount *big.Int) (*types.Transaction, error) {
	return _TradeBot.Contract.MintChiToken(&_TradeBot.TransactOpts, amount)
}

// MintChiToken is a paid mutator transaction binding the contract method 0x84c469f1.
//
// Solidity: function mintChiToken(uint256 amount) returns()
func (_TradeBot *TradeBotTransactorSession) MintChiToken(amount *big.Int) (*types.Transaction, error) {
	return _TradeBot.Contract.MintChiToken(&_TradeBot.TransactOpts, amount)
}

// SendBnbBack is a paid mutator transaction binding the contract method 0xa80db93a.
//
// Solidity: function sendBnbBack() returns()
func (_TradeBot *TradeBotTransactor) SendBnbBack(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradeBot.contract.Transact(opts, "sendBnbBack")
}

// SendBnbBack is a paid mutator transaction binding the contract method 0xa80db93a.
//
// Solidity: function sendBnbBack() returns()
func (_TradeBot *TradeBotSession) SendBnbBack() (*types.Transaction, error) {
	return _TradeBot.Contract.SendBnbBack(&_TradeBot.TransactOpts)
}

// SendBnbBack is a paid mutator transaction binding the contract method 0xa80db93a.
//
// Solidity: function sendBnbBack() returns()
func (_TradeBot *TradeBotTransactorSession) SendBnbBack() (*types.Transaction, error) {
	return _TradeBot.Contract.SendBnbBack(&_TradeBot.TransactOpts)
}

// SendTokenBack is a paid mutator transaction binding the contract method 0x9bbebbc1.
//
// Solidity: function sendTokenBack(address token, uint256 amount) returns()
func (_TradeBot *TradeBotTransactor) SendTokenBack(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TradeBot.contract.Transact(opts, "sendTokenBack", token, amount)
}

// SendTokenBack is a paid mutator transaction binding the contract method 0x9bbebbc1.
//
// Solidity: function sendTokenBack(address token, uint256 amount) returns()
func (_TradeBot *TradeBotSession) SendTokenBack(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TradeBot.Contract.SendTokenBack(&_TradeBot.TransactOpts, token, amount)
}

// SendTokenBack is a paid mutator transaction binding the contract method 0x9bbebbc1.
//
// Solidity: function sendTokenBack(address token, uint256 amount) returns()
func (_TradeBot *TradeBotTransactorSession) SendTokenBack(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TradeBot.Contract.SendTokenBack(&_TradeBot.TransactOpts, token, amount)
}

// SendTokenBackAll is a paid mutator transaction binding the contract method 0xb735b89b.
//
// Solidity: function sendTokenBackAll(address token) returns()
func (_TradeBot *TradeBotTransactor) SendTokenBackAll(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _TradeBot.contract.Transact(opts, "sendTokenBackAll", token)
}

// SendTokenBackAll is a paid mutator transaction binding the contract method 0xb735b89b.
//
// Solidity: function sendTokenBackAll(address token) returns()
func (_TradeBot *TradeBotSession) SendTokenBackAll(token common.Address) (*types.Transaction, error) {
	return _TradeBot.Contract.SendTokenBackAll(&_TradeBot.TransactOpts, token)
}

// SendTokenBackAll is a paid mutator transaction binding the contract method 0xb735b89b.
//
// Solidity: function sendTokenBackAll(address token) returns()
func (_TradeBot *TradeBotTransactorSession) SendTokenBackAll(token common.Address) (*types.Transaction, error) {
	return _TradeBot.Contract.SendTokenBackAll(&_TradeBot.TransactOpts, token)
}

// SetWhite is a paid mutator transaction binding the contract method 0xc03646ba.
//
// Solidity: function setWhite(address account) returns()
func (_TradeBot *TradeBotTransactor) SetWhite(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _TradeBot.contract.Transact(opts, "setWhite", account)
}

// SetWhite is a paid mutator transaction binding the contract method 0xc03646ba.
//
// Solidity: function setWhite(address account) returns()
func (_TradeBot *TradeBotSession) SetWhite(account common.Address) (*types.Transaction, error) {
	return _TradeBot.Contract.SetWhite(&_TradeBot.TransactOpts, account)
}

// SetWhite is a paid mutator transaction binding the contract method 0xc03646ba.
//
// Solidity: function setWhite(address account) returns()
func (_TradeBot *TradeBotTransactorSession) SetWhite(account common.Address) (*types.Transaction, error) {
	return _TradeBot.Contract.SetWhite(&_TradeBot.TransactOpts, account)
}

// Swap is a paid mutator transaction binding the contract method 0xa9678a18.
//
// Solidity: function swap(address router, address token0, address token1, uint256 amount) returns()
func (_TradeBot *TradeBotTransactor) Swap(opts *bind.TransactOpts, router common.Address, token0 common.Address, token1 common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TradeBot.contract.Transact(opts, "swap", router, token0, token1, amount)
}

// Swap is a paid mutator transaction binding the contract method 0xa9678a18.
//
// Solidity: function swap(address router, address token0, address token1, uint256 amount) returns()
func (_TradeBot *TradeBotSession) Swap(router common.Address, token0 common.Address, token1 common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TradeBot.Contract.Swap(&_TradeBot.TransactOpts, router, token0, token1, amount)
}

// Swap is a paid mutator transaction binding the contract method 0xa9678a18.
//
// Solidity: function swap(address router, address token0, address token1, uint256 amount) returns()
func (_TradeBot *TradeBotTransactorSession) Swap(router common.Address, token0 common.Address, token1 common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TradeBot.Contract.Swap(&_TradeBot.TransactOpts, router, token0, token1, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TradeBot *TradeBotTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradeBot.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TradeBot *TradeBotSession) Receive() (*types.Transaction, error) {
	return _TradeBot.Contract.Receive(&_TradeBot.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TradeBot *TradeBotTransactorSession) Receive() (*types.Transaction, error) {
	return _TradeBot.Contract.Receive(&_TradeBot.TransactOpts)
}
