/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-30
* Time: 12:27
 */

package websocket

import (
	"PithyGo/service"
	"PithyGo/service/websocket/cache"
	"PithyGo/service/websocket/models"
	"PithyGo/service/websocket/servers/grpcclient"
	"errors"
	"strconv"

	"github.com/go-redis/redis"

	"time"
)

// 查询所有用户
func UserList(appId uint32) (userList []string) {

	userList = make([]string, 0)
	currentTime := uint64(time.Now().Unix())
	servers, err := cache.GetServerAll(currentTime)
	if err != nil {
		service.LOG.Sugar().Info("给全体用户发消息", err)

		return
	}

	for _, server := range servers {
		var (
			list []string
		)
		if IsLocal(server) {
			list = GetUserList(appId)
		} else {
			list, _ = grpcclient.GetUserList(server, appId)
		}
		userList = append(userList, list...)
	}

	return
}

// 查询用户是否在线
func CheckUserOnline(appId uint32, userId string) (online bool) {
	// 全平台查询
	if appId == 0 {
		for _, appId := range GetAppIds() {
			online, _ = checkUserOnline(appId, userId)
			if online == true {
				break
			}
		}
	} else {
		online, _ = checkUserOnline(appId, userId)
	}

	return
}

// 查询用户 是否在线
func checkUserOnline(appId uint32, userId string) (online bool, err error) {
	key := GetUserKey(appId, userId)
	userOnline, err := cache.GetUserOnlineInfo(key)
	if err != nil {
		if err == redis.Nil {
			service.LOG.Sugar().Info("GetUserOnlineInfo", appId, userId, err)

			return false, nil
		}

		service.LOG.Sugar().Info("GetUserOnlineInfo", appId, userId, err)

		return
	}

	online = userOnline.IsOnline()

	return
}

// 给用户发送消息
func SendUserMessage(appId uint32, userId uint, msgId, cmd, message string) (sendResults bool, err error) {

	userIdStr := strconv.Itoa(int(userId))
	data := models.GetMsgData(userIdStr, msgId, cmd, message)

	client := GetUserClient(appId, userIdStr)

	if client != nil {
		// 在本机发送
		sendResults, err = SendUserMessageLocal(appId, userIdStr, data)
		if err != nil {
			service.LOG.Sugar().Info("给用户发送消息", appId, userId, err)
		}

		return
	}

	key := GetUserKey(appId, userIdStr)
	info, err := cache.GetUserOnlineInfo(key)
	if err != nil {
		service.LOG.Sugar().Info("给用户发送消息失败", key, err)

		return false, nil
	}
	if !info.IsOnline() {
		service.LOG.Sugar().Info("用户不在线", key)
		return false, nil
	}
	server := models.NewServer(info.AccIp, info.AccPort)
	msg, err := grpcclient.SendMsg(server, msgId, appId, userIdStr, models.MessageCmdMsg, models.MessageCmdMsg, message)
	if err != nil {
		service.LOG.Sugar().Info("给用户发送消息失败", key, err)

		return false, err
	}
	service.LOG.Sugar().Info("给用户发送消息成功-rpc", msg)
	sendResults = true

	return
}

// 给本机用户发送消息
func SendUserMessageLocal(appId uint32, userId string, data string) (sendResults bool, err error) {

	client := GetUserClient(appId, userId)
	if client == nil {
		err = errors.New("用户不在线")

		return
	}

	// 发送消息
	client.SendMsg([]byte(data))
	sendResults = true

	return
}

// 给全体用户发消息
func SendUserMessageAll(appId uint32, userId string, msgId, cmd, message string) (sendResults bool, err error) {
	sendResults = true

	currentTime := uint64(time.Now().Unix())
	servers, err := cache.GetServerAll(currentTime)
	if err != nil {
		service.LOG.Sugar().Info("给全体用户发消息", err)

		return
	}

	for _, server := range servers {
		if IsLocal(server) {
			data := models.GetMsgData(userId, msgId, cmd, message)
			AllSendMessages(appId, userId, data)
		} else {
			grpcclient.SendMsgAll(server, msgId, appId, userId, cmd, message)
		}
	}

	return
}
