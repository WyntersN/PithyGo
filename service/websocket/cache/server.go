/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-10-21 21:38:38
 * @LastEditTime: 2022-07-21 14:25:11
 * @FilePath: \PithyGo\service\websocket\cache\server.go
 */
package cache

import (
	"encoding/json"
	"fmt"

	"PithyGo/service"
	"PithyGo/service/websocket/models"
	"strconv"
)

const (
	serversHashKey       = "acc:hash:servers" // 全部的服务器
	serversHashCacheTime = 1 * 60 * 60        // key过期时间
	serversHashTimeout   = 1 * 60             // 超时时间
)

func getServersHashKey() (key string) {
	key = fmt.Sprintf("%s", serversHashKey)

	return
}

// 设置服务器信息
func SetServerInfo(server *models.Server, currentTime uint64) (err error) {
	key := getServersHashKey()

	value := fmt.Sprintf("%d", currentTime)

	number, err := service.RedisClient.Do("hSet", key, server.String(), value).Int()
	if err != nil {
		service.LOG.Sugar().Info("SetServerInfo", key, number, err)

		return
	}

	if number != 1 {

		return
	}

	service.RedisClient.Do("Expire", key, serversHashCacheTime)

	return
}

// 下线服务器信息
func DelServerInfo(server *models.Server) (err error) {
	key := getServersHashKey()
	number, err := service.RedisClient.Do("hDel", key, server.String()).Int()
	if err != nil {
		service.LOG.Sugar().Info("DelServerInfo", key, number, err)

		return
	}

	if number != 1 {

		return
	}

	service.RedisClient.Do("Expire", key, serversHashCacheTime)

	return
}

func GetServerAll(currentTime uint64) (servers []*models.Server, err error) {

	servers = make([]*models.Server, 0)
	key := getServersHashKey()

	val, err := service.RedisClient.Do("hGetAll", key).Result()

	valByte, _ := json.Marshal(val)
	service.LOG.Sugar().Info("GetServerAll", key, string(valByte))

	serverMap, err := service.RedisClient.HGetAll(key).Result()
	if err != nil {
		service.LOG.Sugar().Info("SetServerInfo", key, err)

		return
	}

	for key, value := range serverMap {
		valueUint64, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			service.LOG.Sugar().Info("GetServerAll", key, err)

			return nil, err
		}

		// 超时
		if valueUint64+serversHashTimeout <= currentTime {
			continue
		}

		server, err := models.StringToServer(key)
		if err != nil {
			service.LOG.Sugar().Info("GetServerAll", key, err)

			return nil, err
		}

		servers = append(servers, server)
	}

	return
}
