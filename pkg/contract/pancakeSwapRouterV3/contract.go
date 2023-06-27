// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pancakeSwapRouterV3

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

// IApproveAndCallIncreaseLiquidityParams is an auto generated low-level Go binding around an user-defined struct.
type IApproveAndCallIncreaseLiquidityParams struct {
	Token0     common.Address
	Token1     common.Address
	TokenId    *big.Int
	Amount0Min *big.Int
	Amount1Min *big.Int
}

// IApproveAndCallMintParams is an auto generated low-level Go binding around an user-defined struct.
type IApproveAndCallMintParams struct {
	Token0     common.Address
	Token1     common.Address
	Fee        *big.Int
	TickLower  *big.Int
	TickUpper  *big.Int
	Amount0Min *big.Int
	Amount1Min *big.Int
	Recipient  common.Address
}

// IV3SwapRouterExactInputParams is an auto generated low-level Go binding around an user-defined struct.
type IV3SwapRouterExactInputParams struct {
	Path             []byte
	Recipient        common.Address
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

// IV3SwapRouterExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IV3SwapRouterExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitX96 *big.Int
}

// IV3SwapRouterExactOutputParams is an auto generated low-level Go binding around an user-defined struct.
type IV3SwapRouterExactOutputParams struct {
	Path            []byte
	Recipient       common.Address
	AmountOut       *big.Int
	AmountInMaximum *big.Int
}

// IV3SwapRouterExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IV3SwapRouterExactOutputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	AmountOut         *big.Int
	AmountInMaximum   *big.Int
	SqrtPriceLimitX96 *big.Int
}

// PancakeSwapRouterV3MetaData contains all meta data concerning the PancakeSwapRouterV3 contract.
var PancakeSwapRouterV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factoryV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_deployer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factoryV3\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_positionManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stableFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stableInfo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"info\",\"type\":\"address\"}],\"name\":\"SetStableSwap\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"approveMax\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"approveMaxMinusOne\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"approveZeroThenMax\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"approveZeroThenMaxMinusOne\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"callPositionManager\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"paths\",\"type\":\"bytes[]\"},{\"internalType\":\"uint128[]\",\"name\":\"amounts\",\"type\":\"uint128[]\"},{\"internalType\":\"uint24\",\"name\":\"maximumTickDivergence\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"secondsAgo\",\"type\":\"uint32\"}],\"name\":\"checkOracleSlippage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint24\",\"name\":\"maximumTickDivergence\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"secondsAgo\",\"type\":\"uint32\"}],\"name\":\"checkOracleSlippage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structIV3SwapRouter.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIV3SwapRouter.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"flag\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"exactInputStableSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"}],\"internalType\":\"structIV3SwapRouter.ExactOutputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIV3SwapRouter.ExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"flag\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"exactOutputStableSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getApprovalType\",\"outputs\":[{\"internalType\":\"enumIApproveAndCall.ApprovalType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"}],\"internalType\":\"structIApproveAndCall.IncreaseLiquidityParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"increaseLiquidity\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structIApproveAndCall.MintParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"previousBlockhash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"pancakeV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"positionManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"pull\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowed\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowedIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_info\",\"type\":\"address\"}],\"name\":\"setStableSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stableSwapFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stableSwapInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"sweepTokenWithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"sweepTokenWithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9WithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9WithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"wrapETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PancakeSwapRouterV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use PancakeSwapRouterV3MetaData.ABI instead.
var PancakeSwapRouterV3ABI = PancakeSwapRouterV3MetaData.ABI

// PancakeSwapRouterV3 is an auto generated Go binding around an Ethereum contract.
type PancakeSwapRouterV3 struct {
	PancakeSwapRouterV3Caller     // Read-only binding to the contract
	PancakeSwapRouterV3Transactor // Write-only binding to the contract
	PancakeSwapRouterV3Filterer   // Log filterer for contract events
}

// PancakeSwapRouterV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type PancakeSwapRouterV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeSwapRouterV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PancakeSwapRouterV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeSwapRouterV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PancakeSwapRouterV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeSwapRouterV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PancakeSwapRouterV3Session struct {
	Contract     *PancakeSwapRouterV3 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PancakeSwapRouterV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PancakeSwapRouterV3CallerSession struct {
	Contract *PancakeSwapRouterV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// PancakeSwapRouterV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PancakeSwapRouterV3TransactorSession struct {
	Contract     *PancakeSwapRouterV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PancakeSwapRouterV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type PancakeSwapRouterV3Raw struct {
	Contract *PancakeSwapRouterV3 // Generic contract binding to access the raw methods on
}

// PancakeSwapRouterV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PancakeSwapRouterV3CallerRaw struct {
	Contract *PancakeSwapRouterV3Caller // Generic read-only contract binding to access the raw methods on
}

// PancakeSwapRouterV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PancakeSwapRouterV3TransactorRaw struct {
	Contract *PancakeSwapRouterV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPancakeSwapRouterV3 creates a new instance of PancakeSwapRouterV3, bound to a specific deployed contract.
