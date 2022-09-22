/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-04-19 18:22:01
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2022-09-22 14:55:05
 */
package bootstrap

import (
	"PithyGo/service"
	"PithyGo/service/s_init"
	"PithyGo/service/websocket"
	"PithyGo/service/websocket/servers/grpcserver"
	"PithyGo/service/websocket/servers/task"

	"go.uber.org/zap"
)

func init() {

	service.VP = service.Viper() // 初始化Viper
	service.LOG = service.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(service.LOG)
	//初始化缓存
	service.InitCache()

	defer func() {
		err := recover() // recover() 捕获panic异常，获得程序执行权。
		if err != nil {
			service.LOG.Sugar().Error("recover 错误拦截：", err) // runtime error: index out of range
		}
	}()

	//	database := config.GetString("db.selectDatabase")

	//	if database == "mongo" {
	//使MallGo支持mongoDB
	// dbConfig := service.MongoDBConfig{
	// 	config.GetString("mongo.host"), ////
	// 	config.GetString("mongo.port"),
	// 	config.GetString("mongo.database"),
	// }
	// service.MongoDB = dbConfig.InitMongoDB()
	//}

	//	使MallGo支持MYSQL
	//if database == "mysql" {
	dbConfig := &service.DBConfig{
		DriverName:   service.CONFIG.DB.DriverName, ////
		Host:         service.CONFIG.DB.Host,       ////
		Port:         service.CONFIG.DB.Port,
		Database:     service.CONFIG.DB.Database,
		User:         service.CONFIG.DB.User,
		Password:     service.CONFIG.DB.Password,
		Prefix:       service.CONFIG.DB.Prefix,
		Charset:      service.CONFIG.DB.Charset,
		LogMode:      service.CONFIG.DB.LogMode,
		MaxIdleConns: service.CONFIG.DB.MaxIdleConns,
		MaxOpenConns: service.CONFIG.DB.MaxOpenConns,
	}
	service.DB = dbConfig.InitDB()
	//service.NewRedisClient()
	s_init.Initialization()
	initWSConfig()
	//}
}

func initWSConfig() {

	// websocket 定时任务
	task.Init()

	// websocket 服务注册
	task.ServerInit()

	go websocket.Initialization()
	// websocket grpc
	go grpcserver.Init()
	//fmt.Println("config app:", viper.Get("app"))
	//fmt.Println("config redis:", viper.Get("redis"))

}
