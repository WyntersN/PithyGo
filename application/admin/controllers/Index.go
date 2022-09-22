/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-04-25 16:54:18
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2022-09-22 15:02:20
 */
package controllers

import (
	"PithyGo/application/admin"
	templates_admin "PithyGo/templates/admin"
)

type AdminController struct {
	admin.AdminController
}

func (c *AdminController) Get() {

	tmpl := &templates_admin.Index{
		Vars: map[string]interface{}{"title": ""},
	}
	admin.ExecuteTemplate(c.Ctx, tmpl)
	//model.CreatUser(&c.Ctx)
	//c.Ctx.HTML("ok")

}
