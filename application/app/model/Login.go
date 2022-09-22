/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-04-26 17:24:52
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2022-09-16 09:48:47
 */
package model

import (
	"PithyGo/application/app"
	"PithyGo/common"
	"PithyGo/models/member"
	"PithyGo/models/structure"
	"PithyGo/service"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/kataras/iris/v12"
)

func ReMemberPass(c iris.Context) (int, string) {

	type fromJsonType struct {
		Password   string `json:"password"`
		Repassword string `json:"repassword"`
	}

	var fromJson fromJsonType

	if err := c.ReadJSON(&fromJson); err != nil {
		service.LOG.Sugar().Error("解析[OutDeliverOpen]解析出错 %s", fromJson)
		return 101, err.Error()
	}

	type validateType struct {
		Password   string `json:"password"      form:"password"   validate:"required"`
		Repassword string `json:"repassword"    form:"repassword" validate:"required"`
	}

	validate := &validateType{
		Password:   fromJson.Password,
		Repassword: fromJson.Repassword,
	}

	if err := common.Validate(validate); err != nil {
		return 102, err.Error()
	}

	if fromJson.Password != fromJson.Repassword {
		return 103, "两次密码输入不一致"
	}
	if len(fromJson.Password) < 6 {
		return 103, "密码不得少于六位"
	}

	if fromJson.Password == "111111" {
		return 103, "禁止将密码修改为初始密码"
	}

	if err := member.ReMemberPass(common.SessManager.Start(c).Get("MEMBER_ID").(uint), fromJson.Password); err == nil {
		return 200, "success"
	} else {
		return -1, err.Error()
	}

}

func Register(c iris.Context) (int, string) {

	type fromJsonType struct {
		Mobile   string `json:"mobile"`
		Code     string `json:"code"`
		Password string `json:"password"`
	}

	var fromJson fromJsonType
	if err := c.ReadJSON(&fromJson); err != nil {
		service.LOG.Sugar().Error("用户登陆时解析出错 %s", fromJson)
	}

	code, err := service.RedisClient.Get("REG_" + fromJson.Mobile).Result()
	if err != nil {
		return 103, "请发送验证码"
	}

	type validateType struct {
		Mobile   string `json:"mobile"      form:"mobile"   validate:"required"`
		Code     string `json:"code"    form:"code" validate:"required"`
		Password string `json:"password"    form:"password" validate:"required"`
	}

	validate := &validateType{
		Mobile:   fromJson.Mobile,
		Code:     fromJson.Code,
		Password: fromJson.Password,
	}

	if err := common.Validate(validate); err != nil {
		return 102, err.Error()
	}

	if len(code) < 8 {
		return 103, "请发送验证码"
	}
	codeArry := strings.Split(code, "_")
	codeCount, err := strconv.Atoi(codeArry[1])
	if err == nil && codeCount > 4 {
		return 104, "验证码已失效"
	}
	if codeArry[0] != fromJson.Code {
		service.RedisClient.Set("REG_"+fromJson.Mobile, codeArry[0]+"_"+strconv.Itoa(codeCount+1), time.Minute*30)
		return 105, "验证码有误"
	}
	if len(fromJson.Password) < 6 {
		return 103, "密码不能小于六位"
	}

	if member.MemberRegister(fromJson.Mobile, fromJson.Password) != nil {
		return 106, err.Error()
	} else {
		return 200, "success"
	}
}

func LoginForGet(c iris.Context) (int, string) {

	type fromJsonType struct {
		Mobile   string `json:"mobile"`
		Code     string `json:"code"`
		Password string `json:"password"`
	}

	var fromJson fromJsonType
	if err := c.ReadJSON(&fromJson); err != nil {
		service.LOG.Sugar().Error("用户登陆时解析出错 %s", fromJson)
	}

	code, err := service.RedisClient.Get("FORGET_" + fromJson.Mobile).Result()
	if err != nil {
		return 103, "请发送验证码"
	}

	type validateType struct {
		Mobile   string `json:"mobile"      form:"mobile"   validate:"required"`
		Code     string `json:"code"    form:"code" validate:"required"`
		Password string `json:"password"    form:"password" validate:"required"`
	}

	validate := &validateType{
		Mobile:   fromJson.Mobile,
		Code:     fromJson.Code,
		Password: fromJson.Password,
	}

	if err := common.Validate(validate); err != nil {
		return 102, err.Error()
	}

	if len(code) < 8 {
		return 103, "请发送验证码"
	}
	codeArry := strings.Split(code, "_")
	codeCount, err := strconv.Atoi(codeArry[1])
	if err == nil && codeCount > 4 {
		return 104, "验证码已失效"
	}
	if codeArry[0] != fromJson.Code {
		service.RedisClient.Set("FORGET_"+fromJson.Mobile, codeArry[0]+"_"+strconv.Itoa(codeCount+1), time.Minute*30)
		return 105, "验证码有误"
	}
	if len(fromJson.Password) < 6 {
		return 103, "密码不能小于六位"
	}
	if member.ReMemberPassByPhone(fromJson.Mobile, fromJson.Password) != nil {
		return 106, err.Error()
	} else {
		return 200, "success"
	}
}
func LoginSendCode(c iris.Context) (int, string) {
	var (
		phone = c.URLParam("phone")
		_type = c.URLParam("type")
	)
	if len(phone) != 11 {
		return 101, "请输入正确的手机号码"
	}
	if _type != "reg" && _type != "forget" {
		return 102, "未知的验证码方式"
	}
	if err := member.MemberLoginSendCode(phone, _type); err == nil {
		return 200, "success"
	} else {
		return -1, err.Error()
	}
}

