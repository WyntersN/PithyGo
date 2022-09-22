/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-10-22 15:32:49
 * @LastEditTime: 2021-11-23 02:55:27
 * @FilePath: \PithyGo\service\websocket\servers\task\server_task.go
 */
/**
* Created by GoLand.
* User: link1st
* Date: 2019-08-03
* Time: 15:44
 */

package task

import (
	"PithyGo/service"
	"PithyGo/service/websocket"
	"PithyGo/service/websocket/cache"

	"runtime/debug"
	"time"
)

func ServerInit() {
	Timer(3*time.Second, 60*time.Second, server, "", serverDefer, "")
}

// 服务注册
func server(param interface{}) bool {

	defer func() {
		if r := recover(); r != nil {
			service.LOG.Sugar().Info("服务注册 stop", r, string(debug.Stack()))
		}
	}()

	server := websocket.GetServer()
	currentTime := uint64(time.Now().Unix())
	service.LOG.Sugar().Info("定时任务，服务注册", param, server, currentTime)

	cache.SetServerInfo(server, currentTime)

	return true
}

// 服务下线
func serverDefer(param interface{}) bool {
	defer func() {
		if r := recover(); r != nil {
			service.LOG.Sugar().Info("服务下线 stop", r, string(debug.Stack()))
		}
	}()

	service.LOG.Sugar().Info("服务下线", param)

	server := websocket.GetServer()
	cache.DelServerInfo(server)

	return true
}
