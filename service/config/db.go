/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-09-22 14:13:48
 * @LastEditTime: 2022-09-22 14:23:55
 * @FilePath: \PithyGo\service\config\db.go
 */
package config

type DB struct {
	DriverName   string `mapstructure:"DriverName" json:"DriverName" yaml:"DriverName"`
	Host         string `mapstructure:"host" json:"host" yaml:"host"`
	Port         uint   `mapstructure:"port" json:"port" yaml:"port"`
	Database     string `mapstructure:"database" json:"database" yaml:"database"`
	User         string `mapstructure:"user" json:"user" yaml:"user"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Charset      string `mapstructure:"charset" json:"charset" yaml:"charset"`
	LogMode      uint   `mapstructure:"LogMode" json:"LogMode" yaml:"LogMode"`
	MaxIdleConns uint   `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns uint   `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
}
