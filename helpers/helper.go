package helpers

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/chaosirgit/defi-common/pkg/contract/pancakeSwapRouter"
	"github.com/chaosirgit/defi-common/pkg/contract/token"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v8"
	"io"
	"log"
	"math/big"
	"net/http"
	"reflect"
	"strings"
)

func ParsePath(path []byte) ([]common.Address, []int, error) {
	var addresses []common.Address
	var fees []int

	for i := 0; i < len(path); i += 23 {
		// 从 path 中解析出代币地址
		addresses = append(addresses, common.BytesToAddress(path[i:i+20]))

		// 如果这不是最后一个地址，则还有一个费用字段
		if i+23 < len(path) {
			// 从 path 中解析出费用
			fee := new(big.Int).SetBytes(path[i+20 : i+23]).Int64()
			fees = append(fees, int(fee))
		}
	}

	return addresses, fees, nil
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

// InArray will search element inside array with any type.
// Will return boolean and index for matched element.
// True and index more than 0 if element is exist.
// needle is element to search, haystack is slice of value to be search.
func InArray(needle interface{}, haystack interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(haystack)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(needle, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

func GetBuyCoefficient(amount *big.Int, total *big.Int) *big.Float {
	//计算交易系数
	bFloat := new(big.Float).SetInt(total)
	aFloat := new(big.Float).SetInt(amount)
	f := new(big.Float).Quo(aFloat, bFloat)
	rounded := new(big.Float)
	rounded.SetPrec(64).SetMode(big.ToNearestEven)
	rounded.Mul(f, big.NewFloat(10000))
	rounded.Quo(rounded, big.NewFloat(1))
	rounded.SetMode(big.ToNearestEven).Quo(rounded, big.NewFloat(10000))
	return rounded
}

// 计算两个值的差值和差值占x的系数
// 返回差值和系数
func PercentageDifference(x *big.Int, y *big.Int) (*big.Int, *big.Float) {
	// 计算x与y之间的差值
	difference := new(big.Int).Sub(x, y)
	if difference.Sign() == -1 {
		difference.Neg(difference)
	}

	// 创建big.Float值，用于计算百分比
	differenceFloat := new(big.Float).SetInt(difference)
	xFloat := new(big.Float).SetInt(x)
	//hundred := big.NewFloat(100)
	// 如果 x 或 difference 为零，则直接返回结果
	zero := big.NewInt(0)
	if x.Cmp(zero) == 0 || difference.Cmp(zero) == 0 {
		return difference, big.NewFloat(0)
	}
	// 计算百分比
	percentage := new(big.Float).Quo(differenceFloat, xFloat)
	//percentage.Mul(percentage, hundred)

	return difference, percentage
}

func MultiplyIntFloat(x *big.Int, y *big.Float) *big.Int {
	// 将 *big.Int 转换为 *big.Float
	xFloat := new(big.Float).SetInt(x)

	// 计算乘积
	product := new(big.Float).Mul(xFloat, y)

	// 将 *big.Float 转换回 *big.Int，同时处理舍入
	productInt := new(big.Int)
	product.Int(productInt)

	return productInt
}
func GetCurrentNonce(client *ethclient.Client, account common.Address) (uint64, error) {
	nonce, err := client.PendingNonceAt(context.Background(), account)
	if err != nil {
		return 0, err
	}
	return nonce, nil
}

// 交易签名数据

func GetAuthData(privateKey *ecdsa.PrivateKey, chainId *big.Int, client *ethclient.Client) (*bind.TransactOpts, error) {
	gasPrice, err := GetCurrentGasPrice(client)
	if err != nil {
		return nil, err
	}
	//account := crypto.PubkeyToAddress(privateKey.PublicKey)
	//
	//nonce, err := GetCurrentNonce(client, account)
	//if err != nil {
	//	return nil, err
	//}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, err
	}
	//auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	//auth.GasPrice = new(big.Int).SetInt64(5000000000)
	auth.GasPrice = gasPrice
	return auth, nil
}

func GetCurrentGasPrice(client *ethclient.Client) (*big.Int, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	return gasPrice, nil
}

// 获取所有 redis 键
func GetKeysWithPrefix(rdb *redis.Client, prefix string) ([]string, error) {
	var keys []string
	var cursor uint64
	var err error

	for {
		var result []string
		result, cursor, err = rdb.Scan(context.Background(), cursor, prefix+"*", 10).Result()
		if err != nil {
			return nil, err
		}

		keys = append(keys, result...)

		if cursor == 0 {
			break
		}
	}

	return keys, nil
}

func GetBaseTokenAddressesByChainId(chainId *big.Int, rdb *redis.Client) (result []common.Address, mainToken *common.Address, usdToken *common.Address, err error) {
	baseTokens, err := rdb.SMembers(context.Background(), "BaseTokenList:"+chainId.String()).Result()
	if err != nil {
		return nil, nil, nil, err
	}

	for _, ts := range baseTokens {
		v, err := rdb.HGetAll(context.Background(), fmt.Sprintf("BaseToken:%s:%s", chainId.String(), ts)).Result()
		if err != nil {
			return nil, nil, nil, err
		}
		isMain := v["is_main"] == "1"
		if isMain {
			main := common.HexToAddress(ts)
			mainToken = &main
		}
		isUsd := v["is_usd"] == "1"
		if isUsd {
			u := common.HexToAddress(ts)
			usdToken = &u
		}
		result = append(result, common.HexToAddress(ts))
	}
	return
}

func IsHoneypotToken(token common.Address, chainId *big.Int, rdb *redis.Client) bool {
	key := fmt.Sprintf("HoneyPot:%s:%s", chainId.String(), strings.ToLower(token.String()))
	e, err := rdb.Exists(context.Background(), key).Result()
	if err != nil {
		log.Println(err)
		return true
	}
	if e > 0 {
		a, err := rdb.Get(context.Background(), key).Result()
		if err != nil {
			log.Println(err)
			return true
		}
		if a == "safely" {
			return false
		}
	} else {
		resp, err := http.Get(fmt.Sprintf("https://api.gopluslabs.io/api/v1/token_security/%s?contract_addresses=%s", chainId.String(), token.String()))
		if err != nil {
			log.Println(err)
			return true
		}
		// 关闭响应体
		defer resp.Body.Close()

		// 读取响应体
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return true
		}
		//str := string(body)
		a := make(map[string]interface{})
		err = json.Unmarshal(body, &a)
		if err != nil {
			log.Println(err)
			return true
		}
		res := a["result"].(map[string]interface{})[strings.ToLower(token.String())]
		if res != nil {
			r := "dangerous"
			m := res.(map[string]interface{})
			//开源
			if m["is_open_source"].(string) == "1" {
				//无代理
				if m["is_proxy"] != nil && m["is_proxy"].(string) == "0" {
					if m["is_honeypot"] != nil && m["is_honeypot"].(string) == "0" {
						r = "safely"
					}
				}
			}
			rdb.Set(context.Background(), key, r, redis.KeepTTL)
			if r == "safely" {
				return false
			} else {
				return true
			}
		}
	}
	return true
}

// price 除以基本币的小数位后的价格 price1 没除以小数位价格
func GetPrice(defiAddress common.Address, tokenAddress common.Address, baseAddress common.Address, ec *ethclient.Client) (*big.Float, *big.Int, error) {
	t1, _ := token.NewToken(tokenAddress, ec)
	tb, _ := token.NewToken(baseAddress, ec)
	decimal, _ := t1.Decimals(nil)
	baseDecimal, _ := tb.Decimals(nil)
	pancake, err := pancakeSwapRouter.NewPancakeSwapRouter(defiAddress, ec)
	// 10 的 decimal 次方 的 token1 = 多少 token0 也就是 1个 购买币 等于多少个 基本币
	prices, err := pancake.GetAmountsOut(nil, new(big.Int).Exp(big.NewInt(10), new(big.Int).SetUint64(uint64(decimal)), nil), []common.Address{tokenAddress, baseAddress})
	if err != nil {
		return nil, nil, err
	}
	baseD := new(big.Int).Exp(big.NewInt(10), new(big.Int).SetUint64(uint64(baseDecimal)), nil)
	price := new(big.Float).Quo(new(big.Float).SetInt(prices[1]), new(big.Float).SetInt(baseD))
	return price, prices[1], nil
}
