/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-04-26 17:24:52
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2022-09-13 18:33:34
 */
package user

import (
	"PithyGo/application/admin"
	"PithyGo/common"
	"PithyGo/models/structure"
	"PithyGo/models/user"
	"PithyGo/service"
	"errors"
	"strconv"
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/kataras/iris/v12"
)

func ReUserPass(c iris.Context) (int, string) {

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

	if err := user.ReUserPass(common.SessManager.Start(c).Get("USER_ID").(uint), fromJson.Password); err == nil {
		return 200, "success"
	} else {
		return -1, err.Error()
	}

}

func Login(c iris.Context) (map[string]interface{}, error) {

	type userLogin struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	postJsonData := &userLogin{}
	if err := c.ReadJSON(postJsonData); err != nil {
		service.LOG.Sugar().Error("用户登陆时解析出错 %s", postJsonData)
	}
	userData := user.GetByUserOr(postJsonData.Username)
	//fmt.Println(userData.ID)

	if userData == (structure.User{}) {
		return nil, errors.New("账号或者密码有误")
	}

	//获取加密密码
	sha1Pass := common.GenerateAppPassWord(postJsonData.Password, userData.Salt)
	//println(sha1Pass, "--------------------")
	if sha1Pass != userData.Password {
		return nil, errors.New("账号或者密码有误")
	}
	if userData.IsQuit == 1 {
		return nil, errors.New("您的状态系统无法为您登陆")
	}
	repass := 0
	if postJsonData.Password == "111111" {
		repass = 1
	}

	token, err := admin.GenerateToken(userData.ID, userData.Username)
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

		user.CreatUserLogin(userData.ID, terminal, token, common.GetClientIP(c.Request()))
		c.SetCookieKV("NR_Authorization", token)
		var identity string
		if userData.Jurisdiction == 1 {
			identity = "admin"
		} else {
			identity = "user"
		}

		return map[string]interface{}{
			"userinfo": map[string]interface{}{
				"id":                 userData.ID,
				"username":           userData.Username,
				"nickname":           userData.Nickname,
				"realname":           userData.Realname,
				"avatar":             userData.Avatar,
				"mobile":             userData.Mobile,
				"department":         userData.DepartmentName,
				"department_jobName": userData.DepartmentJobName,
				"job_num":            userData.JobNum,
				"birth_date":         userData.BirthDate,
				"sex":                userData.Sex,
				"mail":               userData.Mail,
			},
			"authorization": token,
			"repass":        repass,
			"identity":      identity,
		}, nil
	} else {
		return nil, err
	}

	// claim, err := common.ParseToken("Bearer " + token)

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(uint(claim["userid"].(float64)))
	// fmt.Println(claim["username"].(string))
	// fmt.Println(uint(claim["exp"].(float64)))
	//return nil, errors.New("未知错误")
}
func Info(c iris.Context) (map[string]interface{}, error) {
	type validateType struct {
		Token string `json:"token" form:"token" validate:"required"`
	}

	validate := &validateType{
		Token: c.GetHeader("authorization"),
	}

	if err := common.Validate(validate); err != nil {
		return nil, err
	}

	userData := user.GetByUserID(common.SessManager.Start(c).Get("USER_ID").(uint))

	return map[string]interface{}{
		"username": userData.Username,
		"nickname": userData.Nickname,
		"realname": userData.Realname,
		"avatar":   userData.Avatar,
		"mobile":   userData.Mobile,
		"roles":    []string{"admin"},
	}, nil
}

func GetUserMessage(c iris.Context) interface{} {
	return user.GetUserMessage(common.SessManager.Start(c).Get("USER_ID").(uint))
}
func UserMessageRead(c iris.Context) interface{} {
	return user.UserMessageRead(common.SessManager.Start(c).Get("USER_ID").(uint))
}

func UserGetMenu(c iris.Context) interface{} {

	json := `{"menu":[{"path":"/dashboard","name":"dashboard","meta":{"title":"控制台","icon":"el-icon-eleme-filled","type":"menu","affix":true},"component":"home"},{"name":"my","path":"/my","meta":{"title":"我","icon":"el-icon-collection-tag","type":"menu"},"children":[{"path":"/my/department","name":"myDepartment","meta":{"title":"我的部门","icon":"el-icon-avatar","type":"menu"},"component":"my/department"},{"path":"/my/project","name":"myProject","meta":{"title":"我的项目","icon":"el-icon-menu","type":"menu"},"component":"my/project"},{"path":"/my/attendance","name":"myAttendance","meta":{"title":"我的考勤","icon":"el-icon-checked","type":"menu"},"component":"my/attendance"},{"path":"/my/account","name":"myAccount","meta":{"title":"我的账号","icon":"el-icon-finished","type":"menu"},"component":"my/account"},{"name":"userCenter","path":"/my/usercenter","meta":{"title":"个人信息","icon":"el-icon-user-filled"},"component":"my/userCenter"},{"path":"/my/resume","name":"myResume","meta":{"title":"个人简历","icon":"el-icon-document-checked","type":"menu"},"component":"my/resume"}]}],"permissions":["list.add","list.edit","list.delete","user.add","user.edit","user.delete"]}`

	j, _ := simplejson.NewJson([]byte(json))

	return j.Interface()
}

func List(c iris.Context) (map[string]interface{}, error) {

	var (
		page, _  = strconv.Atoi(c.URLParam("page"))
		limit, _ = strconv.Atoi(c.URLParam("pageSize"))
		keyword  = c.URLParam("keyword")
		order    = c.URLParam("sort")
	)

	userList, count := user.GetUserList(limit, page, order, keyword)
	return map[string]interface{}{
		"page":     page,
		"pageSize": limit,
		"total":    count,
		"rows":     userList,
	}, nil
}
