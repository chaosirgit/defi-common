package structs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strings"
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

// 第一次买入，存入 redis
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

	// 将 UserBotSettingID 添加到 Redis 集合中
	tokenKey := fmt.Sprintf("FirstPurchaseToken:%s", strings.ToLower(purchase.TokenAddress))
	err = rdb.SAdd(ctx, tokenKey, purchase.UserBotSettingID).Err()
	if err != nil {
		return err
	}

	return nil
}

// 获取 该用户下的所有购买的币
func (purchase *PurchaseInfo) GetAllTokens(rdb *redis.Client) ([]string, error) {
	ctx := context.Background()

	userKey := purchase.getKey()

	tokenAddresses, err := rdb.HKeys(ctx, userKey).Result()
	if err != nil {
		return nil, err
	}

	return tokenAddresses, nil
}

// 获取平台所有已买入的代币
func GetAllUniqueTokensBought(rdb *redis.Client) ([]string, error) {
	ctx := context.Background()

	// 使用一个 map 来存储去重后的代币地址
	uniqueTokens := make(map[string]bool)

	// 获取所有以 "FirstPurchaseToken:" 开头的 key
	keys, err := rdb.Keys(ctx, "FirstPurchaseToken:*").Result()
	if err != nil {
		return nil, err
	}

	// 遍历所有的 FirstPurchaseToken:<TokenAddress> key
	for _, key := range keys {
		// 从 key 中提取 TokenAddress
		tokenAddress := strings.TrimPrefix(key, "FirstPurchaseToken:")

		// 将 TokenAddress 添加到 uniqueTokens map 中
		uniqueTokens[tokenAddress] = true
	}

	// 将 uniqueTokens map 转换为字符串切片
	result := make([]string, 0, len(uniqueTokens))
	for tokenAddress := range uniqueTokens {
		result = append(result, tokenAddress)
	}

	return result, nil
}

// 获取某个 代币下的所有买入用户
func GetUserBotSettingIDsForToken(rdb *redis.Client, tokenAddress string) ([]string, error) {
	ctx := context.Background()

	// 根据 tokenAddress 构建 Redis 集合 key
	tokenKey := fmt.Sprintf("FirstPurchaseToken:%s", tokenAddress)

	// 从 Redis 集合中获取所有 UserBotSettingID
	userBotSettingIDs, err := rdb.SMembers(ctx, tokenKey).Result()
	if err != nil {
		return nil, err
	}

	return userBotSettingIDs, nil
}
