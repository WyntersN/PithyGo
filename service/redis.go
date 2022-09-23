/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-10-21 21:28:28
 * @LastEditTime: 2022-09-23 13:59:36
 * @FilePath: \PithyGo\service\redis.go
 */
package service

import (
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func NewRedisClient() {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:         CONFIG.Redis.Addr,
		Password:     CONFIG.Redis.Password,
		DB:           CONFIG.Redis.DB,
		PoolSize:     CONFIG.Redis.PoolSize,
		MinIdleConns: CONFIG.Redis.MinIdleConns,
	})

	pong, err := RedisClient.Ping().Result()
	if err == nil {
		LOG.Sugar().Info("初始化Redis成功:", pong)
	} else {
		LOG.Sugar().Error("初始化Redis失败:", err)
	}
	// Output: PONG <nil>
}
