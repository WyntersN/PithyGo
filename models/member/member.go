/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-04-26 17:52:37
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2022-09-16 09:43:01
 */
package member

import (
	"PithyGo/addons/tencentyun/sms"
	"PithyGo/common"
	"PithyGo/models/structure"
	"PithyGo/service"
	"errors"
	"strings"

	"time"
)

/**
 * @summary: 根据关键词[keyword]获取用户表信息
 * @description:
 * @param {string} keyword
 * @router:
 */
func GetByMemberOr(keyword string) structure.Member {
	var data structure.Member
	service.DB.Where("username = ? OR mobile = ?", keyword, keyword).Find(&data)
	return data
}

func (c *Content) GetMemberByMID() (structure.MemberInfo, error) {
	var data structure.Member

	if err := service.DB.First(&data, c.MemberID).Error; err != nil {
		return structure.MemberInfo{}, common.ReturnErrorToDB("查询用户信息失败", err)
	} else {
		// priceRoundFloor, _ := decimal.NewFromFloat(rand.Float64() * 100).RoundFloor(2).Float64()
		// amount, _ := strconv.ParseFloat(strconv.FormatFloat(priceRoundFloor, 'f', 2, 64), 10)
		// service.DB.Model(&data).UpdateColumn("amount", gorm.Expr("amount + ?", amount)).First(&data)

		return structure.MemberInfo{
			Username: data.Username,
			Nickname: data.Nickname,
			Mobile:   data.Mobile,
			Avatar:   *data.Avatar,
			Amount:   data.Amount,
			Level:    data.Level,
		}, nil
	}

}

func (c *AdminContent) GetMemberList(pageSize, page int, keyword string) ([]structure.MemberList, int64) {
	var (
		data  []structure.MemberList
		count int64
	)
	keyword = "%" + strings.ToLower(common.CompresStr(keyword)) + "%"
	service.DB.Model(&structure.Member{}).Where("username LIKE ? OR nickname LIKE ? OR id LIKE ? OR realname LIKE ? OR mobile LIKE ? ", keyword, keyword, keyword, keyword, keyword).Count(&count).Order("id ASC").Limit(pageSize).Offset((page - 1) * pageSize).Scan(&data)

	return data, count
}

func (c *AdminContent) GetMemberCash(pageSize, page int, keyword string) ([]structure.MemberCashList, int64) {
	var (
		data  []structure.MemberCashList
		count int64
	)
	keyword = "%" + strings.ToLower(common.CompresStr(keyword)) + "%"
	service.DB.Table(common.DataPrefix+"member_cash mc").Joins("left join nr_member m ON m.id = mc.member_id").Where("m.username LIKE ? OR m.nickname LIKE ? OR m.id LIKE ? OR m.realname LIKE ? OR m.mobile LIKE ? ", keyword, keyword, keyword, keyword, keyword).Count(&count).Order("mc.status ASC,mc.id ASC").Limit(pageSize).Offset((page - 1) * pageSize).Scan(&data)

	return data, count
}

func ReMemberPass(member_id uint, pass string) error {
	var data structure.Member
	if err := service.DB.First(&data, member_id).Error; err != nil {
		return err
	}
	password := common.GenerateAppPassWord(pass, data.Salt)
	return service.DB.Model(&data).Update("password", password).Error

}

func ReMemberPassByPhone(phone, pass string) error {
	service.RedisClient.Del("FORGET_" + phone)
	var data structure.Member
	if err := service.DB.First(&data, "mobile = ?", phone).Error; err != nil {
		return err
	}
	password := common.GenerateAppPassWord(pass, data.Salt)
	return service.DB.Model(&data).Update("password", password).Error

}

func MemberLoginSendCode(phone, _type string) error {

	var code = common.GetRandomNumber(6)
	if _type == "reg" {
		if service.DB.Where("mobile = ?", phone).Find(&structure.Member{}).RowsAffected > 0 {
			return errors.New("该手机号码已注册")
		}
	} else {
		if err := service.DB.First(&structure.Member{}, "mobile = ?", phone).Error; err != nil {
			return common.ReturnErrorToDB("查询用户失败", err)
		}
	}

	if err := sms.CurrencyMinutesSend(phone, code, "30"); err == nil {
		service.RedisClient.Set(strings.ToUpper(_type)+"_"+phone, code+"_0", time.Minute*30)
		return nil
	} else {
		return err
	}

}

func MemberRegister(mobile, password string) error {
	service.RedisClient.Del("REG_" + mobile)
	var (
		data structure.Member
		salt = common.GetRandomString(6)
	)
	service.DB.Where("mobile = ?", mobile).Find(&data)
	if data.ID != 0 {
		return errors.New("该手机号码已注册")
	}

	return service.DB.Create(&structure.Member{
		Mobile:   mobile,
		Username: mobile,
		Nickname: mobile[:3] + "****" + mobile[7:11],
		Salt:     salt,
		Password: common.GenerateAppPassWord(password, salt),
	}).Error

}

/**
 * @summary: 添加用户登陆记录
 * @description:
 * @param {int} userid
 * @param {string} key
 * @param {string} clientIP
 * @router:
 */
func CreatMemberLogin(userid uint, terminal, key, clientIP string) {
	// var dataUserLogin []structure.UserLogin
	// service.DB.Where("user_id = ? AND terminal = ?", userid, terminal).Order("id ASC").Limit(5).Find(&dataUserLogin)
	// if len(dataUserLogin) > 4 {
	// 	service.DB.Delete(&dataUserLogin[0])
	// }

	service.DB.Delete(&structure.MemberLogin{}, "member_id = ? AND terminal = ?", userid, terminal)
	var data structure.MemberLogin
	{
		data.MemberId = userid
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
func GetByMemberID(userid uint) structure.Member {

	var structureUser structure.Member
	service.DB.First(&structureUser, userid)
	return structureUser
}

/**
 * @summary: 获取用户最近登陆数据
 * @description:
 * @param {*}
 * @router:
 */
func GetMemberLoginLogByAuthorization(authorization string) structure.MemberLogin {
	var structureUserLogin structure.MemberLogin

	service.DB.Where("authorization = ?", authorization).Find(&structureUserLogin)
	return structureUserLogin
}

func GetMemberList(pageSize int, page int, order string, keyword string) ([]structure.Member, int64) {

	var (
		structureUser []structure.Member
		count         int64
	)
	keyword = "%" + strings.ToLower(common.CompresStr(keyword)) + "%"
	service.DB.Table("yc_user").Where("deleted_at IS NULL").Where("username LIKE  ? OR mobile LIKE  ? OR nickname LIKE  ? OR realname LIKE  ? OR lower(cnfirstchar(username || nickname)) LIKE  ? OR lower(f_gethzpy(username || nickname)) LIKE  ?", keyword, keyword, keyword, keyword, keyword, keyword).Count(&count).Order("department_job DESC").Limit(pageSize).Offset((page - 1) * pageSize).Scan(&structureUser)

	return structureUser, count
}
