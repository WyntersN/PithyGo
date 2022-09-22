/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-04-25 16:54:18
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2022-09-13 18:25:39
 */
package controllers

import (
	api "PithyGo/application/admin"
	"PithyGo/application/admin/model/user"
)

type AdminUserController struct {
	api.AdminController
}

func (c *AdminUserController) GetInfo() {
	var retData api.ReturnJson
	data, err := user.Info(c.Ctx)
	if err == nil {
		retData.Code = 200
		retData.Message = "success"
		retData.Data = data
	} else {
		retData.Code = -1
		retData.Message = err.Error()
	}

	c.Ctx.JSON(retData)

}

func (c *AdminUserController) PostLoginOut() {
	var retData api.ReturnJson

	retData.Code = 200
	retData.Message = "success"
	c.Ctx.JSON(retData)

}

func (c *AdminUserController) PostRepass() {
	var retData api.ReturnJson
	retData.Code, retData.Message = user.ReUserPass(c.Ctx)
	c.Ctx.JSON(retData)
}

func (c *AdminUserController) GetStaffList() {
	var retData api.ReturnJson

	retData.Data, _ = user.List(c.Ctx)

	retData.Code = 200
	retData.Message = "success"
	c.Ctx.JSON(retData)

}

func (c *AdminUserController) PostMessage() {
	var retData api.ReturnJson
	retData.Data = user.GetUserMessage(c.Ctx)
	retData.Code = 200
	retData.Message = "success"
	c.Ctx.JSON(retData)
}
func (c *AdminUserController) PostMessageRead() {
	var retData api.ReturnJson
	retData.Data = user.UserMessageRead(c.Ctx)
	retData.Code = 200
	retData.Message = "success"
	c.Ctx.JSON(retData)
}
