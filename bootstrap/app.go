/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-04-25 16:54:18
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2022-09-22 15:40:30
 */

package bootstrap

import (
	"PithyGo/Import"
	"PithyGo/common"
	"PithyGo/router"
	"PithyGo/service"
	"net/http"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/sessions"
)

var app = iris.New()

func Run() {

	app.Configure(iris.WithConfiguration(iris.YAML("./config/iris.yaml")))
	app.Logger().SetLevel(service.CONFIG.App.AppMode)
	if service.CONFIG.App.AppMode == "debug" {
		app.Use(recover.New())
		app.Use(logger.New())
	}
	//设置 sessions
	common.SessManager = sessions.New(sessions.Config{
		Cookie:       service.CONFIG.App.SessionCoolieName,
		Expires:      time.Duration(service.CONFIG.App.SessionExpires) * time.Hour,
		AllowReclaim: true,
	})

	//设置错误模版
	// app.OnErrorCode(iris.StatusNotFound, service.NotFound)
	// app.OnErrorCode(iris.StatusInternalServerError, service.InternalServerError)
	// app.Favicon("./favicon.ico")
	app.HandleDir("/", "./public")
	app.HandleDir("/upload", "./public/upload")
	app.HandleDir("/attachment", "./public/attachment")
	//路由
	router.InitRouter(app)
	srv := &http.Server{
		Addr:         service.CONFIG.App.Domain,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 600 * time.Second,
	}
	Import.Test()
	err := app.Run(iris.Server(srv))
	if err != nil {
		service.LOG.Sugar().Fatalf("IRIS服务已关闭 %s", err)
	}

}