func Login(c iris.Context) (map[string]interface{}, error) {

	type memberLogin struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var postJsonData memberLogin
	if err := c.ReadJSON(&postJsonData); err != nil {
		service.LOG.Sugar().Error("用户登陆时解析出错 %s", err)
	}
	userData := member.GetByMemberOr(postJsonData.Username)
	//fmt.Println(userData.ID)

	if userData == (structure.Member{}) {
		return nil, errors.New("账号或者密码有误")
	}

	//获取加密密码
	sha1Pass := common.GenerateAppPassWord(postJsonData.Password, userData.Salt)
	//println(sha1Pass, "--------------------")
	if sha1Pass != userData.Password {
		return nil, errors.New("账号或者密码有误")
	}

	token, err := app.GenerateToken(userData.ID, userData.Username)
	if err == nil && token != "" {

		terminal := "unknown"

		if strings.Contains(c.GetHeader("User-Agent"), "AppleWebKit") && strings.Contains(c.GetHeader("User-Agent"), "Mobile") {
			terminal = "webMobile"
		} else if strings.Contains(c.GetHeader("User-Agent"), "AppleWebKit") {
			terminal = "web"
		} else if strings.Contains(c.GetHeader("User-Agent"), "Dalvik") {
			terminal = "app"
		} else {
			return nil, errors.New(c.GetHeader("User-Agent") + "-无法识别")
		}

		member.CreatMemberLogin(userData.ID, terminal, token, common.GetClientIP(c.Request()))
		c.SetCookieKV("NR_Authorization", token)

		return map[string]interface{}{
			"userinfo": map[string]interface{}{
				"id":       userData.ID,
				"username": userData.Username,
				"nickname": userData.Nickname,
				"avatar":   userData.Avatar,
				"mobile":   userData.Mobile,
				"level":    userData.Level,
				"amount":   userData.Amount,
			},
			"authorization": token,
		}, nil
	} else {
		return nil, err
	}
}

func UserGetMenu(c iris.Context) interface{} {

	json := `{"menu":[{"path":"/dashboard","name":"dashboard","meta":{"title":"控制台","icon":"el-icon-eleme-filled","type":"menu","affix":true},"component":"home"},{"name":"my","path":"/my","meta":{"title":"我","icon":"el-icon-collection-tag","type":"menu"},"children":[{"path":"/my/department","name":"myDepartment","meta":{"title":"我的部门","icon":"el-icon-avatar","type":"menu"},"component":"my/department"},{"path":"/my/project","name":"myProject","meta":{"title":"我的项目","icon":"el-icon-menu","type":"menu"},"component":"my/project"},{"path":"/my/attendance","name":"myAttendance","meta":{"title":"我的考勤","icon":"el-icon-checked","type":"menu"},"component":"my/attendance"},{"path":"/my/account","name":"myAccount","meta":{"title":"我的账号","icon":"el-icon-finished","type":"menu"},"component":"my/account"},{"name":"userCenter","path":"/my/usercenter","meta":{"title":"个人信息","icon":"el-icon-user-filled"},"component":"my/userCenter"},{"path":"/my/resume","name":"myResume","meta":{"title":"个人简历","icon":"el-icon-document-checked","type":"menu"},"component":"my/resume"}]}],"permissions":["list.add","list.edit","list.delete","user.add","user.edit","user.delete"]}`

	j, _ := simplejson.NewJson([]byte(json))

	return j.Interface()
}
