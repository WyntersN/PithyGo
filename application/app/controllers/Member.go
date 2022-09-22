/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-08-05 16:48:22
 * @LastEditTime: 2022-09-16 09:54:41
 * @FilePath: \PithyGo\app\app\controllers\Member.go
 */
package controllers

import (
	"PithyGo/application/app"
	"PithyGo/application/app/model"
)

type AppMemberController struct {
	app.AppController
}

func (c *AppMemberController) GetInfo() {
	var retData app.ReturnJson
	retData.Code, retData.Message, retData.Data = model.GetMemberInfo(c.Ctx)
	c.Ctx.JSON(retData)
}
func (c *AppMemberController) PostCashApply() {
	var retData app.ReturnJson
	retData.Code, retData.Message = model.MemberCashApply(c.Ctx)
	c.Ctx.JSON(retData)
}
func (c *AppMemberController) GetCash() {
	var retData app.ReturnJson
	retData.Code, retData.Message, retData.Data = model.GetMyCash(c.Ctx)
	c.Ctx.JSON(retData)
}
