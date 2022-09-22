/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-05-01 20:50:12
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2022-09-13 18:26:56
 */
package controllers

import (
	api "PithyGo/application/admin"
	"PithyGo/application/admin/model/user"
)

type AdminLoginController struct {
	api.AdminController
}

func (c *AdminLoginController) PostSign() {
	var retData api.ReturnJson
	data, err := user.Login(c.Ctx)

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
