/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-04-26 16:29:35
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2022-09-13 18:38:21
 */
package router

import (
	"PithyGo/application/admin"
	"PithyGo/application/admin/controllers"
	"PithyGo/common"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func routeAdmin(app *iris.Application) {
	//登录路由
	mvc.New(app.Party("/admin")).
		Handle(new(controllers.AdminController))
	//不需要密钥路由
	mvc.New(app.Party("/admin/api/login")).
		Register(common.SessManager.Start).
		Handle(new(controllers.AdminLoginController))

	//需要密钥路由
	mvc.New(app.Party("/admin/api/upload", admin.SessionLoginSignAuthApp)).
		Register(common.SessManager.Start).
		Handle(new(controllers.AdminUploadController))

	mvc.New(app.Party("/admin/api/user", admin.SessionLoginSignAuthApp)).
		Register(common.SessManager.Start).
		Handle(new(controllers.AdminUserController))

	//app.Any("/api/ws", websocket.Handler(v1App.WS))

	//websocket
	// 定时任务
	// task.Init()
	// websocket.WebsocketInit()

	// websocketAPI := app.Party("/ws", api.SessionLoginSignAuthApp)
	// mvc.New(websocketAPI).Register(api.SessManager.Start)
	// websocketAPI.Get("/", websocket.StartWebSocket)

	// // 服务注册
	// task.ServerInit()

	// // grpc
	// go grpcserver.Init()

	//r.Any("/ws", websocket.StartWebSocket, api.SessionLoginSignAuthApp)

}
