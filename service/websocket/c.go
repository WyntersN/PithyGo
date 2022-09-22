/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-10-21 22:08:05
 * @LastEditTime: 2022-09-22 14:43:51
 * @FilePath: \PithyGo\service\websocket\c.go
 */
package websocket

import (
	"PithyGo/service"
	"PithyGo/service/websocket/helper"
	"PithyGo/service/websocket/models"

	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	defaultAppId = 101 // 默认平台Id
)

var (
	clientManager = NewClientManager()                    // 管理者
	appIds        = []uint32{defaultAppId, 102, 103, 104} // 全部的平台

	serverIp   string
	serverPort string
)

func GetAppIds() []uint32 {

	return appIds
}

func GetServer() (server *models.Server) {
	server = models.NewServer(serverIp, serverPort)

	return
}

func IsLocal(server *models.Server) (isLocal bool) {
	if server.Ip == serverIp && server.Port == serverPort {
		isLocal = true
	}

	return
}

func InAppIds(appId uint32) (inAppId bool) {

	for _, value := range appIds {
		if value == appId {
			inAppId = true

			return
		}
	}

	return
}

func GetDefaultAppId() (appId uint32) {
	appId = defaultAppId

	return
}

// 启动程序
func Initialization() {

	//webSocketPort := viper.GetString("app.webSocketPort")
	Register("login", LoginController)
	Register("heartbeat", HeartbeatController)
	Register("ping", PingController)

	serverIp = helper.GetServerIp()
	var (
		webSocketPort = service.CONFIG.WebSocket.Port
		rpcPort       = service.CONFIG.WebSocket.RpcPort
	)

	http.HandleFunc("/", wsPage)

	// 添加处理程序
	go clientManager.start()
	service.LOG.Sugar().Info("WebSocket 启动程序成功", serverIp, rpcPort)

	err := http.ListenAndServe(":"+webSocketPort, nil)
	if err != nil {
		service.LOG.Error(err.Error())
	}
}

func wsPage(w http.ResponseWriter, req *http.Request) {

	// 升级协议
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		//	service.LOG.Sugar().Info("升级协议", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])
		return true
	}}).Upgrade(w, req, nil)
	if err != nil {
		//	http.NotFound(w, req)
		return
	}

	service.LOG.Sugar().Info("webSocket 建立连接:", conn.RemoteAddr().String())

	currentTime := uint64(time.Now().Unix())
	client := NewClient(conn.RemoteAddr().String(), conn, currentTime)

	go client.read()
	go client.write()

	// 用户连接事件
	clientManager.Register <- client
}