func NewPancakeSwapRouterV3(address common.Address, backend bind.ContractBackend) (*PancakeSwapRouterV3, error) {
	contract, err := bindPancakeSwapRouterV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PancakeSwapRouterV3{PancakeSwapRouterV3Caller: PancakeSwapRouterV3Caller{contract: contract}, PancakeSwapRouterV3Transactor: PancakeSwapRouterV3Transactor{contract: contract}, PancakeSwapRouterV3Filterer: PancakeSwapRouterV3Filterer{contract: contract}}, nil
}

// NewPancakeSwapRouterV3Caller creates a new read-only instance of PancakeSwapRouterV3, bound to a specific deployed contract.
func NewPancakeSwapRouterV3Caller(address common.Address, caller bind.ContractCaller) (*PancakeSwapRouterV3Caller, error) {
	contract, err := bindPancakeSwapRouterV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeSwapRouterV3Caller{contract: contract}, nil
}

// NewPancakeSwapRouterV3Transactor creates a new write-only instance of PancakeSwapRouterV3, bound to a specific deployed contract.
func NewPancakeSwapRouterV3Transactor(address common.Address, transactor bind.ContractTransactor) (*PancakeSwapRouterV3Transactor, error) {
	contract, err := bindPancakeSwapRouterV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeSwapRouterV3Transactor{contract: contract}, nil
}

// NewPancakeSwapRouterV3Filterer creates a new log filterer instance of PancakeSwapRouterV3, bound to a specific deployed contract.
func NewPancakeSwapRouterV3Filterer(address common.Address, filterer bind.ContractFilterer) (*PancakeSwapRouterV3Filterer, error) {
	contract, err := bindPancakeSwapRouterV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PancakeSwapRouterV3Filterer{contract: contract}, nil
}

// bindPancakeSwapRouterV3 binds a generic wrapper to an already deployed contract.
func bindPancakeSwapRouterV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PancakeSwapRouterV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeSwapRouterV3.Contract.PancakeSwapRouterV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.PancakeSwapRouterV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.PancakeSwapRouterV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeSwapRouterV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Caller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeSwapRouterV3.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) WETH9() (common.Address, error) {
	return _PancakeSwapRouterV3.Contract.WETH9(&_PancakeSwapRouterV3.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3CallerSession) WETH9() (common.Address, error) {
	return _PancakeSwapRouterV3.Contract.WETH9(&_PancakeSwapRouterV3.CallOpts)
}

// CheckOracleSlippage is a free data retrieval call binding the contract method 0xefdeed8e.
//
// Solidity: function checkOracleSlippage(bytes[] paths, uint128[] amounts, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Caller) CheckOracleSlippage(opts *bind.CallOpts, paths [][]byte, amounts []*big.Int, maximumTickDivergence *big.Int, secondsAgo uint32) error {
	var out []interface{}
	err := _PancakeSwapRouterV3.contract.Call(opts, &out, "checkOracleSlippage", paths, amounts, maximumTickDivergence, secondsAgo)

	if err != nil {
		return err
	}

	return err

}

// CheckOracleSlippage is a free data retrieval call binding the contract method 0xefdeed8e.
//
// Solidity: function checkOracleSlippage(bytes[] paths, uint128[] amounts, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) CheckOracleSlippage(paths [][]byte, amounts []*big.Int, maximumTickDivergence *big.Int, secondsAgo uint32) error {
	return _PancakeSwapRouterV3.Contract.CheckOracleSlippage(&_PancakeSwapRouterV3.CallOpts, paths, amounts, maximumTickDivergence, secondsAgo)
}

// CheckOracleSlippage is a free data retrieval call binding the contract method 0xefdeed8e.
//
// Solidity: function checkOracleSlippage(bytes[] paths, uint128[] amounts, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3CallerSession) CheckOracleSlippage(paths [][]byte, amounts []*big.Int, maximumTickDivergence *big.Int, secondsAgo uint32) error {
	return _PancakeSwapRouterV3.Contract.CheckOracleSlippage(&_PancakeSwapRouterV3.CallOpts, paths, amounts, maximumTickDivergence, secondsAgo)
}

// CheckOracleSlippage0 is a free data retrieval call binding the contract method 0xf25801a7.
//
// Solidity: function checkOracleSlippage(bytes path, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Caller) CheckOracleSlippage0(opts *bind.CallOpts, path []byte, maximumTickDivergence *big.Int, secondsAgo uint32) error {
	var out []interface{}
	err := _PancakeSwapRouterV3.contract.Call(opts, &out, "checkOracleSlippage0", path, maximumTickDivergence, secondsAgo)

	if err != nil {
		return err
	}

	return err

}

// CheckOracleSlippage0 is a free data retrieval call binding the contract method 0xf25801a7.
//
// Solidity: function checkOracleSlippage(bytes path, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) CheckOracleSlippage0(path []byte, maximumTickDivergence *big.Int, secondsAgo uint32) error {
	return _PancakeSwapRouterV3.Contract.CheckOracleSlippage0(&_PancakeSwapRouterV3.CallOpts, path, maximumTickDivergence, secondsAgo)
}

