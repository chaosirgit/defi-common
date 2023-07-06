package main

import (
	"github.com/chaosirgit/defi-common/helpers"
	"github.com/chaosirgit/defi-common/pkg/contract/uniswapQuoterV3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	getAmountOutV3()
}

func getAmountOutV3() {
	client, err := ethclient.Dial("https://rpc.payload.de")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	contractAddress := common.HexToAddress("0xb27308f9F90D607463bb33eA1BeBb41C27CE5AB6")
	//quoterContract, err := uniswapQuoterV3.NewUniswapQuoterV3(contractAddress, client)
	// Prepare the method arguments
	//path := BuildPath([]common.Address{common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"), common.HexToAddress("0x2dfF88A56767223A5529eA5960Da7A3F5f766406")}, []int{1000000}) // replace with your path
	path, _ := BuildPath([]common.Address{common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"), common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"), common.HexToAddress("0x95aD61b0a150d79219dCF64E1E6Cc01f0B64C4cE")}, []int{500, 3000}) // replace with your path

	instance, err := uniswapQuoterV3.NewUniswapQuoterV3(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to create new contract instance: %v", err)
	}

	// Prepare the method arguments
	tokenIn := common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")  // replace with your tokenIn address
	tokenOut := common.HexToAddress("0x95aD61b0a150d79219dCF64E1E6Cc01f0B64C4cE") // replace with your tokenOut address

	fee := big.NewInt(10000)                                       // replace with your fee
	amountIn, _ := new(big.Int).SetString("10000000000000000", 10) // replace with your amount
	sqrtPriceLimitX96, _ := new(big.Int).SetString("0", 10)        // replace with your sqrtPriceLimitX96
	//path, _ := BuildPath([]common.Address{tokenIn, tokenOut}, []int{10000})
	log.Printf("tokenIn: %v\n tokenOut: %v\n path: %v\n", tokenIn.Bytes(), tokenOut.Bytes(), path) // replace with your path
	parse, f, err := helpers.ParsePath(path)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("parse path: %v\nfee: %v\n", parse, f)
	// Prepare a TransactOpts object with no signer
	//opts := &bind.TransactOpts{}

	// Call the method
	amountOut, err := instance.QuoteExactInputSingle(nil, tokenIn, tokenOut, fee, amountIn, sqrtPriceLimitX96)
	if err != nil {
		log.Printf("Failed to call QuoteExactInputSingle method: %v", err)
	}
	amountOutPath, err := instance.QuoteExactInput(nil, path, amountIn)
	if err != nil {
		log.Printf("Failed to call QuoteExactInput method: %v", err)
	}

	log.Printf("Amount out: %v", amountOut)
	log.Printf("Amount out path: %v", amountOutPath)

}
func BuildPath(addresses []common.Address, fees []int) ([]byte, error) {
	var path []byte

	for i, address := range addresses {
		// 将地址添加到路径中
		path = append(path, address.Bytes()...)

		// 如果这不是最后一个地址，则还有一个费用字段
		if i < len(fees) {
			// 将费用转换为字节数组并添加到路径中
			feeBytes := big.NewInt(int64(fees[i])).Bytes()
			// 如果费用字节数组的长度小于3，我们需要在前面添加零字节
			for len(feeBytes) < 3 {
				feeBytes = append([]byte{0}, feeBytes...)
			}
			path = append(path, feeBytes...)
		}
	}

	return path, nil
}
