/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-09-28 02:41:48
 * @LastEditTime: 2022-04-24 16:48:00
 * @FilePath: \PithyGo\models\structure\user_log.go
 */
package structure

import (
	"PithyGo/service"

	"gorm.io/gorm"
)

type UserLog struct {
	gorm.Model
	UserId uint   `json:"user_id"`
	Record string `json:"record"`
	Route  string `json:"route"`
}

func SaveUserLog(user_id uint, record string, route string) {
	var data UserLog
	data.Record = record
	data.Route = route
	data.UserId = user_id
	saveUserLog(data)
}

func saveUserLog(data UserLog) {
	service.DB.Save(&data)
}
