/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-10-22 15:32:49
 * @LastEditTime: 2021-11-23 02:55:24
 * @FilePath: \PithyGo\service\websocket\servers\task\clean_connection _task.go
 */
/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-31
* Time: 15:17
 */

package task

import (
	"PithyGo/service"
	"PithyGo/service/websocket"

	"runtime/debug"
	"time"
)

func Init() {
	Timer(3*time.Second, 60*time.Second, cleanConnection, "", nil, nil)

}

// 清理超时连接
func cleanConnection(param interface{}) bool {

	defer func() {
		if r := recover(); r != nil {
			service.LOG.Sugar().Info("ClearTimeoutConnections stop", r, string(debug.Stack()))
		}
	}()

	service.LOG.Sugar().Info("定时任务，清理超时连接", param)

	websocket.ClearTimeoutConnections()

	return true
}
