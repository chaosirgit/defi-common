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

// 用户配置
type UserBotSetting struct {
	ID              int64          `json:"id"`
	WhaleAddress    common.Address `json:"whale_address"`
	WhaleName       string         `json:"whale_name"`
	ChainID         *big.Int       `json:"chain_id"`
	ListenBase      string         `json:"listen_base"`
	MinWethAmount   *big.Int       `json:"min_weth_amount"`
	MinBusdAmount   *big.Int       `json:"min_busd_amount"`
	MinUsdtAmount   *big.Int       `json:"min_usdt_amount"`
	MinUsdcAmount   *big.Int       `json:"min_usdc_amount"`
	SellAllWethLine *big.Int       `json:"sell_all_weth_line"`
	SellAllBusdLine *big.Int       `json:"sell_all_busd_line"`
	SellAllUsdtLine *big.Int       `json:"sell_all_usdt_line"`
	SellAllUsdcLine *big.Int       `json:"sell_all_usdc_line"`
	StopAt          int64          `json:"stop_at"`
	IsSellAll       int8           `json:"is_sell_all"`
	StopWin         int64          `json:"stop_win"`
	StopWinMin      int64          `json:"stop_win_min"`
	AvoidHoneyPot   int8           `json:"avoid_honey_pot"`
	IsZero          int8           `json:"is_zero"`
	PrivateKey      string         `json:"private_key"`
}

