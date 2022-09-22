/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-10-21 21:28:28
 * @LastEditTime: 2022-07-21 14:31:41
 * @FilePath: \PithyGo\service\redis.go
 */
package service

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var RedisClient *redis.Client

func NewRedisClient() {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConns"),
	})

	pong, err := RedisClient.Ping().Result()
	if err == nil {
		LOG.Sugar().Info("初始化Redis成功:", pong)
	} else {
		LOG.Sugar().Error("初始化Redis失败:", err)
	}
	// Output: PONG <nil>
}
