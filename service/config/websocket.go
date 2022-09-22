/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-09-22 14:41:05
 * @LastEditTime: 2022-09-22 14:42:59
 * @FilePath: \PithyGo\service\config\websocket.go
 */
package config

type WebSocket struct {
	Port    string `mapstructure:"port" json:"port" yaml:"port"`
	RpcPort string `mapstructure:"rpcPort" json:"rpcPort" yaml:"rpcPort"`
}
