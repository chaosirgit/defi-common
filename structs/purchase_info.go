package structs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strings"
)

type PurchaseInfo struct {
	UserTaskId    string `json:"user_task_id"`
	TokenAddress  string `json:"token_address"`
	BaseAddress   string `json:"base_address"`
	DefiAddress   string `json:"defi_address"`
	ChainId       string `json:"chain_id"`
	PurchasePrice string `json:"purchase_price"`
	Symbol        string `json:"symbol"`
	Decimals      int64  `json:"decimals"`
	Time          int64  `json:"time"`
}

func (purchase *PurchaseInfo) getKey() string {
	return fmt.Sprintf("SettingFirstPurchase:%s", purchase.UserTaskId)
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
	tokenKey := fmt.Sprintf("FirstPurchaseToken:%s/%s", strings.ToLower(purchase.TokenAddress), strings.ToLower(purchase.BaseAddress))
	err = rdb.SAdd(ctx, tokenKey, jsonData).Err()
	if err != nil {
		return err
	}

	return nil
}

// 删除买入信息
func (purchase *PurchaseInfo) RemovePurchaseInfoFromRedis(rdb *redis.Client) error {
	ctx := context.Background()

	userKey := purchase.getKey()

	// 从 Redis 散列中删除用户购买信息
	err := rdb.HDel(ctx, userKey, purchase.TokenAddress).Err()
	if err != nil {
		return err
	}

	// 从 Redis 集合中删除 UserBotSettingID
	tokenKey := fmt.Sprintf("FirstPurchaseToken:%s/%s", strings.ToLower(purchase.TokenAddress), strings.ToLower(purchase.BaseAddress))
	err = rdb.SRem(ctx, tokenKey, purchase.UserTaskId).Err()
	if err != nil {
		return err
	}

	return nil
}

// 获取用户的所有币对儿
func GetAllPurchasedTokenPairsForUser(rdb *redis.Client, userBotSettingID string) ([]TokenPair, error) {
	ctx := context.Background()

	userKey := fmt.Sprintf("SettingFirstPurchase:%s", userBotSettingID)

	// 从 Redis 散列中获取 PurchaseInfo JSON 数据
	jsonDataList, err := rdb.HVals(ctx, userKey).Result()
	if err != nil {
		return nil, err
	}

	tokenPairs := make([]TokenPair, 0, len(jsonDataList))

	// 解析 JSON 数据并将其添加到 tokenPairs 列表中
	for _, jsonData := range jsonDataList {
		var purchase PurchaseInfo
		err = json.Unmarshal([]byte(jsonData), &purchase)
		if err != nil {
			return nil, err
		}
		tokenPairs = append(tokenPairs, TokenPair{
			TokenAddress: purchase.TokenAddress,
			BaseAddress:  purchase.BaseAddress,
		})
	}

	return tokenPairs, nil
}

type TokenPair struct {
	TokenAddress string
	BaseAddress  string
}

// 获取平台所有已买入的代币对
func GetAllDistinctTokenPairs(rdb *redis.Client) ([]TokenPair, error) {
	ctx := context.Background()

	// 获取所有以 FirstPurchaseToken: 开头的键
	keys, err := rdb.Keys(ctx, "FirstPurchaseToken:*").Result()
	if err != nil {
		return nil, err
	}

	// 用于存储唯一 tokenAddress/baseAddress 对的集合
	tokenSet := make(map[string]struct{})

	// 遍历所有键，提取 tokenAddress/baseAddress 并将其添加到集合中
	for _, key := range keys {
		parts := strings.Split(key, ":")
		if len(parts) == 2 {
			tokenSet[parts[1]] = struct{}{}
		}
	}

	// 将唯一 tokenAddress/baseAddress 对从集合转换为 TokenPair 切片
	distinctTokenPairs := make([]TokenPair, 0, len(tokenSet))
	for token := range tokenSet {
		pairParts := strings.Split(token, "/")
		if len(pairParts) == 2 {
			distinctTokenPairs = append(distinctTokenPairs, TokenPair{
				TokenAddress: pairParts[0],
				BaseAddress:  pairParts[1],
			})
		}
	}

	return distinctTokenPairs, nil
}

// 获取某个代币对下的所有买入信息
func GetUserPurchasesByTokenPair(rdb *redis.Client, tokenAddress, baseAddress string) ([]PurchaseInfo, error) {
	ctx := context.Background()

	tokenKey := fmt.Sprintf("FirstPurchaseToken:%s/%s", strings.ToLower(tokenAddress), strings.ToLower(baseAddress))

	// 从 Redis 集合中获取 PurchaseInfo JSON 数据
	jsonDataList, err := rdb.SMembers(ctx, tokenKey).Result()
	if err != nil {
		return nil, err
	}

	purchases := make([]PurchaseInfo, 0, len(jsonDataList))

	// 解析 JSON 数据并将其添加到 purchases 列表中
	for _, jsonData := range jsonDataList {
		var purchase PurchaseInfo
		err = json.Unmarshal([]byte(jsonData), &purchase)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	return purchases, nil
}

func (purchase *PurchaseInfo) HasPurchasedToken(rdb *redis.Client) (bool, error) {
	ctx := context.Background()

	userKey := purchase.getKey()

	// 检查用户是否已经购买过这个代币
	exists, err := rdb.HExists(ctx, userKey, purchase.TokenAddress).Result()
	if err != nil {
		return false, err
	}

	return exists, nil
}
