/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-10-20 22:05:02
 * @LastEditTime: 2021-10-21 01:10:27
 * @FilePath: \PithyGo\models\structure\user_message.go
 */
package structure

import (
	"PithyGo/service"
	"time"

	"gorm.io/gorm"
)

type UserMessage struct {
	gorm.Model
	SendUserId    uint   `json:"send_user_id"`    //发送人ID
	ReceiveUserId uint   `json:"receive_user_id"` //接收人ID
	IsRead        uint   `json:"is_read"`         //是否已读
	Content       string `json:"content"`         //发送内容
	Type          string `json:"type"`            //发送内容
	Link          string `json:"link"`            //链接
}

type UserMessageList struct {
	ID         uint `json:"id"`           //ID
	SendUserId uint `json:"send_user_id"` //发送人ID
	//	ReceiveUserId       uint      `json:"receive_user_id"`       //接收人ID
	SendUserAvatar string `json:"send_user_avatar"` //发送人头像
	//ReceiveUserAvatar   string    `json:"receive_user_avatar"`   //接收人头像
	SendUserRealname string `json:"send_user_realname"` //发送人名字
	//ReceiveUserRealname string    `json:"receive_user_realname"` //接收人名字
	Type      string    `json:"type"`    //发送内容
	Link      string    `json:"link"`    //链接
	Content   string    `json:"content"` //发送内容
	CreatedAt time.Time `json:"createdAt"`
}

func NewUserMessage(send_user_id uint, receive_user_id uint, content string, typeA string, link string) {
	var data UserMessage
	data.SendUserId = send_user_id
	data.ReceiveUserId = receive_user_id
	data.Content = content
	data.Link = link
	data.Type = typeA
	service.DB.Save(&data)
}
