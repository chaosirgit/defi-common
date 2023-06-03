package structs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chaosirgit/defi-common/helpers"
	"github.com/chaosirgit/defi-common/pkg/contract/pancakeSwapRouter"
	"github.com/chaosirgit/defi-common/pkg/contract/token"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v8"
	"math/big"
	"strconv"
	"strings"
	"time"
)

func getMethodAndArgsFromInputData(stringData string) (*abi.Method, []interface{}, error) {
	hexData := common.FromHex(stringData)
	methodId := hexData[:4]
	input := hexData[4:]
	pancakeSwapAbi, _ := pancakeSwapRouter.PancakeSwapRouterMetaData.GetAbi()
	method, err := pancakeSwapAbi.MethodById(methodId)
	if err != nil {
		return nil, nil, err
	}
	args, err := method.Inputs.UnpackValues(input)
	if err != nil {
		return nil, nil, err
	}
	return method, args, nil
}

func GetUserTaskById(chainId string, id string, rds *redis.Client) (*UserTask, error) {
	settingKey := fmt.Sprintf("FollowSetting:%s:%s", chainId, id)
	setting, err := rds.HGetAll(context.Background(), settingKey).Result()
	if err != nil {
		return nil, err
	}
	us, err := MapToUserTask(setting)
	if err != nil {
		return nil, err
	}
	return us, nil
}

// 用户配置
type UserTask struct {
	ID                     int64          `json:"id"`
	UserId                 int64          `json:"user_id"`
	Name                   string         `json:"name"`
	ChainId                *big.Int       `json:"chain_id"`
	FollowAddress          common.Address `json:"follow_address"`
	BuyType                int64          `json:"buy_type"`
	Amount                 *big.Int       `json:"amount"`
	MinAmountAction        int64          `json:"min_amount_action"`
	MinAmount              *big.Int       `json:"min_amount"`
	MaxAmountAction        int64          `json:"max_amount_action"`
	MaxAmount              *big.Int       `json:"max_amount"`
	StopWin                int64          `json:"stop_win"`
	StopWinMin             int64          `json:"stop_win_min"`
	StopLost               int64          `json:"stop_lost"`
	StopLostMin            int64          `json:"stop_lost_min"`
	StopAt                 int64          `json:"stop_at"`
	IsSellAll              int64          `json:"is_sell_all"`
	IsSell                 int64          `json:"is_sell"`
	AvoidHoneyPot          int64          `json:"avoid_honey_pot"`
	IsZero                 int64          `json:"is_zero"`
	PreventSandwichAttacks int64          `json:"prevent_sandwich_attacks"`
	PrivateKey             string         `json:"private_key"`
	IsVip                  int64          `json:"is_vip"`
}

