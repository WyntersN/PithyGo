/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-04-26 17:17:48
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2022-08-23 17:45:15
 */
package structure

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Username  string  `gorm:"type:varchar(20);NOT NULL;"  		json:"username"`
	Realname  *string `json:"realname"`
	Nickname  string  `gorm:"type:varchar(30);NOT NULL;"		json:"nickname"`
	Mobile    string  `gorm:"type:char(11);NOT NULL;"			json:"mobile"`
	Password  string  `gorm:"type:char(40);NOT NULL"			json:"password"`
	Salt      string  `gorm:"type:char(6);NOT NULL"				json:"salt"`
	Avatar    *string `gorm:"type:varchar(255);NOT NULL"		json:"avatar"`
	Amount    float32 `json:"amount"`
	Level     uint    `json:"level"`
	AccAlipay *string `json:"acc_alipay"`
}
type MemberList struct {
	Id        uint      `json:"id"` //
	Username  string    `json:"username"`
	Realname  string    `json:"realname"`
	Nickname  string    `json:"nickname"`
	Mobile    string    `json:"mobile"`
	Avatar    string    `json:"avatar"`
	Amount    float32   `json:"amount"`
	Level     uint      `json:"level"`
	AccAlipay *string   `json:"acc_alipay"`
	CreatedAt time.Time `json:"createdAt"` //False
}
type MemberInfo struct {
	Username string  `json:"username"`
	Nickname string  `json:"nickname"`
	Mobile   string  `json:"mobile"`
	Avatar   string  `json:"avatar"`
	Amount   float32 `json:"amount"`
	Level    uint    `json:"level"`
}