// CheckOracleSlippage0 is a free data retrieval call binding the contract method 0xf25801a7.
//
// Solidity: function checkOracleSlippage(bytes path, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3CallerSession) CheckOracleSlippage0(path []byte, maximumTickDivergence *big.Int, secondsAgo uint32) error {
	return _PancakeSwapRouterV3.Contract.CheckOracleSlippage0(&_PancakeSwapRouterV3.CallOpts, path, maximumTickDivergence, secondsAgo)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Caller) Deployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeSwapRouterV3.contract.Call(opts, &out, "deployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) Deployer() (common.Address, error) {
	return _PancakeSwapRouterV3.Contract.Deployer(&_PancakeSwapRouterV3.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3CallerSession) Deployer() (common.Address, error) {
	return _PancakeSwapRouterV3.Contract.Deployer(&_PancakeSwapRouterV3.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeSwapRouterV3.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) Factory() (common.Address, error) {
	return _PancakeSwapRouterV3.Contract.Factory(&_PancakeSwapRouterV3.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3CallerSession) Factory() (common.Address, error) {
	return _PancakeSwapRouterV3.Contract.Factory(&_PancakeSwapRouterV3.CallOpts)
}

// FactoryV2 is a free data retrieval call binding the contract method 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Caller) FactoryV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeSwapRouterV3.contract.Call(opts, &out, "factoryV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryV2 is a free data retrieval call binding the contract method 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) FactoryV2() (common.Address, error) {
	return _PancakeSwapRouterV3.Contract.FactoryV2(&_PancakeSwapRouterV3.CallOpts)
}

// FactoryV2 is a free data retrieval call binding the contract method 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3CallerSession) FactoryV2() (common.Address, error) {
	return _PancakeSwapRouterV3.Contract.FactoryV2(&_PancakeSwapRouterV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeSwapRouterV3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) Owner() (common.Address, error) {
	return _PancakeSwapRouterV3.Contract.Owner(&_PancakeSwapRouterV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3CallerSession) Owner() (common.Address, error) {
	return _PancakeSwapRouterV3.Contract.Owner(&_PancakeSwapRouterV3.CallOpts)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Caller) PositionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeSwapRouterV3.contract.Call(opts, &out, "positionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) PositionManager() (common.Address, error) {
	return _PancakeSwapRouterV3.Contract.PositionManager(&_PancakeSwapRouterV3.CallOpts)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3CallerSession) PositionManager() (common.Address, error) {
	return _PancakeSwapRouterV3.Contract.PositionManager(&_PancakeSwapRouterV3.CallOpts)
}

// StableSwapFactory is a free data retrieval call binding the contract method 0x57c79961.
//
// Solidity: function stableSwapFactory() view returns(address)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Caller) StableSwapFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeSwapRouterV3.contract.Call(opts, &out, "stableSwapFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StableSwapFactory is a free data retrieval call binding the contract method 0x57c79961.
//
// Solidity: function stableSwapFactory() view returns(address)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) StableSwapFactory() (common.Address, error) {
	return _PancakeSwapRouterV3.Contract.StableSwapFactory(&_PancakeSwapRouterV3.CallOpts)
}

// StableSwapFactory is a free data retrieval call binding the contract method 0x57c79961.
//
// Solidity: function stableSwapFactory() view returns(address)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3CallerSession) StableSwapFactory() (common.Address, error) {
	return _PancakeSwapRouterV3.Contract.StableSwapFactory(&_PancakeSwapRouterV3.CallOpts)
}

// StableSwapInfo is a free data retrieval call binding the contract method 0xb85aa7af.
//
// Solidity: function stableSwapInfo() view returns(address)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Caller) StableSwapInfo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeSwapRouterV3.contract.Call(opts, &out, "stableSwapInfo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StableSwapInfo is a free data retrieval call binding the contract method 0xb85aa7af.
//
// Solidity: function stableSwapInfo() view returns(address)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) StableSwapInfo() (common.Address, error) {
	return _PancakeSwapRouterV3.Contract.StableSwapInfo(&_PancakeSwapRouterV3.CallOpts)
}

// StableSwapInfo is a free data retrieval call binding the contract method 0xb85aa7af.
//
// Solidity: function stableSwapInfo() view returns(address)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3CallerSession) StableSwapInfo() (common.Address, error) {
	return _PancakeSwapRouterV3.Contract.StableSwapInfo(&_PancakeSwapRouterV3.CallOpts)
}

// ApproveMax is a paid mutator transaction binding the contract method 0x571ac8b0.
//
// Solidity: function approveMax(address token) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) ApproveMax(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "approveMax", token)
}

// ApproveMax is a paid mutator transaction binding the contract method 0x571ac8b0.
//
// Solidity: function approveMax(address token) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) ApproveMax(token common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.ApproveMax(&_PancakeSwapRouterV3.TransactOpts, token)
}

// ApproveMax is a paid mutator transaction binding the contract method 0x571ac8b0.
//
// Solidity: function approveMax(address token) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) ApproveMax(token common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.ApproveMax(&_PancakeSwapRouterV3.TransactOpts, token)
}

// ApproveMaxMinusOne is a paid mutator transaction binding the contract method 0xcab372ce.
//
// Solidity: function approveMaxMinusOne(address token) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) ApproveMaxMinusOne(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "approveMaxMinusOne", token)
}

// ApproveMaxMinusOne is a paid mutator transaction binding the contract method 0xcab372ce.
//
// Solidity: function approveMaxMinusOne(address token) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) ApproveMaxMinusOne(token common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.ApproveMaxMinusOne(&_PancakeSwapRouterV3.TransactOpts, token)
}

