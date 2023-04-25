package structs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type PurchaseInfo struct {
	UserBotSettingID string `json:"user_bot_setting_id"`
	TokenAddress     string `json:"token_address"`
	BaseAddress      string `json:"base_address"`
	PurchasePrice    string `json:"purchase_price"`
	Time             int64  `json:"time"`
}

func (purchase *PurchaseInfo) getKey() string {
	return fmt.Sprintf("SettingFirstPurchase:%s", purchase.UserBotSettingID)
}

func (purchase *PurchaseInfo) SavePurchaseInfoToRedis(rdb *redis.Client) error {
	ctx := context.Background()

	userKey := purchase.getKey()

	// 检查用户是否已经购买过这个代币
	exists, err := rdb.HExists(ctx, userKey, purchase.TokenAddress).Result()
	if err != nil {
		return err
	}

	// 如果用户已经购买过这个代币，直接返回
	if exists {
		return nil
	}

	jsonData, err := json.Marshal(purchase)
	if err != nil {
		return err
	}

	// 将用户购买信息存储到 Redis 散列中
	err = rdb.HSet(ctx, userKey, purchase.TokenAddress, jsonData).Err()
	if err != nil {
		return err
	}

	return nil
}

func (purchase *PurchaseInfo) GetAllTokens(rdb *redis.Client) ([]string, error) {
	ctx := context.Background()

	userKey := purchase.getKey()

	tokenAddresses, err := rdb.HKeys(ctx, userKey).Result()
	if err != nil {
		return nil, err
	}

	return tokenAddresses, nil
}
