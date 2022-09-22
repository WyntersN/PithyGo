/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-10-10 13:23:03
 * @LastEditTime: 2022-09-13 18:26:20
 * @FilePath: \PithyGo\app\admin\controllers\Upload.go
 */
package controllers

import (
	api "PithyGo/application/admin"
	"PithyGo/application/admin/model"
)

type AdminUploadController struct {
	api.AdminController
}

type returnJson struct {
	Code     int                    `json:"code"`
	Message  string                 `json:"message"`
	Data     map[string]interface{} `json:"data"`
	Url      interface{}            `json:"url"`
	Uploaded bool                   `json:"uploaded"`
}

func (c *AdminUploadController) Post() {
	var retData returnJson

	retData.Code, retData.Message, retData.Data = model.Upload(c.Ctx)
	if retData.Code == 200 {
		retData.Uploaded = true
		retData.Url = retData.Data["src"]
	}

	c.Ctx.JSON(retData)
}
