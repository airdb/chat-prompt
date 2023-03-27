package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"server/config"
)

var (
	CHPT_VP                  *viper.Viper
	CHPT_CONFIG              config.Server
	CHPT_LOG                 *zap.Logger
	CHPT_DB                  *gorm.DB
	CHPT_REDIS               *redis.Client
	CHPT_Concurrency_Control = &singleflight.Group{}
)
