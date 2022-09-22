/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-09-23 00:06:47
 * @LastEditTime: 2022-09-13 18:26:50
 * @FilePath: \PithyGo\app\admin\controllers\System.go
 */
package controllers

import (
	api "PithyGo/application/admin"

	"github.com/bitly/go-simplejson"
)

type AdminSystemController struct {
	api.AdminController
}

func (c *AdminSystemController) GetVer() {
	var retData api.ReturnJson

	retData.Code = 200
	retData.Message = "success"
	retData.Data = api.Version

	c.Ctx.JSON(retData)
}

func (c *AdminSystemController) GetMenu() {
	var retData api.ReturnJson

	retData.Code = 200
	retData.Message = "success"

	json := ``

	j, _ := simplejson.NewJson([]byte(json))

	retData.Data = j.Interface()
	c.Ctx.JSON(retData)
}