// ApproveMaxMinusOne is a paid mutator transaction binding the contract method 0xcab372ce.
//
// Solidity: function approveMaxMinusOne(address token) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) ApproveMaxMinusOne(token common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.ApproveMaxMinusOne(&_PancakeSwapRouterV3.TransactOpts, token)
}

// ApproveZeroThenMax is a paid mutator transaction binding the contract method 0x639d71a9.
//
// Solidity: function approveZeroThenMax(address token) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) ApproveZeroThenMax(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "approveZeroThenMax", token)
}

// ApproveZeroThenMax is a paid mutator transaction binding the contract method 0x639d71a9.
//
// Solidity: function approveZeroThenMax(address token) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) ApproveZeroThenMax(token common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.ApproveZeroThenMax(&_PancakeSwapRouterV3.TransactOpts, token)
}

// ApproveZeroThenMax is a paid mutator transaction binding the contract method 0x639d71a9.
//
// Solidity: function approveZeroThenMax(address token) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) ApproveZeroThenMax(token common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.ApproveZeroThenMax(&_PancakeSwapRouterV3.TransactOpts, token)
}

// ApproveZeroThenMaxMinusOne is a paid mutator transaction binding the contract method 0xab3fdd50.
//
// Solidity: function approveZeroThenMaxMinusOne(address token) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) ApproveZeroThenMaxMinusOne(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "approveZeroThenMaxMinusOne", token)
}

// ApproveZeroThenMaxMinusOne is a paid mutator transaction binding the contract method 0xab3fdd50.
//
// Solidity: function approveZeroThenMaxMinusOne(address token) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) ApproveZeroThenMaxMinusOne(token common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.ApproveZeroThenMaxMinusOne(&_PancakeSwapRouterV3.TransactOpts, token)
}

// ApproveZeroThenMaxMinusOne is a paid mutator transaction binding the contract method 0xab3fdd50.
//
// Solidity: function approveZeroThenMaxMinusOne(address token) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) ApproveZeroThenMaxMinusOne(token common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.ApproveZeroThenMaxMinusOne(&_PancakeSwapRouterV3.TransactOpts, token)
}

// CallPositionManager is a paid mutator transaction binding the contract method 0xb3a2af13.
//
// Solidity: function callPositionManager(bytes data) payable returns(bytes result)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) CallPositionManager(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "callPositionManager", data)
}

// CallPositionManager is a paid mutator transaction binding the contract method 0xb3a2af13.
//
// Solidity: function callPositionManager(bytes data) payable returns(bytes result)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) CallPositionManager(data []byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.CallPositionManager(&_PancakeSwapRouterV3.TransactOpts, data)
}

// CallPositionManager is a paid mutator transaction binding the contract method 0xb3a2af13.
//
// Solidity: function callPositionManager(bytes data) payable returns(bytes result)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) CallPositionManager(data []byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.CallPositionManager(&_PancakeSwapRouterV3.TransactOpts, data)
}

// ExactInput is a paid mutator transaction binding the contract method 0xb858183f.
//
// Solidity: function exactInput((bytes,address,uint256,uint256) params) payable returns(uint256 amountOut)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) ExactInput(opts *bind.TransactOpts, params IV3SwapRouterExactInputParams) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "exactInput", params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xb858183f.
//
// Solidity: function exactInput((bytes,address,uint256,uint256) params) payable returns(uint256 amountOut)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) ExactInput(params IV3SwapRouterExactInputParams) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.ExactInput(&_PancakeSwapRouterV3.TransactOpts, params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xb858183f.
//
// Solidity: function exactInput((bytes,address,uint256,uint256) params) payable returns(uint256 amountOut)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) ExactInput(params IV3SwapRouterExactInputParams) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.ExactInput(&_PancakeSwapRouterV3.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x04e45aaf.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) ExactInputSingle(opts *bind.TransactOpts, params IV3SwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "exactInputSingle", params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x04e45aaf.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) ExactInputSingle(params IV3SwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.ExactInputSingle(&_PancakeSwapRouterV3.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x04e45aaf.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) ExactInputSingle(params IV3SwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.ExactInputSingle(&_PancakeSwapRouterV3.TransactOpts, params)
}

// ExactInputStableSwap is a paid mutator transaction binding the contract method 0xb4554231.
//
// Solidity: function exactInputStableSwap(address[] path, uint256[] flag, uint256 amountIn, uint256 amountOutMin, address to) payable returns(uint256 amountOut)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) ExactInputStableSwap(opts *bind.TransactOpts, path []common.Address, flag []*big.Int, amountIn *big.Int, amountOutMin *big.Int, to common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "exactInputStableSwap", path, flag, amountIn, amountOutMin, to)
}

// ExactInputStableSwap is a paid mutator transaction binding the contract method 0xb4554231.
//
// Solidity: function exactInputStableSwap(address[] path, uint256[] flag, uint256 amountIn, uint256 amountOutMin, address to) payable returns(uint256 amountOut)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) ExactInputStableSwap(path []common.Address, flag []*big.Int, amountIn *big.Int, amountOutMin *big.Int, to common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.ExactInputStableSwap(&_PancakeSwapRouterV3.TransactOpts, path, flag, amountIn, amountOutMin, to)
}

