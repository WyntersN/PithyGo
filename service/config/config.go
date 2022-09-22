/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-09-22 13:55:11
 * @LastEditTime: 2022-09-22 14:47:25
 * @FilePath: \PithyGo\service\config\config.go
 */
package config

const (
	ConfigEnv  = "GVA_CONFIG"
	ConfigFile = "config/conf.yaml"
)

type Server struct {
	Zap       Zap       `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis     Redis     `mapstructure:"redis" json:"redis" yaml:"redis"`
	DB        DB        `mapstructure:"db" json:"db" yaml:"db"`
	AutoCode  Autocode  `mapstructure:"autoCode" json:"autoCode" yaml:"autoCode"`
	App       APP       `mapstructure:"app" json:"app" yaml:"app"`
	WebSocket WebSocket `mapstructure:"websocket" json:"websocket" yaml:"websocket"`
}
