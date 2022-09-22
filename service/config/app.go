/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-09-22 14:36:00
 * @LastEditTime: 2022-09-22 14:56:04
 * @FilePath: \PithyGo\service\config\app.go
 */
package config

type APP struct {
	Domain             string `mapstructure:"domain" json:"domain" yaml:"domain"`
	Port               string `mapstructure:"port" json:"port" yaml:"port"`
	DefaultLayout      string `mapstructure:"DefaultLayout" json:"DefaultLayout" yaml:"DefaultLayout"`
	DefaultError       string `mapstructure:"DefaultError" json:"DefaultError" yaml:"DefaultError"`
	SessionCoolieName  string `mapstructure:"SessionCoolieName" json:"SessionCoolieName" yaml:"SessionCoolieName"`
	SessionExpires     int64  `mapstructure:"SessionExpires" json:"SessionExpires" yaml:"SessionExpires"`
	AppMode            string `mapstructure:"AppMode" json:"AppMode" yaml:"AppMode"`
	UploadSize         uint   `mapstructure:"UploadSize" json:"UploadSize" yaml:"UploadSize"`
	UploadSuffixExists string `mapstructure:"UploadSuffixExists" json:"UploadSuffixExists" yaml:"UploadSuffixExists"`
	MaxOpenConns       string `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
}