func (s *UserBotSetting) UnmarshalJSON(data []byte) error {
	aux := &struct {
		ID              int64  `json:"id"`
		WhaleAddress    string `json:"whale_address"`
		WhaleName       string `json:"whale_name"`
		ChainID         int64  `json:"chain_id"`
		ListenBase      string `json:"listen_base"`
		MinWethAmount   string `json:"min_weth_amount"`
		MinBusdAmount   string `json:"min_busd_amount"`
		MinUsdtAmount   string `json:"min_usdt_amount"`
		MinUsdcAmount   string `json:"min_usdc_amount"`
		SellAllWethLine string `json:"sell_all_weth_line"`
		SellAllBusdLine string `json:"sell_all_busd_line"`
		SellAllUsdtLine string `json:"sell_all_usdt_line"`
		SellAllUsdcLine string `json:"sell_all_usdc_line"`
		StopAt          int64  `json:"stop_at"`
		IsSellAll       int8   `json:"is_sell_all"`
		StopWin         int64  `json:"stop_win"`
		StopWinMin      int64  `json:"stop_win_min"`
		AvoidHoneyPot   int8   `json:"avoid_honey_pot"`
		IsZero          int8   `json:"is_zero"`
		PrivateKey      string `json:"private_key"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	chainId := new(big.Int).SetInt64(aux.ChainID)

	minWeth, ok := new(big.Int).SetString(aux.MinWethAmount, 10)
	if !ok {
		return fmt.Errorf("failed to unmarshal MinWethAmount from string: %s", aux.MinWethAmount)
	}
	minBusd, ok := new(big.Int).SetString(aux.MinBusdAmount, 10)
	if !ok {
		return fmt.Errorf("failed to unmarshal MinBusdAmount from string: %s", aux.MinBusdAmount)
	}
	minUsdt, ok := new(big.Int).SetString(aux.MinUsdtAmount, 10)
	if !ok {
		return fmt.Errorf("failed to unmarshal MinUsdtAmount from string: %s", aux.MinUsdtAmount)
	}
	minUsdc, ok := new(big.Int).SetString(aux.MinUsdcAmount, 10)
	if !ok {
		return fmt.Errorf("failed to unmarshal MinUsdcAmount from string: %s", aux.MinUsdcAmount)
	}
	sellWeth, ok := new(big.Int).SetString(aux.SellAllWethLine, 10)
	if !ok {
		return fmt.Errorf("failed to unmarshal SellAllWethLine from string: %s", aux.SellAllWethLine)
	}
	sellBusd, ok := new(big.Int).SetString(aux.SellAllBusdLine, 10)
	if !ok {
		return fmt.Errorf("failed to unmarshal SellAllBusdLine from string: %s", aux.SellAllBusdLine)
	}
	sellUsdt, ok := new(big.Int).SetString(aux.SellAllUsdtLine, 10)
	if !ok {
		return fmt.Errorf("failed to unmarshal SellAllUsdtLine from string: %s", aux.SellAllUsdtLine)
	}
	sellUsdc, ok := new(big.Int).SetString(aux.SellAllUsdcLine, 10)
	if !ok {
		return fmt.Errorf("failed to unmarshal SellAllUsdcLine from string: %s", aux.SellAllUsdcLine)
	}

	s.ID = aux.ID
	s.WhaleAddress = common.HexToAddress(aux.WhaleAddress)
	s.WhaleName = aux.WhaleName
	s.ChainID = chainId
	s.ListenBase = aux.ListenBase
	s.MinWethAmount = minWeth
	s.MinBusdAmount = minBusd
	s.MinUsdtAmount = minUsdt
	s.MinUsdcAmount = minUsdc
	s.SellAllWethLine = sellWeth
	s.SellAllBusdLine = sellBusd
	s.SellAllUsdtLine = sellUsdt
	s.SellAllUsdcLine = sellUsdc
	s.StopAt = aux.StopAt
	s.IsSellAll = aux.IsSellAll
	s.StopWin = aux.StopWin
	s.StopWinMin = aux.StopWinMin
	s.AvoidHoneyPot = aux.AvoidHoneyPot
	s.IsZero = aux.IsZero
	s.PrivateKey = aux.PrivateKey
	return nil
}

func MapToUserBotSetting(setting map[string]string) (*UserBotSetting, error) {
	chainID, ok := new(big.Int).SetString(setting["chain_id"], 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse chain_id: %s", setting["chain_id"])
	}
	minWethAmount, ok := new(big.Int).SetString(setting["min_weth_amount"], 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse min_weth_amount: %s", setting["min_weth_amount"])
	}
	minBusdAmount, ok := new(big.Int).SetString(setting["min_busd_amount"], 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse min_busd_amount: %s", setting["min_busd_amount"])
	}
	minUsdtAmount, ok := new(big.Int).SetString(setting["min_usdt_amount"], 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse min_usdt_amount from string: %s", setting["min_usdt_amount"])
	}
	minUsdcAmount, ok := new(big.Int).SetString(setting["min_usdc_amount"], 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse min_usdc_amount from string: %s", setting["min_usdc_amount"])
	}
	sellAllWethLine, ok := new(big.Int).SetString(setting["sell_all_weth_line"], 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse sell_all_weth_line from string: %s", setting["sell_all_weth_line"])
	}
	sellAllBusdLine, ok := new(big.Int).SetString(setting["sell_all_busd_line"], 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse sell_all_busd_line from string: %s", setting["sell_all_busd_line"])
	}
	sellAllUsdtLine, ok := new(big.Int).SetString(setting["sell_all_usdt_line"], 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse sell_all_usdt_line from string: %s", setting["sell_all_usdt_line"])
	}
	sellAllUsdcLine, ok := new(big.Int).SetString(setting["sell_all_usdc_line"], 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse sell_all_usdc_line from string: %s", setting["sell_all_usdc_line"])
	}
	stopAt, err := strconv.ParseInt(setting["stop_at"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse stop_at: %s", setting["stop_at"])
	}
	isSellAll, err := strconv.ParseInt(setting["is_sell_all"], 10, 8)
	if err != nil {
		return nil, fmt.Errorf("failed to parse is_sell_all: %s", setting["is_sell_all"])
	}
	stopWin, err := strconv.ParseInt(setting["stop_win"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse stop_win: %s", setting["stop_win"])
	}
	stopWinMin, err := strconv.ParseInt(setting["stop_win_min"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse stop_win_min: %s", setting["stop_win_min"])
	}
	avoidHoneyPot, err := strconv.ParseInt(setting["avoid_honey_pot"], 10, 8)
	if err != nil {
		return nil, fmt.Errorf("failed to parse avoid_honey_pot: %s", setting["avoid_honey_pot"])
	}
	isZero, err := strconv.ParseInt(setting["is_zero"], 10, 8)
	if err != nil {
		return nil, fmt.Errorf("failed to parse is_zero: %s", setting["is_zero"])
	}
	id, err := strconv.ParseInt(setting["id"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse id: %s", setting["id"])
	}
	userBotSetting := &UserBotSetting{
		ID:              id,
		WhaleAddress:    common.HexToAddress(setting["whale_address"]),
		WhaleName:       setting["whale_name"],
		ChainID:         chainID,
		ListenBase:      setting["listen_base"],
		MinWethAmount:   minWethAmount,
		MinBusdAmount:   minBusdAmount,
		MinUsdtAmount:   minUsdtAmount,
		MinUsdcAmount:   minUsdcAmount,
		SellAllWethLine: sellAllWethLine,
		SellAllBusdLine: sellAllBusdLine,
		SellAllUsdtLine: sellAllUsdtLine,
		SellAllUsdcLine: sellAllUsdcLine,
		StopAt:          stopAt,
		IsSellAll:       int8(isSellAll),
		StopWin:         stopWin,
		StopWinMin:      stopWinMin,
		AvoidHoneyPot:   int8(avoidHoneyPot),
		IsZero:          int8(isZero),
		PrivateKey:      setting["private_key"],
	}
	return userBotSetting, nil
}

func (s *UserBotSetting) AnalyzeTransaction(t TransactionData, client *ethclient.Client, rds *redis.Client) (*ReadyTrade, error) {
	var ready ReadyTrade
	ready.IsBuy = false
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
	baseAddress, err := helpers.GetBaseTokenAddressesByChainId(t.ChainId)
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
		token0 = common.HexToAddress(baseAddress["WETH"])
		token1 = pathRouter[len(pathRouter)-1]
	} else if ok, _ = helpers.InArray(method.RawName, TokensForETH); ok {
		pathRouter := args[2].([]common.Address)
		token0 = pathRouter[0]
		token1 = common.HexToAddress(baseAddress["WETH"])

	} else if ok, _ = helpers.InArray(method.RawName, TokensForTokens); ok {
		pathRouter := args[2].([]common.Address)
		token0 = pathRouter[len(pathRouter)-2]
		token1 = pathRouter[len(pathRouter)-1]

	} else if ok, _ = helpers.InArray(method.RawName, RemoveLiquidity); ok {
		token0 = args[0].(common.Address)
		token1 = args[1].(common.Address)
		if is, _ := helpers.InArray(args[0].(common.Address), []common.Address{common.HexToAddress(baseAddress["WETH"]), common.HexToAddress(baseAddress["USDT"]), common.HexToAddress(baseAddress["BUSD"]), common.HexToAddress(baseAddress["USDC"])}); is {
			token0 = args[1].(common.Address)
			token1 = args[0].(common.Address)
		}

	} else if ok, _ = helpers.InArray(method.RawName, RemoveETHLiquidity); ok {
		token0 = args[0].(common.Address)
		token1 = common.HexToAddress(baseAddress["WETH"])
	} else if ok, _ = helpers.InArray(method.RawName, addLiquidity); ok {
		return nil, errors.New("添加流动性,暂时跳过")
	} else if ok, _ = helpers.InArray(method.RawName, addETHLiquidity); ok {
		return nil, errors.New("移除流动性,暂时跳过")
	}

	//确定买卖以及金额
	type A struct {
		Address common.Address
		Amount  *big.Int
		Line    *big.Int
	}
	baseArr := []A{}
	for i := 0; i < len(s.ListenBase); i++ {
		char := s.ListenBase[i]
		if char == byte('1') {
			if i == 0 {
				baseArr = append(baseArr, A{Address: common.HexToAddress(baseAddress["WETH"]), Amount: s.MinWethAmount, Line: s.SellAllWethLine})
			} else if i == 1 {
				baseArr = append(baseArr, A{Address: common.HexToAddress(baseAddress["BUSD"]), Amount: s.MinBusdAmount, Line: s.SellAllBusdLine})
			} else if i == 2 {
				baseArr = append(baseArr, A{Address: common.HexToAddress(baseAddress["USDT"]), Amount: s.MinUsdtAmount, Line: s.SellAllUsdtLine})
			} else if i == 3 {
				baseArr = append(baseArr, A{Address: common.HexToAddress(baseAddress["USDC"]), Amount: s.MinUsdcAmount, Line: s.SellAllUsdcLine})
			}
		}
	}
	var kBuy int
	var kSell int
	var hasBuy bool
	var hasSell bool
	for index, b := range baseArr {
		if token0 == b.Address {
			kBuy = index
			hasBuy = true
			break
		}
	}
	for index, b := range baseArr {
		if token1 == b.Address {
			kSell = index
			hasSell = true
			break
		}
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

	pancakeSwapRouterV2Address := common.HexToAddress(baseAddress["DEFI"])
	pancake, err := pancakeSwapRouter.NewPancakeSwapRouter(pancakeSwapRouterV2Address, client)
	if err != nil {
		return nil, err
	}
	//
	if hasBuy {
		//TODO 判断用户是否开启防貔貅

		//TODO 是否 40 卖过的
		ready.IsBuy = true
		isNot, err := rds.Exists(context.Background(), fmt.Sprintf("notBuy-%d-%s", s.ID, token1.String())).Result()
		if err != nil {
			return nil, errors.New(fmt.Sprintf("读取 Reids 失败..."))
		}
		if isNot > 0 {
			return nil, errors.New("该代币超过止盈线,目前处于观察期..")
		}

		//百分比买
		beforeCall := new(bind.CallOpts)
		beforeCall.BlockNumber = new(big.Int).Sub(t.BlockNumber, new(big.Int).SetInt64(1))
		before, _ := t0.BalanceOf(beforeCall, t.From)
		after, _ := t0.BalanceOf(nil, t.From)
		useAmount, per := helpers.PercentageDifference(before, after)
		//log.Printf("买入前 (%s) 余额 %v, 买入后 (%s) 余额 %v , 使用 %v, 系数: %v\n", t0symbol, before, t0symbol, after, useAmount, per)
		// 如果设置从 0 开始跟的币种
		if s.IsZero == int8(1) {
			// 是不是以前买过的
			isBuyer, err := rds.Exists(context.Background(), fmt.Sprintf("buyTokens-%d-%s", s.ID, strings.ToLower(token1.String()))).Result()
			if err != nil {
				return nil, err
			}
			// 自己以前没买过才判断大哥是否买过,已买过直接跟
			if isBuyer == 0 {
				bt1call := new(bind.CallOpts)
				bt1call.BlockNumber = new(big.Int).Sub(t.BlockNumber, new(big.Int).SetInt64(1))
				beforeT1, _ := t1.BalanceOf(bt1call, t.From)
				if beforeT1.Cmp(big.NewInt(0)) > 0 {
					return nil, errors.New("设置从最新代币开始跟单,已不是...跳过")
				}
			}
		}
		myT0balance, _ := t0.BalanceOf(nil, account)

		if useAmount.Cmp(baseArr[kBuy].Amount) < 0 {
			amount = useAmount
		} else {
			amount = helpers.MultiplyIntFloat(myT0balance, per)
			if amount.Cmp(baseArr[kBuy].Amount) < 0 {
				amount = baseArr[kBuy].Amount
			}
		}

		//如果余额小于最小余额，跳过
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
		if amount.Cmp(big.NewInt(0)) > 0 {
			surplus := new(big.Int).Sub(myT0balance, amount)
			surplusAmountsOut, err := pancake.GetAmountsOut(nil, surplus, []common.Address{token0, token1})
			if err != nil {
				return nil, errors.New("获取剩余代币价值错误:" + err.Error())
			}
			if surplusAmountsOut[1].Cmp(baseArr[kSell].Line) != 1 {
				amount = myT0balance
			}
		}
	} else {
		amount = new(big.Int).SetInt64(0)
	}
	//如果余额小于等于0 跳过
	if amount.Cmp(new(big.Int).SetInt64(0)) <= 0 {
		return nil, errors.New("使用零交易,跳过")
	}

	ready.ChainId = s.ChainID.String()
	ready.Amount = amount.String()
	ready.Account = account.String()
	ready.DefiAddress = pancakeSwapRouterV2Address.String()
	ready.Token0 = token0.String()
	ready.Token1 = token1.String()
	ready.SettingId = strconv.FormatInt(s.ID, 10)
	return &ready, nil
}

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

func GetUserBotSettingById(id string, rds *redis.Client) (*UserBotSetting, error) {
	settingKey := "FollowSetting:" + id
	setting, err := rds.HGetAll(context.Background(), settingKey).Result()
	if err != nil {
		return nil, err
	}
	us, err := MapToUserBotSetting(setting)
	if err != nil {
		return nil, err
	}
	return us, nil
}
