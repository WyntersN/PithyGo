/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-04-19 20:46:35
 * @LastEditTime: 2022-08-20 18:02:54
 * @FilePath: \PithyGo\service\s_init\c.go
 */
package s_init

import (
	"PithyGo/addons/tencentyun/cos"
	"PithyGo/service"
)

func Initialization() {
	service.NewRedisClient()
	cos.NewCos()
	go service.StartCron()
	cronTask()

}
