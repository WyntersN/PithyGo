/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-04-26 17:52:37
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2022-07-16 17:37:36
 */
package user

import (
	"PithyGo/common"
	"PithyGo/models/structure"
	"PithyGo/service"
	"strings"

	"time"
)

/**
 * @summary: 根据关键词[keyword]获取用户表信息
 * @description:
 * @param {string} keyword
 * @router:
 */
func GetByUserOr(keyword string) structure.User {
	var data structure.User
	service.DB.Where("username = ? OR mobile = ?  ", keyword, keyword).Find(&data)
	return data
}

func ReUserPass(user_id uint, pass string) error {
	var data structure.User
	if err := service.DB.Where("id = ?", user_id).First(&data).Error; err != nil {
		return err
	}
	password := common.GenerateAppPassWord(pass, data.Salt)
	return service.DB.Model(&data).Update("password", password).Error

}

/**
 * @summary: 添加用户登陆记录
 * @description:
 * @param {int} userid
 * @param {string} key
 * @param {string} clientIP
 * @router:
 */
func CreatUserLogin(userid uint, terminal, key, clientIP string) {
	// var dataUserLogin []structure.UserLogin
	// service.DB.Where("user_id = ? AND terminal = ?", userid, terminal).Order("id ASC").Limit(5).Find(&dataUserLogin)
	// if len(dataUserLogin) > 4 {
	// 	service.DB.Delete(&dataUserLogin[0])
	// }

	service.DB.Where("expire_at < ?", time.Now()).Delete(&structure.UserLogin{})
	var data structure.UserLogin
	{
		data.UserId = userid
		data.Terminal = terminal
		data.Authorization = key
		data.LoginIp = clientIP
		data.ExpireAt = time.Unix(time.Now().Add(time.Hour*720).Unix(), 0)

		service.DB.Create(&data)
	}

}

/**
 * @summary: 根据用户UID获取用户信息
 * @description:
 * @param {int} userid
 * @router:
 */
func GetByUserID(userid uint) structure.User {

	var structureUser structure.User
	service.DB.First(&structureUser, userid)
	return structureUser
}

/**
 * @summary: 获取用户最近登陆数据
 * @description:
 * @param {*}
 * @router:
 */
func GetLoginLogByAuthorization(authorization string) structure.UserLogin {
	var structureUserLogin structure.UserLogin

	service.DB.Where("authorization = ?", authorization).Find(&structureUserLogin)
	return structureUserLogin
}

func GetUserList(pageSize int, page int, order string, keyword string) ([]structure.UserList, int64) {

	var (
		structureUser []structure.UserList
		count         int64
	)
	keyword = "%" + strings.ToLower(common.CompresStr(keyword)) + "%"
	service.DB.Table("yc_user").Where("deleted_at IS NULL").Where("username LIKE  ? OR mobile LIKE  ? OR nickname LIKE  ? OR realname LIKE  ? OR lower(cnfirstchar(username || nickname)) LIKE  ? OR lower(f_gethzpy(username || nickname)) LIKE  ?", keyword, keyword, keyword, keyword, keyword, keyword).Count(&count).Order("department_job DESC").Limit(pageSize).Offset((page - 1) * pageSize).Scan(&structureUser)

	return structureUser, count
}

func GetUserMessage(user_id uint) []structure.UserMessageList {
	var data []structure.UserMessageList
	service.DB.Table("nr_user_message um").Select("um.*,us.realname as send_user_realname,us.avatar as send_user_avatar").Joins("left join nr_user us on us.id = um.send_user_id").Where("is_read = 0 AND  receive_user_id = ?", user_id).Order("id DESC").Find(&data)
	return data
}
func UserMessageRead(user_id uint) error {
	var data structure.UserMessage
	return service.DB.Model(&data).Where("receive_user_id = ?", user_id).Update("is_read", 1).Error
}
