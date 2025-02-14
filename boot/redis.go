package boot

import (
	"strings"

	"github.com/redis/go-redis/v9"
)

func SetupRedis() *redis.Client {
	cfg := Config.Redis
	Redis = redis.NewClient(&redis.Options{
		Addr:     strings.Join([]string{cfg.Addr, cfg.Port}, ":"),
		Password: cfg.Password, // 没有密码，默认值
		DB:       cfg.DB,       // 默认DB 0
	})

	return Redis
}
