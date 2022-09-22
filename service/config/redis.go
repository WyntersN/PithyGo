/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-09-22 14:13:48
 * @LastEditTime: 2022-09-22 14:27:40
 * @FilePath: \PithyGo\service\config\redis.go
 */
package config

type Redis struct {
	DB           int    `mapstructure:"db" json:"db" yaml:"db"`
	Addr         string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	PoolSize     int    `mapstructure:"PoolSize" json:"PoolSize" yaml:"PoolSize"`
	MinIdleConns int    `mapstructure:"MinIdleConns" json:"MinIdleConns" yaml:"MinIdleConns"`
}