// ExactInputStableSwap is a paid mutator transaction binding the contract method 0xb4554231.
//
// Solidity: function exactInputStableSwap(address[] path, uint256[] flag, uint256 amountIn, uint256 amountOutMin, address to) payable returns(uint256 amountOut)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) ExactInputStableSwap(path []common.Address, flag []*big.Int, amountIn *big.Int, amountOutMin *big.Int, to common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.ExactInputStableSwap(&_PancakeSwapRouterV3.TransactOpts, path, flag, amountIn, amountOutMin, to)
}

// ExactOutput is a paid mutator transaction binding the contract method 0x09b81346.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256) params) payable returns(uint256 amountIn)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) ExactOutput(opts *bind.TransactOpts, params IV3SwapRouterExactOutputParams) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "exactOutput", params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0x09b81346.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256) params) payable returns(uint256 amountIn)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) ExactOutput(params IV3SwapRouterExactOutputParams) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.ExactOutput(&_PancakeSwapRouterV3.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0x09b81346.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256) params) payable returns(uint256 amountIn)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) ExactOutput(params IV3SwapRouterExactOutputParams) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.ExactOutput(&_PancakeSwapRouterV3.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0x5023b4df.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) ExactOutputSingle(opts *bind.TransactOpts, params IV3SwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "exactOutputSingle", params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0x5023b4df.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) ExactOutputSingle(params IV3SwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.ExactOutputSingle(&_PancakeSwapRouterV3.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0x5023b4df.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) ExactOutputSingle(params IV3SwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.ExactOutputSingle(&_PancakeSwapRouterV3.TransactOpts, params)
}

// ExactOutputStableSwap is a paid mutator transaction binding the contract method 0xb4c4e555.
//
// Solidity: function exactOutputStableSwap(address[] path, uint256[] flag, uint256 amountOut, uint256 amountInMax, address to) payable returns(uint256 amountIn)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) ExactOutputStableSwap(opts *bind.TransactOpts, path []common.Address, flag []*big.Int, amountOut *big.Int, amountInMax *big.Int, to common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "exactOutputStableSwap", path, flag, amountOut, amountInMax, to)
}

// ExactOutputStableSwap is a paid mutator transaction binding the contract method 0xb4c4e555.
//
// Solidity: function exactOutputStableSwap(address[] path, uint256[] flag, uint256 amountOut, uint256 amountInMax, address to) payable returns(uint256 amountIn)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) ExactOutputStableSwap(path []common.Address, flag []*big.Int, amountOut *big.Int, amountInMax *big.Int, to common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.ExactOutputStableSwap(&_PancakeSwapRouterV3.TransactOpts, path, flag, amountOut, amountInMax, to)
}

// ExactOutputStableSwap is a paid mutator transaction binding the contract method 0xb4c4e555.
//
// Solidity: function exactOutputStableSwap(address[] path, uint256[] flag, uint256 amountOut, uint256 amountInMax, address to) payable returns(uint256 amountIn)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) ExactOutputStableSwap(path []common.Address, flag []*big.Int, amountOut *big.Int, amountInMax *big.Int, to common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.ExactOutputStableSwap(&_PancakeSwapRouterV3.TransactOpts, path, flag, amountOut, amountInMax, to)
}

// GetApprovalType is a paid mutator transaction binding the contract method 0xdee00f35.
//
// Solidity: function getApprovalType(address token, uint256 amount) returns(uint8)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) GetApprovalType(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "getApprovalType", token, amount)
}

// GetApprovalType is a paid mutator transaction binding the contract method 0xdee00f35.
//
// Solidity: function getApprovalType(address token, uint256 amount) returns(uint8)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) GetApprovalType(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.GetApprovalType(&_PancakeSwapRouterV3.TransactOpts, token, amount)
}

// GetApprovalType is a paid mutator transaction binding the contract method 0xdee00f35.
//
// Solidity: function getApprovalType(address token, uint256 amount) returns(uint8)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) GetApprovalType(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.GetApprovalType(&_PancakeSwapRouterV3.TransactOpts, token, amount)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0xf100b205.
//
// Solidity: function increaseLiquidity((address,address,uint256,uint256,uint256) params) payable returns(bytes result)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) IncreaseLiquidity(opts *bind.TransactOpts, params IApproveAndCallIncreaseLiquidityParams) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "increaseLiquidity", params)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0xf100b205.
//
// Solidity: function increaseLiquidity((address,address,uint256,uint256,uint256) params) payable returns(bytes result)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) IncreaseLiquidity(params IApproveAndCallIncreaseLiquidityParams) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.IncreaseLiquidity(&_PancakeSwapRouterV3.TransactOpts, params)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0xf100b205.
//
// Solidity: function increaseLiquidity((address,address,uint256,uint256,uint256) params) payable returns(bytes result)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) IncreaseLiquidity(params IApproveAndCallIncreaseLiquidityParams) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.IncreaseLiquidity(&_PancakeSwapRouterV3.TransactOpts, params)
}

// Mint is a paid mutator transaction binding the contract method 0x11ed56c9.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,address) params) payable returns(bytes result)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) Mint(opts *bind.TransactOpts, params IApproveAndCallMintParams) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "mint", params)
}

// Mint is a paid mutator transaction binding the contract method 0x11ed56c9.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,address) params) payable returns(bytes result)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) Mint(params IApproveAndCallMintParams) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.Mint(&_PancakeSwapRouterV3.TransactOpts, params)
}