func (ut *UserTask) UnmarshalJSON(data []byte) error {
	aux := struct {
		ID                     int64  `json:"id"`
		UserId                 int64  `json:"user_id"`
		Name                   string `json:"name"`
		ChainId                string `json:"chain_id"`
		FollowAddress          string `json:"follow_address"`
		BuyType                int64  `json:"buy_type"`
		Amount                 string `json:"amount"`
		MinAmountAction        int64  `json:"min_amount_action"`
		MinAmount              string `json:"min_amount"`
		MaxAmountAction        int64  `json:"max_amount_action"`
		MaxAmount              string `json:"max_amount"`
		StopWin                int64  `json:"stop_win"`
		StopWinMin             int64  `json:"stop_win_min"`
		StopLost               int64  `json:"stop_lost"`
		StopLostMin            int64  `json:"stop_lost_min"`
		StopAt                 int64  `json:"stop_at"`
		IsSellAll              int64  `json:"is_sell_all"`
		IsSell                 int64  `json:"is_sell"`
		AvoidHoneyPot          int64  `json:"avoid_honey_pot"`
		IsZero                 int64  `json:"is_zero"`
		PreventSandwichAttacks int64  `json:"prevent_sandwich_attacks"`
		PrivateKey             string `json:"private_key"`
		IsVip                  int64  `json:"is_vip"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	chainId, ok := new(big.Int).SetString(aux.ChainId, 10)
	if !ok {
		return fmt.Errorf("failed to unmarshal chain id from string: %s", aux.ChainId)
	}
	amount, ok := new(big.Int).SetString(aux.Amount, 10)
	if !ok {
		return fmt.Errorf("failed to unmarshal amount from string: %s", aux.Amount)
	}
	minAmount, ok := new(big.Int).SetString(aux.MinAmount, 10)
	if !ok {
		return fmt.Errorf("failed to unmarshal minAmount from string: %s", aux.MinAmount)
	}
	maxAmount, ok := new(big.Int).SetString(aux.MaxAmount, 10)
	if !ok {
		return fmt.Errorf("failed to unmarshal maxAmount from string: %s", aux.MaxAmount)
	}
	ut.ID = aux.ID
	ut.UserId = aux.UserId
	ut.Name = aux.Name
	ut.ChainId = chainId
	ut.FollowAddress = common.HexToAddress(aux.FollowAddress)
	ut.BuyType = aux.BuyType
	ut.Amount = amount
	ut.MinAmountAction = aux.MinAmountAction
	ut.MinAmount = minAmount
	ut.MaxAmountAction = aux.MaxAmountAction
	ut.MaxAmount = maxAmount
	ut.StopWin = aux.StopWin
	ut.StopWinMin = aux.StopWinMin
	ut.StopLost = aux.StopLost
	ut.StopLostMin = aux.StopLostMin
	ut.StopAt = aux.StopAt
	ut.IsSellAll = aux.IsSellAll
	ut.IsSell = aux.IsSell
	ut.AvoidHoneyPot = aux.AvoidHoneyPot
	ut.IsZero = aux.IsZero
	ut.PreventSandwichAttacks = aux.PreventSandwichAttacks
	ut.PrivateKey = aux.PrivateKey
	ut.IsVip = aux.IsVip
	return nil
}

// 分析交易
func (s *UserTask) AnalyzeTransaction(t TransactionData, client *ethclient.Client, rds *redis.Client) (*ReadyTrade, error) {
	var ready ReadyTrade
	ready.IsBuy = false
	ready.Type = 1
	if s.StopAt <= time.Now().Unix() {
		return nil, errors.New("用户脚本运行到期")
	}

	// 判断巨鲸交易状态
	r, err := client.TransactionReceipt(context.Background(), t.Hash)
	if err != nil || r.Status != 1 {
		return nil, errors.New("监听到的交易为错误交易")
	}
	method, args, err := getMethodAndArgsFromInputData(t.Data)
	if err != nil {
		return nil, err
	}
	// 确定 token0 token1
	var token0, token1 common.Address
	baseAddress, mainAddress, usdAddress, err := helpers.GetBaseTokenAddressesByChainId(t.ChainId, rds)
	if err != nil {
		return nil, err
	}
	var ETHForTokens = []string{"swapETHForExactTokens", "swapExactETHForTokens", "swapExactETHForTokensSupportingFeeOnTransferTokens"}
	var TokensForETH = []string{"swapExactTokensForETH", "swapTokensForExactETH", "swapExactTokensForETHSupportingFeeOnTransferTokens"}
	var TokensForTokens = []string{"swapExactTokensForTokens", "swapTokensForExactTokens", "swapExactTokensForTokensSupportingFeeOnTransferTokens"}
	var RemoveLiquidity = []string{"removeLiquidity", "removeLiquidityWithPermit"}
	var RemoveETHLiquidity = []string{"removeLiquidityETH", "removeLiquidityETHWithPermit", "removeLiquidityETHSupportingFeeOnTransferTokens", "removeLiquidityETHWithPermitSupportingFeeOnTransferTokens"}
	var addLiquidity = []string{"addLiquidity"}
	var addETHLiquidity = []string{"addLiquidityETH"}
	if ok, _ := helpers.InArray(method.RawName, ETHForTokens); ok {
		pathRouter := args[1].([]common.Address)
		token0 = *mainAddress
		token1 = pathRouter[len(pathRouter)-1]
	} else if ok, _ = helpers.InArray(method.RawName, TokensForETH); ok {
		pathRouter := args[2].([]common.Address)
		token0 = pathRouter[0]
		token1 = *mainAddress

	} else if ok, _ = helpers.InArray(method.RawName, TokensForTokens); ok {
		pathRouter := args[2].([]common.Address)
		token0 = pathRouter[len(pathRouter)-2]
		token1 = pathRouter[len(pathRouter)-1]

	} else if ok, _ = helpers.InArray(method.RawName, RemoveLiquidity); ok {
		token0 = args[0].(common.Address)
		token1 = args[1].(common.Address)
		if is, _ := helpers.InArray(args[0].(common.Address), baseAddress); is {
			token0 = args[1].(common.Address)
			token1 = args[0].(common.Address)
		}

	} else if ok, _ = helpers.InArray(method.RawName, RemoveETHLiquidity); ok {
		token0 = args[0].(common.Address)
		token1 = *mainAddress
	} else if ok, _ = helpers.InArray(method.RawName, addLiquidity); ok {
		return nil, errors.New("添加流动性,暂时跳过")
	} else if ok, _ = helpers.InArray(method.RawName, addETHLiquidity); ok {
		return nil, errors.New("移除流动性,暂时跳过")
	}

	//确定买卖以及金额
	var hasBuy bool
	var hasSell bool

	if ok, _ := helpers.InArray(token0, baseAddress); ok {
		hasBuy = true
	} else {
		hasSell = true
	}

	if !hasBuy && !hasSell {
		return nil, errors.New("不是需要处理的币对")
	}
	if hasBuy && hasSell {
		return nil, errors.New("不是需要处理的币对")
	}
	//买卖
	var amount *big.Int
	t0, _ := token.NewToken(token0, client)
	t0symbol, _ := t0.Symbol(nil)
	t1, _ := token.NewToken(token1, client)
	//t1symbol, _ := t1.Symbol(nil)
	//账号
	privateKey, err := crypto.HexToECDSA(s.PrivateKey)
	if err != nil {
		return nil, errors.New(strconv.FormatInt(s.ID, 10) + " 私钥解析失败")
	}
	account := crypto.PubkeyToAddress(privateKey.PublicKey)

	pancakeSwapRouterV2Address := t.To
	//pancake, err := pancakeSwapRouter.NewPancakeSwapRouter(pancakeSwapRouterV2Address, client)
	if err != nil {
		return nil, err
	}
	//
	if hasBuy {
		//判断用户是否开启防貔貅
		if s.AvoidHoneyPot == 1 {
			if helpers.IsHoneypotToken(token1, s.ChainId, rds) {
				return nil, errors.New("设置了防蜜罐,该代币疑蜜罐,跳过...")
			}
		}

		//是否处于观察期
		ready.IsBuy = true
		isNot, err := rds.Exists(context.Background(), fmt.Sprintf("OB:%d:%s", s.ID, token1.String())).Result()
		if err != nil {
			return nil, errors.New(fmt.Sprintf("读取 Reids 失败..."))
		}
		if isNot > 0 {
			return nil, errors.New("该代币超过止盈线,目前处于观察期...")
		}

		//百分比买
		beforeCall := new(bind.CallOpts)
		beforeCall.BlockNumber = new(big.Int).Sub(t.BlockNumber, new(big.Int).SetInt64(1))
		before, _ := t0.BalanceOf(beforeCall, t.From)
		after, _ := t0.BalanceOf(nil, t.From)
		//useAmount, per := helpers.PercentageDifference(before, after)
		_, per := helpers.PercentageDifference(before, after)
		//log.Printf("买入前 (%s) 余额 %v, 买入后 (%s) 余额 %v , 使用 %v, 系数: %v\n", t0symbol, before, t0symbol, after, useAmount, per)
		// 如果设置从 0 开始跟的币种
		if s.IsZero == 1 {
			// 自己以前有没有买过
			var purchase PurchaseInfo
			purchase.UserTaskId = fmt.Sprintf("%d", s.ID)
			purchase.TokenAddress = strings.ToLower(token1.String())
			isBuy, err := purchase.HasPurchasedToken(rds)
			if err != nil {
				return nil, errors.New("redis 获取买入记录失败")
			}
			// 如果没买过
			if !isBuy {
				// 判断 策略交易者 是不是第一次购买
				beforeAmount, _ := t1.BalanceOf(beforeCall, t.From)
				if beforeAmount.Cmp(new(big.Int).SetInt64(0)) > 0 {
					return nil, errors.New("设置从最新代币开始跟单,已不是...跳过")
				}
			}
			// 买过则继续执行
		}
		myT0balance, _ := t0.BalanceOf(nil, account)
		// 判断使用金额
		// 找出设置金额
		settingAmount := new(big.Int)
		settingMinAmount := new(big.Int)
		settingMaxAmount := new(big.Int)
		// 如果是主网币 换算为主网币的价格
		if token0 == *mainAddress {
			priceF, _, err := helpers.GetPrice(pancakeSwapRouterV2Address, *mainAddress, *usdAddress, client)
			if err != nil {
				return nil, err
			}
			priceInt := new(big.Int)
			priceF.Int(priceInt)
			//设置的金额 / priceInt
			settingAmount.Quo(new(big.Int).Mul(s.Amount, big.NewInt(1e18)), priceInt)
			settingMinAmount.Quo(new(big.Int).Mul(s.MinAmount, big.NewInt(1e18)), priceInt)
			settingMaxAmount.Quo(new(big.Int).Mul(s.MaxAmount, big.NewInt(1e18)), priceInt)
			// 否则，VIP 使用美元金额
		} else {
			if s.IsVip == 1 {
				t0decimal, _ := t0.Decimals(nil)
				t0Exp := new(big.Int).Exp(big.NewInt(10), new(big.Int).SetUint64(uint64(t0decimal)), nil)

				settingAmount.Mul(s.Amount, t0Exp)
				settingMinAmount.Mul(s.MinAmount, t0Exp)
				settingMaxAmount.Mul(s.MaxAmount, t0Exp)
			}
		}
		// 如果是固定金额买入
		if s.BuyType == 0 {
			amount = settingAmount
		} else {
			//百分比
			amount = helpers.MultiplyIntFloat(myT0balance, per)
			// 小于最小
			if amount.Cmp(settingMinAmount) < 0 {
				if s.MinAmountAction == 0 {
					//放弃购买
					return nil, errors.New("小于最小金额，设置为放弃购买...")
				} else {
					amount = settingMinAmount
				}
			} else if amount.Cmp(settingMaxAmount) > 0 {
				if s.MaxAmountAction == 0 {
					return nil, errors.New("大于最大金额, 设置为放弃购买...")
				} else {
					amount = settingMaxAmount
				}
			}
		}
		//如果余额不足，跳过
		if myT0balance.Cmp(amount) < 0 {
			return nil, errors.New(fmt.Sprintf("(%s) 余额不足 %v , 跳过...", t0symbol, amount))
		}
	} else if hasSell {
		//百分比卖
		beforeCall := new(bind.CallOpts)
		beforeCall.BlockNumber = new(big.Int).Sub(t.BlockNumber, new(big.Int).SetInt64(1))
		before, _ := t0.BalanceOf(beforeCall, t.From)
		after, _ := t0.BalanceOf(nil, t.From)
		_, per := helpers.PercentageDifference(before, after)
		//log.Printf("卖出前 (%s) 余额 %v, 卖出后 (%s) 余额 %v , 系数: %v\n", t0symbol, before, t0symbol, after, per)

		myT0balance, _ := t0.BalanceOf(nil, account)

		amount = helpers.MultiplyIntFloat(myT0balance, per)

		// 如果剩余的代币价值小于等于设置的底线。则全部卖出 如果 amount 大于 0 才会
		//if amount.Cmp(big.NewInt(0)) > 0 {
		//	surplus := new(big.Int).Sub(myT0balance, amount)
		//	surplusAmountsOut, err := pancake.GetAmountsOut(nil, surplus, []common.Address{token0, token1})
		//	if err != nil {
		//		return nil, errors.New("获取剩余代币价值错误:" + err.Error())
		//	}
		//	if surplusAmountsOut[1].Cmp(baseArr[kSell].Line) != 1 {
		//		amount = myT0balance
		//	}
		//}
	} else {
		amount = new(big.Int).SetInt64(0)
	}
	//如果余额小于等于0 跳过
	if amount.Cmp(new(big.Int).SetInt64(0)) <= 0 {
		return nil, errors.New("使用零交易,跳过")
	}

	ready.ChainId = s.ChainId.String()
	ready.Amount = amount.String()
	ready.Account = account.String()
	ready.DefiAddress = pancakeSwapRouterV2Address.String()
	ready.Token0 = token0.String()
	ready.Token1 = token1.String()
	ready.SettingId = strconv.FormatInt(s.ID, 10)
	return &ready, nil
}

// 止盈交易  price 现在价格 purchase 购买信息 token 要卖出的代币
func (s *UserTask) StopWinTransaction(price *big.Float, purchase PurchaseInfo, token *token.Token) (*ReadyTrade, error) {
	var ready ReadyTrade
	ready.IsBuy = false
	ready.Type = 2
	beforePrice, ok := new(big.Float).SetString(purchase.PurchasePrice)
	if !ok {
		return nil, errors.New("string 转换 bigFloat 价格 失败")
	}
	// 如果设置了止盈
	if s.StopWin > 0 {
		// 指定涨幅百分比
		percentage := new(big.Float).SetInt64(s.StopWin)
		// 将涨幅百分比转换为小数
		divisor := big.NewFloat(100)
		percentageQuotient := new(big.Float).Quo(percentage, divisor)
		multiplier := new(big.Float).Add(big.NewFloat(1), percentageQuotient)

		// 计算 结果
		targetPrice := new(big.Float).Mul(beforePrice, multiplier)
		//高于 stopWin% 价格.全部卖出,并设置有过期时限的key.禁止买入
		if price.Cmp(targetPrice) >= 0 {
			//账号
			privateKey, err := crypto.HexToECDSA(s.PrivateKey)
			if err != nil {
				return nil, errors.New("Private key parse fail")
			}
			account := crypto.PubkeyToAddress(privateKey.PublicKey)
			amount, err := token.BalanceOf(nil, account)
			if err != nil {
				return nil, err
			}
			if amount.Cmp(big.NewInt(0)) < 1 {
				return nil, errors.New("没有该代币,跳过...")
			}
			if err != nil {
				return nil, err
			}
			ready.ChainId = s.ChainId.String()
			ready.Amount = amount.String()
			ready.Account = account.String()
			ready.DefiAddress = purchase.DefiAddress
			ready.Token0 = purchase.TokenAddress
			ready.Token1 = purchase.BaseAddress
			ready.SettingId = strconv.FormatInt(s.ID, 10)
			return &ready, nil
		}
	}
	// 如果设置了止损
	if s.StopLost > 0 {
		// 指定涨幅百分比
		percentage := new(big.Float).SetInt64(s.StopLost)
		// 将涨幅百分比转换为小数
		divisor := big.NewFloat(100)
		percentageQuotient := new(big.Float).Quo(percentage, divisor)
		multiplier := new(big.Float).Sub(big.NewFloat(1), percentageQuotient)

		// 计算 结果
		targetPrice := new(big.Float).Mul(beforePrice, multiplier)
		//低于 stopLost% 价格.全部卖出,并设置有过期时限的key.禁止买入
		if price.Cmp(targetPrice) < 0 {
			//账号
			privateKey, err := crypto.HexToECDSA(s.PrivateKey)
			if err != nil {
				return nil, errors.New("Private key parse fail")
			}
			account := crypto.PubkeyToAddress(privateKey.PublicKey)
			amount, err := token.BalanceOf(nil, account)
			if err != nil {
				return nil, err
			}
			if amount.Cmp(big.NewInt(0)) < 1 {
				return nil, errors.New("没有该代币,跳过...")
			}
			if err != nil {
				return nil, err
			}
			ready.ChainId = s.ChainId.String()
			ready.Amount = amount.String()
			ready.Account = account.String()
			ready.DefiAddress = purchase.DefiAddress
			ready.Token0 = purchase.TokenAddress
			ready.Token1 = purchase.BaseAddress
			ready.SettingId = strconv.FormatInt(s.ID, 10)
			return &ready, nil
		}
	}

	return nil, errors.New("价格未过止盈或止损点")
}

// 全部卖出- 停止跟单
func (s *UserTask) SellAllTransaction(defiAddress common.Address, rdb *redis.Client, client *ethclient.Client) ([]*ReadyTrade, error) {
	//获取用户所有买入的币
	pairs, err := GetAllPurchasedTokenPairsForUser(rdb, fmt.Sprintf("%d", s.ID))
	if err != nil {
		return nil, err
	}
	var readies []*ReadyTrade
	if err != nil {
		return nil, err
	}
	for _, pair := range pairs {
		var ready ReadyTrade
		t0, _ := token.NewToken(common.HexToAddress(pair.TokenAddress), client)
		//账号
		privateKey, err := crypto.HexToECDSA(s.PrivateKey)
		if err != nil {
			return nil, errors.New("Private key parse fail")
		}
		account := crypto.PubkeyToAddress(privateKey.PublicKey)
		amount, err := t0.BalanceOf(nil, account)
		if err != nil {
			return nil, err
		}
		if amount.Cmp(big.NewInt(0)) < 1 {
			continue
		}
		ready.ChainId = s.ChainId.String()
		ready.Amount = amount.String()
		ready.Account = account.String()
		ready.DefiAddress = defiAddress.String()
		ready.Token0 = pair.TokenAddress
		ready.Token1 = pair.BaseAddress
		ready.SettingId = strconv.FormatInt(s.ID, 10)
		ready.Type = 3
		ready.IsBuy = false
		readies = append(readies, &ready)
	}
	return readies, nil
}

func MapToUserTask(setting map[string]string) (*UserTask, error) {
	chainID, ok := new(big.Int).SetString(setting["chain_id"], 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse chain_id: %s", setting["chain_id"])
	}
	userId, err := strconv.ParseInt(setting["user_id"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user_id: %s", setting["user_id"])
	}
	id, err := strconv.ParseInt(setting["id"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse id: %s", setting["id"])
	}
	buyType, err := strconv.ParseInt(setting["buy_type"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse buy_type: %s", setting["buy_type"])
	}
	amount, ok := new(big.Int).SetString(setting["amount"], 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse amount: %s", setting["amount"])
	}
	minAmountAction, err := strconv.ParseInt(setting["min_amount_action"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse min_amount_action: %s", setting["min_amount_action"])
	}
	minAmount, ok := new(big.Int).SetString(setting["min_amount"], 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse min_amount: %s", setting["min_amount"])
	}
	maxAmountAction, err := strconv.ParseInt(setting["max_amount_action"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse max_amount_action: %s", setting["max_amount_action"])
	}
	maxAmount, ok := new(big.Int).SetString(setting["max_amount"], 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse max_amount: %s", setting["max_amount"])
	}
	stopWin, err := strconv.ParseInt(setting["stop_win"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse stop_win: %s", setting["stop_win"])
	}
	stopWinMin, err := strconv.ParseInt(setting["stop_win_min"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse stop_win_min: %s", setting["stop_win_min"])
	}
	stopLost, err := strconv.ParseInt(setting["stop_lost"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse stop_lost: %s", setting["stop_lost"])
	}
	stopLostMin, err := strconv.ParseInt(setting["stop_lost_min"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse stop_lost_min: %s", setting["stop_lost_min"])
	}
	stopAt, err := strconv.ParseInt(setting["stop_at"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse stop_at: %s", setting["stop_at"])
	}
	isSellAll, err := strconv.ParseInt(setting["is_sell_all"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse is_sell_all: %s", setting["is_sell_all"])
	}
	isSell, err := strconv.ParseInt(setting["is_sell"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse is_sell: %s", setting["is_sell"])
	}
	avoidHoneyPot, err := strconv.ParseInt(setting["avoid_honey_pot"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse avoid_honey_pot: %s", setting["avoid_honey_pot"])
	}
	isZero, err := strconv.ParseInt(setting["is_zero"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse is_zero: %s", setting["is_zero"])
	}
	preventSandwichAttacks, err := strconv.ParseInt(setting["prevent_sandwich_attacks"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse prevent_sandwich_attacks: %s", setting["prevent_sandwich_attacks"])
	}

	userTask := &UserTask{
		ID:                     id,
		UserId:                 userId,
		Name:                   setting["name"],
		ChainId:                chainID,
		FollowAddress:          common.HexToAddress(setting["follow_address"]),
		BuyType:                buyType,
		Amount:                 amount,
		MinAmountAction:        minAmountAction,
		MinAmount:              minAmount,
		MaxAmountAction:        maxAmountAction,
		MaxAmount:              maxAmount,
		StopWin:                stopWin,
		StopWinMin:             stopWinMin,
		StopLost:               stopLost,
		StopLostMin:            stopLostMin,
		StopAt:                 stopAt,
		IsSellAll:              isSellAll,
		IsSell:                 isSell,
		AvoidHoneyPot:          avoidHoneyPot,
		IsZero:                 isZero,
		PreventSandwichAttacks: preventSandwichAttacks,
		PrivateKey:             setting["private_key"],
	}
	return userTask, nil
}
