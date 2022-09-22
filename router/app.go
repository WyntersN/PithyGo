/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-04-25 16:54:18
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2022-09-16 11:49:29
 */
package router

import (
	api "PithyGo/application/app"
	"PithyGo/application/app/controllers"
	"PithyGo/common"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func routeApp(app *iris.Application) {

	mvc.New(app.Party("/")).
		Handle(new(controllers.AppController))
	mvc.New(app.Party("/api/login")).
		Register(common.SessManager.Start).
		Handle(new(controllers.AppLoginController))

	mvc.New(app.Party("/api/upload", api.SessionLoginSignAuthApp)).
		Register(common.SessManager.Start).
		Handle(new(controllers.AppLoginController))
	mvc.New(app.Party("/api/member", api.SessionLoginSignAuthApp)).
		Register(common.SessManager.Start).
		Handle(new(controllers.AppMemberController))

}
