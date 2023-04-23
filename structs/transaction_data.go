package structs

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// 交易
type TransactionData struct {
	ChainId     *big.Int
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
	f := "ChainId:%s\nHash: %s\nFrom: %s\nTo: %s\nValue: %s\nGas: %d\nGasPrice: %s\nData: 0x%x\n"
	return fmt.Sprintf(f, t.ChainId.String(), t.Hash.String(), t.From.String(), t.To.String(), t.Value.String(), t.Gas, t.GasPrice.String(), string(t.Data))
}

func (t *TransactionData) UnmarshalJSON(data []byte) error {
	aux := &struct {
		ChainId     string `json:"chainId"`
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
	chainId, ok := new(big.Int).SetString(aux.ChainId, 10)
	if !ok {
		return fmt.Errorf("failed to unmarshal ChainId from string: %s", aux.ChainId)
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
	t.ChainId = chainId
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
		ChainId     string `json:"chainId"`
		Hash        string `json:"hash"`
		From        string `json:"from"`
		To          string `json:"to"`
		Value       string `json:"value"`
		Gas         uint64 `json:"gas"`
		GasPrice    string `json:"gasPrice"`
		Data        string `json:"data"`
		BlockNumber string `json:"blockNumber"`
	}{
		ChainId:     t.ChainId.String(),
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
