package initialize

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"server/global"
)

func Redis() {
	redisCfg := global.CHPT_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.CHPT_CONFIG.Redis.Host, global.CHPT_CONFIG.Redis.Port),
		//Password: redisCfg.Password, // no password set
		DB: redisCfg.DB, // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.CHPT_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.CHPT_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.CHPT_REDIS = client
	}
}
