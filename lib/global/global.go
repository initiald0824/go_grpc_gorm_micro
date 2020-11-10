package global

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"go_grpc_gorm_micro/config"
	"github.com/spf13/viper"
)

var (
	CURD_DB     *gorm.DB
	CURD_CONFIG config.Server
	CURD_VP		*viper.Viper
	CURD_LOG    *zap.Logger
)
