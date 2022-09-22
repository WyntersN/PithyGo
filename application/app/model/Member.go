/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-08-05 16:45:03
 * @LastEditTime: 2022-09-16 09:49:22
 * @FilePath: \PithyGo\app\app\model\Member.go
 */
package model

import (
	"PithyGo/common"
	"PithyGo/models/member"
	"PithyGo/models/structure"
	"PithyGo/service"

	"github.com/kataras/iris/v12"
)

var _member member.Interface

func GetMemberInfo(c iris.Context) (int, string, interface{}) {

	_member = &member.Content{
		MemberID: common.SessManager.Start(c).Get("MEMBER_ID").(uint),
	}
	if data, err := _member.GetMemberByMID(); err == nil {
		return 200, "success", data
	} else {
		return -1, err.Error(), nil
	}
}
func GetMyCash(c iris.Context) (int, string, interface{}) {
	_member = &member.Content{
		MemberID: common.SessManager.Start(c).Get("MEMBER_ID").(uint),
	}
	return 200, "success", _member.GetMyCash()
}

func MemberCashApply(c iris.Context) (int, string) {

	var fromJson structure.MemberCash
	if err := c.ReadJSON(&fromJson); err != nil {
		service.LOG.Sugar().Error("MemberCashApply %s", fromJson)
	}

	type validateType struct {
		Realname       *string `json:"realname"  form:"realname" validate:"required"`
		PaymentAccount *string `json:"payment_account"   form:"payment_account"  validate:"required"`
		PaymentAmount  float32 `json:"payment_amount"    form:"payment_amount"   validate:"required"`
	}

	validate := &validateType{
		Realname:       fromJson.Realname,
		PaymentAccount: fromJson.PaymentAccount,
		PaymentAmount:  fromJson.PaymentAmount,
	}
	if err := common.Validate(validate); err != nil {
		return 102, err.Error()
	}

	_member = &member.Content{
		MemberID: common.SessManager.Start(c).Get("MEMBER_ID").(uint),
	}
	if err := _member.CashApply(fromJson); err == nil {
		return 200, "success"
	} else {
		return -1, err.Error()
	}
}
