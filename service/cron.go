/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-06-17 15:14:16
 * @LastEditTime: 2022-06-17 15:14:35
 * @FilePath: \PithyGo\service\cron.go
 */
package service

import "github.com/robfig/cron/v3"

var Cron = cron.New(cron.WithSeconds())

func StartCron() {
	Cron.Start()
	defer Cron.Stop()
	select {}
}
