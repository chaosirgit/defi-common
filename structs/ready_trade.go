package structs

import (
	"context"
	"errors"
	"fmt"
	"github.com/chaosirgit/defi-common/helpers"
	"github.com/chaosirgit/defi-common/pkg/contract/pancakeSwapRouter"
	"github.com/chaosirgit/defi-common/pkg/contract/token"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v8"
	"log"
	"math/big"
	"time"
)

type ReadyTrade struct {
	Token0      string `json:"token0"`
	Token1      string `json:"token1"`
	DefiAddress string `json:"defi_address"`
	ChainId     string `json:"chain_id"`
	Amount      string `json:"amount"`
	Account     string `json:"account"`
	IsBuy       bool   `json:"is_buy"`
	SettingId   string `json:"setting_id"`
	Type        int8   `json:"type"` // 1-正常分析(百分比买卖) 2-止盈分析(涨幅超过用户设置的全部卖出)
}

func (r *ReadyTrade) Send(rdb *redis.Client, ec *ethclient.Client) (*types.Transaction, error) {
	//读取用户设置
	var operation string
	if r.IsBuy {
		operation = "购买"
	} else {
		operation = "出售"
	}

	us, err := GetUserBotSettingById(r.SettingId, rdb)
	if err != nil {
		return nil, err
	}
	token0 := common.HexToAddress(r.Token0)
	token1 := common.HexToAddress(r.Token1)
	//发送交易
	t0, err := token.NewToken(token0, ec)
	if err != nil {
		return nil, err
	}
	t1, err := token.NewToken(token1, ec)
	if err != nil {
		return nil, err
	}
	t0symbol, _ := t0.Symbol(nil)
	t1symbol, _ := t1.Symbol(nil)
	account := common.HexToAddress(r.Account)
	defiRouterAddress := common.HexToAddress(r.DefiAddress)
	//首先查询 Approve
	allowance, err := t0.Allowance(nil, account, defiRouterAddress)
	if err != nil {
		return nil, err
	}
	amount, ok := new(big.Int).SetString(r.Amount, 10)
	if !ok {
		return nil, errors.New("Approve Error amount converting string to *big.Int")
	}
	chainId, ok := new(big.Int).SetString(r.ChainId, 10)
	if !ok {
		return nil, errors.New("Approve Error chainId converting string to *big.Int")
	}

	privateKey, err := crypto.HexToECDSA(us.PrivateKey)
	if err != nil {
		return nil, errors.New("Parse private error")
	}

	if amount.Cmp(allowance) > 0 {
		max := new(big.Int).Lsh(big.NewInt(2), 256)
		max.Sub(max, big.NewInt(1))
		maxApprove := new(big.Int).Sub(max, allowance)
		authData, _ := helpers.GetAuthData(privateKey, chainId, ec)
		approveTx, err := t0.Approve(authData, defiRouterAddress, maxApprove)
		if err != nil {
			return nil, err
		}
		// Wait for the transaction to be mined.
		receipt, err := bind.WaitMined(context.Background(), ec, approveTx)
		if err != nil {
			return nil, err
		}
		log.Printf("授权 (%s) 成功, 授权交易hash: %s ,继续交易...\n", t0symbol, receipt.TxHash.String())
	} else {
		log.Printf("代币 (%s) 已授权给 Defi , 继续交易...\n", t0symbol)
	}

	//开始交易
	//买卖
	pancake, err := pancakeSwapRouter.NewPancakeSwapRouter(defiRouterAddress, ec)
	if err != nil {
		return nil, err
	}
	path := []common.Address{token0, token1}
	// 调用 getAmountsOut 方法计算输入代币数量。
	amountsOut, err := pancake.GetAmountsOut(
		nil,
		amount, // 输入代币数量
		path,
	)
	if err != nil {
		return nil, err
	}
	deadline := big.NewInt(time.Now().Add(1 * time.Minute).Unix())
	// 百分之10的滑点
	amountOut := new(big.Int).Sub(amountsOut[1], new(big.Int).Div(amountsOut[1], big.NewInt(100)))
	authPancakeData, _ := helpers.GetAuthData(privateKey, chainId, ec)
	b, err := pancake.SwapExactTokensForTokens(authPancakeData, amount, amountOut, path, account, deadline)
	if err != nil {
		return nil, err
	}
	log.Printf("开始广播(%s) (%s:%s) To (%s:%s),Hash: %v\n", operation, t0symbol, amount.String(), t1symbol, amountOut.String(), b.Hash().String())
	bReceipt, err := bind.WaitMined(context.Background(), ec, b)
	if err != nil {
		return nil, err
	}
	if bReceipt.Status == 1 {
		if r.IsBuy {
			if b != nil {
				decimal, _ := t1.Decimals(nil)
				// 10 的 decimal 次方 的 token1 = 多少 token0 也就是 1个 购买币 等于多少个 基本币
				prices, err := pancake.GetAmountsOut(nil, new(big.Int).Exp(big.NewInt(10), new(big.Int).SetUint64(uint64(decimal)), nil), []common.Address{token1, token0})
				if err != nil {
					return nil, err
				}
				price := new(big.Float).Quo(new(big.Float).SetInt(prices[1]), big.NewFloat(1e18))
				var purchase PurchaseInfo
				purchase.UserBotSettingID = r.SettingId
				purchase.Time = time.Now().Unix()
				purchase.TokenAddress = r.Token1
				purchase.BaseAddress = r.Token0
				purchase.PurchasePrice = price.String()
				err = purchase.SavePurchaseInfoToRedis(rdb)
				if err != nil {
					return b, err
				}
			}
		} else {
			// 到达止盈线的出售
			if r.Type == 2 {
				// 删除买入 key
				var purchase PurchaseInfo
				purchase.UserBotSettingID = r.SettingId
				// 这里时卖出，记得 token0 和 token1 的顺序
				purchase.TokenAddress = r.Token0
				purchase.BaseAddress = r.Token1
				err = purchase.RemovePurchaseInfoFromRedis(rdb)
				if err != nil {
					return b, err
				}
				//设置几分钟之内不买
				err = rdb.Set(context.Background(), fmt.Sprintf("OB:%d:%s", r.SettingId, r.Token0), "1", time.Minute*time.Duration(us.StopWinMin)).Err()
				if err != nil {
					return b, err
				}
			}
		}
		log.Printf("发送(%s)成功 (%s:%s) To (%s:%s),Hash: %v\n", operation, t0symbol, amount.String(), t1symbol, amountOut.String(), b.Hash().String())
		return b, nil
	}
	return nil, errors.New("交易未成功...")
}
