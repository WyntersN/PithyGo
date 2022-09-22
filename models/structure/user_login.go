/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-09-25 14:04:05
 * @LastEditTime: 2021-11-15 03:34:18
 * @FilePath: \PithyGo\models\structure\user_login.go
 */
package structure

import (
	"time"

	"gorm.io/gorm"
)

type UserLogin struct {
	gorm.Model
	UserId        uint      `gorm:"type:int(4);NOT NULL;"				json:"user_id"`
	Authorization string    `gorm:"type:varchar(255);NOT NULL;unique;"	json:"secret_key"`
	Terminal      string    `json:"terminal"`
	LoginIp       string    `gorm:"type:varchar(20);NOT NULL"			json:"login_ip"`
	ExpireAt      time.Time `gorm:"type:timestamp(6);NOT NULL"			json:"exp"`
}
