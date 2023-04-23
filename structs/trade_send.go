package structs

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type TradeSend struct {
	Transaction *types.Transaction
	Token0      common.Address
	Token1      common.Address
	DefiAddress common.Address
	Price       *big.Float
	Setting     *UserBotSetting
	IsBuy       bool
	Error       error
}
