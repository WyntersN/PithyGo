/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-04-26 17:17:48
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2022-05-14 00:52:15
 */
package structure

import (
	"PithyGo/common"
	"PithyGo/service"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username            string     `gorm:"type:varchar(20);NOT NULL;"  		json:"username"`
	Nickname            string     `gorm:"type:varchar(30);NOT NULL;"		json:"nickname"`
	Realname            string     `gorm:"type:varchar(20);NOT NULL;"		json:"realname"`
	Mobile              *string    `gorm:"type:char(11);NOT NULL;"			json:"mobile"`
	Password            string     `gorm:"type:char(40);NOT NULL"			json:"password"`
	Salt                string     `gorm:"type:char(6);NOT NULL"				json:"salt"`
	Avatar              string     `gorm:"type:varchar(255);NOT NULL"		json:"avatar"`
	UserDingtalkUserid  *string    `gorm:"type:varchar(30);NOT NULL;unique;"	json:"userDingtalkUserid"`
	IsQuit              uint       `gorm:"type:int2(16);"			json:"isQuit"`
	DepartmentName      *string    `json:"department_name"`      //所在部门
	DepartmentJob       uint       `json:"department_job"`       //所在部门的职务
	InductionTime       *time.Time `json:"induction_time"`       //入职时间
	DepartmentJobName   string     `json:"department_job_name"`  //职务别称
	JobNum              *string    `json:"job_num"`              //工号
	Sex                 *string    `json:"sex"`                  //性别
	BirthDate           *time.Time `json:"birth_date"`           //出身日期
	NativePlace         *string    `json:"native_place"`         //籍贯
	Nation              *string    `json:"nation"`               //名族
	IdentityNum         *string    `json:"identity_num"`         //身份证号码
	EducationBackground *string    `json:"education_background"` //学历
	GraduateInstitution *string    `json:"graduate_institution"` //毕业学校
	Major               *string    `json:"major"`                //专业
	UprightDate         *time.Time `json:"upright_date"`         //转正日期
	WeworkUserid        *string    `json:"wework_userid"`        //企业微信ID
	Mail                *string    `json:"mail"`                 //邮箱
	Jurisdiction        uint       `gorm:"type:int2(16);"			json:"jurisdiction"`
	BaseSalary          *float64   `json:"base_salary"` //基本薪资
	AlipayAcc           *string    `json:"alipay_acc"`  //支付宝账号
}

type UserList struct {
	ID                  uint       `json:"id"`
	Username            string     `json:"username"`
	Nickname            string     `json:"nickname"`
	Realname            string     `json:"realname"`
	Mobile              string     `json:"mobile"`
	Avatar              string     `json:"avatar"`
	IsQuit              uint       `json:"isQuit"`
	InductionTime       *time.Time `json:"induction_time"` //入职时间
	CreatedAt           time.Time  `json:"createdAt"`
	DepartmentName      *string    `json:"department_name"`      //所在部门
	DepartmentJob       uint       `json:"department_job"`       //所在部门的职务
	DepartmentJobName   string     `json:"department_job_name"`  //职务别称
	JobNum              *string    `json:"job_num"`              //工号
	Sex                 *string    `json:"sex"`                  //性别
	BirthDate           *time.Time `json:"birth_date"`           //出身日期
	NativePlace         *string    `json:"native_place"`         //籍贯
	Nation              *string    `json:"nation"`               //名族
	IdentityNum         *string    `json:"identity_num"`         //身份证号码
	EducationBackground *string    `json:"education_background"` //学历
	GraduateInstitution *string    `json:"graduate_institution"` //毕业学校
	Major               *string    `json:"major"`                //专业
	UprightDate         *time.Time `json:"upright_date"`         //转正日期
	WeworkUserid        *string    `json:"wework_userid"`        //企业微信ID
	Mail                *string    `json:"mail"`                 //邮箱
	BaseSalary          *float64   `json:"base_salary"`          //基本薪资
	AlipayAcc           *string    `json:"alipay_acc"`           //支付宝账号
}

type UserByIdRealname struct {
	ID       uint   `json:"id"`
	Realname string `json:"realname"`
}

