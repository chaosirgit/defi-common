package main

import (
	"context"
	"encoding/json"
	"github.com/chaosirgit/defi-common/helpers"
	"github.com/chaosirgit/defi-common/pkg/contract/uniSwapRouterV3"
	"github.com/chaosirgit/defi-common/structs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
)

func main() {
	hash := "0x24300892dc970262fa4452ecce1eef87b06bdbd89a7abbd8897c78b8372fdae6"
	//hash := "0x6990d920b6189ac5a5b7d0f9d3ddc775df1179b70e6b7a501b76683157b6a54f"
	log.Println(hash)
	c, err := rpc.DialOptions(context.Background(), "https://rpc.payload.de")
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
	transaction, pending, err := ec.TransactionByHash(context.Background(), common.HexToHash(hash))
	log.Println(pending)
	log.Println(err)
	log.Println(transaction)
	data := transaction.Data()
	d, err := structs.GetMethodsAndParamsFromInputData(common.Bytes2Hex(data), 3, "uniswap")
	println(d.Method.RawName)
	log.Println(d.Params[0])
	p, err := json.Marshal(d.Params[0])
	log.Println(string(p))

	var temp uniSwapRouterV3.ISwapRouterExactInputParams
	json.Unmarshal(p, &temp)
	path, _, _ := helpers.ParsePath(temp.Path)
	log.Println(path)
	log.Println(temp)

}
