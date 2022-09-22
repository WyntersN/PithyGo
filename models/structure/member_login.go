/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-09-25 14:04:05
 * @LastEditTime: 2022-08-03 16:58:54
 * @FilePath: \tborder_wxq\models\structure\member_login.go
 */
package structure

import (
	"time"

	"gorm.io/gorm"
)

type MemberLogin struct {
	gorm.Model
	MemberId      uint      `gorm:"type:int(4);NOT NULL;"				json:"member_id"`
	Authorization string    `gorm:"type:varchar(255);NOT NULL;unique;"	json:"secret_key"`
	Terminal      string    `json:"terminal"`
	LoginIp       string    `gorm:"type:varchar(20);NOT NULL"			json:"login_ip"`
	ExpireAt      time.Time `gorm:"type:timestamp(6);NOT NULL"			json:"exp"`
}