type WeworkUserInfo struct {
	ToUsername   string `xml:"ToUserName"`
	FromUsername string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	ChangeType   string `xml:"ChangeType"`
	UserID       string `xml:"UserID"`
	Name         string `xml:"Name"`
	Mobile       string `xml:"Mobile"`
	Avatar       string `xml:"Avatar"`
	Alias        string `xml:"Alias"` //别名
	BizMail      string `xml:"BizMail"`
	Department   uint32 `xml:"Department"`
	Gender       uint32 `xml:"Gender"`   //性别
	Status       uint32 `xml:"Status"`   //状态
	Position     string `xml:"Position"` //职务
}

func DelUserWeWork(data *WeworkUserInfo) {

	var user User

	if err := service.DB.Where("wework_userid = ?", data.UserID).First(&user).Error; err == nil {
		tx := service.DB.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		if err := tx.Delete(&user).Error; err != nil {
			service.LOG.Sugar().Error("Sql Begin Error：删除用户信息事务出错 %s", err)
			tx.Rollback()
		}
		if err := tx.Delete(&UserLogin{}, "user_id = ?", user.ID).Error; err != nil {
			service.LOG.Sugar().Error("Sql Begin Error：删除用户登陆事务出错 %s", err)
			tx.Rollback()
		}
		tx.Commit()
	}

}
func UpdateUserWeWork(data *WeworkUserInfo) {

	if data.Status == 5 {
		var user User

		if err := service.DB.Where("wework_userid = ?", data.UserID).First(&user).Error; err == nil {

			tx := service.DB.Begin()
			defer func() {
				if r := recover(); r != nil {
					tx.Rollback()
				}
			}()
			user.IsQuit = 1
			if err := tx.Save(&user).Error; err != nil {
				service.LOG.Sugar().Error("Sql Begin Error：保存用户信息事务出错 %s", err)
				tx.Rollback()
			}
			if err := tx.Delete(&UserLogin{}, "user_id = ?", user.ID).Error; err != nil {
				service.LOG.Sugar().Error("Sql Begin Error：删除用户登陆事务出错 %s", err)
				tx.Rollback()
			}
			tx.Commit()
		}
	}

	user := service.DB.Model(&User{})
	if data.BizMail != "" {
		user.Where("wework_userid = ?", data.UserID).UpdateColumn("mail", data.BizMail)
	}
	if data.UserID != "" {
		user.Where("wework_userid = ?", data.UserID).UpdateColumn("wework_userid", data.UserID)
	}
	if data.Position != "" {
		user.Where("wework_userid = ?", data.UserID).UpdateColumn("department_job_name", data.Position)
	}
	if data.Alias != "" {
		user.Where("wework_userid = ?", data.UserID).UpdateColumn("nickname", data.Alias)
	}
	if data.Mobile != "" {
		user.Where("wework_userid = ?", data.UserID).UpdateColumn("mobile", data.Mobile)
	}
	if data.Name != "" {
		user.Where("wework_userid = ?", data.UserID).UpdateColumn("realname", data.Name)
	}
	if data.Gender == 1 {
		user.Where("wework_userid = ?", data.UserID).UpdateColumn("sex", "男")
	}
	if data.Gender == 2 {
		user.Where("wework_userid = ?", data.UserID).UpdateColumn("sex", "女")
	}
}

func CreatUserWeWork(data *WeworkUserInfo) {

	var (
		structureUser User
		sex           string
	)

	tx := service.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	{
		rand := common.GetRandomString(6)
		structureUser.Avatar = data.Avatar
		structureUser.Username = data.Name
		structureUser.Nickname = data.Alias
		structureUser.Realname = data.Name
		structureUser.Mobile = &data.Mobile
		structureUser.DepartmentJobName = "成员"
		timeAt := time.Now()
		structureUser.InductionTime = &timeAt
		structureUser.Password = common.GenerateAppPassWord("111111", rand)
		structureUser.Salt = rand
		structureUser.WeworkUserid = &data.UserID
		if data.Gender == 1 {
			sex = "男"
		} else if data.Gender == 2 {
			sex = "女"
		}
		structureUser.Sex = &sex
		structureUser.Mail = &data.BizMail

		if err := tx.Create(&structureUser).Error; err != nil {
			service.LOG.Sugar().Error("创建事务出错 %s", err)
			tx.Rollback()
		}
	}
	tx.Commit()

}
