/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-08-03 17:10:16
 * @LastEditTime: 2022-09-16 09:58:28
 * @FilePath: \PithyGo\application\app\controllers\Login.go
 */
package controllers

import (
	"PithyGo/application/app"
	"PithyGo/application/app/model"
)

type AppLoginController struct {
	app.AppController
}

func (c *AppLoginController) PostSign() {
	var retData app.ReturnJson
	data, err := model.Login(c.Ctx)

	if err == nil {
		c.Session.Set("appAuthorization", data["authorization"])
		retData.Code = 200
		retData.Message = "success"
		retData.Data = data
	} else {
		retData.Code = -1
		retData.Message = err.Error()
	}

	c.Ctx.JSON(retData)
}

func (c *AppLoginController) PostRegister() {
	var retData app.ReturnJson
	retData.Code, retData.Message = model.Register(c.Ctx)
	c.Ctx.JSON(retData)
}

func (c *AppLoginController) PostForget() {
	var retData app.ReturnJson
	retData.Code, retData.Message = model.LoginForGet(c.Ctx)
	c.Ctx.JSON(retData)
}

func (c *AppLoginController) GetSendCode() {
	var retData app.ReturnJson
	retData.Code, retData.Message = model.LoginSendCode(c.Ctx)
	c.Ctx.JSON(retData)
}
