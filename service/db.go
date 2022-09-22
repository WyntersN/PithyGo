/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2021-04-19 18:22:02
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2022-09-22 14:29:03
 */
package service

import (
	"PithyGo/common"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DBConfig struct {
	DriverName   string
	Host         string
	Port         uint
	Database     string
	User         string
	Password     string
	Prefix       string
	Charset      string
	LogMode      uint
	MaxIdleConns uint
	MaxOpenConns uint
}

func (c *DBConfig) InitDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4,utf8&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.Prefix,
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.LogLevel(c.LogMode)),
	})
	if err != nil {
		LOG.Sugar().Fatal(err)
		os.Exit(-1)
	}
	//println("------------", c.MaxIdleConns, c.MaxOpenConns)
	// db.SingularTable(true)                  //全局设置表名不可以为复数形式。
	// db.DB().SetMaxIdleConns(c.MaxIdleConns) //空闲时最大的连接数
	// db.DB().SetMaxOpenConns(c.MaxOpenConns) //最大的连接数

	//指定表前缀，修改默认表名
	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return c.Prefix + defaultTableName
	// }
	common.DataPrefix = c.Prefix
	LOG.Info("DB服务连接成功")
	return db
}
