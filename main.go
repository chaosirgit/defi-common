package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
)

func main() {
	hash := "0xf73e4e060cc8c4de810b5ffda1b0a38cd5f3cb90e27891875ed3e2dbf7605131"
	//hash := "0x6990d920b6189ac5a5b7d0f9d3ddc775df1179b70e6b7a501b76683157b6a54f"
	log.Println(hash)
	//c, err := rpc.DialOptions(context.Background(), "https://rpc.payload.de")
	c, err := rpc.DialOptions(context.Background(), "https://data-seed-prebsc-1-s1.binance.org:8545")
	//c, err := rpc.DialOptions(context.Background(), "https://bsc-dataseed.binance.org/")
	if err != nil {
		log.Fatalf("Client Node Error: %v\n", err)
	}
	defer c.Close()
	ec := ethclient.NewClient(c)

	// 连接节点
	if ec == nil {
		log.Fatal("Failed to create ethclient.Client")
	}
	//订阅
	chainId, err := ec.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(chainId)
	receipt, err := ec.TransactionReceipt(context.Background(), common.HexToHash(hash))
	log.Println(receipt)
	log.Println(err)
	for _, l := range receipt.Logs {
		fmt.Println("Log Block Number: ", l.BlockNumber)
		fmt.Println("Log Index: ", l.Index)
		fmt.Println("Log Removed: ", l.Removed)
		//fmt.Println("Log Type: ", )
		// Data is a byte array containing the non-indexed log data
		fmt.Println("Log Data: ", l.Data)
		// Topics are a list of 32-byte hash which are the indexed event parameters
		fmt.Println("Log Topics: ", l.Topics)
		// Etc...
	}

}