// Mint is a paid mutator transaction binding the contract method 0x11ed56c9.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,address) params) payable returns(bytes result)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) Mint(params IApproveAndCallMintParams) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.Mint(&_PancakeSwapRouterV3.TransactOpts, params)
}

// Multicall is a paid mutator transaction binding the contract method 0x1f0464d1.
//
// Solidity: function multicall(bytes32 previousBlockhash, bytes[] data) payable returns(bytes[])
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) Multicall(opts *bind.TransactOpts, previousBlockhash [32]byte, data [][]byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "multicall", previousBlockhash, data)
}

// Multicall is a paid mutator transaction binding the contract method 0x1f0464d1.
//
// Solidity: function multicall(bytes32 previousBlockhash, bytes[] data) payable returns(bytes[])
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) Multicall(previousBlockhash [32]byte, data [][]byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.Multicall(&_PancakeSwapRouterV3.TransactOpts, previousBlockhash, data)
}

// Multicall is a paid mutator transaction binding the contract method 0x1f0464d1.
//
// Solidity: function multicall(bytes32 previousBlockhash, bytes[] data) payable returns(bytes[])
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) Multicall(previousBlockhash [32]byte, data [][]byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.Multicall(&_PancakeSwapRouterV3.TransactOpts, previousBlockhash, data)
}

// Multicall0 is a paid mutator transaction binding the contract method 0x5ae401dc.
//
// Solidity: function multicall(uint256 deadline, bytes[] data) payable returns(bytes[])
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) Multicall0(opts *bind.TransactOpts, deadline *big.Int, data [][]byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "multicall0", deadline, data)
}

// Multicall0 is a paid mutator transaction binding the contract method 0x5ae401dc.
//
// Solidity: function multicall(uint256 deadline, bytes[] data) payable returns(bytes[])
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) Multicall0(deadline *big.Int, data [][]byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.Multicall0(&_PancakeSwapRouterV3.TransactOpts, deadline, data)
}

// Multicall0 is a paid mutator transaction binding the contract method 0x5ae401dc.
//
// Solidity: function multicall(uint256 deadline, bytes[] data) payable returns(bytes[])
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) Multicall0(deadline *big.Int, data [][]byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.Multicall0(&_PancakeSwapRouterV3.TransactOpts, deadline, data)
}

// Multicall1 is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) Multicall1(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "multicall1", data)
}

// Multicall1 is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) Multicall1(data [][]byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.Multicall1(&_PancakeSwapRouterV3.TransactOpts, data)
}

// Multicall1 is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) Multicall1(data [][]byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.Multicall1(&_PancakeSwapRouterV3.TransactOpts, data)
}

// PancakeV3SwapCallback is a paid mutator transaction binding the contract method 0x23a69e75.
//
// Solidity: function pancakeV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) PancakeV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "pancakeV3SwapCallback", amount0Delta, amount1Delta, _data)
}

// PancakeV3SwapCallback is a paid mutator transaction binding the contract method 0x23a69e75.
//
// Solidity: function pancakeV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) PancakeV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.PancakeV3SwapCallback(&_PancakeSwapRouterV3.TransactOpts, amount0Delta, amount1Delta, _data)
}

