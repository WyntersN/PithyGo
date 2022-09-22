/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-10-22 15:33:45
 * @LastEditTime: 2021-11-14 21:39:33
 * @FilePath: \PithyGo\service\websocket\models\request_model.go
 */
/**
 * Created by GoLand.
 * User: link1st
 * Date: 2019-07-27
 * Time: 14:41
 */

package models

/************************  请求数据  **************************/
// 通用请求数据格式
type Request struct {
	Seq  string      `json:"seq"`            // 消息的唯一Id
	Cmd  string      `json:"cmd"`            // 请求命令字
	Data interface{} `json:"data,omitempty"` // 数据 json
}

// 登录请求数据
type Login struct {
	Token  string `json:"token"` // 验证用户是否登录
	AppId  uint32 `json:"appid,omitempty"`
	UserId string `json:"userid,omitempty"`
}

// 心跳请求数据
type HeartBeat struct {
	UserId string `json:"userId,omitempty"`
}
