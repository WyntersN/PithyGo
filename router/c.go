/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-08-03 17:05:02
 * @LastEditTime: 2022-09-13 18:35:08
 * @FilePath: \PithyGo\router\c.go
 */
package router

import (
	"PithyGo/common"
	"PithyGo/service"
	"net/http"

	"github.com/kataras/iris/v12"
)

func InitRouter(app *iris.Application) {

	app.AllowMethods(iris.MethodOptions)

	app.Use(common.SessManager.Handler(), func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization,Token,User-Agent")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Content-Type,User-Agent")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		if ctx.Method() == "OPTIONS" {
			ctx.StatusCode(http.StatusOK)
			return
		}
		ctx.Next()
	})
	{
		routeAdmin(app)
		service.LOG.Info("Admin 路由注册成功")
	}
	{
		routeApp(app)
		service.LOG.Sugar().Info("App 路由注册成功")

	}

}
