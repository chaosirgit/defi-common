package structs

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// 交易
type TransactionData struct {
	Hash        common.Hash
	From        common.Address
	To          common.Address
	Value       *big.Int
	Gas         uint64
	GasPrice    *big.Int
	Data        string
	BlockNumber *big.Int
}

func (t *TransactionData) ToMessage() string {
	f := "Hash: %s\nFrom: %s\nTo: %s\nValue: %s\nGas: %d\nGasPrice: %s\nData: 0x%x\n"
	return fmt.Sprintf(f, t.Hash.String(), t.From.String(), t.To.String(), t.Value.String(), t.Gas, t.GasPrice.String(), string(t.Data))
}

func (t *TransactionData) UnmarshalJSON(data []byte) error {
	aux := &struct {
		Hash        string `json:"hash"`
		From        string `json:"from"`
		To          string `json:"to"`
		Value       string `json:"value"`
		Gas         uint64 `json:"gas"`
		GasPrice    string `json:"gasPrice"`
		Data        string `json:"data"`
		BlockNumber string `json:"blockNumber"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	value, ok := new(big.Int).SetString(aux.Value, 10)
	if !ok {
		return fmt.Errorf("failed to unmarshal Value from string: %s", aux.Value)
	}
	gasPrice, ok := new(big.Int).SetString(aux.GasPrice, 10)
	if !ok {
		return fmt.Errorf("failed to unmarshal GasPrice from string: %s", aux.GasPrice)
	}
	blockNumber, ok := new(big.Int).SetString(aux.BlockNumber, 10)
	if !ok {
		return fmt.Errorf("failed to unmarshal BlockNumber from string: %s", aux.BlockNumber)
	}
	t.Hash = common.HexToHash(aux.Hash)
	t.From = common.HexToAddress(aux.From)
	t.To = common.HexToAddress(aux.To)
	t.Value = value
	t.Gas = aux.Gas
	t.GasPrice = gasPrice
	t.Data = aux.Data
	t.BlockNumber = blockNumber
	return nil
}

func (t TransactionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Hash        string `json:"hash"`
		From        string `json:"from"`
		To          string `json:"to"`
		Value       string `json:"value"`
		Gas         uint64 `json:"gas"`
		GasPrice    string `json:"gasPrice"`
		Data        string `json:"data"`
		BlockNumber string `json:"blockNumber"`
	}{
		Hash:        t.Hash.String(),
		From:        t.From.String(),
		To:          t.To.String(),
		Value:       t.Value.String(),
		Gas:         t.Gas,
		GasPrice:    t.GasPrice.String(),
		Data:        t.Data,
		BlockNumber: t.BlockNumber.String(),
	})
}

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