// PancakeV3SwapCallback is a paid mutator transaction binding the contract method 0x23a69e75.
//
// Solidity: function pancakeV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) PancakeV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.PancakeV3SwapCallback(&_PancakeSwapRouterV3.TransactOpts, amount0Delta, amount1Delta, _data)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address token, uint256 value) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) Pull(opts *bind.TransactOpts, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "pull", token, value)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address token, uint256 value) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) Pull(token common.Address, value *big.Int) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.Pull(&_PancakeSwapRouterV3.TransactOpts, token, value)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address token, uint256 value) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) Pull(token common.Address, value *big.Int) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.Pull(&_PancakeSwapRouterV3.TransactOpts, token, value)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) RefundETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "refundETH")
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) RefundETH() (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.RefundETH(&_PancakeSwapRouterV3.TransactOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) RefundETH() (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.RefundETH(&_PancakeSwapRouterV3.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) RenounceOwnership() (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.RenounceOwnership(&_PancakeSwapRouterV3.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.RenounceOwnership(&_PancakeSwapRouterV3.TransactOpts)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) SelfPermit(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "selfPermit", token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.SelfPermit(&_PancakeSwapRouterV3.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.SelfPermit(&_PancakeSwapRouterV3.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) SelfPermitAllowed(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "selfPermitAllowed", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.SelfPermitAllowed(&_PancakeSwapRouterV3.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.SelfPermitAllowed(&_PancakeSwapRouterV3.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) SelfPermitAllowedIfNecessary(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.SelfPermitAllowedIfNecessary(&_PancakeSwapRouterV3.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.SelfPermitAllowedIfNecessary(&_PancakeSwapRouterV3.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) SelfPermitIfNecessary(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "selfPermitIfNecessary", token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.SelfPermitIfNecessary(&_PancakeSwapRouterV3.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.SelfPermitIfNecessary(&_PancakeSwapRouterV3.TransactOpts, token, value, deadline, v, r, s)
}

// SetStableSwap is a paid mutator transaction binding the contract method 0x24dec034.
//
// Solidity: function setStableSwap(address _factory, address _info) returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) SetStableSwap(opts *bind.TransactOpts, _factory common.Address, _info common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "setStableSwap", _factory, _info)
}

// SetStableSwap is a paid mutator transaction binding the contract method 0x24dec034.
//
// Solidity: function setStableSwap(address _factory, address _info) returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) SetStableSwap(_factory common.Address, _info common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.SetStableSwap(&_PancakeSwapRouterV3.TransactOpts, _factory, _info)
}

// SetStableSwap is a paid mutator transaction binding the contract method 0x24dec034.
//
// Solidity: function setStableSwap(address _factory, address _info) returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) SetStableSwap(_factory common.Address, _info common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.SetStableSwap(&_PancakeSwapRouterV3.TransactOpts, _factory, _info)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x472b43f3.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to) payable returns(uint256 amountOut)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x472b43f3.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to) payable returns(uint256 amountOut)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.SwapExactTokensForTokens(&_PancakeSwapRouterV3.TransactOpts, amountIn, amountOutMin, path, to)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x472b43f3.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to) payable returns(uint256 amountOut)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.SwapExactTokensForTokens(&_PancakeSwapRouterV3.TransactOpts, amountIn, amountOutMin, path, to)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x42712a67.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to) payable returns(uint256 amountIn)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x42712a67.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to) payable returns(uint256 amountIn)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.SwapTokensForExactTokens(&_PancakeSwapRouterV3.TransactOpts, amountOut, amountInMax, path, to)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x42712a67.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to) payable returns(uint256 amountIn)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.SwapTokensForExactTokens(&_PancakeSwapRouterV3.TransactOpts, amountOut, amountInMax, path, to)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) SweepToken(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "sweepToken", token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.SweepToken(&_PancakeSwapRouterV3.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.SweepToken(&_PancakeSwapRouterV3.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken0 is a paid mutator transaction binding the contract method 0xe90a182f.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) SweepToken0(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "sweepToken0", token, amountMinimum)
}

// SweepToken0 is a paid mutator transaction binding the contract method 0xe90a182f.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) SweepToken0(token common.Address, amountMinimum *big.Int) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.SweepToken0(&_PancakeSwapRouterV3.TransactOpts, token, amountMinimum)
}

// SweepToken0 is a paid mutator transaction binding the contract method 0xe90a182f.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) SweepToken0(token common.Address, amountMinimum *big.Int) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.SweepToken0(&_PancakeSwapRouterV3.TransactOpts, token, amountMinimum)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0x3068c554.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) SweepTokenWithFee(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "sweepTokenWithFee", token, amountMinimum, feeBips, feeRecipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0x3068c554.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) SweepTokenWithFee(token common.Address, amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.SweepTokenWithFee(&_PancakeSwapRouterV3.TransactOpts, token, amountMinimum, feeBips, feeRecipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0x3068c554.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) SweepTokenWithFee(token common.Address, amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.SweepTokenWithFee(&_PancakeSwapRouterV3.TransactOpts, token, amountMinimum, feeBips, feeRecipient)
}

// SweepTokenWithFee0 is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) SweepTokenWithFee0(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "sweepTokenWithFee0", token, amountMinimum, recipient, feeBips, feeRecipient)
}

// SweepTokenWithFee0 is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) SweepTokenWithFee0(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.SweepTokenWithFee0(&_PancakeSwapRouterV3.TransactOpts, token, amountMinimum, recipient, feeBips, feeRecipient)
}

// SweepTokenWithFee0 is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) SweepTokenWithFee0(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.SweepTokenWithFee0(&_PancakeSwapRouterV3.TransactOpts, token, amountMinimum, recipient, feeBips, feeRecipient)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.TransferOwnership(&_PancakeSwapRouterV3.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.TransferOwnership(&_PancakeSwapRouterV3.TransactOpts, newOwner)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) UnwrapWETH9(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "unwrapWETH9", amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.UnwrapWETH9(&_PancakeSwapRouterV3.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.UnwrapWETH9(&_PancakeSwapRouterV3.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) UnwrapWETH9WithFee(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "unwrapWETH9WithFee", amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) UnwrapWETH9WithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.UnwrapWETH9WithFee(&_PancakeSwapRouterV3.TransactOpts, amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) UnwrapWETH9WithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.UnwrapWETH9WithFee(&_PancakeSwapRouterV3.TransactOpts, amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee0 is a paid mutator transaction binding the contract method 0xd4ef38de.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) UnwrapWETH9WithFee0(opts *bind.TransactOpts, amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "unwrapWETH9WithFee0", amountMinimum, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee0 is a paid mutator transaction binding the contract method 0xd4ef38de.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) UnwrapWETH9WithFee0(amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.UnwrapWETH9WithFee0(&_PancakeSwapRouterV3.TransactOpts, amountMinimum, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee0 is a paid mutator transaction binding the contract method 0xd4ef38de.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) UnwrapWETH9WithFee0(amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.UnwrapWETH9WithFee0(&_PancakeSwapRouterV3.TransactOpts, amountMinimum, feeBips, feeRecipient)
}

// WrapETH is a paid mutator transaction binding the contract method 0x1c58db4f.
//
// Solidity: function wrapETH(uint256 value) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) WrapETH(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.Transact(opts, "wrapETH", value)
}

// WrapETH is a paid mutator transaction binding the contract method 0x1c58db4f.
//
// Solidity: function wrapETH(uint256 value) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) WrapETH(value *big.Int) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.WrapETH(&_PancakeSwapRouterV3.TransactOpts, value)
}

