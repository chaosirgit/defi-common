package helpers

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v8"
	"math/big"
	"reflect"
)

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

func GetBaseTokenAddressesByChainId(chainId *big.Int) (map[string]string, error) {
	result := map[string]string{}
	bnbChain := big.NewInt(56)
	ethChain := big.NewInt(1)
	switch {
	case chainId.Cmp(bnbChain) == 0:
		result = map[string]string{
			"WETH": "0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c",
			"BUSD": "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56",
			"USDT": "0x55d398326f99059fF775485246999027B3197955",
			"USDC": "0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d",
			"DEFI": "0x10ED43C718714eb63d5aA57B78B54704E256024E",
		}
	case chainId.Cmp(ethChain) == 0:
		result = map[string]string{
			"WETH": "0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c",
			"BUSD": "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56",
			"USDT": "0x55d398326f99059fF775485246999027B3197955",
			"USDC": "0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d",
			"DEFI": "",
		}
	default:
		return nil, errors.New("暂不支持的链")
	}
	return result, nil
}
