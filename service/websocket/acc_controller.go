/**
 * Created by GoLand.
 * User: link1st
 * Date: 2019-07-27
 * Time: 13:12
 */

package websocket

import (
	"PithyGo/application/admin"
	"PithyGo/service"
	"PithyGo/service/websocket/cache"
	"PithyGo/service/websocket/common"
	"PithyGo/service/websocket/models"
	"encoding/json"
	"strconv"

	"github.com/go-redis/redis"

	"time"
)

// ping
func PingController(client *Client, seq string, message []byte) (code uint32, msg string, data interface{}) {

	code = common.OK
	service.LOG.Sugar().Info("webSocket_request ping接口", client.Addr, seq, message)

	data = "pong"

	return
}

// 用户登录
func LoginController(client *Client, seq string, message []byte) (code uint32, msg string, data interface{}) {

	code = common.OK
	currentTime := uint64(time.Now().Unix())

	request := &models.Login{}
	if err := json.Unmarshal(message, request); err != nil {
		code = common.ParameterIllegal
		service.LOG.Sugar().Info("用户登录 解析数据失败", seq, err)

		return
	}

	service.LOG.Sugar().Info("webSocket_request 用户登录-", seq, "ServiceToken-", request.Token)
	userid, realname, err := admin.AuthApp(request.Token)
	if err != nil {
		code = common.UnauthorizedUserId
		msg = err.Error()
		return
	}

	request.UserId = strconv.Itoa(int(userid))

	// TODO::进行用户权限认证，一般是客户端传入TOKEN，然后检验TOKEN是否合法，通过TOKEN解析出来用户ID
	// 本项目只是演示，所以直接过去客户端传入的用户ID
	if request.UserId == "" || len(request.UserId) >= 20 {
		code = common.UnauthorizedUserId
		service.LOG.Sugar().Info("用户登录 非法的用户", seq, request.UserId)
		return
	}

	if !InAppIds(request.AppId) {
		code = common.Unauthorized
		service.LOG.Sugar().Info("用户登录 不支持的平台", seq, request.AppId)

		return
	}

	if client.IsLogin() {
		service.LOG.Sugar().Info("用户登录 用户已经登录", client.AppId, client.UserId, seq)
		code = common.OperationFailure

		return
	}

	client.Login(request.AppId, request.UserId, realname, currentTime)
	data = admin.WebVersion

	// 存储数据
	userOnline := models.UserLogin(serverIp, serverPort, request.AppId, request.UserId, realname, client.Addr, currentTime)
	err = cache.SetUserOnlineInfo(client.GetKey(), userOnline)
	if err != nil {
		code = common.ServerError
		service.LOG.Sugar().Info("用户登录 SetUserOnlineInfo", seq, err)

		return
	}

	// 用户登录
	login := &login{
		AppId:  request.AppId,
		UserId: request.UserId,
		Client: client,
	}
	clientManager.Login <- login

	service.LOG.Sugar().Info("用户登录 成功", seq, client.Addr, request.UserId)

	return
}

// 心跳接口
func HeartbeatController(client *Client, seq string, message []byte) (code uint32, msg string, data interface{}) {

	code = common.OK
	currentTime := uint64(time.Now().Unix())

	request := &models.HeartBeat{}
	if err := json.Unmarshal(message, request); err != nil {
		code = common.ParameterIllegal
		service.LOG.Sugar().Info("心跳接口 解析数据失败", seq, err)

		return
	}

	service.LOG.Sugar().Info("webSocket_request 心跳接口", client.AppId, client.UserId)

	if !client.IsLogin() {
		service.LOG.Sugar().Info("心跳接口 用户未登录", client.AppId, client.UserId, seq)
		code = common.NotLoggedIn

		return
	}

	userOnline, err := cache.GetUserOnlineInfo(client.GetKey())
	if err != nil {
		if err == redis.Nil {
			code = common.NotLoggedIn
			service.LOG.Sugar().Info("心跳接口 用户未登录", seq, client.AppId, client.UserId)

			return
		} else {
			code = common.ServerError
			service.LOG.Sugar().Info("心跳接口 GetUserOnlineInfo", seq, client.AppId, client.UserId, err)

			return
		}
	}

	client.Heartbeat(currentTime)
	userOnline.Heartbeat(currentTime)
	err = cache.SetUserOnlineInfo(client.GetKey(), userOnline)
	if err != nil {
		code = common.ServerError
		service.LOG.Sugar().Info("心跳接口 SetUserOnlineInfo", seq, client.AppId, client.UserId, err)

		return
	}
	data = admin.WebVersion
	return
}
