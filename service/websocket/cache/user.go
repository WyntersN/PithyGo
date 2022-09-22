/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-10-21 21:36:55
 * @LastEditTime: 2022-08-29 14:55:16
 * @FilePath: \PithyGo\service\websocket\cache\user.go
 */
package cache

import (
	"encoding/json"
	"fmt"

	"PithyGo/service"
	"PithyGo/service/websocket/models"

	"github.com/go-redis/redis"
)

const (
	userOnlinePrefix    = "acc:user:online:" // 用户在线状态
	userOnlineCacheTime = 180
)

/*********************  查询用户是否在线  ************************/
func getUserOnlineKey(userKey string) (key string) {
	key = fmt.Sprintf("%s%s", userOnlinePrefix, userKey)

	return
}

func GetUserOnlineInfo(userKey string) (userOnline *models.UserOnline, err error) {

	key := getUserOnlineKey(userKey)

	data, err := service.RedisClient.Get(key).Bytes()
	if err != nil {
		if err == redis.Nil {
			service.LOG.Sugar().Info("GetUserOnlineInfo", userKey, err)

			return
		}

		service.LOG.Sugar().Info("GetUserOnlineInfo", userKey, err)

		return
	}

	userOnline = &models.UserOnline{}
	err = json.Unmarshal(data, userOnline)
	if err != nil {
		service.LOG.Sugar().Info("获取用户在线数据 json Unmarshal", userKey, err)

		return
	}

	service.LOG.Sugar().Info("获取用户在线数据", userKey, "time", userOnline.LoginTime, userOnline.HeartbeatTime, "AccIp", userOnline.AccIp, userOnline.IsLogoff)

	return
}

// 设置用户在线数据
func SetUserOnlineInfo(userKey string, userOnline *models.UserOnline) (err error) {

	key := getUserOnlineKey(userKey)

	valueByte, err := json.Marshal(userOnline)
	if err != nil {
		service.LOG.Sugar().Info("设置用户在线数据 json Marshal", key, err)

		return
	}

	_, err = service.RedisClient.Do("setEx", key, userOnlineCacheTime, string(valueByte)).Result()
	if err != nil {
		service.LOG.Sugar().Info("设置用户在线数据 ", key, err)

		return
	}

	return
}
