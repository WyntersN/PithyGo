package service

import (
	"PithyGo/service/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	LOG    *zap.Logger
	VP     *viper.Viper
	CONFIG config.Server
)