// WrapETH is a paid mutator transaction binding the contract method 0x1c58db4f.
//
// Solidity: function wrapETH(uint256 value) payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) WrapETH(value *big.Int) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.WrapETH(&_PancakeSwapRouterV3.TransactOpts, value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeSwapRouterV3.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Session) Receive() (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.Receive(&_PancakeSwapRouterV3.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3TransactorSession) Receive() (*types.Transaction, error) {
	return _PancakeSwapRouterV3.Contract.Receive(&_PancakeSwapRouterV3.TransactOpts)
}

// PancakeSwapRouterV3OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PancakeSwapRouterV3 contract.
type PancakeSwapRouterV3OwnershipTransferredIterator struct {
	Event *PancakeSwapRouterV3OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PancakeSwapRouterV3OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeSwapRouterV3OwnershipTransferred)
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
		it.Event = new(PancakeSwapRouterV3OwnershipTransferred)
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
func (it *PancakeSwapRouterV3OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeSwapRouterV3OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeSwapRouterV3OwnershipTransferred represents a OwnershipTransferred event raised by the PancakeSwapRouterV3 contract.
type PancakeSwapRouterV3OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PancakeSwapRouterV3OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PancakeSwapRouterV3.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PancakeSwapRouterV3OwnershipTransferredIterator{contract: _PancakeSwapRouterV3.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PancakeSwapRouterV3OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PancakeSwapRouterV3.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeSwapRouterV3OwnershipTransferred)
				if err := _PancakeSwapRouterV3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Filterer) ParseOwnershipTransferred(log types.Log) (*PancakeSwapRouterV3OwnershipTransferred, error) {
	event := new(PancakeSwapRouterV3OwnershipTransferred)
	if err := _PancakeSwapRouterV3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeSwapRouterV3SetStableSwapIterator is returned from FilterSetStableSwap and is used to iterate over the raw logs and unpacked data for SetStableSwap events raised by the PancakeSwapRouterV3 contract.
type PancakeSwapRouterV3SetStableSwapIterator struct {
	Event *PancakeSwapRouterV3SetStableSwap // Event containing the contract specifics and raw log

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
func (it *PancakeSwapRouterV3SetStableSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeSwapRouterV3SetStableSwap)
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
		it.Event = new(PancakeSwapRouterV3SetStableSwap)
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
func (it *PancakeSwapRouterV3SetStableSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeSwapRouterV3SetStableSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeSwapRouterV3SetStableSwap represents a SetStableSwap event raised by the PancakeSwapRouterV3 contract.
type PancakeSwapRouterV3SetStableSwap struct {
	Factory common.Address
	Info    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetStableSwap is a free log retrieval operation binding the contract event 0x26e41379222b54b0470031bc11852ad23058ffb8983f7cc0e18257d6f7afca9d.
//
// Solidity: event SetStableSwap(address indexed factory, address indexed info)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Filterer) FilterSetStableSwap(opts *bind.FilterOpts, factory []common.Address, info []common.Address) (*PancakeSwapRouterV3SetStableSwapIterator, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}
	var infoRule []interface{}
	for _, infoItem := range info {
		infoRule = append(infoRule, infoItem)
	}

	logs, sub, err := _PancakeSwapRouterV3.contract.FilterLogs(opts, "SetStableSwap", factoryRule, infoRule)
	if err != nil {
		return nil, err
	}
	return &PancakeSwapRouterV3SetStableSwapIterator{contract: _PancakeSwapRouterV3.contract, event: "SetStableSwap", logs: logs, sub: sub}, nil
}

// WatchSetStableSwap is a free log subscription operation binding the contract event 0x26e41379222b54b0470031bc11852ad23058ffb8983f7cc0e18257d6f7afca9d.
//
// Solidity: event SetStableSwap(address indexed factory, address indexed info)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Filterer) WatchSetStableSwap(opts *bind.WatchOpts, sink chan<- *PancakeSwapRouterV3SetStableSwap, factory []common.Address, info []common.Address) (event.Subscription, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}
	var infoRule []interface{}
	for _, infoItem := range info {
		infoRule = append(infoRule, infoItem)
	}

	logs, sub, err := _PancakeSwapRouterV3.contract.WatchLogs(opts, "SetStableSwap", factoryRule, infoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeSwapRouterV3SetStableSwap)
				if err := _PancakeSwapRouterV3.contract.UnpackLog(event, "SetStableSwap", log); err != nil {
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

// ParseSetStableSwap is a log parse operation binding the contract event 0x26e41379222b54b0470031bc11852ad23058ffb8983f7cc0e18257d6f7afca9d.
//
// Solidity: event SetStableSwap(address indexed factory, address indexed info)
func (_PancakeSwapRouterV3 *PancakeSwapRouterV3Filterer) ParseSetStableSwap(log types.Log) (*PancakeSwapRouterV3SetStableSwap, error) {
	event := new(PancakeSwapRouterV3SetStableSwap)
	if err := _PancakeSwapRouterV3.contract.UnpackLog(event, "SetStableSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
